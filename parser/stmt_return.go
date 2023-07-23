package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

// parseReturnStatement parses a return statement from the lexers token stream
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	stmt.Vaule = p.parseExpression(LowestPresedence)

	if !p.curTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}
