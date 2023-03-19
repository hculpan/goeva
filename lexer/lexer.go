package lexer

import (
	"fmt"
	"unicode"
)

func Tokenize(expr string) (Tokens, error) {
	result := Tokens{}
	currPos := 0
	currLine := 1
	currCol := 1
	buff := ""
	var token *Token = nil

	for currPos < len(expr) {
		if expr[currPos] == '\n' {
			currLine++
			currPos++
			currCol = 1
			continue
		}
		if unicode.IsSpace(rune(expr[currPos])) {
			currPos++
			currCol++
			continue
		}

		buff = string(expr[currPos:])
		token = matches(buff, currCol, currLine)
		if token != nil {
			result = append(result, *token)
			currPos += len(token.Literal)
			currCol += len(token.Literal)
		} else {
			return nil, fmt.Errorf("unmatched token '%s'", buff)
		}
	}

	return result, nil
}

func matches(buff string, currCol int, currLine int) *Token {
	var result *Token = nil
	for _, m := range TokenMatchers {
		matched := m.Matcher.Find([]byte(buff))
		if matched != nil {
			result = &Token{
				Type:    m.Identifier,
				Literal: string(matched),
				Col:     currCol,
				Line:    currLine,
			}
			break
		}
	}
	return result
}
