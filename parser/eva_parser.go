package parser

import (
	"github.com/hculpan/goeva/lexer"
)

type EvaParser struct {
	Tokens    lexer.Tokens
	CurrToken int
}

func (p *EvaParser) Parse(tokens lexer.Tokens) *Node {
	p.CurrToken = 0
	p.Tokens = tokens

	return p.Expression()
}

func (p *EvaParser) Expression() *Node {
	var result *Node = nil

	if p.MoreTokens() {
		if p.CheckToken(lexer.PLUS, lexer.MINUS, lexer.STAR, lexer.SLASH) {
			result = p.BinaryOpExpression()
		} else if p.CheckToken(lexer.LEFTPAREN) {
			p.NextToken()
			result = p.Expression()
			if p.CheckToken(lexer.RIGHTPAREN) {
				p.NextToken()
			}
		} else {
			result = &Node{T: p.NextToken()}
		}
	}

	return result
}

func (p *EvaParser) BinaryOpExpression() *Node {
	result := &Node{T: p.NextToken()}

	if p.MoreTokens() {
		result.Left = p.Expression()
		result.Right = p.Expression()
	}

	return result
}

func (p *EvaParser) MoreTokens() bool {
	return p.CurrToken < len(p.Tokens)
}

func (p *EvaParser) NextToken() *lexer.Token {
	result := p.CurrentToken()
	p.CurrToken++
	return result
}

func (p *EvaParser) CheckToken(tokenTypes ...lexer.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.CurrentToken().Type == tokenType {
			return true
		}
	}

	return false
}

func (p *EvaParser) PeekToken(t lexer.TokenType) bool {
	return p.CurrentToken().Type == t
}

func (p *EvaParser) StepBackToken() {
	if p.CurrToken > 0 {
		p.CurrToken--
	}
}

func (p *EvaParser) CurrentToken() *lexer.Token {
	return &p.Tokens[p.CurrToken]
}
