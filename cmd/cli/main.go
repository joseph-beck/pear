package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joseph-beck/pear/pkg/lexer"
	"github.com/joseph-beck/pear/pkg/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, err := os.ReadFile("examples/struct.pr")
	if err != nil {
		panic(err)
	}

	source := string(bytes)
	fmt.Printf("pear: {\n%s\n}\n", strings.TrimSuffix(source, "\n"))

	tokens := lexer.Lex(source)
	for _, t := range tokens {
		t.Debug()
	}

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
