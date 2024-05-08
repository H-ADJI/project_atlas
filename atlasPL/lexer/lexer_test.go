package lexer

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	sourceCode := "let plus_mult = fn(x){ return x + x , x * x if x >= 0 else - (x + x), x * x  };"
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LET, "let"},
		{IDENT, "plus_mult"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "x"},
		{ASTERISK, "*"},
		{IDENT, "x"},
		{IF, "if"},
		{IDENT, "x"},
		{GTE, ">="},
		{INT, "0"},
		{ELSE, "else"},
		{MINUS, "-"},
		{LPAREN, "("},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "x"},
		{RPAREN, ")"},
		{COMMA, ","},
		{IDENT, "x"},
		{ASTERISK, "*"},
		{IDENT, "x"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	lexer := NewLexer(readSource(sourceCode))
	for i, tc := range tests {
		token := lexer.NextToken()
		fmt.Printf("type : %d ----- value : %s \n", token.Type, token.Literal)
		if tc.expectedType != token.Type {
			t.Errorf("Test failed [%d], should be : %v, but got %v ", i, tc, token)
		}
	}

	sourceCode = "@ ;"
	tests = []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{ILLEGAL, "@"},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	lexer = NewLexer(readSource(sourceCode))
	for i, tc := range tests {
		token := lexer.NextToken()
		fmt.Printf("type : %d ----- value : %s \n", token.Type, token.Literal)
		if tc.expectedType != token.Type {
			t.Errorf("Test failed [%d], should be : %v, but got %v ", i, tc, token)
		}
	}
}
