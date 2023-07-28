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
// File Contains -> This file contains functions that can create new type structures to better categorize and represent a specific token that was scanned during the scanner
//                  This file also contains new method functions which return a scanner structure for the start of a new scanner.

package SkyLine_Configuration_Engine_Backend_Source

func New(input string) *ScannerStructureRegister {
	Scanner := &ScannerStructureRegister{CharInput: input}
	Scanner.readChar()
	return Scanner
}

func RegisterToken(TT TokenDataType, char byte) TokenRegistry {
	return TokenRegistry{
		TokenDataType: TT,
		Literal:       string(char),
	}
}
