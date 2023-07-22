package ast

import (
	"bytes"

	"github.com/indeedhat/monkey-lang/token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// String implements Expression
func (n *InfixExpression) String() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(n.Left.String())
	buf.WriteString(" " + n.Operator + " ")
	buf.WriteString(n.Right.String())
	buf.WriteString(")")

	return buf.String()
}

// TokenLiteral implements Expression
func (n *InfixExpression) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*InfixExpression) expressionNode() {
}

var _ Expression = (*InfixExpression)(nil)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// String implements Expression
func (n *IntegerLiteral) String() string {
	return n.Token.Literal
}

// TokenLiteral implements Expression
func (n *IntegerLiteral) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*IntegerLiteral) expressionNode() {
}

var _ Expression = (*IntegerLiteral)(nil)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// String implements Expression
func (n *PrefixExpression) String() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(n.Operator)
	buf.WriteString(n.Right.String())
	buf.WriteString(")")

	return buf.String()
}

// TokenLiteral implements Expression
func (n *PrefixExpression) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*PrefixExpression) expressionNode() {
}

var _ Expression = (*PrefixExpression)(nil)

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

// String implements Expression
func (n *BooleanLiteral) String() string {
	return n.Token.Literal
}

// TokenLiteral implements Expression
func (n *BooleanLiteral) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*BooleanLiteral) expressionNode() {
}

var _ Expression = (*BooleanLiteral)(nil)
