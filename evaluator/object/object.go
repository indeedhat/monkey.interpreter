package object

import (
	"fmt"
)

type ObjectType string

const (
	IntegerObj ObjectType = "int"
	BoolObj    ObjectType = "bool"
	StringObj  ObjectType = "string"
	NullObj    ObjectType = "null"
	ReturnObj  ObjectType = "return"
	ErrObj     ObjectType = "error"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

// Inspect implements Object
func (i *Integer) Inspect() string {
	return fmt.Sprint(i.Value)
}

// Type implements Object
func (*Integer) Type() ObjectType {
	return IntegerObj
}

var _ Object = (*Integer)(nil)

type Boolean struct {
	Value bool
}

// Inspect implements Object
func (i *Boolean) Inspect() string {
	if i.Value {
		return "true"
	}

	return "false"
}

// Type implements Object
func (*Boolean) Type() ObjectType {
	return BoolObj
}

var _ Object = (*Boolean)(nil)

type String struct {
	Value string
}

// Inspect implements Object
func (s *String) Inspect() string {
	return s.Value
}

// Type implements Object
func (*String) Type() ObjectType {
	return StringObj
}

var _ Object = (*String)(nil)

type Null struct{}

// Inspect implements Object
func (*Null) Inspect() string {
	return "null"
}

// Type implements Object
func (*Null) Type() ObjectType {
	return NullObj
}

var _ Object = (*Null)(nil)

type ReturnValue struct {
	Value Object
}

// Inspect implements Object
func (r *ReturnValue) Inspect() string {
	return r.Value.Inspect()
}

// Type implements Object
func (*ReturnValue) Type() ObjectType {
	return ReturnObj
}

var _ Object = (*ReturnValue)(nil)

type Error struct {
	Message string
}

// Inspect implements Object
func (e *Error) Inspect() string {
	return "Error: " + e.Message
}

// Type implements Object
func (*Error) Type() ObjectType {
	return ErrObj
}

var _ Object = (*Error)(nil)
