package main

import (
	gobpmn_builder "github.com/deemount/gobpmnBuilder"
	gobpmn_counter "github.com/deemount/gobpmnCounter"
	gobpmn_hash "github.com/deemount/gobpmnHash"
	gobpmn_reader "github.com/deemount/gobpmnReader"
)

// Global variables ...
var (
	build  gobpmn_builder.Builder
	count  gobpmn_counter.Quantities
	hash   gobpmn_hash.Injection
	reader gobpmn_reader.Reader
)
