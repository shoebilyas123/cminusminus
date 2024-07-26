package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT" // Identifiers
	INT       = "INT"
	ADD       = "=" // Operators
	PLUS      = "+"
	COMMA     = "," // Delimiters
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "{"
	FUNCTION  = "FUNCTION" // Keywords
	LET       = "LET"
)
