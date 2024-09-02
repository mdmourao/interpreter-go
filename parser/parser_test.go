package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

// what if no Semicolon is added to the end of the program: infinite loop?
func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`

	myLexer := lexer.New(input)
	myParser := New(myLexer)

	program := myParser.ParseProgram()
	testParserErrors(t, myParser)
	if program == nil {
		t.Fatalf("ParseProgram() retuned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program Statementes does not contain 3 statements. current: %d", len(program.Statements))
	}

	expectedIdentifiers := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, expectedIdentifier := range expectedIdentifiers {
		stmt := program.Statements[i]

		if !testLetState(t, stmt, expectedIdentifier.expectedIdentifier) {
			return
		}
	}
}

func TestErrorLetStatements(t *testing.T) {
	input := `
	let x 5;
	let = 10;
	let 838383;
`

	myLexer := lexer.New(input)
	myParser := New(myLexer)

	myParser.ParseProgram()

	if len(myParser.errors) != 3 {
		t.Fatalf("should have 3 error but got %d", len(myParser.errors))
	}
}

func testLetState(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("TokenLiteral is not 'let' got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("statement if not a LetStatement, got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("let statement name value is not: '%s' got: '%s' ", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s' got: '%s'", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func testParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}

	for _, error := range p.errors {
		t.Errorf("parser error detected: %s", error)
	}

	t.FailNow()
}
