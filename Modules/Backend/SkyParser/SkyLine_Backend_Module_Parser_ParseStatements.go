///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseStatements
// Extension         | .go ( golang source code file )
// Purpose           | Defines all the statement parser functions for the language's parser
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
//
//
//
package SkyLine_Backend_Module_Parser

import (
	SLAST "SkyLine/Modules/Backend/SkyAST"
	SkyErr "SkyLine/Modules/Backend/SkyErrorSystem"
	SkyLine_Backend_Tokens "SkyLine/Modules/Backend/SkyTokens"
	"fmt"
)

func (SLP *SkyLine_Parser) SkyLine_Parser_Satements_Statement() SLAST.SL_Statement {
	switch SLP.SL_CurrentToken.Token_Type {
	case SkyLine_Backend_Tokens.TOKEN_KEYWORD_ENGINE:
		return SLP.SkyLine_Parser_Statements_KeywordENGINE()
	case SkyLine_Backend_Tokens.TOKEN_RETURN:
		return SLP.SkyLine_Parser_Statements_KeywordReturn()
	case SkyLine_Backend_Tokens.TOKEN_CONSTANT:
		return SLP.SkyLine_Parser_Statements_KeywordConst()
	case SkyLine_Backend_Tokens.TOKEN_ALLOW:
		return SLP.SkyLine_Parser_Statements_KeywordSetLetAllowCause()
	case SkyLine_Backend_Tokens.TOKEN_REGISTER:
		return SLP.SkyLine_Parser_Statements_KeywordRegister()
	default:
		return SLP.SkyLine_Parser_Statements_ExpressionStatement()
	}
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Statements_ExpressionStatement will parse an expression statement
//::
//:: that was declared.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Statements_ExpressionStatement() *SLAST.Expression_Statement {
	Statement := &SLAST.Expression_Statement{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	Statement.Expression = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	for SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return Statement
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Statements_KeywordReturn will parse any statement that holds a return
//::
//:: keyword that was found by the scanner
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Statements_KeywordReturn() *SLAST.Return_Ret_Return_Information {
	Return := &SLAST.Return_Ret_Return_Information{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Return.Expression = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	for !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_EOF) {
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, "Unterminated return statement (ret|return) ")
			return nil
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return Return
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Statements_KeywordConst will parse any statements that are constants
//::
//:: which means it uses constant or const keywords.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Statements_KeywordConst() *SLAST.Assignment_Constant_Const {
	Constant := &SLAST.Assignment_Constant_Const{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_IDENT) {
		return nil
	}
	Constant.Name = &SLAST.SL_Identifier{
		TokenConstruct: SLP.SL_CurrentToken,
		Value:          SLP.SL_CurrentToken.Literal,
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_ASSIGN) {
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Constant.Value = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	for !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_EOF) {
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, "unterminated assignment and variable creation (constant|const)")
			return nil
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return Constant
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Statements_KeywordSetLetAllowCause will parse any statement that
//::
//:: was found using let, set, cause or allow keywords which indicate variable creation.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Statements_KeywordSetLetAllowCause() *SLAST.Assignment_Cause_Set_Allow {
	Allow := &SLAST.Assignment_Cause_Set_Allow{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_IDENT) {
		return nil
	}
	Allow.Name = &SLAST.SL_Identifier{
		TokenConstruct: SLP.SL_CurrentToken,
		Value:          SLP.SL_CurrentToken.Literal,
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_PEEKASSIGN) && !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_ASSIGN) {
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Allow.Value = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	for !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_EOF) {
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, "unterminated creation or variable assignment (let|set|cause|allow)")
			return nil
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return Allow
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Statements_KeywordRegister will parse any statement that
//::
//:: was found involving the registry statement.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Statements_KeywordRegister() *SLAST.SL_Register {
	statement := &SLAST.SL_Register{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if !SLP.SkyLine_Parser_Helper_PeekTokenCmp("(") {
		fmt.Println("Missing ( in import statement")
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	statement.RegistryValue = SLP.SkyLine_Parser_Functions_Parse_ExpressionList(")")
	for SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return statement
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Statements_KeywordENGINE will parse engine statements. For context
//::
//:: engine statements defined as `ENGINE(...)` are statements that were developed to
//::
//:: work internally with the SkyLine configuration engine to load projects, start projects,
//::
//:: modify SkyLine's internal environment and structure while also allowing for ease of use
//::
//:: and easy library registration.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Statements_KeywordENGINE() *SLAST.SL_ENGINE {
	ENGINE := &SLAST.SL_ENGINE{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if !SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_LPAREN) {
		msg := SkyErr.ErrorMap[SkyErr.SkyLine_Parser_Missing_RightFacing_Parenthesis_Token]
		tree := SkyErr.CreateStandardErrorTree(
			Error,
			Technology,
			fmt.Sprint(SkyErr.SkyLine_Parser_Missing_RightFacing_Parenthesis_Token),
			msg,
			fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
			"",
			SLP.GenErrorLine(),
		)
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	ENGINE.EngineValue = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	for SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return ENGINE
}
