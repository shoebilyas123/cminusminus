package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENTIFIERS AND LITERALS
	IDENT = "IDENT"
	INT   = "INT"

	// OPERATORS
	ASSIGN = "="
	ADD    = "+"

	// DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// KEYWORDS
	FN  = "FN"
	LET = "LET"
)
