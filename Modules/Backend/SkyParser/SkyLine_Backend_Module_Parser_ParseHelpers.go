///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseHelpers
// Extension         | .go ( golang source code file )
// Purpose           | Defines all helper functions for the parser
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Module_Parser
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// The second major part of a programming language is parsing. Parsing does not necessarily execute the tokens but rather `parses` the tokens themselves and can pass them onto
//
// the evaluation step. In this step, we parse statements and expressions such as let, set, cause, engine, allow, call, etc.
//
//
//
//
package SkyLine_Backend_Module_Parser

import (
	SkyErr "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyErrorSystem"
	SkyFs "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyFS"
	SLTK "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens"

	"fmt"
)

// Error setup
var (
	Technology = "SkyLine Parser"
	Warn       = "W"
	Error      = "E"
)

func (SLP *SkyLine_Parser) GenErrorLine() string {
	CodeLine := SkyFs.Current.FileSystem_IndexLine(SLP.SL_Scanner.ReadLineNum() + 1) // Index the line of code
	Box := SkyFs.Current.DrawBoxWithinLineRange(CodeLine)
	return Box
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: Precedence functions and helper functions
//::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_Peek_Precedence() int {
	if prec, x := Precedences[SLP.SL_PeekToken.Token_Type]; x {
		return prec
	}
	return LOWEST
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_Current_Precedence() int {
	if prec, x := Precedences[SLP.SL_CurrentToken.Token_Type]; x {
		return prec
	}
	return LOWEST
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: Error functions
//::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_Ret_Errors() []string {
	return SLP.SL_Parser_Errors
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_Log_Errors(pref bool, TKTYPE SLTK.SL_TokenDataType) {
	var message string
	if pref {
		message += "No prefix parsing function found for %s"
	} else {
		message += "No infix parsing function found for %s"
	}
	message = fmt.Sprintf(message, TKTYPE)
	SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, message)
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: Peek, token and programatic parsing functions.
//::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

// Current Token Is What?
func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_CurrentTokenIs(TKTYPE SLTK.SL_TokenDataType) bool {
	return SLP.SL_CurrentToken.Token_Type == TKTYPE
}

// Peek Token Is What?
func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_PeekTokenCmp(TKTYPE SLTK.SL_TokenDataType) bool {
	return SLP.SL_PeekToken.Token_Type == TKTYPE
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_PeekTokenErr(TKTYPE SLTK.SL_TokenDataType) {
	Error2 := fmt.Sprintf(SkyErr.UseErrorMap(SkyErr.SkyLine_Parser_Unexpected_Token_Expected_Spec_Token), "< "+TKTYPE+" > ", SLP.SL_CurrentToken.Token_Type)
	tree := SkyErr.CreateStandardErrorTree(
		Error,
		Technology,
		fmt.Sprint(SkyErr.SkyLine_Parser_Unexpected_Token_Expected_Spec_Token),
		Error2,
		fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
		"",
		SLP.GenErrorLine(),
	)
	SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_LoadNextToken() {
	SLP.SL_PreviousToken = SLP.SL_CurrentToken // Call
	SLP.SL_CurrentToken = SLP.SL_PeekToken
	SLP.SL_PeekToken = SLP.SL_Scanner.NT()
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Helper_ExpectPeek(TKTYPE SLTK.SL_TokenDataType) bool {
	if SLP.SkyLine_Parser_Helper_PeekTokenCmp(TKTYPE) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		return true
	}
	SLP.SkyLine_Parser_Helper_PeekTokenErr(TKTYPE)
	return false
}
