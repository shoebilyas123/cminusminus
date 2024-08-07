package ast

import "github.com/shoebilyas123/monkeylang/monkey/token"

type Node interface {
	TokenLiteral() string
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

// identifier is an expression node
// because in some parts of the program it might produce a Value
// e.g; let x = valueProducingIdentifier;
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode()       {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

//	====Step 0 done: Next Step -> constructing an AST=========
//...
//...
