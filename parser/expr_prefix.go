package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
)

// parsePrefixExpression parses a prefix expression
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(Prefix)

	return expression
}
