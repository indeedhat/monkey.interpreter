package lexer

import (
	"testing"

	"github.com/indeedhat/monkey-lang/token"
)

var nextTokenTests = []struct {
	expectedType    token.TokenType
	expectedLiteral string
}{
	{token.Let, "let"},
	{token.Ident, "five"},
	{token.Assign, "="},
	{token.Int, "5"},
	{token.Semicolon, ";"},
	{token.Let, "let"},
	{token.Ident, "ten"},
	{token.Assign, "="},
	{token.Int, "10"},
	{token.Semicolon, ";"},

	{token.Let, "let"},
	{token.Ident, "add"},
	{token.Assign, "="},
	{token.Function, "fn"},
	{token.LParen, "("},
	{token.Ident, "x"},
	{token.Comma, ","},
	{token.Ident, "y"},
	{token.RParen, ")"},
	{token.LBrace, "{"},
	{token.Ident, "x"},
	{token.Plus, "+"},
	{token.Ident, "y"},
	{token.Semicolon, ";"},
	{token.RBrace, "}"},
	{token.Semicolon, ";"},

	{token.Let, "let"},
	{token.Ident, "result"},
	{token.Assign, "="},
	{token.Ident, "add"},
	{token.LParen, "("},
	{token.Ident, "five"},
	{token.Comma, ","},
	{token.Ident, "ten"},
	{token.RParen, ")"},
	{token.Semicolon, ";"},

	{token.Bang, "!"},
	{token.Minus, "-"},
	{token.Slash, "/"},
	{token.Asterisk, "*"},
	{token.Int, "5"},
	{token.Semicolon, ";"},
	{token.Int, "5"},
	{token.LessThan, "<"},
	{token.Int, "10"},
	{token.GreaterThan, ">"},
	{token.Int, "5"},
	{token.Semicolon, ";"},

	{token.If, "if"},
	{token.LParen, "("},
	{token.Int, "5"},
	{token.LessThan, "<"},
	{token.Int, "10"},
	{token.RParen, ")"},
	{token.LBrace, "{"},
	{token.Return, "return"},
	{token.True, "true"},
	{token.Semicolon, ";"},
	{token.RBrace, "}"},
	{token.Else, "else"},
	{token.LBrace, "{"},
	{token.Return, "return"},
	{token.False, "false"},
	{token.Semicolon, ";"},
	{token.RBrace, "}"},

	{token.Int, "10"},
	{token.Equal, "=="},
	{token.Int, "10"},
	{token.Semicolon, ";"},
	{token.Int, "10"},
	{token.NotEqual, "!="},
	{token.Int, "9"},
	{token.Semicolon, ";"},
	{token.Int, "10"},
	{token.LessOrEqual, "<="},
	{token.Int, "10"},
	{token.Semicolon, ";"},
	{token.Int, "10"},
	{token.GreaterOrEqual, ">="},
	{token.Int, "10"},
	{token.Semicolon, ";"},

	{token.String, "this is a string"},
	{token.Semicolon, ";"},
	{token.String, `this is a string with "quotes"`},
	{token.Semicolon, ";"},
	{token.String, `this is a string with a \ (backslash)`},
	{token.Semicolon, ";"},

	{token.LBracket, "["},
	{token.Int, "0"},
	{token.Comma, ","},
	{token.Int, "1"},
	{token.RBracket, "]"},
	{token.Semicolon, ";"},

	{token.Eof, ""},
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
x + y;
};

let result = add(five, ten);

!-/*5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}

10 == 10;
10 != 9;
10 <= 10;
10 >= 10;

"this is a string";
"this is a string with \"quotes\"";
"this is a string with a \ (backslash)";

[0, 1];
`
	//
	//data[0];
	//`

	// TODO: this test is broken for strings, if they have a ; then it reads them as ; tokenres
	lex := New(input)

	for i, tst := range nextTokenTests {
		tok := lex.NextToken()

		if tok.Type != tst.expectedType {
			t.Errorf("tests[%d] - wrong token type. expected(%q) found(%q)",
				i,
				tst.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tst.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal.  expected(%q) found(%q)",
				i,
				tst.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
