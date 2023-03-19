package executor

import (
	"strconv"

	"github.com/hculpan/goeva/lexer"
	"github.com/hculpan/goeva/parser"
	"github.com/hculpan/goeva/util"
)

func Execute(node *parser.Node) *util.Value {
	if node == nil {
		return &util.Value{}
	}

	if IsOperator(node.T.Type) {
		lv := Execute(node.Left)
		rv := Execute(node.Right)

		switch node.T.Type {
		case lexer.PLUS:
			lv.Add(rv)
		case lexer.MINUS:
			lv.Sub(rv)
		case lexer.STAR:
			lv.Mult(rv)
		case lexer.SLASH:
			lv.Div(rv)
		}

		return lv
	} else {
		switch node.T.Type {
		case lexer.INTEGER:
			n, _ := strconv.ParseInt(node.T.Literal, 10, 64)
			return util.NewIntegerValue(n)
		case lexer.FLOAT:
			n, _ := strconv.ParseFloat(node.T.Literal, 64)
			return util.NewFloatValue(n)
		case lexer.STRING:
			return util.NewStringValue(node.T.Literal)
		}
	}

	return &util.Value{}
}

func IsOperator(tokenType lexer.TokenType) bool {
	return tokenType == lexer.PLUS || tokenType == lexer.MINUS || tokenType == lexer.STAR || tokenType == lexer.SLASH
}
