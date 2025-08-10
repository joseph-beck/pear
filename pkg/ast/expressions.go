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
