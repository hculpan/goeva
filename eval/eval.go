package eval

import (
	"github.com/hculpan/goeva/executor"
	"github.com/hculpan/goeva/lexer"
	"github.com/hculpan/goeva/parser"
	"github.com/hculpan/goeva/util"
)

func Eval(expr string) (*util.Value, error) {
	tokens, err := lexer.Tokenize(expr)
	if err != nil {
		return nil, err
	}

	/*	if len(tokens) == 1 {
			if tokens[0].Type == lexer.INTEGER {
				n, _ := strconv.ParseInt(tokens[0].Literal, 10, 64)
				return util.NewIntegerValue(n), nil
			} else if tokens[0].Type == lexer.FLOAT {
				n, _ := strconv.ParseFloat(tokens[0].Literal, 64)
				return util.NewFloatValue(n), nil
			} else if tokens[0].Type == lexer.STRING {
				return util.NewStringValue(tokens[0].Literal), nil
			}
		} else if tokens[0].Type == lexer.PLUS {
			v1 := util.NewValueFromToken(tokens[1])
			v2 := util.NewValueFromToken(tokens[2])

			v1.Add(v2)

			return v1, nil
		} else if tokens[0].Type == lexer.MINUS {
			v1 := util.NewValueFromToken(tokens[1])
			v2 := util.NewValueFromToken(tokens[2])

			v1.Sub(v2)

			return v1, nil
		} else if tokens[0].Type == lexer.STAR {
			v1 := util.NewValueFromToken(tokens[1])
			v2 := util.NewValueFromToken(tokens[2])

			v1.Mult(v2)

			return v1, nil
		} else if tokens[0].Type == lexer.SLASH {
			v1 := util.NewValueFromToken(tokens[1])
			v2 := util.NewValueFromToken(tokens[2])

			v1.Div(v2)

			return v1, nil
		}
	*/

	parser := parser.EvaParser{}
	topNode := parser.Parse(tokens)
	v := executor.Execute(topNode)

	return v, nil
}
