package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

const (
	_ = iota
	LowestPresedence
	Equals
	LessGreater
	Sum
	Product
	Prefix
	Call
)

// parseExpressionStatement parses an expression into a statement wrapper
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LowestPresedence)

	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(presedence int) ast.Expression {
	prefix := p.prefixParsers[p.curToken.Type]
	if prefix == nil {
		p.errorf("no prefix parser found for %s", p.curToken.Type)
		return nil
	}

	leftExpr := prefix()
	return leftExpr
}
