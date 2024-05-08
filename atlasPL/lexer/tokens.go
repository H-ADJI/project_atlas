package lexer

const (
	// literals
	INT = iota
	// identifiers
	IDENT

	// keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	// operators
	PLUS
	MINUS
	ASSIGN
	SLASH
	ASTERISK
	BANG
	EQ
	NOT_EQ
	GT
	GTE
	LT
	LTE
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

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
