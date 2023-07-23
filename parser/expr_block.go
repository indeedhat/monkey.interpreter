package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}

	// advance cursor bast { token
	p.nextToken()

	// loop until either a } or EOF token is found
	for !p.curTokenIs(token.RBrace) && !p.curTokenIs(token.Eof) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		// p.parseStatement() leaves the cursor on the final token of the statement so we need to advance
		// the cursor before the next parse
		p.nextToken()
	}

	return block
}
