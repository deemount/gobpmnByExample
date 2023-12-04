package main

import (
	_ "embed"
	"log"
	"time"

	factory "github.com/deemount/gobpmnByExample/pkg/factory"
)

// bpmnFactory ...
var bpmnFactory factory.BpmnFactory

// init ...
func init() {
	bpmnFactory = factory.NewBpmnFactory()
}

// main ...
func main() {

	start := time.Now()

	_, err := bpmnFactory.Build()
	if err != nil {
		panic(err)
	}

	log.Println("total time:", time.Since(start))

}
