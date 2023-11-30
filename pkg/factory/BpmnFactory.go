package repository

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/deemount/gobpmnByExample/example"
	"github.com/deemount/gobpmnByExample/pkg/models/bpmn/core"
)

// BpmnFactory ...
type BpmnFactory interface {
	Build() (bpmnFactory, error)
	GetCurrentlyCreatedFile() string
}

// bpmnFactory ...
type bpmnFactory struct {
	Options BpmnOptions
	Repo    core.DefinitionsRepository
}

type BpmnFactoryOption func(o BpmnOptions) BpmnOptions

// NewBpmnFactory ...
func NewBpmnFactory(opt ...BpmnFactoryOption) BpmnFactory {
	//
	options := BpmnOptions{}
	for _, o := range opt {
		options = o(options)
	}
	// path to bpmn files
	path := "files/bpmn"
	// read the dir for created bpmn files
	files, err := os.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}

	// set number and count up for each created file
	if options.Counter == 0 {
		options.Counter = len(files)
	} else {
		options.Counter += 1
	}

	// set default name for bpmn-file
	options.CurrentFile = "diagram_" + fmt.Sprintf("%d", options.Counter)

	return &bpmnFactory{Options: options}
}

// set is a private method and is called inside Create().
// It calls the Create() method in the written business model,
// when it fits to the correct expectations
func (factory *bpmnFactory) set() {

	factory.Repo = example.New().Build().Call()

}

// Build...
func (factory bpmnFactory) Build() (bpmnFactory, error) {

	var err error

	factory.set()

	// create .bpmn
	err = factory.toBPMN()
	if err != nil {
		return bpmnFactory{}, err
	}

	return factory, nil

}

// GetCurrentlyCreatedFilename ...
func (factory bpmnFactory) GetCurrentlyCreatedFile() string {
	return factory.Options.CurrentFile
}

// toBPMN ...
func (factory *bpmnFactory) toBPMN() error {

	var err error

	// marshal xml to byte slice
	b, _ := xml.MarshalIndent(&factory.Repo, " ", "  ")

	// path to bpmn files
	path := "files/bpmn"
	// create .bpmn file
	f, err := os.Create(path + "/" + factory.Options.CurrentFile + ".bpmn")
	if err != nil {
		return err
	}
	defer f.Close()

	// add xml header
	w := []byte(fmt.Sprintf("%v", xml.Header+string(b)))

	// write bytes to file
	_, err = f.Write(w)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}

	// create .json
	err = factory.toJSON()
	if err != nil {
		return err
	}

	return nil

}
