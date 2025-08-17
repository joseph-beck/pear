package ast

import "github.com/joseph-beck/pear/pkg/lexer"

type NumberExpression struct {
	Value float64
}

func (e NumberExpression) expression() {}

type StringExpression struct {
	Value string
}

func (e StringExpression) expression() {}

type SymbolExpression struct {
	Value string
}

func (e SymbolExpression) expression() {}

type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator lexer.Token
}

func (e BinaryExpression) expression() {}

type PrefixExpression struct {
	Operator lexer.Token
	Right    Expression
}

func (e PrefixExpression) expression() {}

type AssignmentExpression struct {
	Assignee Expression
	Operator lexer.Token
	Right    Expression
}

func (e AssignmentExpression) expression() {}

type StructInstantationExpression struct {
	Name   string
	Fields map[string]Expression
}

func (e StructInstantationExpression) expression() {}

type ArrayInstantationExpression struct {
	Underlying Type
	Contents   []Expression
}

func (e ArrayInstantationExpression) expression() {}
