package token

type TokenType string // Easy to debug, using int or byte would increase performance

type Token struct {
	Type    TokenType
	Literal string
}

// Token types
const (
	ILLEGEL = "ILLEGAL" // token/char we dont know
	EOF     = "EOF"     // inform parser to stop

	IDENT = "IDENT" // variables names
	INT   = "INT"   // integer values (12345)

	ASSIGN = "=" // Operator
	PLUS   = "+" // Operator

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
