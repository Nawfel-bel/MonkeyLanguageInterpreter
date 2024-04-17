package Lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) (l *Lexer) {
	l = &Lexer{input: input}
	l.ReadChar()
	return
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.EQ, Literal: string(l.ch) + string(ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.ReadChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.ADD, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.ReadChar()
	return
}

func newToken(TokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: TokenType, Literal: string(literal)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || '_' == ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}

func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}
