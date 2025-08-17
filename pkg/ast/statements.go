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

type StructField struct {
	Type   Type
	Static bool
}

type StructMethod struct {
	Type   Type
	Static bool
}

type StructStatement struct {
	Name    string
	Fields  map[string]StructField
	Methods map[string]StructMethod
}

func (s StructStatement) statement() {}
