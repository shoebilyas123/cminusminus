package lexer

import (
	"fmt"

	"github.com/shoebilyas123/monkeylang/monkey/token"
)

type Lexer struct {
	input        string
	nextPosition int
	currPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input, currPosition: -1, nextPosition: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition++

}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
		break
	case '+':
		tok = newToken(token.PLUS, l.ch)
		break
	case '(':
		tok = newToken(token.LPAREN, l.ch)
		break
	case ')':
		tok = newToken(token.RPAREN, l.ch)
		break
	case '{':
		tok = newToken(token.LBRACE, l.ch)
		break
	case '}':
		tok = newToken(token.RBRACE, l.ch)
		break
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		break
	case ',':
		tok = newToken(token.COMMA, l.ch)
		break
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		break
	}
	fmt.Printf("CHAR:%s\n", string(l.ch))
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
