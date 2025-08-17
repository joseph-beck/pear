package parser

import (
	"fmt"
	"strconv"

	"github.com/joseph-beck/pear/pkg/ast"
	"github.com/joseph-beck/pear/pkg/lexer"
	"github.com/joseph-beck/pear/pkg/reflections"
)

func parseExpression(p *parser, bp bindingPower) ast.Expression {
	k := p.kind()
	nud, exists := nullDenotationLookup[k]

	if !exists {
		panic(fmt.Sprintf("No null denotation for %s", k))
	}

	l := nud(p)
	// this is for the next token, will not be the same as k
	for bindingPowerLookup[p.kind()] > bp {
		k = p.kind()
		led, exists := leftDenotationLookup[k]

		if !exists {
			panic(fmt.Sprintf("No left denotation for %s", k))
		}

		l = led(p, l, bindingPowerLookup[k])
	}

	return l
}

func parsePrimaryExpression(p *parser) ast.Expression {
	switch p.kind() {
	case lexer.Number:
		n, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpression{
			Value: n,
		}
	case lexer.String:
		return ast.StringExpression{
			Value: p.advance().Value,
		}
	case lexer.Identifier:
		return ast.SymbolExpression{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Unable to create primary expression from %s\n", p.kind()))
	}
}

func parseBinaryExpression(p *parser, l ast.Expression, bp bindingPower) ast.Expression {
	t := p.advance()
	r := parseExpression(p, bp)

	return ast.BinaryExpression{
		Left:     l,
		Right:    r,
		Operator: t,
	}
}

func parsePrefixExpression(p *parser) ast.Expression {
	t := p.advance()
	r := parseExpression(p, defaultBindingPower)

	return ast.PrefixExpression{
		Operator: t,
		Right:    r,
	}
}

func parseGroupingExpression(p *parser) ast.Expression {
	p.advance()
	e := parseExpression(p, defaultBindingPower)
	p.expect(lexer.CloseParen, "Expected expression to be closed by paren")

	return e
}

func parseAssignmentExpression(p *parser, l ast.Expression, bp bindingPower) ast.Expression {
	t := p.advance()
	r := parseExpression(p, defaultBindingPower)

	return ast.AssignmentExpression{
		Operator: t,
		Right:    r,
		Assignee: l,
	}
}

func parseStructInstantiationExpression(p *parser, l ast.Expression, bp bindingPower) ast.Expression {
	fmt.Print(1)
	n := reflections.ExpectType[ast.SymbolExpression](l).Value
	fmt.Print(2)
	var flds = map[string]ast.Expression{}

	p.expect(lexer.OpenCurly)

	for !p.eof() && p.kind() != lexer.CloseCurly {
		fmt.Print(3)
		fld := p.expect(lexer.Identifier, "Unable to find field name in struct instantiation").Value

		p.expect(lexer.Colon)

		e := parseExpression(p, logical)

		flds[fld] = e

		if p.kind() != lexer.CloseCurly {
			p.expect(lexer.Comma)
		}
	}

	p.expect(lexer.CloseCurly)

	return ast.StructInstantationExpression{
		Name:   n,
		Fields: flds,
	}
}

func parseArrayInstantiationExpression(p *parser) ast.Expression {
	var c = []ast.Expression{}

	p.expect(lexer.OpenBracket)
	p.expect(lexer.CloseBracket)

	ut := parseType(p, defaultBindingPower)

	p.expect(lexer.OpenCurly)

	for !p.eof() && p.kind() != lexer.CloseCurly {
		e := parseExpression(p, logical)
		c = append(c, e)

		if p.kind() != lexer.CloseCurly {
			p.expect(lexer.Comma)
		}
	}

	p.expect(lexer.CloseCurly)

	return ast.ArrayInstantationExpression{
		Underlying: ut,
		Contents:   c,
	}
}
