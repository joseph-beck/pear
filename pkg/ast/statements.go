package ast

type BlockStatement struct {
	Body []Statement
}

func (s BlockStatement) statement() {}

type ExpressionStatement struct {
	Expression Expression
}

func (s ExpressionStatement) statement() {}
