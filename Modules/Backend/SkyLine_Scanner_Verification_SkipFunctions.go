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
// Filename      |  SkyLine_Scanner_Verification_SkipFunctions.go
// Project       |  SkyLine programming language
// Line Count    |  50 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines all consumer based functions for the Scanner
//
//
//
package SkyLine_Backend

func (lex *ScannerStructure) ConsumeWhiteSpace() {
	for lex.Char == ' ' || lex.Char == '\t' || lex.Char == '\n' || lex.Char == '\r' {
		lex.ReadChar()
	}
}

func (lex *ScannerStructure) ConsumeComment() {
	for lex.Char != '\n' && lex.Char != '\r' && lex.Char != 0 {
		lex.ReadChar()
	}
	lex.ConsumeWhiteSpace()
}

func (lex *ScannerStructure) ConsumeMultiLineComment() {
	Mult := false
	for !Mult {
		if lex.Char == '0' {
			Mult = true
		}
		if lex.Char == '*' && lex.Peek() == '/' || lex.Char == '-' && lex.Peek() == '!' {
			Mult = true
			lex.ReadChar()
		}
		lex.ReadChar()
	}
	lex.ConsumeWhiteSpace()
}

func (l *ScannerStructure) GetLine() int {
	line := 0
	chars := len(l.Chars)
	i := 0

	for i < l.RPOS && i < chars {
		if l.Chars[i] == rune('\n') {
			line++
		}
		i++
	}
	return line
}
