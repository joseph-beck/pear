package parser

import (
	"github.com/joseph-beck/pear/pkg/ast"
	"github.com/joseph-beck/pear/pkg/lexer"
)

type parser struct {
	tokens   []lexer.Token
	position int
}

func new(t []lexer.Token) *parser {
	return &parser{
		tokens:   t,
		position: 0,
	}
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

func (p *parser) token() lexer.Token {
	return p.tokens[p.position]
}

func (p *parser) kind() lexer.TokenKind {
	return p.token().Kind
}

func (p *parser) advance() lexer.Token {
	t := p.token()
	p.position++
	return t
}

func (p *parser) eof() bool {
	return p.position >= len(p.tokens) && p.kind() == lexer.EndOfFile
}
