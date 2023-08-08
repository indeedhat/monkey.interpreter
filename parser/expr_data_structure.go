package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

func (p *Parser) parseArrayLiteral() ast.Expression {
	return &ast.ArrayLiteral{
		Token:    p.curToken,
		Elements: p.parseExpressionList(token.RBracket),
	}
}

func (p *Parser) parseIndexExpression(subject ast.Expression) ast.Expression {
	expr := &ast.IndexExpression{
		Token:   p.curToken,
		Subject: subject,
	}

	p.nextToken()

	expr.Index = p.parseExpression(LowestPresedence)

	if !p.expectPeek(token.RBracket) {
		return nil
	}

	return expr
}
