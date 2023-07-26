package parser

import (
	"strconv"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/token"
)

// parseNullLiteral parses a token as a null literal expression
func (p *Parser) parseNullLiteral() ast.Expression {
	return &ast.NullLiteral{Token: p.curToken}
}

// parseBooleanLiteral parses a token as a boolean literal expression
func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(token.True)}
}

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
