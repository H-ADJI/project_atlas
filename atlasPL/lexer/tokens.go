package lexer

import "slices"

const (
	// literals
	INT = iota
	// identifiers
	IDENT
	// keywords
	FUNCTION
	LET

	// operations
	PLUS
	MINUS
	ASSIGN
	DIV
	MULT

	// DELIMITER
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	// SPECIAL
	EOF
	ILLEGAL
)

type TokenType int8

type Token struct {
	Literal string
	Type    TokenType
}

var INDENTIFIERS = []string{"let", "fn"}

func isKeyword(literal string) bool {
	return slices.Contains(INDENTIFIERS, literal)
}
