package ast

import (
	"bytes"

	"github.com/indeedhat/monkey-lang/token"
)

type ReturnStatement struct {
	Token token.Token
	Vaule Expression
}

// String implements Statement
func (n *ReturnStatement) String() string {
	var buf bytes.Buffer

	buf.WriteString(n.TokenLiteral() + " ")

	if n.Vaule != nil {
		buf.WriteString(n.Vaule.String())
	}

	buf.WriteString(";")

	return buf.String()
}

// statementNode implements Statement
func (n *ReturnStatement) statementNode() {
}

// TokenLiteral implements Node
func (n *ReturnStatement) TokenLiteral() string {
	return n.Token.Literal
}

var _ Statement = (*ReturnStatement)(nil)
