package parser

import (
	"github.com/panzerit/lumi/ast"
	"github.com/panzerit/lumi/lexer"
	"github.com/panzerit/lumi/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// parse two tokens so that curToken and peekToken are initialized
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
