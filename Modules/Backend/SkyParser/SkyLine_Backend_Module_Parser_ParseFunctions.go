///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseFunctions
// Extension         | .go ( golang source code file )
// Purpose           | Defines all the parser functions for function calls, function declarations and even object call functions or invokes
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
	SkyErr "SkyLine/Modules/Backend/SkyErrorSystem"
	SkyLine_Backend_Tokens "SkyLine/Modules/Backend/SkyTokens"
	"fmt"
)

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Functions_Parse_ExpressionList is a function to parse expression list s
//::
//:: or elements within a given end. This will parse until the end of a given list or character
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Functions_Parse_ExpressionList(EndExpect SkyLine_Backend_Tokens.SL_TokenDataType) []SLAST.SL_Expression {
	Expressions := make([]SLAST.SL_Expression, 0)

	if SLP.SkyLine_Parser_Helper_PeekTokenCmp(EndExpect) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		return Expressions
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()

	Expressions = append(Expressions, SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST))

	for SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_COMMA) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		Expressions = append(Expressions, SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST))
	}

	if !SLP.SkyLine_Parser_Helper_ExpectPeek(EndExpect) {
		return nil
	}

	return Expressions
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Functions_Parse_FunctionObjectCall is a function that parses a object call
//::
//:: or invoke which works by calling object.call().
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Functions_Parse_FunctionObjectCall(ObjectC SLAST.SL_Expression) SLAST.SL_Expression {
	Method := &SLAST.SL_EN_Object_Call_Expression{
		TokenConstruct: SLP.SL_CurrentToken,
		Object:         ObjectC,
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Name := SLP.SkyLine_Parser_Function_DataType_Identifier()
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Method.Call = SLP.SkyLine_Parser_Functions_Parse_FunctionCall(Name)
	return Method
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Functions_Parse_FunctionCall is a function that parses a function-call expression.
//::
//:: This means that a developer has called a function with `OBJECT->Function(...)``
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Functions_Parse_FunctionCall(Function SLAST.SL_Expression) SLAST.SL_Expression {
	Expression := &SLAST.SL_EN_Call_Expression{
		TokenConstruct: SLP.SL_CurrentToken,
		Function:       Function,
	}

	Expression.Arguments = SLP.SkyLine_Parser_Functions_Parse_ExpressionList(SkyLine_Backend_Tokens.TOKEN_RPAREN)
	return Expression
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Functions_Parse_Function_ArgumentsParams is a function that parses
//::
//:: a functions arguments and implements default arguments
//::
//:: - Demo of default arguments
//::
//:: func Result(x=10, x) {...};
//::
//:: Note: x=10 is the default argument
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// SkyLine_Parser_Functions_Parse_FunctionCall parses a function-call expression.
func (SLP *SkyLine_Parser) SkyLine_Parser_Functions_Parse_Function_ArgumentsParams() (map[string]SLAST.SL_Expression, []*SLAST.SL_Identifier) {
	Defaults := make(map[string]SLAST.SL_Expression)
	Identifiers := make([]*SLAST.SL_Identifier, 0)

	if SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_RPAREN) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		return Defaults, Identifiers
	}

	SLP.SkyLine_Parser_Helper_LoadNextToken()

	for !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_RPAREN) {

		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_EOF) {
			msg := SkyErr.ErrorMap[SkyErr.SkyLine_Parser_UUnterminated_Function_Parameter]
			tree := SkyErr.CreateStandardErrorTree(
				Error,
				Technology,
				fmt.Sprint(SkyErr.SkyLine_Parser_UUnterminated_Function_Parameter),
				msg,
				fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
				"Definitions should be like this -> define Name(){};",
				SLP.GenErrorLine(),
			)
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
			return nil, nil
		}

		Identifier := &SLAST.SL_Identifier{
			TokenConstruct: SLP.SL_CurrentToken,
			Value:          SLP.SL_CurrentToken.Literal,
		}

		Identifiers = append(Identifiers, Identifier)

		SLP.SkyLine_Parser_Helper_LoadNextToken()

		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_ASSIGN) {
			SLP.SkyLine_Parser_Helper_LoadNextToken()
			Defaults[Identifier.Value] = SLP.SkyLine_Parser_Statements_ExpressionStatement().Expression
			SLP.SkyLine_Parser_Helper_LoadNextToken()
		}

		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_COMMA) {
			SLP.SkyLine_Parser_Helper_LoadNextToken()
		}
	}

	return Defaults, Identifiers
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Lables_Parse_Lable_Definition will parse any lables kind of like your
//::
//:: Goto statements or your `@lable:` in ASM. This just makes it easier during loops
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Lables_Parse_Lable_Definition() SLAST.SL_Expression {
	SLP.SkyLine_Parser_Helper_LoadNextToken()

	Def := &SLAST.SL_EN_Lable_JumpDef{
		TokenConstruct: SLP.SL_CurrentToken,
	}

	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
		return nil
	}
	Def.Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
	return Def
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Functions_Parse_Function_Definition is a function that parses a .
//::
//:: functions definition / decleration using define or func keywords.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Functions_Parse_Function_Definition() SLAST.SL_Expression {
	SLP.SkyLine_Parser_Helper_LoadNextToken()

	Definition := &SLAST.SL_EN_Function_Definition{
		TokenConstruct: SLP.SL_CurrentToken,
	}

	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LPAREN) {
		return nil
	}

	Definition.Defaults, Definition.FunctionArguments = SLP.SkyLine_Parser_Functions_Parse_Function_ArgumentsParams()
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
		return nil
	}

	Definition.Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
	return Definition
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Functions_Parse_Function_LiteralToken is a function that parses a
//::
//:: function that is assigned via variable like `cause x := func(...){...};`
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// SkyLine_Parser_Functions_Parse_Function_LiteralToken will parse any function or rather compacted function
func (SLP *SkyLine_Parser) SkyLine_Parser_Functions_Parse_Function_LiteralToken() SLAST.SL_Expression {
	Function := &SLAST.SL_EN_Function_Literal{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LPAREN) {
		return nil
	}

	Function.Defaults, Function.Parameters = SLP.SkyLine_Parser_Functions_Parse_Function_ArgumentsParams()
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
		return nil
	}
	Function.Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
	return Function
}
