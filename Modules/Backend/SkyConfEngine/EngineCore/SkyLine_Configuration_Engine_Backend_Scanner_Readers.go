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
// Contains -> This file contains all necessary functions within the class of readers. Reader functions are functions that reads and return data to the scanner.
//             Functions like this allow us to read characters, integers, strings, special lists and much more among that. These functions help in aiding the scanner.

package SkyLine_Configuration_Engine_Backend_Source

import "strings"

// Helps peek the next token
func (Scanner *ScannerStructureRegister) Peek() byte {
	if Scanner.RPOS >= len(Scanner.CharInput) {
		return 0
	}
	return Scanner.CharInput[Scanner.RPOS]
}

// Helps to read various data based on its input checking function
func (Scanner *ScannerStructureRegister) READ(VerificationFunc func(byte) bool) string {
	POS := Scanner.RPOS
	for VerificationFunc(Scanner.Char) {
		Scanner.readChar()
	}
	return Scanner.CharInput[POS:Scanner.RPOS]
}

//Helps to read the identifiers within the code
func (Scanner *ScannerStructureRegister) RIDENT() string {
	return Scanner.READ(VerifyLetterAlpha)
}

//Helps to read the digits within the code
func (Scanner *ScannerStructureRegister) RDIGIT() string {
	return Scanner.READ(VerifyDigit)
}

//Helps to read the integers being parsed
func (Scanner *ScannerStructureRegister) RINT() TokenRegistry {
	Integer := Scanner.RDIGIT()
	return TokenRegistry{
		TokenDataType: INTEGER_Token,
		Literal:       Integer,
	}
}

//Helps to read strings and return control characters
func (Scanner *ScannerStructureRegister) RSTR() string {
	breader := &strings.Builder{}
	for {
		Scanner.readChar()
		if Scanner.Char == 0 || Scanner.Char == '"' {
			break
		}
		if Scanner.Char == '\\' {
			Scanner.readChar()
			switch Scanner.Char {
			case 'n':
				breader.WriteByte('\n')
			case 'r':
				breader.WriteByte('\r')
			case 't':
				breader.WriteByte('\t')
			case 'f':
				breader.WriteByte('\f')
			case 'v':
				breader.WriteByte('\v')
			case '\\':
				breader.WriteByte('\\')
			case '"':
				breader.WriteByte('"')
			}
			Scanner.readChar()
			continue
		}
		breader.WriteByte(Scanner.Char)
	}
	return breader.String()
}

func (Scanner *ScannerStructureRegister) RNTCS() TokenRegistry {
	intPart := Scanner.readNumber()
	return TokenRegistry{
		TokenDataType: INTEGER_Token,
		Literal:       intPart,
	}
}

func (l *ScannerStructureRegister) readNumber() string {
	position := l.POS
	for VerifyDigit(l.Char) {
		l.readChar()
	}
	return l.CharInput[position:l.POS]
}
