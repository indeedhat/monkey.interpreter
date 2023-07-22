package parser

import (
	"github.com/indeedhat/monkey-lang/ast"
)

// parseIdentifier parses an identifier token into an expression
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}
