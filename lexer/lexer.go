package lexer

import "interpreter/token"

// Only supports ASCII characters

type Lexer struct {
	input        string
	position     int  // current position (current char)
	readPosition int  // next position (after current char)
	char         byte // current char (needs to be a rune to support unicode)
}

// reads the next character and increases our position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // ASCII - "NUL" (represents: EOF)
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	var lexer = &Lexer{input: input}
	lexer.readChar()
	return lexer
}
