package parser

import (
	"fmt"

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

func parseVariableDeclarationStatement(p *parser) ast.Statement {
	var t ast.Type
	var v ast.Expression

	cnst := p.advance().Kind == lexer.Const
	nme := p.expect(lexer.Identifier, "Unable to find variable name in decloration").Value

	// when we have a let thing: type = 123.
	if p.kind() == lexer.Colon {
		p.advance()
		t = parseType(p, defaultBindingPower)
	}

	// if we wanted to have a let thing: type; without assignment.
	if p.kind() != lexer.SemiColon {
		p.expect(lexer.Assignment, "Unable to find assignment")
		v = parseExpression(p, assignment)
	} else if t == nil {
		panic(fmt.Sprintf("Variable {%s} missing type decloration or assignment", nme))
	}

	p.expect(lexer.SemiColon, "Unable to find semi-colon at end of variable decloration")

	if cnst && v == nil {
		panic(fmt.Sprintf("Const {%s} must be assigned to a value", nme))
	}

	return ast.VariableDeclorationStatement{
		VariableName:  nme,
		Constant:      cnst,
		VariableValue: v,
		Type:          t,
	}
}
