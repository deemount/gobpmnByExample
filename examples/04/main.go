package main

import (
	gobpmn_builder "github.com/deemount/gobpmnBuilder"
)

// Main ...
func main() {

	exampleProcess := New().Build().Call()

	builder, err := gobpmn_builder.New(
		gobpmn_builder.WithPath(),
		gobpmn_builder.WithCounter(),
		gobpmn_builder.WithDefinitions(exampleProcess))
	if err != nil {
		panic(err)
	}
	builder.Build()

	file := builder.GetCurrentlyCreatedFile()

	// Read the file and unmarshal the data
	_, err = reader.ReadFileAndUnmarshal(file)
	if err != nil {
		panic(err)
	}

}
