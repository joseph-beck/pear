package ast

type Statement interface {
	statement()
}

type Expression interface {
	expression()
}

type Type interface {
	// type is a reserved key word go so use this instead...
	archetype()
}
