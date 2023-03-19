package main

import (
	"fmt"

	"github.com/hculpan/goeva/eval"
)

func main() {
	str := `(* (/ 12 2) (+ (- 7 2) 5))`
	v, err := eval.Eval(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v.String())
	}
}
