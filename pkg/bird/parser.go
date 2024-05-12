package bird

import (
	"fmt"
	"strconv"
)

type nodeType int

const (
	nodeBinaryExpr nodeType = iota
	nodeNumber
)

type Node interface {
	Type() nodeType
	String() string
}

type NumberNode struct {
	value float64
}

func (n *NumberNode) Type() nodeType {
	return nodeNumber
}

func (n *NumberNode) String() string {
	return fmt.Sprintf("%f", n.value)
}

type BinaryExprNode struct {
	operator string
	left     Node
	right    Node
}

func (b *BinaryExprNode) Type() nodeType {
	return nodeBinaryExpr
}

func (b *BinaryExprNode) String() string {
	return fmt.Sprintf("(%s %s %s)", b.left.String(), b.operator, b.right.String())
}

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) ParseExpression() Node {
	return p.parseAddition()
}

func (p *Parser) parseAddition() Node {
	node := p.parseMultiplication()

	for p.match(OPERATOR) && (p.peek().Literal == "+" || p.peek().Literal == "-") {
		operator := p.advance().Literal
		right := p.parseMultiplication()
		node = &BinaryExprNode{operator: operator, left: node, right: right}
	}

	return node
}

func (p *Parser) parseMultiplication() Node {
	node := p.parsePrimary()

	for p.match(OPERATOR) && (p.peek().Literal == "*" || p.peek().Literal == "/") {
		operator := p.advance().Literal
		right := p.parsePrimary()
		node = &BinaryExprNode{operator: operator, left: node, right: right}
	}

	return node
}

func (p *Parser) parsePrimary() Node {
	if p.match(NUMBER) {
		value := p.advance()
		v, _ := strconv.ParseFloat(value.Literal, 64)
		return &NumberNode{value: v}
	}
	// Handling error by panicking, which should be replaced with proper error handling
	panic("syntax error: expected a number or an expression")
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.tokens[p.current-1]
}

func (p *Parser) match(t TokenType) bool {
	if !p.isAtEnd() && p.peek().Type == t {
		return true
	}
	return false
}

func (p *Parser) peek() Token {
	if !p.isAtEnd() {
		return p.tokens[p.current]
	}
	return Token{Type: EOF, Literal: ""}
}

func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens)
}
