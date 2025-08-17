package lexer

import (
	"fmt"
	"slices"
)

type TokenKind int

const (
	EndOfFile TokenKind = iota

	Null
	True
	False
	Number
	String
	Identifier

	OpenBracket
	CloseBracket
	OpenCurly
	CloseCurly
	OpenParen
	CloseParen

	Assignment
	Equals
	Not
	NotEquals
	LessThan
	LessThanEqual
	GreaterThan
	GreaterThanEqual
	Or
	And

	Dot
	Comma
	Range
	Spread
	SemiColon
	Colon
	Question
	Exclamation

	PlusPlus
	PlusEquals
	MinusMinus
	MinusEquals

	Plus
	Minus
	Divide
	Multiply
	Modulus

	Let
	Const
	Class
	New
	Import
	From
	Fn
	If
	ElseIf
	Else
	For
	ForEach
	While
	Export
	TypeOf
	In
	Struct
	Static
)

var Keywords map[string]TokenKind = map[string]TokenKind{
	"let":     Let,
	"const":   Const,
	"class":   Class,
	"new":     New,
	"import":  Import,
	"from":    From,
	"fn":      Fn,
	"if":      If,
	"elseif":  ElseIf,
	"else":    Else,
	"for":     For,
	"foreach": ForEach,
	"while":   While,
	"export":  Export,
	"typeof":  TypeOf,
	"in":      In,
	"struct":  Struct,
	"static":  Static,
}

func (t TokenKind) String() string {
	switch t {
	case EndOfFile:
		return "end_of_file"
	case Null:
		return "null"
	case True:
		return "true"
	case False:
		return "false"
	case Number:
		return "number"
	case String:
		return "string"
	case Identifier:
		return "identifier"
	case OpenBracket:
		return "open_bracket"
	case CloseBracket:
		return "close_bracket"
	case OpenCurly:
		return "open_curly"
	case CloseCurly:
		return "close_curly"
	case OpenParen:
		return "open_paren"
	case CloseParen:
		return "close_paren"
	case Assignment:
		return "assignment"
	case Equals:
		return "equals"
	case NotEquals:
		return "not_equals"
	case Not:
		return "not"
	case LessThan:
		return "less_than"
	case LessThanEqual:
		return "less_than_equals"
	case GreaterThan:
		return "greater_than"
	case GreaterThanEqual:
		return "greater_than_equals"
	case Or:
		return "or"
	case And:
		return "and"
	case Dot:
		return "dot"
	case Range:
		return "range"
	case Spread:
		return "spread"
	case SemiColon:
		return "semi_colon"
	case Colon:
		return "colon"
	case Question:
		return "question"
	case Comma:
		return "comma"
	case PlusPlus:
		return "plus_plus"
	case MinusMinus:
		return "minus_minus"
	case PlusEquals:
		return "plus_equals"
	case MinusEquals:
		return "minus_equals"
	case Plus:
		return "plus"
	case Minus:
		return "minus"
	case Divide:
		return "divide"
	case Multiply:
		return "multiply"
	case Modulus:
		return "modulus"
	case Let:
		return "let"
	case Const:
		return "const"
	case Class:
		return "class"
	case New:
		return "new"
	case Import:
		return "import"
	case From:
		return "from"
	case Fn:
		return "fn"
	case If:
		return "if"
	case Else:
		return "else"
	case For:
		return "for"
	case ForEach:
		return "foreach"
	case While:
		return "while"
	case Export:
		return "export"
	case In:
		return "in"
	case Struct:
		return "struct"
	case Static:
		return "static"
	default:
		return fmt.Sprintf("unknown(%d)", t)
	}
}

type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(k TokenKind, v string) Token {
	return Token{
		Kind:  k,
		Value: v,
	}
}

func (t Token) OfKind(k ...TokenKind) bool {
	return slices.Contains(k, t.Kind)
}

func (t Token) Debug() {
	if t.OfKind(Identifier, String, Number) {
		fmt.Printf("%s (%s)\n", t.Kind, t.Value)
	} else {
		fmt.Printf("%s ()\n", t.Kind)
	}
}
