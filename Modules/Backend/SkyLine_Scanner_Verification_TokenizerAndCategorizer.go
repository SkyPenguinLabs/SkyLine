/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Scanner_Verification_TokenizerAndCategorizer.go
// Project       |  SkyLine programming language
// Line Count    |  50+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines prime categorization functions for the scanner
//
//
package SkyLine_Backend

import (
	"regexp"
)

type TokenizerFunc func(*ScannerStructure) Token
type TokenizerPeeked func(*ScannerStructure) Token

func Categorize_EOF(lex *ScannerStructure) Token {
	var tk Token
	tk.Literal = ""
	tk.Token_Type = TOKEN_EOF
	return tk
}

func Categorize_String(lex *ScannerStructure) Token {
	var tk Token
	tk.Token_Type = TOKEN_STRING
	tk.Literal = lex.ReadString()
	return tk
}

const (
	MODULECALL = "::"
)

// A work in progress
// Not yet implemented
var ScannerTokenizationRegularExpressions = map[*regexp.Regexp]Token_Type{
	regexp.MustCompile(`^DEFAULT$`):                               TOKEN_DEFAULT,
	regexp.MustCompile(`^FOR$`):                                   TOKEN_FOR,
	regexp.MustCompile(`^IMPORT$`):                                TOKEN_IMPORT,
	regexp.MustCompile(`^module$`):                                TOKEN_MODULE,
	regexp.MustCompile(`^ENGINE$`):                                TOKEN_KEYWORD_ENGINE,
	regexp.MustCompile(`^ENGINE::ENVIRONMENT_MODIFIER->CALL:::>`): TOKEN_ENGINE_TYPE,
	regexp.MustCompile(`^FOREACH$`):                               TOKEN_FOREACH,
	regexp.MustCompile(`^IN$`):                                    TOKEN_INSIDE,
	regexp.MustCompile(`^REGISTER$`):                              TOKEN_REGISTER,
	regexp.MustCompile(`^ILLEGAL$`):                               TOKEN_ILLEGAL,
	regexp.MustCompile(`^EOF$`):                                   TOKEN_EOF,
	regexp.MustCompile(`^TOKEN_IDENT$`):                           TOKEN_IDENT,
	regexp.MustCompile(`^INT$`):                                   TOKEN_INT,
	regexp.MustCompile(`^FLOAT$`):                                 TOKEN_FLOAT,
	regexp.MustCompile(`^STRING$`):                                TOKEN_STRING,
	regexp.MustCompile(`^CONST$`):                                 TOKEN_CONSTANT,
	regexp.MustCompile(`^FUNCTION$`):                              TOKEN_FUNCTION,
	regexp.MustCompile(`^LET$`):                                   TOKEN_LET,
	regexp.MustCompile(`^SET$`):                                   TOKEN_LET,
	regexp.MustCompile(`^TOKEN_TRUE$`):                            TOKEN_TRUE,
	regexp.MustCompile(`^FALSE$`):                                 TOKEN_FALSE,
	regexp.MustCompile(`^IF$`):                                    TOKEN_IF,
	regexp.MustCompile(`^ELSE$`):                                  TOKEN_ELSE,
	regexp.MustCompile(`^RETURN$`):                                TOKEN_RETURN,
	regexp.MustCompile(`^SWITCH$`):                                TOKEN_SWITCH,
	regexp.MustCompile(`^CASE$`):                                  TOKEN_CASE,
	regexp.MustCompile(`^<=$`):                                    TOKEN_LTEQ,
	regexp.MustCompile(`^>=$`):                                    TOKEN_GTEQ,
	regexp.MustCompile(`^\*=$`):                                   TOKEN_ASTERISK_EQUALS,
	regexp.MustCompile(`^!$`):                                     TOKEN_BANG,
	regexp.MustCompile(`^=$`):                                     TOKEN_ASSIGN,
	regexp.MustCompile(`^\+$`):                                    TOKEN_PLUS,
	regexp.MustCompile(`^\-$`):                                    TOKEN_MINUS,
	regexp.MustCompile(`^\*$`):                                    TOKEN_ASTARISK,
	regexp.MustCompile(`^/$`):                                     TOKEN_SLASH,
	regexp.MustCompile(`^<$`):                                     TOKEN_LT,
	regexp.MustCompile(`^>$`):                                     TOKEN_GT,
	regexp.MustCompile(`^==$`):                                    TOKEN_EQ,
	regexp.MustCompile(`^-=\$`):                                   TOKEN_MINUS_EQUALS,
	regexp.MustCompile(`^!=$`):                                    TOKEN_NEQ,
	regexp.MustCompile(`^/=$`):                                    TOKEN_DIVEQ,
	regexp.MustCompile(`^\.$`):                                    TOKEN_PERIOD,
	regexp.MustCompile(`^\+=$`):                                   TOKEN_PLUS_EQUALS,
	regexp.MustCompile(`^,$`):                                     TOKEN_COMMA,
	regexp.MustCompile(`^;$`):                                     TOKEN_SEMICOLON,
	regexp.MustCompile(`^:$`):                                     TOKEN_COLON,
	regexp.MustCompile(`^\($`):                                    TOKEN_LPAREN,
	regexp.MustCompile(`^\)$`):                                    TOKEN_RPAREN,
	regexp.MustCompile(`^\|$`):                                    TOKEN_LINE,
	regexp.MustCompile(`^{`):                                      TOKEN_LBRACE,
	regexp.MustCompile(`^}$`):                                     TOKEN_RBRACE,
	regexp.MustCompile(`^\[$`):                                    TOKEN_LBRACKET,
	regexp.MustCompile(`^\]$`):                                    TOKEN_RBRACKET,
	regexp.MustCompile(`^\|\|$`):                                  TOKEN_OROR,
	regexp.MustCompile(`^&&$`):                                    TOKEN_ANDAND,
	regexp.MustCompile(`^\^*$`):                                   TOKEN_POWEROF,
	regexp.MustCompile(`^%$`):                                     TOKEN_MODULO,
	regexp.MustCompile(`^[a-zA-Z_]\w*$`):                          TOKEN_IDENT,
	regexp.MustCompile(`^-?\d+$`):                                 TOKEN_INT,
	regexp.MustCompile(`^-?\d+\.\d*$`):                            TOKEN_FLOAT,
	regexp.MustCompile(`^\/.*\/[gim]*$`):                          TOKEN_REGEXP,
}

func (lex *ScannerStructure) Scan_NT() Token {
	lex.ConsumeWhiteSpace()
	if lex.Char == '/' {
		switch lex.Peek() {
		case '/':
			lex.ConsumeComment()
			return lex.Scan_NT()
		case '*':
			lex.ConsumeMultiLineComment()
			return lex.Scan_NT()
		}
	}
	if lex.Char == '!' && lex.Peek() == '-' {
		lex.ConsumeMultiLineComment()
		return lex.Scan_NT()
	}
	var tok Token
	switch lex.Char {
	case '`':
		tok.Token_Type = TOKEN_STRING
		tok.Literal = lex.ReadBacktick()
	case '.':
		if lex.Peek() == '.' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_DOTDOT,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_PERIOD, lex.Char)
		}
	case '&':
		if lex.Peek() == '&' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_ANDAND,
				Literal:    string(ch) + string(lex.Char),
			}
		}
	case '=':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_EQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_ASSIGN, lex.Char)
		}
	case '!':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_NEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_BANG, lex.Char)
		}
	case ';':
		tok = ScanNewToken(TOKEN_SEMICOLON, lex.Char)
	case ':':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_ASSIGN,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == ':' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: MODULECALL,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_COLON, lex.Char)
		}
	case '(':
		tok = ScanNewToken(TOKEN_LPAREN, lex.Char)
	case '|':
		if lex.Peek() == '|' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_OROR,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_LINE, lex.Char)
		}
	case ')':
		tok = ScanNewToken(TOKEN_RPAREN, lex.Char)
	case ',':
		tok = ScanNewToken(TOKEN_COMMA, lex.Char)
	case '+':
		tok = ScanNewToken(TOKEN_PLUS, lex.Char)
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_PLUS_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == '+' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_PLUS_PLUS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_PLUS, lex.Char)
		}
	case '-':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_MINUS_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == '-' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_MINUS_MINUS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_MINUS, lex.Char)
		}
	case '*':
		if lex.Peek() == '*' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_POWEROF,
				Literal:    string(ch) + string(lex.Char),
			}
		} else if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_ASTERISK_EQUALS,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_ASTARISK, lex.Char)
		}
	case '%':
		tok = ScanNewToken(TOKEN_MODULO, lex.Char)
	case '/':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_DIVEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_SLASH, lex.Char)
		}
	case '<':
		if lex.Peek() == '=' {
			ch := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_LTEQ,
				Literal:    string(ch) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_LT, lex.Char)
		}
	case '>':
		if lex.Peek() == '=' {
			cha := lex.Char
			lex.ReadChar()
			tok = Token{
				Token_Type: TOKEN_GTEQ,
				Literal:    string(cha) + string(lex.Char),
			}
		} else {
			tok = ScanNewToken(TOKEN_GT, lex.Char)
		}
	case '{':
		tok = ScanNewToken(TOKEN_LBRACE, lex.Char)
	case '}':
		tok = ScanNewToken(TOKEN_RBRACE, lex.Char)
	case '[':
		tok = ScanNewToken(TOKEN_LBRACKET, lex.Char)
	case ']':
		tok = ScanNewToken(TOKEN_RBRACKET, lex.Char)
	case '"':
		tok.Token_Type = TOKEN_STRING
		tok.Literal = lex.ReadString()
	case 0:
		tok.Literal = ""
		tok.Token_Type = TOKEN_EOF
	default:
		if CharIsDigit(lex.Char) {
			return lex.ReadIntToken()
		}

		if CharIsLetter(lex.Char) {
			tok.Literal = lex.ReadIdentifier()
			tok.Token_Type = LookupIdentifier(tok.Literal)
			lex.PrevTok = tok
			return tok
		}

		tok = ScanNewToken(TOKEN_ILLEGAL, lex.Char)
	}
	lex.ReadChar()
	lex.PrevTok = tok
	return tok
}
