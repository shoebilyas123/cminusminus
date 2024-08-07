package parser

import (
	"github.com/shoebilyas123/monkeylang/monkey/ast"
	"github.com/shoebilyas123/monkeylang/monkey/lexer"
	"github.com/shoebilyas123/monkeylang/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	// curToken is the current token under examination
	// Based on the peek token we will identify whether there's more ops
	// e.g; 5;
	// In the above expression the curToken=5;
	// The peekToken will tell us whether we are at the end of a line
	// or at the start of an arithmetic expression
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
