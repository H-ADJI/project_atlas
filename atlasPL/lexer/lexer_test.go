package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	sourceCode := "=*+{},;"
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{ASSIGN, "="},
		{MULT, "*"},
		{PLUS, "+"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	lexer := NewLexer(readSource(sourceCode))
	for i, tc := range tests {
		token := lexer.NextToken()
		if tc.expectedType != token.Type {
			t.Errorf("Test failed [%d], should be : %v, but got %v ", i, tc, token)
		}
	}
}
