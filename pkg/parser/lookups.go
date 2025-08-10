package parser

import (
	"github.com/joseph-beck/pear/pkg/ast"
	"github.com/joseph-beck/pear/pkg/lexer"
)

type bindingPower int

const (
	defaultBindingPower bindingPower = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type statementHandler func(p *parser) ast.Statement
type StatementLookup map[lexer.TokenKind]statementHandler

var statementLookup = StatementLookup{}

func statement(k lexer.TokenKind, fn statementHandler) {
	statementLookup[k] = fn
	bindingPowerLookup[k] = defaultBindingPower
}

type nullDenotationHandler func(p *parser) ast.Expression
type NullDenotationLookup map[lexer.TokenKind]nullDenotationHandler

var nullDenotationLookup = NullDenotationLookup{}

func nullDenotation(k lexer.TokenKind, bp bindingPower, fn nullDenotationHandler) {
	nullDenotationLookup[k] = fn
	bindingPowerLookup[k] = bp
}

type leftDenotationHandler func(p *parser, l ast.Expression, bp bindingPower) ast.Expression
type LeftDenotationLookup map[lexer.TokenKind]leftDenotationHandler

var leftDenotationLookup = LeftDenotationLookup{}

func leftDenotation(k lexer.TokenKind, bp bindingPower, fn leftDenotationHandler) {
	leftDenotationLookup[k] = fn
	bindingPowerLookup[k] = bp
}

type BindingPowerLookup map[lexer.TokenKind]bindingPower

var bindingPowerLookup = BindingPowerLookup{}
