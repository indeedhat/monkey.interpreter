package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

// parseLetStatement parses a let statement from the lexers token stream
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.Assign) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression(LowestPresedence)

	if !p.curTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}
