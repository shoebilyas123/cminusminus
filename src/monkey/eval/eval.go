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
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, right, left)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	case *ast.IfExpression:
		return evalIfExpression(node)
	}

	return nil
}

func evalIfExpression(node *ast.IfExpression) object.Object {
	condition := Eval(node.Condition)

	if isTruthy(condition) {
		return Eval(node.Consequence)
	} else if node.Alternative != nil {
		return Eval(node.Alternative)
	}

	return NULL
}

func isTruthy(condition object.Object) bool {
	switch condition {
	case TRUE:
		return true
	case FALSE:
		return false
	case NULL:
		return false
	default:
		return true
	}
}

func evalInfixExpression(op string, right, left object.Object) object.Object {
	switch {
	case right.Type() == object.INTEGER_OBJ && left.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(op, right, left)
	case op == "==":
		return getBooleanObject(left == right)
	case op == "!=":
		return getBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(op string, right, left object.Object) object.Object {
	// 1.0: Assuming that the right and left expressions are integers;

	le_val := left.(*object.IntegerObject).Value
	re_val := right.(*object.IntegerObject).Value

	switch op {
	case "+":
		return &object.IntegerObject{Value: le_val + re_val}
	case "-":
		return &object.IntegerObject{Value: le_val - re_val}
	case "/":
		return &object.IntegerObject{Value: le_val / re_val}
	case "*":
		return &object.IntegerObject{Value: le_val * re_val}
	case "<":
		return getBooleanObject(le_val < re_val)
	case ">":
		return getBooleanObject(le_val > re_val)
	case "==":
		return getBooleanObject(left == right)
	case "!=":
		return getBooleanObject(left != right)
	default:
		return NULL
	}

}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return NULL
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	num, ok := right.(*object.IntegerObject)

	if !ok {
		return NULL
	}

	return &object.IntegerObject{Value: -num.Value}
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
