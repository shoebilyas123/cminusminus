package eval

import (
	"fmt"

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
		return evalProgram(node)
	case *ast.IntegerLiteral:
		return &object.IntegerObject{Value: node.Value}
	case *ast.BooleanExpression:
		return getBooleanObject(node.Value)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)

		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
		if isError(left) {
			return left
		}

		right := Eval(node.Right)

		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, right, left)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
	case *ast.IfExpression:
		return evalIfExpression(node)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue)
		if isError(val) {
			return val
		}
		return &object.ReturnObject{Value: val}
	}

	return nil
}

func evalProgram(node *ast.Program) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Eval(statement)
		switch result := result.(type) {
		case *object.ReturnObject:
			return result.Value
		case *object.ErrorObject:
			return result
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}
	return result
}

func evalIfExpression(node *ast.IfExpression) object.Object {
	condition := Eval(node.Condition)

	switch condition.(type) {
	case *object.ErrorObject:
		return condition
	default:
		if isTruthy(condition) {
			return Eval(node.Consequence)
		} else if node.Alternative != nil {
			return Eval(node.Alternative)
		}
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
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), op, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), op, right.Type())

	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
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
		return newError("unknown operator: %s %s %s",
			left.Type(), op, right.Type())
	}

}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", op, right.Type())
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	num, ok := right.(*object.IntegerObject)

	if !ok {
		return NULL
	}

	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
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

func getBooleanObject(i bool) *object.BooleanObject {
	if i {
		return TRUE
	}

	return FALSE
}

func newError(format string, a ...interface{}) *object.ErrorObject {
	return &object.ErrorObject{Message: fmt.Sprintf(format, a...)}
}
