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
	bindingPowerLookup[k] = defaultBindingPower
	statementLookup[k] = fn
}

type nullDenotationHandler func(p *parser) ast.Expression
type NullDenotationLookup map[lexer.TokenKind]nullDenotationHandler

var nullDenotationLookup = NullDenotationLookup{}

func nullDenotation(k lexer.TokenKind, fn nullDenotationHandler) {
	nullDenotationLookup[k] = fn
}

type leftDenotationHandler func(p *parser, l ast.Expression, bp bindingPower) ast.Expression
type LeftDenotationLookup map[lexer.TokenKind]leftDenotationHandler

var leftDenotationLookup = LeftDenotationLookup{}

func leftDenotation(k lexer.TokenKind, bp bindingPower, fn leftDenotationHandler) {
	bindingPowerLookup[k] = bp
	leftDenotationLookup[k] = fn
}

type BindingPowerLookup map[lexer.TokenKind]bindingPower

var bindingPowerLookup = BindingPowerLookup{}

// createLookups must be called before anything in the parser package is ran
func createLookups() {
	leftDenotation(lexer.Assignment, assignment, parseAssignmentExpression)
	leftDenotation(lexer.PlusEquals, assignment, parseAssignmentExpression)
	leftDenotation(lexer.MinusEquals, assignment, parseAssignmentExpression)

	leftDenotation(lexer.And, logical, parseBinaryExpression)
	leftDenotation(lexer.Or, logical, parseBinaryExpression)
	leftDenotation(lexer.Range, logical, parseBinaryExpression)

	leftDenotation(lexer.LessThan, relational, parseBinaryExpression)
	leftDenotation(lexer.LessThanEqual, relational, parseBinaryExpression)
	leftDenotation(lexer.GreaterThan, relational, parseBinaryExpression)
	leftDenotation(lexer.GreaterThanEqual, relational, parseBinaryExpression)
	leftDenotation(lexer.Equals, relational, parseBinaryExpression)
	leftDenotation(lexer.NotEquals, relational, parseBinaryExpression)

	leftDenotation(lexer.Plus, additive, parseBinaryExpression)
	leftDenotation(lexer.Minus, additive, parseBinaryExpression)

	leftDenotation(lexer.Multiply, multiplicative, parseBinaryExpression)
	leftDenotation(lexer.Divide, multiplicative, parseBinaryExpression)
	leftDenotation(lexer.Modulus, multiplicative, parseBinaryExpression)

	nullDenotation(lexer.Number, parsePrimaryExpression)
	nullDenotation(lexer.String, parsePrimaryExpression)
	nullDenotation(lexer.Identifier, parsePrimaryExpression)

	nullDenotation(lexer.OpenParen, parseGroupingExpression)

	nullDenotation(lexer.Minus, parsePrefixExpression)

	leftDenotation(lexer.OpenCurly, call, parseStructInstantiationExpression)
	nullDenotation(lexer.OpenBracket, parseArrayInstantiationExpression)

	statement(lexer.Const, parseVariableDeclarationStatement)
	statement(lexer.Let, parseVariableDeclarationStatement)
	statement(lexer.Struct, parseStructDeclorationStatement)
}
