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
	Public        bool
}

func (s VariableDeclorationStatement) statement() {}

type StructField struct {
	Type   Type
	Public bool
	Static bool
}

type StructMethod struct {
	Type   Type
	Public bool
	Static bool
}

type StructStatement struct {
	Name    string
	Public  bool
	Fields  map[string]StructField
	Methods map[string]StructMethod
}

func (s StructStatement) statement() {}
