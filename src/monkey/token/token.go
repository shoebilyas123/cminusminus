package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // unsupported token
	EOF     = "EOF"     // end of file

	// IDENTIFIERS AND LITERALS
	IDENT = "IDENT"
	INT   = "INT"

	// OPERATORS
	ASSIGN = "="
	PLUS   = "+"

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
