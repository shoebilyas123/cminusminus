package parser

import (
	"fmt"

	"github.com/shoebilyas123/monkeylang/monkey/ast"
	"github.com/shoebilyas123/monkeylang/monkey/lexer"
	"github.com/shoebilyas123/monkeylang/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s. Got %s instead", p.peekToken.Type, t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// Construct the root node of our AST named program
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program

}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	_st := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// We are skipping expressions until we encounter a semicolon;
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return _st
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	_statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	_statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We will evaluate expressions later
	// We're skipping expressions until we encounter a semicolon

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return _statement

}

func (p *Parser) curTokenIs(expectedToken token.TokenType) bool {
	return p.curToken.Type == expectedToken
}

func (p *Parser) peekTokenIs(expectedToken token.TokenType) bool {
	// fmt.Printf("%+v::%+v\n", expectedToken, p.peekToken)
	return p.peekToken.Type == expectedToken
}

func (p *Parser) expectPeek(_expectedToken token.TokenType) bool {
	if p.peekTokenIs(_expectedToken) {
		p.nextToken()
		return true
	} else {
		p.peekError(_expectedToken)
		return false
	}
}
