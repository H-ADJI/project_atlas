package lexer

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	sourceCode := "let plus_mult = fn(x){ return x + x , x * x };"
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{KEYWORD, "let"},
		{IDENT, "plus_mult"},
		{ASSIGN, "="},
		{KEYWORD, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{KEYWORD, "return"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "x"},
		{MULT, "*"},
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
