package parser

import (
	"github.com/joseph-beck/pear/pkg/ast"
	"github.com/joseph-beck/pear/pkg/lexer"
)

func parseStatement(p *parser) ast.Statement {
	fn, exists := statementLookup[p.kind()]

	if exists {
		return fn(p)
	}

	e := parseExpression(p, defaultBindingPower)
	p.expect(lexer.SemiColon)

	return ast.ExpressionStatement{
		Expression: e,
	}
}
