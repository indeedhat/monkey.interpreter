package parser

import (
	"fmt"

	"github.com/indeedhat/monkey-lang/ast"
	"github.com/indeedhat/monkey-lang/lexer"
	"github.com/indeedhat/monkey-lang/token"
)

type prefixParser func() ast.Expression
type infixParser func(ast.Expression) ast.Expression

// operatorPresedence defines the presedence values for each of the operator types
var operatorPresedence = map[token.TokenType]int{
	token.Equal:       Equals,
	token.NotEqual:    Equals,
	token.LessThan:    LessGreater,
	token.GreaterThan: LessGreater,
	token.Plus:        Sum,
	token.Minus:       Sum,
	token.Slash:       Product,
	token.Asterisk:    Product,
}

type Parser struct {
	lex *lexer.Lexer

	errors    []error
	curToken  token.Token
	peekToken token.Token

	prefixParsers map[token.TokenType]prefixParser
	infixParsers  map[token.TokenType]infixParser
}

// New creates a new parser for the provided lexer
func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		lex:           lex,
		prefixParsers: make(map[token.TokenType]prefixParser),
		infixParsers:  make(map[token.TokenType]infixParser),
	}

	p.registerPrefixParser(token.Ident, p.parseIdentifier)
	p.registerPrefixParser(token.Int, p.parseIntegerLiteral)
	p.registerPrefixParser(token.Bang, p.parsePrefixExpression)
	p.registerPrefixParser(token.Minus, p.parsePrefixExpression)
	p.registerPrefixParser(token.True, p.parseBooleanLiteral)
	p.registerPrefixParser(token.False, p.parseBooleanLiteral)
    p.registerPrefixParser(token.LParen, p.parseGroupedExpressions)

	p.registerInfixParser(token.Equal, p.parseInfixExpression)
	p.registerInfixParser(token.NotEqual, p.parseInfixExpression)
	p.registerInfixParser(token.LessThan, p.parseInfixExpression)
	p.registerInfixParser(token.GreaterThan, p.parseInfixExpression)
	p.registerInfixParser(token.Plus, p.parseInfixExpression)
	p.registerInfixParser(token.Minus, p.parseInfixExpression)
	p.registerInfixParser(token.Slash, p.parseInfixExpression)
	p.registerInfixParser(token.Asterisk, p.parseInfixExpression)

	// populate cur/next token fields
	p.nextToken()
	p.nextToken()

	return p
}

// ParseProgram parses the tokens in the lexer into an AST
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Type != token.Eof {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

// Errors returns a slice of errors generated by the parser
func (p *Parser) Errors() []error {
	return p.errors
}

// registerPrefixParser registers a prefix parser funcion for the token type
func (p *Parser) registerPrefixParser(tknType token.TokenType, parser prefixParser) {
	p.prefixParsers[tknType] = parser
}

// registerInfixParser registers an infix parser funtion for the token type
func (p *Parser) registerInfixParser(tknType token.TokenType, parser infixParser) {
	p.infixParsers[tknType] = parser
}

func (p *Parser) errorf(format string, args ...any) {
	p.errors = append(p.errors, fmt.Errorf(format, args...))
}

// nextToken advances the lexer to the next token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

// parseStatement parses the next statement in the lexers token stream
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// curPresedence gets the presedence of the current token from the operator table
func (p *Parser) curPresedence() int {
	if presedence, ok := operatorPresedence[p.curToken.Type]; ok {
		return presedence
	}

	return LowestPresedence
}

// peekPresedence gets the presedence of the next token from the operator table
func (p *Parser) peekPresedence() int {
	if presedence, ok := operatorPresedence[p.peekToken.Type]; ok {
		return presedence
	}

	return LowestPresedence
}

// peekTokenIs checks if the next token in the stream is of the provided type
func (p *Parser) peekTokenIs(tknType token.TokenType) bool {
	return p.peekToken.Type == tknType
}

// curTokenIs checks if the current token is of the given type
func (p *Parser) curTokenIs(tknType token.TokenType) bool {
	return p.curToken.Type == tknType
}

// peekError checks that the next token is of the given type
// if not then an error will be generated
func (p *Parser) peekError(tknType token.TokenType) (err error) {
	if !p.peekTokenIs(tknType) {
		p.errorf("Unexpected token type: expected(%s) found(%s)", tknType, p.peekToken.Type)
	}

	return
}

// expectPeek runs the peekError method and advances the token streabm if no erro is found
func (p *Parser) expectPeek(tknType token.TokenType) bool {
	if p.peekError(tknType) != nil {
		return false
	}

	p.nextToken()
	return true
}
