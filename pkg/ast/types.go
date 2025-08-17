package ast

type SymbolType struct {
	Name string
}

func (t SymbolType) archetype() {}

type ArrayType struct {
	UnderlyingType Type
}

func (t ArrayType) archetype() {}
