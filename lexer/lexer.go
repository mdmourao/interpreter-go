package lexer

import (
	"interpreter/token"
	"interpreter/utils"
)

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

// returns a full identifier (eg: foo)
func (l *Lexer) readIdentifier() string {
	initialPosition := l.position
	for utils.IsLetter(l.char) {
		l.readChar()
	}
	return l.input[initialPosition:l.position]
}

// returns a full number (eg: 12121)
func (l *Lexer) readNumber() string {
	position := l.position
	for utils.IsDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skip() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skip()

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
	default:
		if utils.IsLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // need eary return because we use readIdentifier to last char
		} else if utils.IsDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok // need eary return because we use readIdentifier to last char
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
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
