package example

import "github.com/deemount/gobpmnByExample/pkg/builder"

var Builder builder.Builder

type Proxy interface{ Build() Process }
