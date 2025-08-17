package parser

import (
	"fmt"

	"github.com/joseph-beck/pear/pkg/ast"
	"github.com/joseph-beck/pear/pkg/lexer"
)

type typeNullDenotationHandler func(p *parser) ast.Type
type TypeNullDenotationLookup map[lexer.TokenKind]typeNullDenotationHandler

var typeNullDenotationLookup = TypeNullDenotationLookup{}

func typeNullDenotation(k lexer.TokenKind, fn typeNullDenotationHandler) {
	typeNullDenotationLookup[k] = fn
}

type typeLeftDenotationHandler func(p *parser, l ast.Type, bp bindingPower) ast.Type
type TypeLeftDenotationLookup map[lexer.TokenKind]typeLeftDenotationHandler

var typeLeftDenotationLookup = TypeLeftDenotationLookup{}

func typeLeftDenotation(k lexer.TokenKind, bp bindingPower, fn typeLeftDenotationHandler) {
	typeBindingPowerLookup[k] = bp
	typeLeftDenotationLookup[k] = fn
}

type TypeBindingPowerLookup map[lexer.TokenKind]bindingPower

var typeBindingPowerLookup = TypeBindingPowerLookup{}

func parseType(p *parser, bp bindingPower) ast.Type {
	k := p.kind()
	nud, exists := typeNullDenotationLookup[k]

	if !exists {
		panic(fmt.Sprintf("No type null denotation for %s", k))
	}

	l := nud(p)
	// this is for the next token, will not be the same as k
	for typeBindingPowerLookup[p.kind()] > bp {
		k = p.kind()
		led, exists := typeLeftDenotationLookup[k]

		if !exists {
			panic(fmt.Sprintf("No type left denotation for %s", k))
		}

		l = led(p, l, typeBindingPowerLookup[k])
	}

	return l
}

func parseSymbolType(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.Identifier).Value,
	}
}

func parseArrayType(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.CloseBracket, "Unable to find closing bracket for array type")

	t := parseType(p, defaultBindingPower)

	return ast.ArrayType{
		UnderlyingType: t,
	}
}

func createTypeLookups() {
	typeNullDenotation(lexer.Identifier, parseSymbolType)
	typeNullDenotation(lexer.OpenBracket, parseArrayType)
}
