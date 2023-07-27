package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{Token: p.curToken}

	// advance to the next token so long as it is a (
	if !p.expectPeek(token.LParen) {
		return nil
	}

	// advance to first token of the expression
	p.nextToken()
	expr.Condition = p.parseExpression(LowestPresedence)

	// ensure and advance to the closing ) and opening {
	if !p.expectPeek(token.RParen) {
		return nil
	}
	if !p.expectPeek(token.LBrace) {
		return nil
	}

	// parse the block statment for the if consiquence clause
	// this will consume all statements until the closing }
	expr.IfBlock = p.parseBlockStatement()

	if p.peekTokenIs(token.Else) {
		p.nextToken()

		if !p.expectPeek(token.LBrace) {
			return nil
		}

		expr.ElseBlock = p.parseBlockStatement()
	}

	return expr
}
