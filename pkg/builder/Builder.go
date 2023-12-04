package factory

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"reflect"
	"strings"

	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/core"
	"github.com/deemount/gobpmnByExample/pkg/utils"
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

// Build receives the definitions repository by the app in the interface {} p argument
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
