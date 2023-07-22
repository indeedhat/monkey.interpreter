package ast

import (
	"bytes"

	"github.com/indeedhat/monkey-lang/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// String implements Statement
func (n *LetStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString(n.TokenLiteral() + " ")
	buf.WriteString(n.Name.Value)
	buf.WriteString(" = ")

	if n.Value != nil {
		buf.WriteString(n.Value.String())
	}

	buf.WriteString(";")

	return buf.String()
}

// statementNode implements Statement
func (n *LetStatement) statementNode() {
}

// TokenLiteral implements Node
func (n *LetStatement) TokenLiteral() string {
	return n.Token.Literal
}

var _ Statement = (*LetStatement)(nil)
