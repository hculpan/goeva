package lexer

import "regexp"

type TokenType int

const (
	INTEGER TokenType = iota
	FLOAT
	STRING
	PLUS
	MINUS
	TRUE
	FALSE
	IDENTIFIER
	STAR
	SLASH
	RIGHTPAREN
	LEFTPAREN
)

//go:generate stringer -type=TokenType

type Token struct {
	Type     TokenType
	Literal  string
	Line     int
	Col      int
	Filename string
}

type Tokens []Token

type TokenMatcher struct {
	Matcher    *regexp.Regexp
	Identifier TokenType
}

var TokenMatchers []TokenMatcher = []TokenMatcher{
	{Matcher: regexp.MustCompile(`^\d+\.\d+`), Identifier: FLOAT},
	{Matcher: regexp.MustCompile(`^\d+`), Identifier: INTEGER},
	{Matcher: regexp.MustCompile(`^\+`), Identifier: PLUS},
	{Matcher: regexp.MustCompile(`^\-`), Identifier: MINUS},
	{Matcher: regexp.MustCompile(`^\*`), Identifier: STAR},
	{Matcher: regexp.MustCompile(`^\/`), Identifier: SLASH},
	{Matcher: regexp.MustCompile(`^\(`), Identifier: LEFTPAREN},
	{Matcher: regexp.MustCompile(`^\)`), Identifier: RIGHTPAREN},
	{Matcher: regexp.MustCompile(`^"(.*?)"`), Identifier: STRING},
	{Matcher: regexp.MustCompile(`^true`), Identifier: TRUE},
	{Matcher: regexp.MustCompile(`^false`), Identifier: FALSE},
	{Matcher: regexp.MustCompile(`^[A-Za-z_][A-Za-z_0-9]+`), Identifier: IDENTIFIER},
}

func (t *Token) IsOperator() bool {
	return t.Type == PLUS || t.Type == MINUS || t.Type == STAR || t.Type == SLASH
}
