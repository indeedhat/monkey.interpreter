package parser

import (
	"strconv"

	"github.com/indeedhat/monkey-lang/ast"
)

// parseIntegerLiteral parses a token as an integer literal expression
func (p *Parser) parseIntegerLiteral() ast.Expression {
	expr := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errorf("could not parse %q as integer", p.curToken.Literal)
		return nil
	}

	expr.Value = value

	return expr
}
