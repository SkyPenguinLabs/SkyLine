////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
//
// The SkyLine configuration language is a language and engine designed to act as a modification language to the SkyLine programming language. This language is
// very minimal and contains a regex base lexer, a very basic parser, a few keywords, a base interpreter and that is all as well as some backend engine code. This
// language is purely modified to be an extension to the SkyLine programming language, something that can be a pre processor language post processing for the main
// SkyLine script. Below is more technical information for the language
//
// Lexer       : Regex based lexer with minimal tokens and base syntax
// Parser      : Base parser with minimal tokens and base syntax with simple error systems
// REPL        : Does not exist
// Environment : Extremely minimal
// Types       : String, Boolean, Integer
// Statements  : set, import, use, errors, output, system, constant/const
//
// File Contains -> This file contains all main scanner functions or the main scanner component to tokenize functions

package SkyLine_Configuration_Engine_Backend_Source

func (Scanner *ScannerStructureRegister) NT() TokenRegistry {
	var tok TokenRegistry
	Scanner.skipWhitespace()
	for re, tokType := range ScannerTokenizationRegularExpressions {
		if re.MatchString(string(Scanner.Char)) {
			tok = newToken(tokType, Scanner.Char)
			Scanner.readChar()
			return tok
		}
	}
	switch Scanner.Char {
	case '-':
		if Scanner.peekChar() == '>' {

			ch := Scanner.Char
			Scanner.readChar()
			tok = TokenRegistry{
				TokenDataType: MODIFY_Token,
				Literal:       string(ch) + string(Scanner.Char),
			}
		}
	case '"':
		tok.TokenDataType = STRING_Token
		tok.Literal = Scanner.readString()
	case 0:
		tok.Literal = ""
		tok.TokenDataType = END_OF_FILE
	default:
		if isLetter(Scanner.Char) {
			tok.Literal = Scanner.readIdentifier()
			tok.TokenDataType = SL_IDENT_LOOKUP(tok.Literal)
			return tok
		} else if isDigit(Scanner.Char) {
			tok.TokenDataType = INTEGER_Token
			tok.Literal = Scanner.readNumber()
			return tok
		} else {
			tok = newToken(ALIENATED_UNVERSED, Scanner.Char)
		}
	}

	Scanner.readChar()
	return tok
}

func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
}

func (Scanner *ScannerStructureRegister) skipWhitespace() {
	for isWhitespace(rune(Scanner.Char)) {
		Scanner.readChar()
	}
}

func (Scanner *ScannerStructureRegister) readChar() {
	if Scanner.RPOS >= len(Scanner.CharInput) {
		Scanner.Char = 0
	} else {
		Scanner.Char = Scanner.CharInput[Scanner.RPOS]
	}
	Scanner.POS = Scanner.RPOS
	Scanner.RPOS += 1
}

func (Scanner *ScannerStructureRegister) peekChar() byte {
	if Scanner.RPOS >= len(Scanner.CharInput) {
		return 0
	} else {
		return Scanner.CharInput[Scanner.RPOS]
	}
}

func (Scanner *ScannerStructureRegister) readIdentifier() string {
	position := Scanner.POS
	for isLetter(Scanner.Char) {
		Scanner.readChar()
	}
	return Scanner.CharInput[position:Scanner.POS]
}

func (Scanner *ScannerStructureRegister) readString() string {
	out := ""

	for {
		Scanner.readChar()
		if Scanner.Char == '"' {
			break
		}
		if Scanner.Char == '\\' {
			Scanner.readChar()

			if Scanner.Char == 'n' {
				Scanner.Char = '\n'
			}
			if Scanner.Char == 'r' {
				Scanner.Char = '\r'
			}
			if Scanner.Char == 't' {
				Scanner.Char = '\t'
			}
			if Scanner.Char == '"' {
				Scanner.Char = '"'
			}
			if Scanner.Char == '\\' {
				Scanner.Char = '\\'
			}
		}
		out = out + string(Scanner.Char)
	}

	return out
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenDataType, ch byte) TokenRegistry {
	return TokenRegistry{TokenDataType: tokenType, Literal: string(ch)}
}
