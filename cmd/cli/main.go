package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joseph-beck/pear/pkg/lexer"
)

func main() {
	bytes, err := os.ReadFile("examples/fn_expr.pr")
	if err != nil {
		panic(err)
	}

	source := string(bytes)
	fmt.Printf("pear: {\n%s\n}\n", strings.TrimSuffix(source, "\n"))

	tokens := lexer.Lex(source)
	for _, t := range tokens {
		t.Debug()
	}
}
