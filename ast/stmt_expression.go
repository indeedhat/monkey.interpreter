package ast

import "github.com/indeedhat/monkey-lang/token"

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
