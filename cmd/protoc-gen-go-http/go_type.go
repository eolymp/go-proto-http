package main

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func isRepeated(field protoreflect.FieldDescriptor) bool {
	return field.Cardinality() == protoreflect.Repeated
}

func needsStar(kind protoreflect.Kind) bool {
	switch kind {
	case protoreflect.GroupKind, protoreflect.MessageKind, protoreflect.BytesKind:
		return true
	}
	return false
}

func GoType(field protoreflect.FieldDescriptor) (typ string) {
	switch field.Kind() {
	case protoreflect.DoubleKind:
		typ = "float64"
	case protoreflect.FloatKind:
		typ = "float32"
	case protoreflect.Int64Kind:
		typ = "int64"
	case protoreflect.Uint64Kind:
		typ = "uint64"
	case protoreflect.Int32Kind:
		typ = "int32"
	case protoreflect.Uint32Kind:
		typ = "uint32"
	case protoreflect.Fixed64Kind:
		typ = "uint64"
	case protoreflect.Fixed32Kind:
		typ = "uint32"
	case protoreflect.BoolKind:
		typ = "bool"
	case protoreflect.StringKind:
		typ = "string"
	case protoreflect.BytesKind:
		typ = "[]byte"
	case protoreflect.Sfixed32Kind:
		typ = "int32"
	case protoreflect.Sfixed64Kind:
		typ = "int64"
	case protoreflect.Sint32Kind:
		typ = "int32"
	case protoreflect.Sint64Kind:
		typ = "int64"
	default:
		panic(fmt.Errorf("unknown type for %v", field.Name()))
	}

	if isRepeated(field) {
		typ = "[]" + typ
	} else if needsStar(field.Kind()) {
		typ = "*" + typ
	}

	return
}
