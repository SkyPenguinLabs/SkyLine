///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_CreateTKConstruct
// Extension         | .go ( golang source code file )
// Purpose           | Define constant definitions for string values of Tokens
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines a new method to scan a new token or rather categorize a new token that the Scanner comes across
//
//
//
//
package SkyLine_Backend_Scanner

import SLTK "SkyLine/Modules/Backend/SkyTokens"

func ScanNewToken(TokenDataType SLTK.SL_TokenDataType, Character rune) SLTK.SL_TokenConstruct {
	return SLTK.SL_TokenConstruct{
		Token_Type: TokenDataType,
		Literal:    string(Character),
	}
}
