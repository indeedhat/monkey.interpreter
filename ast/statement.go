package ast

import (
	"bytes"

	"github.com/indeedhat/monkey-lang/token"
)

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// String implements Statement
func (n *ExpressionStatement) String() string {
	if n.Expression == nil {
		return ""
	}

	return n.Expression.String()
}

// TokenLiteral implements Statement
func (n *ExpressionStatement) TokenLiteral() string {
	return n.Token.Literal
}

// statementNode implements Statement
func (*ExpressionStatement) statementNode() {
}

var _ Statement = (*ExpressionStatement)(nil)

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
