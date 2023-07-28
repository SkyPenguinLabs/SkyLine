///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_Extras
// Extension         | .go ( golang source code file )
// Purpose           | Deifnes all extra and necessary functions for the parser utilities
// Directory         | Modules/Backend/SkyEvaluator
// Modular Directory | SkyLine/Modules/Backend/SkyEvaluator
// Package Name      | SkyLine_Backend_Module_Evaluation
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
//
package SkyLine_Backend_Module_Parser

import (
	SLAST "SkyLine/Modules/Backend/SkyAST"
	SkyLine_Backend_Tokens "SkyLine/Modules/Backend/SkyTokens"
)

func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_And_Statements_ExtraUnit_Broken() SLAST.SL_Expression {
	return nil
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_And_Statements_ExtraUnit_IndexExpression(LeftUnit SLAST.SL_Expression) SLAST.SL_Expression {
	Idx := &SLAST.SL_EN_Index_Expression{
		TokenConstruct: SLP.SL_CurrentToken,
		Left:           LeftUnit,
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Idx.Index = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_RBRACKET) {
		return nil
	}
	return Idx
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse() *SLAST.SL_Prog {
	program := &SLAST.SL_Prog{}

	program.ProgramStatements = []SLAST.SL_Statement{}

	for SLP.SL_CurrentToken.Token_Type != SkyLine_Backend_Tokens.TOKEN_EOF {
		statement := SLP.SkyLine_Parser_Satements_Statement()
		if statement != nil {
			program.ProgramStatements = append(program.ProgramStatements, statement)
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return program
}
