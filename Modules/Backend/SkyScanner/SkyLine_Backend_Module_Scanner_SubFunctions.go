///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_SubFunctions
// Extension         | .go ( golang source code file )
// Purpose           | Defines a function for string characters
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

func CheckEscapedCharacter(QUOTE, PEEK rune) string {
	if QUOTE == '"' {
		switch PEEK {
		case 'n':
			return "\n"
		case 't':
			return "\t"
		case 'v':
			return "\v"
		case 'f':
			return "\f"
		case 'r':
			return "\r"
		case '\\':
			return "\\"
		case '"':
			return "\""
		case '\'':
			return "'"
		default:
			return "\\" + string(PEEK)
		}
	}
	switch PEEK {
	case '"':
		return "\\\""
	case '\'':
		return "'"
	default:
		return "\\" + string(PEEK)
	}
}
