package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatemts(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let z = 20;
  let foobar = 838383;
  `
	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statments) != 3 {
		t.Fatalf("program.Statments does not contai 3 statments. got=%d", len(program.Statments))
	}

	tests := []struct {
		expectedIdentfier string
	}{
		{"x"},
		{"y"},
		{"z"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statments[i]

		if !testLetStatemt(t, stmt, tt.expectedIdentfier) {
			return
		}
	}
}

func testLetStatemt(t *testing.T, s ast.Statment, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatment)
	if !ok {
		t.Errorf("s not *ast.LetStatment. Got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s. Got=%s", name, letStmt.Name.Value)
		return false
	}

	if (letStmt.Name.TokenLiteral()) != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s. Got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}
