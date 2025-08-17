package parser

import (
	"fmt"

	"github.com/joseph-beck/pear/pkg/ast"
	"github.com/joseph-beck/pear/pkg/lexer"
)

type parser struct {
	tokens   []lexer.Token
	position int
}

func new(t []lexer.Token) *parser {
	createLookups()
	createTypeLookups()

	return &parser{
		tokens:   t,
		position: 0,
	}
}

// Get the current token of the parser.
func (p *parser) token() lexer.Token {
	return p.tokens[p.position]
}

// Get the kind of the current token of the parser.
func (p *parser) kind() lexer.TokenKind {
	return p.token().Kind
}

// Peak at the next token without advancing the parser.
func (p *parser) peak() lexer.Token {
	return p.tokens[p.position+1]
}

// Advanced the parse, and get the token.
func (p *parser) advance() lexer.Token {
	t := p.token()
	p.position++
	return t
}

func (p *parser) expect(e lexer.TokenKind, err ...any) lexer.Token {
	t := p.token()
	k := t.Kind

	if k != e {
		if len(err) == 0 {
			err = append(err, fmt.Sprintf("expected %s but got %s instead\n", e, k))
		}

		panic(fmt.Sprintf("unable to parse the token and got {%s} {%s} %T", t, k, err))
	}

	return p.advance()
}

// Is the parser at the end of the file?
func (p *parser) eof() bool {
	return p.position >= len(p.tokens) || p.kind() == lexer.EndOfFile
}

func Parse(t []lexer.Token) ast.BlockStatement {
	body := make([]ast.Statement, 0)
	p := new(t)

	for !p.eof() {
		body = append(body, parseStatement(p))
	}

	return ast.BlockStatement{
		Body: body,
	}
}
