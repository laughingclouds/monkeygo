package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// special identifiers
	ILLEGAL = "ILLEGAL" // A token/character we don't know
	EOF     = "EOF"     // end of file - parser to stop parsing from here

	// Idenifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1, 123, 345

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUCTION = "FUNCTION"
	LET     = "LET"
)
