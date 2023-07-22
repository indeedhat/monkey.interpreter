package ast

import (
	"bytes"

	"github.com/indeedhat/monkey-lang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// String implements Node
func (n *Program) String() string {
	var buf bytes.Buffer

	for _, stmt := range n.Statements {
		buf.WriteString(stmt.String())
	}

	return buf.String()
}

// TokenLiteral implements Node
func (n *Program) TokenLiteral() string {
	if len(n.Statements) == 0 {
		return ""
	}

	return n.Statements[0].TokenLiteral()
}

var _ Node = (*Program)(nil)

type Identifier struct {
	Token token.Token
	Value string
}

// String implements Statement
func (n *Identifier) String() string {
	return n.Value
}

// expressionNode implements Expression
func (*Identifier) expressionNode() {
}

// TokenLiteral implements Node
func (n *Identifier) TokenLiteral() string {
	return n.Token.Literal
}

var _ Expression = (*Identifier)(nil)
