///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseUnits
// Extension         | .go ( golang source code file )
// Purpose           | Defines functions dedicated to implementing parser rules for units and code blocks
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
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
package SkyLine_Backend_Module_Parser

import (
	SLAST "SkyLine/Modules/Backend/SkyAST"
	SkyLine_Backend_Tokens "SkyLine/Modules/Backend/SkyTokens"
)

func (SLP *SkyLine_Parser) SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions() *SLAST.SL_UnitBlockStatement {
	UNIT := &SLAST.SL_UnitBlockStatement{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	UNIT.Statements = []SLAST.SL_Statement{}
	SLP.SkyLine_Parser_Helper_LoadNextToken()

	for !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_RBRACE) && !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_EOF) {
		Statements := SLP.SkyLine_Parser_Satements_Statement()
		if Statements != nil {
			UNIT.Statements = append(UNIT.Statements, Statements)
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return UNIT
}
