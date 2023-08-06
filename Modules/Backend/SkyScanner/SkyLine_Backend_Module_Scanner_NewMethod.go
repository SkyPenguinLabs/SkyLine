///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_NewMethod
// Extension         | .go ( golang source code file )
// Purpose           | Defines a new method for the input scanner
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines a set of functions known as readers. The reader functions main jobs are to read specific values such as characters, read integers, read floats and data types
//
// like string values as well and data within strings. This can improve the performance of the language by not having to use constant conditional expressions under each time we
//
// want to parse a value or read a new character into the scanners stream.
package SkyLine_Backend_Scanner

func New(InputStream string) *SL_ScannerStructure {
	Scanner := &SL_ScannerStructure{Scanner_Characters: []rune(InputStream)}
	Scanner.RCHAR()
	return Scanner
}
