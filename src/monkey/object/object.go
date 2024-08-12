package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
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
