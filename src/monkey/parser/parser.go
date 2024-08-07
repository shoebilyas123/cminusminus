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

// CURRENTLY: We are creating parse programs for parsing let statements
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// TODO: we are skipping expressions for now.
	// We will get back to it after we have expression parsing working
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}

	return false
}

func (p *Parser) curTokenIs(tt token.TokenType) bool {
	if p.curToken.Type == tt {
		return true
	}

	return false
}
