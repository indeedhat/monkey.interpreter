package token

type TokenType string

const (
	// Special
	Illegal TokenType = "illegal"
	Eof     TokenType = "eof"

	// Identifiers & literals
	Ident TokenType = "ident"
	Int   TokenType = "int"

	// Operators
	Assign      TokenType = "assign"
	Plus        TokenType = "plus"
	Minus       TokenType = "minus"
	Bang        TokenType = "bang"
	Asterisk    TokenType = "asterisk"
	Slash       TokenType = "slash"
	LessThan    TokenType = "less-than"
	GreaterThan TokenType = "greater-than"

	Equal          TokenType = "equal"
	NotEqual       TokenType = "not-equal"
	GreaterOrEqual TokenType = "greater-or-equal"
	LessOrEqual    TokenType = "less-or-equal"

	// Delimiters
	Comma     TokenType = "comma"
	Semicolon TokenType = "semicolon"

	LParen TokenType = "lparen"
	RParen TokenType = "rparen"
	LBrace TokenType = "lbrace"
	RBrace TokenType = "rbrace"

	// Keywords
	Function TokenType = "function"
	Let      TokenType = "let"
	If       TokenType = "if"
	Else     TokenType = "else"
	Return   TokenType = "return"
	True     TokenType = "true"
	False    TokenType = "false"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Pos     int
}

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"if":     If,
	"else":   Else,
	"return": Return,
	"true":   True,
	"false":  False,
}

// LookupIdent looks up the identifer in the map of keywords and returns the appropriate
// token type for the string
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return Ident
}
