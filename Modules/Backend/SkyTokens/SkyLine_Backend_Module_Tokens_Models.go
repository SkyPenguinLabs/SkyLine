///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Token_KeywordModels
// Extension         | .go ( golang source code file )
// Purpose           | Defines all models / structures for thids module
// Directory         | Modules/Backend/SkyTokens
// Modular Directory | SkyLine/Modules/Backend/SkyTokens
// Package Name      | SkyLine_Backend_Tokens
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines two type's; one which is an alias to the string data type known as SL_TokenDataType ( The data type of a token such as '=') and a structure defining what consists
//
// of a token within the language. In this case it contains of a literal ( a string value of the keyword or representation '=') and the data type.
//
//
package SkyLine_Backend_Tokens

type SL_TokenDataType string

type SL_TokenConstruct struct {
	Token_Type SL_TokenDataType
	Literal    string
}
