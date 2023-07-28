///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseTokenExpressions
// Extension         | .go ( golang source code file )
// Purpose           | Defines a function to parse bracket based tokens or expressions
// Directory         | Modules/Backend/SkyEvaluator
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEvaluator
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
	"fmt"

	SLAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyErr "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyErrorSystem"
	SkyLine_Backend_Tokens "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens"
)

func (SLP *SkyLine_Parser) SkyLine_Parser_Token_Parse_Brackets() SLAST.SL_Expression {
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LPAREN) {
		Error2 := fmt.Sprintf(SkyErr.UseErrorMap(SkyErr.SkyLine_Parser_Missing_RightFacing_Parenthesis_Token), SLP.SL_CurrentToken.Token_Type)
		tree := SkyErr.CreateStandardErrorTree(
			Error,
			Technology,
			fmt.Sprint(SkyErr.SkyLine_Parser_Missing_RightFacing_Parenthesis_Token),
			Error2,
			fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
			"",
			SLP.GenErrorLine(),
		)
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		return nil
	}

	SLP.SkyLine_Parser_Helper_LoadNextToken()

	Expr := SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)

	if Expr == nil {
		return nil
	}

	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_RPAREN) {
		Error2 := fmt.Sprintf(SkyErr.UseErrorMap(SkyErr.SkyLine_Parser_Missing_LeftFacing_Parenthesis_Token), SLP.SL_CurrentToken.Token_Type)
		tree := SkyErr.CreateStandardErrorTree(
			Error,
			Technology,
			fmt.Sprint(SkyErr.SkyLine_Parser_Missing_LeftFacing_Parenthesis_Token),
			Error2,
			fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
			"",
			SLP.GenErrorLine(),
		)
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		return nil
	}

	return Expr
}
