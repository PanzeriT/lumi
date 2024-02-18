package parser

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/panzerit/lumi/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `a = 1
              ab = 12
              abc = 123`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"ab"},
		{"abc"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		assert.Equal(t, "let", stmt.TokenLiteral())
		log.Println(tt)
	}
}
