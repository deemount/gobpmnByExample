package factory

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"reflect"
	"strings"

	"github.com/deemount/gobpmnByExamples/pkg/models/bpmn/core"
	"github.com/deemount/gobpmnByExamples/pkg/utils"
)

type Def core.DefinitionsRepository

// Builder ...
type Builder struct {
	Reflect
	Settings
	Map       map[string][]interface{}
	Suffix    string
	HashTable []string
}

// Hash ...
func (h *Builder) Hash() string {
	if h.isZero() {
		r, _ := h.hash()
		h.Suffix = r.Suffix
	}
	return h.Suffix
}

// Inject all anonymous fields with a hash value by fields with the type Builder
// It also setup reflected fields with boolean type and checks out the configuration by wording
func (h *Builder) Inject(p interface{}) interface{} { return h.inject(p) }

// Create receives the definitions repository by the app in the interface {} p argument
// and calls the main elements to set the maps, including process parameters
// n of process.
func (h *Builder) Build(p interface{}) { h.build(p) }

// hash ...
func (h Builder) hash() (Builder, error) {

	n := 8
	b := make([]byte, n)
	c := fnv.New32a()

	if _, err := rand.Read(b); err != nil {
		return Builder{}, err
	}
	s := fmt.Sprintf("%x", b)

	if _, err := c.Write([]byte(s)); err != nil {
		return Builder{}, err
	}
	defer c.Reset()

	r := Builder{Suffix: fmt.Sprintf("%x", string(c.Sum(nil)))}

	return r, nil
}

// inject itself reflects a given struct and inject
// signed fields with hash values.
// There are two conditions to assign fields of a strcut:
// a) The struct has anonymous fields
// b) The struct has no anymous fields
// It also counts the element in their specification to know
// how much elements of each package needs to be mapped later then.
func (h *Builder) inject(p interface{}) interface{} {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("factory.builder: inject recovered", r)
		}
	}()

	ref := NewReflect(p)
	ref.Interface().New().Maps().Reflection()

	// anonymous field are reflected
	if ref.hasAnonymous() {

		length := len(ref.Anonym)

		// create processMap, anonymMap and hashMap
		processMap := make(map[string]map[int][]interface{})
		anonymMap := make(map[int][]interface{}, length)
		hashMap := make(map[string][]interface{}, length)

		// walk through the map with names of anonymous fields
		for index, field := range ref.Anonym {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(field)

			// create the field map
			fieldMap := make(map[int][]interface{}, n.NumField())
			// create the hash slice
			hashSlice := make([]interface{}, n.NumField())

			// append to anonymMap the name of anonymous field
			anonymMap[index] = append(anonymMap[index], n.Type().Name())

			// walk through the values of fields assigned to the interface {}
			for i := 0; i < n.NumField(); i++ {

				// get the name of field by index of reflected value above and
				// append to fieldMap the name of field by index
				name := n.Type().Field(i).Name
				fieldMap[i] = append(fieldMap[i], name)

				// switch by kind of field value by index of reflected value above and
				// assign the hash value to the hash slice by index of field value above
				switch n.Field(i).Kind() {

				// kind is a bool
				case reflect.Bool:

					// only the first field, which IsExecutable is set to true,
					// means, only one process in a collaboration can be executed at runtime
					if strings.Contains(name, boolIsExecutable) && i == 0 {
						n.Field(0).SetBool(true)
						hashSlice[i] = bool(true)
					} else {
						hashSlice[i] = bool(false)
					}

					break

				// kind is a struct
				case reflect.Struct:

					h.countStartEvent(name) // counts startevents
					h.countTask(name)       // counts tasks
					h.countEndEvent(name)   // counts endevent

					// Injecting by each index of the given process structs
					// e.g. starts with customer support, customer,
					strHash := fmt.Sprintf("%s", n.Field(i).FieldByName("Suffix"))
					if strHash == "" {
						hash, _ := h.hash() // generate hash value
						hashSlice[i] = hash.Suffix
						n.Field(i).Set(reflect.ValueOf(hash)) // inject the field
					}

					if i+1 < n.NumField() {
						nexthash, _ := h.hash() // generate hash value
						hashSlice[i+1] = nexthash.Suffix
						n.Field(i + 1).Set(reflect.ValueOf(nexthash)) // inject the field
					}

					break

				}

			}
			mergeStringSliceToMap(hashMap, n.Type().Name(), hashSlice)
			processMap[n.Type().Name()] = MergeMaps(anonymMap, fieldMap)
		}
	}

	// anonymous field aren't reflected
	if ref.hasNotAnonymous() {

		// walk through the map with names of builder fields
		for _, builderField := range ref.Builder {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(builderField)

			if strings.Contains(builderField, fieldProcess) {
				h.NumProc++
			}

			if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == fieldStartEvent {
				h.NumStartEvent++
				h.NumShape++
			}

			if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == fieldTask {
				h.NumTask++
				h.NumShape++
			}

			if strings.Contains(builderField, fieldEndEvent) {
				h.NumEndEvent++
				h.NumShape++
			}

			hash, _ := h.hash()          // generate hash value
			n.Set(reflect.ValueOf(hash)) // inject the field

		}

		// walk through the map with names of boolean fields
		for _, configField := range ref.Config {

			// get the reflected value of field
			n2 := ref.Temporary.FieldByName(configField)

			// only the first field, which IsExecutable is set to true
			n2.SetBool(true)

		}

	}

	p = ref.Set()
	ref.countWords()

	return p

}

// build contains the reflected process definition (p interface{})
// and can call SetMainElements from the core package or
// calls it by the reflected method name.
// This method hides some of the setters by building the BPMN
// with reflection
func (h *Builder) build(p interface{}) {

	// el is the interface {}
	el := reflect.ValueOf(&p).Elem()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	definitions := reflect.New(el.Elem().Type()).Elem() // *core.Definitions
	definitions.Set(el.Elem())                          // reflected process definitions el will be assigned to the core definitions

	// set process and diagram
	process := definitions.MethodByName(MethodSetProcess)
	process.Call([]reflect.Value{reflect.ValueOf(h.NumProc)}) // h.numProc represents number of processes

	diagram := definitions.MethodByName(MethodSetDiagram)
	diagram.Call([]reflect.Value{reflect.ValueOf(1)})

}

// isStartEvent ...
func (h *Builder) countStartEvent(builderField string) {
	if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == "" {
		h.NumStartEvent++
		h.NumShape++
	}
}

// isTask ...
func (h *Builder) countTask(builderField string) {
	if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == "" {
		h.NumTask++
		h.NumShape++
	}
}

// isEndEvent ...
func (h *Builder) countEndEvent(builderField string) {
	if strings.Contains(builderField, fieldEndEvent) {
		h.NumEndEvent++
		h.NumShape++
	}
}

// isZero ...
func (h *Builder) isZero() bool { return h.Suffix == "" }

// isUnknown ...
func (h *Builder) isUnknown(field string) bool { return true }

func MergeMaps[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

func mergeStringSliceToMap(m map[string][]interface{}, k string, v []interface{}) {
	if m[k] == nil {
		m[k] = make([]interface{}, len(v))
		for i, s := range v {
			m[k][i] = interface{}(s)
		}
	} else {
		m[k] = append(m[k], v...)
	}
}
