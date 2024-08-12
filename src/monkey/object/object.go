package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
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
	return fmt.Sprintf("%t", iob.Value)
}
