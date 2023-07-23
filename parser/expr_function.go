package parser

import (
	"log"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	expr := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LParen) {
		return nil
	}

	expr.Parameters = p.parseFuncionParameters()

	log.Print(p.curToken.Literal)
	if !p.expectPeek(token.LBrace) {
		return nil
	}

	expr.Body = p.parseBlockStatement()

	return expr
}

func (p *Parser) parseFuncionParameters() (params []*ast.Identifier) {
	// advance past the (
	p.nextToken()

	for !p.curTokenIs(token.RParen) && !p.curTokenIs(token.Eof) {
		ident := p.parseIdentifier()
		if ident == nil {
			return nil
		}

		params = append(params, ident.(*ast.Identifier))
		p.nextToken()

		if p.curTokenIs(token.Comma) {
			p.nextToken()
		}
	}

	return params
}
