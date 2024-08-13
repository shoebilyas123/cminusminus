package eval

import (
	"github.com/shoebilyas123/monkeylang/monkey/ast"
	"github.com/shoebilyas123/monkeylang/monkey/object"
)

var (
	NULL  = &object.NullObject{}
	TRUE  = &object.BooleanObject{Value: true}
	FALSE = &object.BooleanObject{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.IntegerLiteral:
		return &object.IntegerObject{Value: node.Value}
	case *ast.BooleanExpression:
		return getBooleanObject(node.Value)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}

	return nil
}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func getBooleanObject(i bool) *object.BooleanObject {
	if i {
		return TRUE
	}

	return FALSE
}
