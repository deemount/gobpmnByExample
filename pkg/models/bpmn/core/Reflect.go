package core

import (
	"reflect"
	"strings"
)

var (
	fieldLong  = "definitions"
	fieldShort = "def"
)

// IsDefinitions ...
func IsDefinitions(field reflect.StructField) bool {
	return strings.ToLower(field.Name) == fieldShort || strings.ToLower(field.Name) == fieldLong
}

// IsNotDefinitions ...
func IsNotDefinitions(field reflect.StructField) bool {
	return strings.ToLower(field.Name) != fieldShort && strings.ToLower(field.Name) != fieldLong
}
