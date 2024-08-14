package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/shoebilyas123/monkeylang/monkey/ast"
)

type ObjectType string

const (
	INTEGER_OBJ  = "INTEGER"
	BOOLEAN_OBJ  = "BOOLEAN"
	NULL_OBJ     = "NULL"
	RETURN_OBJ   = "RETURN VALUE"
	ERROR_OBJ    = "ERROR_OBJ"
	FUNCTION_OBJ = "FUNCTION"
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

type Function struct {
	Env        *Environment
	Body       *ast.BlockStatement
	Parameters []*ast.Identifier
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
