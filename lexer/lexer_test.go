package lexer

import (
	"github.com/panzerit/lumi/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextTokenForOneCharTokens(t *testing.T) {
	input := "=+-*/(),!"

	tests := []struct {
		wantType    token.Type
		wantLiteral string
	}{
		{token.ASSIGN, "="},
		{token.ADD, "+"},
		{token.SUB, "-"},
		{token.MUL, "*"},
		{token.DIV, "/"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.BANG, "!"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.wantType, tok.Type, "position %d", i)
		assert.Equal(t, tt.wantLiteral, tok.Literal, "position %d", i)
	}
}

func TestNextTokenForCode(t *testing.T) {
	input := `let five = 5
              let ten = 10
              let add = fn x y do
                 x + y
              end
              
              let result = add five ten
			  
              5 < 10
              10 > 5

              10 == 10
              10 != 5
              
              if 5 < 10 do
                  return true
              else
                  return false
              end`

	tests := []struct {
		wantType    token.Type
		wantLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.IDENT, "x"},
		{token.IDENT, "y"},
		{token.DO, "do"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.END, "end"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.IDENT, "five"},
		{token.IDENT, "ten"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NE, "!="},
		{token.INT, "5"},
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.DO, "do"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.ELSE, "else"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.END, "end"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.wantType, tok.Type, "position %d", i)
		assert.Equal(t, tt.wantLiteral, tok.Literal, "position %d", i)
	}
}
