package eval

import (
	"testing"

	"github.com/shoebilyas123/monkeylang/monkey/lexer"
	"github.com/shoebilyas123/monkeylang/monkey/object"
	"github.com/shoebilyas123/monkeylang/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"64", 64},
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestMinusOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"-5", -5},
		{"-10", -10},
		{"20", 20},
		{"-20", -20},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(i string) object.Object {
	l := lexer.New(i)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
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
