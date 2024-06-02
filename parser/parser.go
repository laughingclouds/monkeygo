package parser

import (
	"github.com/laughingclouds/monkeygo/ast"
	"github.com/laughingclouds/monkeygo/lexer"
	"github.com/laughingclouds/monkeygo/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken() // set peekToken
	p.nextToken() // set curToken, peekToken has new value

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
