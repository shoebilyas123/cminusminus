package lexer

import (
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

func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespace()
	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
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
	case '!':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.EXCLAIM, l.ch)
		}
		break
	case '<':
		tok = newToken(token.SMALLER, l.ch)
		break
	case '>':
		tok = newToken(token.GREATER, l.ch)
		break
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
		break
	case '/':
		tok = newToken(token.FOR_SLASH, l.ch)
		break
	case '-':
		tok = newToken(token.MINUS, l.ch)
		break
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		break
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)

			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

func (l *Lexer) readNumber() string {
	pos := l.currPosition

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.currPosition]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func (l *Lexer) readIdentifier() string {
	pos := l.currPosition

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.currPosition]
}

func (l *Lexer) peakChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}

	return l.input[l.nextPosition]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
