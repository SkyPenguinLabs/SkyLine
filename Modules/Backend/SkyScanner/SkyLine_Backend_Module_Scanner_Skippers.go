///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_Skippers
// Extension         | .go ( golang source code file )
// Purpose           | Define all skipper functions for the scanner / lexical analysis step of the interpreter.
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines something the developers of SkyLine like to call skipper functions. Skipper functions are functions that have a purpose to consume or skip over specific sets
//
// of data during lexican analysis. This kind of function will help skip over comments or white space. Most other people call these consumer functions or consume but we call them
//
// skippers because they skip and read the characters until the end of the information that needs to be skipped such as white space. We refuse to use consumer because consuming it would
//
// imply that it is "eating" or "erasing" the data and the information from the scanner / file which is not true. Skipper is a much more better term for this.
//
package SkyLine_Backend_Scanner

func (SL_Scanner *SL_ScannerStructure) CharacterPeek() rune {
	if SL_Scanner.Scanner_RPOS >= len(SL_Scanner.Scanner_Characters) {
		return rune(0)
	}
	return SL_Scanner.Scanner_Characters[SL_Scanner.Scanner_RPOS]
}

func (SL_Scanner *SL_ScannerStructure) SkipWhitespace() {
	for CharacterIsWhiteSpace(SL_Scanner.Scanner_Character) {
		SL_Scanner.RCHAR()
	}
}

func (SL_Scanner *SL_ScannerStructure) SkipSingleLineComment() {
	for SL_Scanner.Scanner_Character != '\n' && SL_Scanner.Scanner_Character != 0 {
		SL_Scanner.RCHAR()
	}
	SL_Scanner.SkipWhitespace()
}

func (SL_Scanner *SL_ScannerStructure) SkipMultiLineComment() {
	Mult := false
	for !Mult {
		if SL_Scanner.Scanner_Character == 0 {
			Mult = true
		}
		if SL_Scanner.Scanner_Character == '*' && SL_Scanner.CharacterPeek() == '/' {
			Mult = true
			SL_Scanner.RCHAR()
		}
		SL_Scanner.RCHAR()
	}
	SL_Scanner.SkipWhitespace()
}
