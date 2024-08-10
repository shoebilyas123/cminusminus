package ast

import (
	"bytes"

	"github.com/shoebilyas123/monkeylang/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// expressionNode() and statementNode() are dummy methods to guide the go compiler
// TokenLiteral() stores the literal value of the token assocciated with a statement
type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

// Program node is the root node
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	//Name identifier is a Pointer because:
	// one identifier can be used at multiple places
	// so we need to keep track of it's value globally
	Value Expression
}

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }
func (l *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")
	return out.String()

}

// identifier is an expression node
// because in some parts of the program it might produce a Value
// e.g; let x = valueProducingIdentifier;
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode()       {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }
func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.Token.Literal + " ")

	if r.Value != nil {
		out.WriteString(r.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// expressions that may or may not produce value
// e.g; x+10 is an expression statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

//	====Step 0 done: Next Step -> constructing an AST=========
//...
//...
