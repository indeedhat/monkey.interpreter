package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
)

// parseInfixExpression parses an infix expression relative to the left expression
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	presedence := p.curPresedence()

	p.nextToken()
	expression.Right = p.parseExpression(presedence)

	return expression
}
