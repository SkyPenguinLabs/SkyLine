///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_InterfacePlug
// Extension         | .go ( golang source code file )
// Purpose           | Define constant definitions for string values of Tokens
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines the main scanner, the function that is called to start lexical analysis which will help categorize and read specific byte orders or byte values. The scanner
//
// will take some input code like `INTEGER value := 10;` and seperate it into readable information like the following below
//
//  {
//		{'INTEGER': TYPE}
//      {'value': IDENTIFIER}
//      {':=': SLTK.TOKEN_ASSIGNMENT}
//      {'10': VALUE}
//      {';': SLTK.TOKEN_SEMICOLON}
//	}
//
// which can then be passed onto the parser for further dissection and 'parsing'. This file will define the function to scan those tokens, split the input and run other functions
//
// which can be used to better help and better parse functions.
//
package SkyLine_Backend_Scanner

import (
	SLTK "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens"
)

func (SL_Scanner *SL_ScannerStructure) NT() SLTK.SL_TokenConstruct {
	SL_Scanner.SkipWhitespace()
	if SL_Scanner.Scanner_Character == '/' {
		switch SL_Scanner.CharacterPeek() {
		case '/':
			SL_Scanner.SkipSingleLineComment()
			return SL_Scanner.NT()
		case '*':
			SL_Scanner.SkipMultiLineComment()
			return SL_Scanner.NT()
		}
	}
	if SL_Scanner.Scanner_Character == '!' && SL_Scanner.CharacterPeek() == '-' {
		SL_Scanner.SkipMultiLineComment()
		return SL_Scanner.NT()
	}
	var tok SLTK.SL_TokenConstruct
	switch SL_Scanner.Scanner_Character {
	case '`':
		tok.Token_Type = SLTK.TOKEN_STRING
		tok.Literal = SL_Scanner.ReadBacktick()
	case '.':
		if SL_Scanner.CharacterPeek() == '.' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_DOTDOT,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_PERIOD, SL_Scanner.Scanner_Character)
		}
	case '@':
		tok = ScanNewToken(SLTK.TOKEN_Lable, SL_Scanner.Scanner_Character)
	case '&':
		if SL_Scanner.CharacterPeek() == '&' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_ANDAND,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_BITWISE_OP_AND, SL_Scanner.Scanner_Character)
		}
	case '=':
		if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_EQ,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_ASSIGN, SL_Scanner.Scanner_Character)
		}
	case rune('!'):
		if SL_Scanner.CharacterPeek() == rune('=') {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{Token_Type: SLTK.TOKEN_NEQ, Literal: string(ch) + string(SL_Scanner.Scanner_Character)}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_BANG, SL_Scanner.Scanner_Character)
		}
	case ';':
		tok = ScanNewToken(SLTK.TOKEN_SEMICOLON, SL_Scanner.Scanner_Character)
	case ':':
		if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_PEEKASSIGN,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else if SL_Scanner.CharacterPeek() == ':' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_MODULECALL,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_COLON, SL_Scanner.Scanner_Character)
		}
	case '(':
		tok = ScanNewToken(SLTK.TOKEN_LPAREN, SL_Scanner.Scanner_Character)
	case '|':
		if SL_Scanner.CharacterPeek() == '|' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_OROR,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_BITWISE_OP_OR, SL_Scanner.Scanner_Character)
		}
	case ')':
		tok = ScanNewToken(SLTK.TOKEN_RPAREN, SL_Scanner.Scanner_Character)
	case ',':
		tok = ScanNewToken(SLTK.TOKEN_COMMA, SL_Scanner.Scanner_Character)
	case '+':
		tok = ScanNewToken(SLTK.TOKEN_PLUS, SL_Scanner.Scanner_Character)
		if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_PLUS_EQUALS,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else if SL_Scanner.CharacterPeek() == '+' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_PLUS_PLUS,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_PLUS, SL_Scanner.Scanner_Character)
		}
	case '-':
		if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_MINUS_EQUALS,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else if SL_Scanner.CharacterPeek() == '-' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_MINUS_MINUS,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_MINUS, SL_Scanner.Scanner_Character)
		}
	case '*':
		if SL_Scanner.CharacterPeek() == '*' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_POWEROF,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_ASTERISK_EQUALS,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_ASTARISK, SL_Scanner.Scanner_Character)
		}
	case '%':
		tok = ScanNewToken(SLTK.TOKEN_MODULO, SL_Scanner.Scanner_Character)
	case '/':
		if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_DIVEQ,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_SLASH, SL_Scanner.Scanner_Character)
		}
	case '<':
		if SL_Scanner.CharacterPeek() == '=' {
			ch := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_LTEQ,
				Literal:    string(ch) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_LT, SL_Scanner.Scanner_Character)
		}
	case '>':
		if SL_Scanner.CharacterPeek() == '=' {
			cha := SL_Scanner.Scanner_Character
			SL_Scanner.RCHAR()
			tok = SLTK.SL_TokenConstruct{
				Token_Type: SLTK.TOKEN_GTEQ,
				Literal:    string(cha) + string(SL_Scanner.Scanner_Character),
			}
		} else {
			tok = ScanNewToken(SLTK.TOKEN_GT, SL_Scanner.Scanner_Character)
		}
	case '{':
		tok = ScanNewToken(SLTK.TOKEN_LBRACE, SL_Scanner.Scanner_Character)
	case '}':
		tok = ScanNewToken(SLTK.TOKEN_RBRACE, SL_Scanner.Scanner_Character)
	case '[':
		tok = ScanNewToken(SLTK.TOKEN_LBRACKET, SL_Scanner.Scanner_Character)
	case ']':
		tok = ScanNewToken(SLTK.TOKEN_RBRACKET, SL_Scanner.Scanner_Character)
	case '"':
		tok.Token_Type = SLTK.TOKEN_STRING
		tok.Literal = SL_Scanner.R_STRING(SL_Scanner.Scanner_Character)
	case 0:
		tok.Literal = ""
		tok.Token_Type = SLTK.TOKEN_EOF
	default:
		if CharacterInputIsDigit(SL_Scanner.Scanner_Character) {
			return SL_Scanner.R_DECIMAL()
		}

		if CharacterInputIsLetter(SL_Scanner.Scanner_Character) {
			tok.Literal = SL_Scanner.R_IDENTIFIER()
			tok.Token_Type = ReadVerifyIdentifier(tok.Literal)
			SL_Scanner.Scanner_PreviousToken = tok
			return tok
		}

		tok = ScanNewToken(SLTK.TOKEN_ILLEGAL, SL_Scanner.Scanner_Character)
	}
	SL_Scanner.RCHAR()
	SL_Scanner.Scanner_PreviousToken = tok
	return tok
}
