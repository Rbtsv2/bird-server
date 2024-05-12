package bird

import (
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for the "NUL" character
	} else {
		l.ch, _ = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(OPERATOR, l.ch)
	case '+':
		tok = newToken(OPERATOR, l.ch)
	case '-':
		tok = newToken(OPERATOR, l.ch)
	case '*':
		tok = newToken(OPERATOR, l.ch)
	case '/':
		tok = newToken(OPERATOR, l.ch)
	case ',':
		tok = newToken(PUNCTUATION, l.ch)
	case ';':
		tok = newToken(PUNCTUATION, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if unicode.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = IDENT
			return tok
		} else if unicode.IsDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = NUMBER
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func newToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
