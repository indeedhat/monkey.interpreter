package ast

import (
	"bytes"
	"strings"

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

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consiquence *BlockStatement
	Alternative *BlockStatement
}

// String implements Expression
func (n *IfExpression) String() string {
	var buf bytes.Buffer

	buf.WriteString("if ")
	buf.WriteString(n.Condition.String())
	buf.WriteString(" ")
	buf.WriteString(n.Consiquence.String())

	if n.Alternative != nil {
		buf.WriteString(" else ")
		buf.WriteString(n.Alternative.String())
	}

	return buf.String()
}

// TokenLiteral implements Expression
func (n *IfExpression) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*IfExpression) expressionNode() {
}

var _ Expression = (*IfExpression)(nil)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

// String implements Expression
func (n *FunctionLiteral) String() string {
	var buf bytes.Buffer

	buf.WriteString(n.TokenLiteral())
	buf.WriteString("(")

	params := make([]string, 0, len(n.Parameters))
	for _, param := range n.Parameters {
		params = append(params, param.String())
	}
	buf.WriteString(strings.Join(params, ", "))

	buf.WriteString(")")
	buf.WriteString(n.Body.String())

	return buf.String()
}

// TokenLiteral implements Expression
func (n *FunctionLiteral) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*FunctionLiteral) expressionNode() {
}

var _ Expression = (*FunctionLiteral)(nil)

type FunctionCallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

// String implements Expression
func (n *FunctionCallExpression) String() string {
	var buf bytes.Buffer

	buf.WriteString(n.Function.String())
	buf.WriteString("(")

	args := make([]string, 0, len(n.Arguments))
	for _, arg := range n.Arguments {
		args = append(args, arg.String())
	}
	buf.WriteString(strings.Join(args, ", "))

	buf.WriteString(")")

	return buf.String()
}

// TokenLiteral implements Expression
func (n *FunctionCallExpression) TokenLiteral() string {
	return n.Token.Literal
}

// expressionNode implements Expression
func (*FunctionCallExpression) expressionNode() {
}

var _ Expression = (*FunctionCallExpression)(nil)
