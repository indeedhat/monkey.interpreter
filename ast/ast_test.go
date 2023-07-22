package ast

import (
	"testing"

	"github.com/indeedhat/monkey-lang/token"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.Let, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.Ident, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	assert.Equal(t, `let myVar = anotherVar;`, program.String())
}
