package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

// parseGroupedExpressions is probably the coolest bit of code in this entire parser
func (p *Parser) parseGroupedExpressions() ast.Expression {
	// advance past the ( token
	p.nextToken()

	// parse the given expression recursively resetting the presedence to the lowest state
	exp := p.parseExpression(LowestPresedence)

	// ensure there is a closing paren
	if !p.expectPeek(token.RParen) {
		return nil
	}

	return exp
}
