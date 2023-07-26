package token

type TokenType string

const (
	// Special
	Illegal TokenType = "ILLEGAL"
	Eof     TokenType = "EOF"

	// Identifiers & literals
	Ident  TokenType = "IDENT"
	Int    TokenType = "INT"
	String TokenType = "STRING"

	// Operators
	Assign      TokenType = "="
	Plus        TokenType = "+"
	Minus       TokenType = "-"
	Bang        TokenType = "!"
	Asterisk    TokenType = "*"
	Slash       TokenType = "/"
	LessThan    TokenType = "<"
	GreaterThan TokenType = ">"

	Equal          TokenType = "=="
	NotEqual       TokenType = "!="
	GreaterOrEqual TokenType = ">="
	LessOrEqual    TokenType = "<="

	// Delimiters
	Comma     TokenType = ","
	Semicolon TokenType = ";"

	LParen   TokenType = "("
	RParen   TokenType = ")"
	LBrace   TokenType = "{"
	RBrace   TokenType = "}"
	LBracket TokenType = "["
	RBracket TokenType = "]"

	// Keywords
	Function TokenType = "FUNCTION"
	Let      TokenType = "LET"
	If       TokenType = "IF"
	Else     TokenType = "ELSE"
	Return   TokenType = "RETURN"
	True     TokenType = "TRUE"
	False    TokenType = "FALSE"
	Null     TokenType = "NULL"
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
	"null":   Null,
}

// LookupIdent looks up the identifer in the map of keywords and returns the appropriate
// token type for the string
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return Ident
}
