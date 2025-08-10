package lexer

import (
	"fmt"
	"regexp"
)

type handler func(l *lexer, r *regexp.Regexp)

type pattern struct {
	regex *regexp.Regexp
	handler
}

func defaultHandler(k TokenKind, v string) handler {
	return func(l *lexer, r *regexp.Regexp) {
		l.advance(len(v))
		l.push(NewToken(k, v))

	}
}

func numberHandler() handler {
	return func(l *lexer, r *regexp.Regexp) {
		match := r.FindString(l.remainder())
		l.push(NewToken(Number, match))
		l.advance(len(match))
	}
}

func skipHandler() handler {
	return func(l *lexer, r *regexp.Regexp) {
		// lets just advanced all the way through whitespace
		match := r.FindStringIndex(l.remainder())
		l.advance(match[1])
	}
}

func stringHandler() handler {
	return func(l *lexer, r *regexp.Regexp) {
		match := r.FindStringIndex(l.remainder())
		str := l.remainder()[match[0]+1 : match[1]-1]
		// push the tokens and advance through the whole string
		l.push(NewToken(String, str))
		l.advance(match[1])
	}
}

func commentHandler() handler {
	return func(l *lexer, r *regexp.Regexp) {
		// advance through the whole comment
		match := r.FindStringIndex(l.remainder())
		l.advance(match[1])
	}
}

func symbolHandler() handler {
	return func(l *lexer, r *regexp.Regexp) {
		match := r.FindString(l.remainder())

		kind, exists := Keywords[match]
		if exists {
			l.push(NewToken(kind, match))
		} else {
			l.push(NewToken(Identifier, match))
		}

		l.advance(len(match))
	}
}

type lexer struct {
	patterns []pattern
	tokens   []Token
	source   string
	position int
}

func new(s string) lexer {
	return lexer{
		position: 0,
		source:   s,
		tokens:   make([]Token, 0),
		patterns: []pattern{
			{
				regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`),
				symbolHandler(),
			},
			{
				regexp.MustCompile(`[0-9]+(\.[0-9]+)?`),
				numberHandler(),
			},
			{
				regexp.MustCompile(`"(?:\\.|[^"\\])*"|'(?:\\.|[^'\\])*'`),
				stringHandler(),
			},
			{
				regexp.MustCompile(`\s+`),
				skipHandler(),
			},
			{
				regexp.MustCompile(`\/\/.*`),
				commentHandler(),
			},
			{
				regexp.MustCompile(`\[`),
				defaultHandler(OpenBracket, "["),
			},
			{
				regexp.MustCompile(`\]`),
				defaultHandler(CloseBracket, "]"),
			},
			{
				regexp.MustCompile(`\{`),
				defaultHandler(OpenCurly, "{"),
			},
			{
				regexp.MustCompile(`\}`),
				defaultHandler(CloseCurly, "}"),
			},
			{
				regexp.MustCompile(`\(`),
				defaultHandler(OpenParen, "("),
			},
			{
				regexp.MustCompile(`\)`),
				defaultHandler(CloseParen, ")"),
			},
			{
				regexp.MustCompile(`==`),
				defaultHandler(Equals, "=="),
			},
			{
				regexp.MustCompile(`!=`),
				defaultHandler(NotEquals, "!="),
			},
			{
				regexp.MustCompile(`=`),
				defaultHandler(Assignment, "="),
			},
			{
				regexp.MustCompile(`!`),
				defaultHandler(Not, "!"),
			},
			{
				regexp.MustCompile(`<`),
				defaultHandler(LessThan, "<"),
			},
			{
				regexp.MustCompile(`<=`),
				defaultHandler(LessThanEqual, "<="),
			},
			{
				regexp.MustCompile(`>`),
				defaultHandler(GreaterThan, ">"),
			},
			{
				regexp.MustCompile(`>=`),
				defaultHandler(GreaterThanEqual, ">="),
			},
			{
				regexp.MustCompile(`\|\|`),
				defaultHandler(Or, "||"),
			},
			{
				regexp.MustCompile(`&&`),
				defaultHandler(And, "&&"),
			},
			{
				regexp.MustCompile(`\.`),
				defaultHandler(Dot, "."),
			},
			{
				regexp.MustCompile(`\.\.`),
				defaultHandler(Range, ".."),
			},
			{
				regexp.MustCompile(`\.\.`),
				defaultHandler(Spread, "..."),
			},
			{
				regexp.MustCompile(`;`),
				defaultHandler(SemiColon, ";"),
			},
			{
				regexp.MustCompile(`:`),
				defaultHandler(Colon, ":"),
			},
			{
				regexp.MustCompile(`\?`),
				defaultHandler(Question, "?"),
			},
			{
				regexp.MustCompile(`,`),
				defaultHandler(Comma, ","),
			},
			{
				regexp.MustCompile(`\+\+`),
				defaultHandler(PlusPlus, "++"),
			},
			{
				regexp.MustCompile(`--`),
				defaultHandler(MinusMinus, "--"),
			},
			{
				regexp.MustCompile(`\+=`),
				defaultHandler(PlusEquals, "+="),
			},
			{
				regexp.MustCompile(`-=`),
				defaultHandler(MinusEquals, "-="),
			},
			{
				regexp.MustCompile(`\+`),
				defaultHandler(Plus, "+"),
			},
			{
				regexp.MustCompile(`-`),
				defaultHandler(Minus, "-"),
			},
			{
				regexp.MustCompile(`/`),
				defaultHandler(Divide, "/"),
			},
			{
				regexp.MustCompile(`\*`),
				defaultHandler(Multiply, "*"),
			},
			{
				regexp.MustCompile(`%`),
				defaultHandler(Modulus, "%"),
			},
		},
	}
}

// Advance the positon, by n amoumt, within the lexer.
func (l *lexer) advance(n int) {
	l.position += n
}

// Push a new token into the lexer.
func (l *lexer) push(t Token) {
	l.tokens = append(l.tokens, t)
}

// Get data from the source at the current position.
func (l lexer) at() byte {
	return l.source[l.position]
}

// Get the remainding code from the source
// from the current position to the end of file.
func (l lexer) remainder() string {
	return l.source[l.position:]
}

// Are we at the end of the file?
func (l lexer) eof() bool {
	return l.position >= len(l.source)
}

func Lex(s string) []Token {
	l := new(s)

	for !l.eof() {
		matched := false

		for _, p := range l.patterns {
			loc := p.regex.FindStringIndex(l.remainder())

			if loc != nil && loc[0] == 0 {
				p.handler(&l, p.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Unexpected token -> %q, %s", l.at(), l.remainder()))
		}
	}

	l.push(NewToken(EndOfFile, "EndOfFile"))

	return l.tokens
}
