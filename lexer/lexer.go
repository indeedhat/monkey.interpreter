package lexer

import (
	"bytes"

	"github.com/indeedhat/monkey-lang/token"
)

type Lexer struct {
	input string
	// current position of input (position of char)
	pos int
	// current reading pos (char + 1)
	readPos int
	// current char under exam
	char byte

	// line of the input
	line int
	// cursor pos on current line
	linePos int
}

// New creates a new Lexer instance with the provided input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	defer l.readChar()
	l.consumeWhitespace()

	switch l.char {
	case ';':
		return l.token(token.Semicolon, string(l.char))
	case ',':
		return l.token(token.Comma, string(l.char))
	case '(':
		return l.token(token.LParen, string(l.char))
	case ')':
		return l.token(token.RParen, string(l.char))
	case '{':
		return l.token(token.LBrace, string(l.char))
	case '}':
		return l.token(token.RBrace, string(l.char))
	case '[':
		return l.token(token.LBracket, string(l.char))
	case ']':
		return l.token(token.RBracket, string(l.char))
	case '+':
		return l.token(token.Plus, string(l.char))
	case '-':
		return l.token(token.Minus, string(l.char))
	case '=':
		if l.peekChar() == '=' {
			return l.token(token.Equal, l.multiReadChar(2))
		}
		return l.token(token.Assign, string(l.char))
	case '!':
		if l.peekChar() == '=' {
			return l.token(token.NotEqual, l.multiReadChar(2))
		}
		return l.token(token.Bang, string(l.char))
	case '*':
		return l.token(token.Asterisk, string(l.char))
	case '/':
		return l.token(token.Slash, string(l.char))
	case '<':
		if l.peekChar() == '=' {
			return l.token(token.LessOrEqual, l.multiReadChar(2))
		}
		return l.token(token.LessThan, string(l.char))
	case '>':
		if l.peekChar() == '=' {
			return l.token(token.GreaterOrEqual, l.multiReadChar(2))
		}
		return l.token(token.GreaterThan, string(l.char))
	case '"':
		str := l.readStringLiteral()
		if str == nil {
			return l.token(token.Illegal, string(l.char))
		}
		return l.token(token.String, *str)
	case 0:
		return token.Token{Type: token.Eof}
	default:
		if isIdentChar(l.char) {
			ident := l.readIdentifier()
			return l.token(token.LookupIdent(ident), ident)
		}
		if isDigit(l.char) {
			return l.token(token.Int, l.readNumber())
		}

		return l.token(token.Illegal, string(l.char))
	}
}

// New initializes a new token
func (l *Lexer) token(tokenType token.TokenType, char string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: char,
		Line:    l.line,
		Pos:     l.linePos - len(char),
	}
}

// readChar reads the next char in the input string
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos++
	l.linePos++
}

// readChar reads the next char in the input string
func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}

	return l.input[l.readPos]
}

// multiReadChar reads multiple characters as a string
func (l *Lexer) multiReadChar(n int) string {
	pos := l.pos
	for i := 0; i < n-1; i++ {
		l.readChar()
	}

	return l.input[pos:l.readPos]
}

// consumeWhitespace keeps reading characters until the current char is not a valid whitespace character
func (l *Lexer) consumeWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
			l.linePos = 0
		}
		l.readChar()
	}
}

// readIdentifier reads a string of consecutive valid identifier characters
func (l *Lexer) readIdentifier() string {
	pos := l.pos

	for isIdentChar(l.peekChar()) {
		l.readChar()
	}

	return l.input[pos:l.readPos]
}

// readStringLiteral reads a string literal
// this will only accept double quoted strings and can have double quotes escaped with a \
func (l *Lexer) readStringLiteral() *string {
	var buf bytes.Buffer

	for {
		char := l.input[l.pos]
		peekChar := l.peekChar()
		// if we dont find a closing " then its an invalid string
		if peekChar == 0 {
			return nil
		}

		if char != '\\' || peekChar != '"' {
			buf.WriteString(string(char))
		}

		l.readChar()
		if peekChar == '"' && char != '\\' {
			break
		}
	}

	// skip final "
	if l.peekChar() == '"' {
		l.readChar()
	}

	// dont include the " on each side in the strings value
	str := buf.String()[1:]
	return &str
}

// readNumber reads a numeric value from the input string
func (l *Lexer) readNumber() string {
	pos := l.pos

	for isDigit(l.peekChar()) {
		l.readChar()
	}

	return l.input[pos:l.readPos]
}
