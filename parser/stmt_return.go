package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

// parseReturnStatement parses a return statement from the lexers token stream
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	if !p.expectPeek(token.Int) {
		return nil
	}

	// TODO: assigning of expressin not yet implemented
	for !p.curTokenIs(token.Semicolon) {
		// skip expression tokens
		p.nextToken()
	}

	return stmt
}
