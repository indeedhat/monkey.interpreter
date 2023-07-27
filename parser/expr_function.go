package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	expr := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LParen) {
		return nil
	}

	expr.Parameters = p.parseFuncionParameters()

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

func (p *Parser) parseFunctionCallExpression(fn ast.Expression) ast.Expression {
	return &ast.FunctionCallExpression{
		Token:     p.curToken,
		Function:  fn,
		Arguments: p.parseFuncitonCallArguments(),
	}
}

func (p *Parser) parseFuncitonCallArguments() (args []ast.Expression) {
	p.nextToken()

	if p.curTokenIs(token.RParen) {
		return args
	}

	// i dont like this procedure but i cant think of a better way atm
	args = append(args, p.parseExpression(LowestPresedence))
	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LowestPresedence))
	}

	if !p.expectPeek(token.RParen) {
		return nil
	}

	return args
}
