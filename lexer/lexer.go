package lexer

import (
	"monkey/m/v2/token"
)

type Lexer struct {
	input        string
	position     int  //current position in input (points to the current character)
	readPosition int  //current reading position in input (after the current char)
	character    byte // current character that is under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//Returns the current token & move to the next possible token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "=="
			tok.Type = token.EQ
		} else {
			tok = newToken(token.ASSIGN, l.character)
		}
	case '+':
		tok = newToken(token.PLUS, l.character)
	case '-':
		tok = newToken(token.MINUS, l.character)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
		} else {
			tok = newToken(token.BANG, l.character)
		}
	case '/':
		tok = newToken(token.SLASH, l.character)
	case '*':
		tok = newToken(token.ASTERISK, l.character)
	case '<':
		tok = newToken(token.LT, l.character)
	case '>':
		tok = newToken(token.GT, l.character)
	case '(':
		tok = newToken(token.LPAREN, l.character)
	case ')':
		tok = newToken(token.RPAREN, l.character)
	case '{':
		tok = newToken(token.LBRACE, l.character)
	case '}':
		tok = newToken(token.RBRACE, l.character)
	case ',':
		tok = newToken(token.COMMA, l.character)
	case ';':
		tok = newToken(token.SEMICOLON, l.character)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.character) {
			literal := l.readIdentifier()
			tok.Literal = literal
			tok.Type = token.LookupKeywords(literal)
			return tok
		} else if isDigit(l.character) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.character)
	}
	l.readChar()
	return tok
}

// Update character, position & readPosition
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// Reads a string of characters & return the text
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.character) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.character == '"' || l.character == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.character) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Skips any different kind of white space
func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

// Peeks at the next character
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
