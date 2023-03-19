package parser

import "github.com/hculpan/goeva/lexer"

type Node struct {
	Left  *Node
	Right *Node

	T *lexer.Token
}

type Parser interface {
	Parse(tokens lexer.Tokens) *Node
}
