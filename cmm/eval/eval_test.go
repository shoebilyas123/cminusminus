package eval

import (
	"testing"

	"github.com/shoebilyas123/cminusminus/cmm/lexer"
	"github.com/shoebilyas123/cminusminus/cmm/object"
	"github.com/shoebilyas123/cminusminus/cmm/parser"
)

// func TestEvalIntegerExpression(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected int64
// 	}{
// 		{"64", 64},
// 		{"5", 5},
// 		{"10", 10},
// 		{"-5", -5},
// 		{"-10", -10},
// 		{"5 + 5 + 5 + 5 - 10", 10},
// 		{"2 * 2 * 2 * 2 * 2", 32},
// 		{"-50 + 100 + -50", 0},
// 		{"5 * 2 + 10", 20},
// 		{"5 + 2 * 10", 25},
// 		{"20 + 2 * -10", 0},
// 		{"50 / 2 * 2 + 10", 60},
// 		{"2 * (5 + 10)", 30},
// 		{"3 * 3 * 3 + 10", 37},
// 		{"3 * (3 * 3) + 10", 37},
// 		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
// 	}

// 	for _, tt := range tests {
// 		evaluated := testEval(tt.input)
// 		fmt.Printf("RES: %+v\n", evaluated)
// 		testIntegerObject(t, evaluated, tt.expected)
// 	}
// }

// func TestBangOperator(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected bool
// 	}{
// 		{"!true", false},
// 		{"!false", true},
// 		{"!5", false},
// 		{"!!true", true},
// 		{"!!false", false},
// 		{"!!5", true},
// 		{"true", true},
// 		{"false", false},
// 		{"1 < 2", true},
// 		{"1 > 2", false},
// 		{"1 < 1", false},
// 		{"1 > 1", false},
// 		{"1 == 1", true},
// 		{"1 != 1", false},
// 		{"1 == 2", false},
// 		{"1 != 2", true},
// 	}
// 	for _, tt := range tests {
// 		evaluated := testEval(tt.input)
// 		testBooleanObject(t, evaluated, tt.expected)
// 	}
// }

// func TestMinusOperator(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected int64
// 	}{
// 		{"-5", -5},
// 		{"-10", -10},
// 		{"20", 20},
// 		{"-20", -20},
// 	}
// 	for _, tt := range tests {
// 		evaluated := testEval(tt.input)
// 		testIntegerObject(t, evaluated, tt.expected)
// 	}
// }

func testEval(i string) object.Object {
	l := lexer.New(i)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testIntegerObject(t *testing.T, ev object.Object, exp int64) bool {
	res, ok := ev.(*object.IntegerObject)

	if !ok {
		t.Errorf("Object is not integer. Got=%t (%+v)\n", ev, ev)
	}

	if res.Value != exp {
		t.Errorf("Object is not integer. Got=%d, Want=%d\n", res.Value, exp)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, ev object.Object, exp bool) bool {
	res, ok := ev.(*object.BooleanObject)

	if !ok {
		t.Errorf("Object is not Boolean. Got=%t (%+v)\n", ev, ev)
	}

	if res.Value != exp {
		t.Errorf("Object has the wrong value assigned. Got=%t, Want=%t\n", res.Value, exp)
		return false
	}

	return true
}

// Testing conditionals
// func TestIfElseExpressions(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected interface{}
// 	}{
// 		{"if (true) { 10 }", 10},
// 		{"if (false) { 10 }", nil},
// 		{"if (1) { 10 }", 10},
// 		{"if (1 < 2) { 10 }", 10},
// 		{"if (1 > 2) { 10 }", nil},
// 		{"if (1 > 2) { 10 } else { 20 }", 20},
// 		{"if (1 < 2) { 10 } else { 20 }", 10},
// 	}

//		for _, tt := range tests {
//			evaluated := testEval(tt.input)
//			integer, ok := tt.expected.(int)
//			if ok {
//				testIntegerObject(t, evaluated, int64(integer))
//			} else {
//				testNullObject(t, evaluated)
//			}
//		}
//	}
func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

// func TestReturnStatements(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected int64
// 	}{
// 		{"return 10;", 10},
// 		{"return 10; 9;", 10},
// 		{"return 2 * 5; 9;", 10},
// 		{"9; return 2 * 5; 9;", 10},
// 		{" if (10 > 1) { if (10 > 1) {return 10;} return 1;}", 10},
// 	}
// 	for _, tt := range tests {
// 		evaluated := testEval(tt.input)
// 		testIntegerObject(t, evaluated, tt.expected)
// 	}
// }

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"ttrue + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`if (10 > 1) {if (10 > 1) {return true + false;}return 1;}`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.ErrorObject)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestClosures(t *testing.T) {
	input := `
let newAdder = fn(x) {
fn(y) { x + y };
};
let addTwo = newAdder(2);
addTwo(2);`
	testIntegerObject(t, testEval(input), 4)
}
