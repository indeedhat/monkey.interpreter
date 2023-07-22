package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

func (p *Parser) parseGroupedExpressions() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LowestPresedence)

	if !p.expectPeek(token.RParen) {
		return nil
	}

	return exp
}
