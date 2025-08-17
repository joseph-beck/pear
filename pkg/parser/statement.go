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
	var pub bool

	if p.kind() == lexer.Pub {
		pub = true
		p.expect(lexer.Pub, "Unable to find public keyword in variable decloration")
	}

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
		Public:        pub,
	}
}

func parseStructDeclorationStatement(p *parser) ast.Statement {
	var flds = map[string]ast.StructField{}
	var mthds = map[string]ast.StructMethod{}
	var pub bool

	if p.kind() == lexer.Pub {
		pub = true
		p.expect(lexer.Pub, "Unable to find public keyword in struct decloration")
	}

	p.expect(lexer.Struct)

	n := p.expect(lexer.Identifier).Value

	p.expect(lexer.OpenCurly, "Unable to find opening curly for struct")

	for !p.eof() && p.kind() != lexer.CloseCurly {
		var static bool
		var public bool
		var name string

		if p.kind() == lexer.Pub {
			public = true
			p.expect(lexer.Pub, "Unable to find public keyword in struct field or method")
		}

		if p.kind() == lexer.Static {
			static = true
			p.expect(lexer.Static, "Unable to find static keyword in struct field or method")
		}

		if p.kind() == lexer.Identifier {
			name = p.expect(lexer.Identifier, "Unable to find field in struct").Value

			p.expect(lexer.Colon, "Unable to find colon after field in struct")

			tp := parseType(p, defaultBindingPower)

			p.expect(lexer.SemiColon, "Unable to find semi colon after type declaration")

			_, exists := flds[name]
			if exists {
				panic(fmt.Sprintf("Duplicate field name {%s} in struct {%s}", name, n))
			}

			flds[name] = ast.StructField{
				Type:   tp,
				Static: static,
				Public: public,
			}

			continue
		}

		panic("unable to parse struct methods for now...")
	}

	p.expect(lexer.CloseCurly, "Unable to find closing curly for struct")

	return ast.StructStatement{
		Name:    n,
		Public:  pub,
		Fields:  flds,
		Methods: mthds,
	}
}

func parsePubStatement(p *parser) ast.Statement {
	switch p.peak().Kind {
	case lexer.Struct:
		return parseStructDeclorationStatement(p)
	case lexer.Let, lexer.Const:
		return parseVariableDeclarationStatement(p)
	}

	panic("Unable to parse public statement")
}
