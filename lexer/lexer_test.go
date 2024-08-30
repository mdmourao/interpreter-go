package lexer

import (
	"testing"

	"interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	expectedResults := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.EOF, Literal: ""},
	}

	l := New(input)

	for _, expectedResult := range expectedResults {
		tokenResult := l.NextToken()

		if expectedResult.Type != tokenResult.Type {
			t.Fatalf("lexer test - tokentype wrong. expected:%q, got %q", expectedResult.Type, tokenResult.Type)
		}

		if expectedResult.Literal != tokenResult.Literal {
			t.Fatalf("lexer test - literal wrong. expected:%q, got %q", expectedResult.Literal, tokenResult.Literal)
		}
	}
}
