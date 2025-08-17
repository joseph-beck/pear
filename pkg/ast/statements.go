package ast

type BlockStatement struct {
	Body []Statement
}

func (s BlockStatement) statement() {}

type ExpressionStatement struct {
	Expression Expression
}

func (s ExpressionStatement) statement() {}

type VariableDeclorationStatement struct {
	VariableName  string
	Constant      bool
	VariableValue Expression
	Type          Type
}

func (s VariableDeclorationStatement) statement() {}
