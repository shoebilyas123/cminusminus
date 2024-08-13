package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
	RETURN_OBJ  = "RETURN VALUE"
	ERROR_OBJ   = "ERROR_OBJ"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type IntegerObject struct {
	Value int64
}

func (iob *IntegerObject) Type() ObjectType {
	return INTEGER_OBJ
}
func (iob *IntegerObject) Inspect() string {
	return fmt.Sprintf("%d", iob.Value)
}

type BooleanObject struct {
	Value bool
}

func (bo *BooleanObject) Type() ObjectType { return BOOLEAN_OBJ }
func (bo *BooleanObject) Inspect() string  { return fmt.Sprintf("%t", bo.Value) }

type NullObject struct{}

func (nullo *NullObject) Type() ObjectType { return NULL_OBJ }
func (nullo *NullObject) Inspect() string  { return "null" }

type ReturnObject struct {
	Value Object
}

func (ro *ReturnObject) Type() ObjectType { return RETURN_OBJ }
func (ro *ReturnObject) Inspect() string  { return ro.Value.Inspect() }

type ErrorObject struct {
	Message string
}

func (eo *ErrorObject) Type() ObjectType { return ERROR_OBJ }
func (eo *ErrorObject) Inspect() string  { return "ERROR: " + eo.Message }
