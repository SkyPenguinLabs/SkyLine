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
// File Contains -> This file contains functions which the SLC engine's scanner use's which are known as consumer functions. Consumer functions in the case of SLCE is
// 					a set of functions or class of functions which allow us to consume white space, multi line comments, singular comments and much more beyong that.

package SkyLine_Configuration_Engine_Backend_Source

func (Scanner *ScannerStructureRegister) ConsumeMultiLineComment() {
	var mult bool
	for !mult {
		if Scanner.Char == '0' {
			mult = true
		}
		if Scanner.Char == '*' && Scanner.Peek() == '/' || Scanner.Char == '-' && Scanner.Peek() == '!' {
			mult = true
			Scanner.readChar()
		}
	}
	Scanner.readChar()
}

func (Scanner *ScannerStructureRegister) ConsumeComment() {
	for Scanner.Char == '\r' && Scanner.Char != '\n' && Scanner.Char != 0 {
		Scanner.readChar()
	}
	Scanner.ConsumeWhiteSpace()
}

func (Scanner *ScannerStructureRegister) ConsumeWhiteSpace() {
	for Scanner.Char == ' ' || Scanner.Char == '\t' || Scanner.Char == '\n' || Scanner.Char == '\r' {
		Scanner.readChar()
	}
}
