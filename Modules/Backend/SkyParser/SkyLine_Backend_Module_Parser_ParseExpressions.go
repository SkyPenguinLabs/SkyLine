///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseExpressions
// Extension         | .go ( golang source code file )
// Purpose           | Defines all expression parsing functions for the language
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
//
package SkyLine_Backend_Module_Parser

import (
	"fmt"

	SLAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyErr "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyErrorSystem"
	SkyLine_Backend_Tokens "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens"
)

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Postfix is a function to parse post fix expressions
//::
//:: that are found by the parsers main function and type finder.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_Postfix() SLAST.SL_Expression {
	EXPRESSION := &SLAST.SL_EN_Postfix{
		TokenConstruct: SLP.SL_PreviousToken,
		Operator:       SLP.SL_CurrentToken.Literal,
	}
	return EXPRESSION
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Infix is a function to parse infix expressions
//::
//:: that are found by the parsers main function and type finder.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_Infix(LOP SLAST.SL_Expression) SLAST.SL_Expression {
	EXPRESSION := &SLAST.SL_EN_Infix{
		TokenConstruct: SLP.SL_CurrentToken,
		Operator:       SLP.SL_CurrentToken.Literal,
		Left:           LOP,
	}
	PREC := SLP.SkyLine_Parser_Helper_Current_Precedence()
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	EXPRESSION.Right = SLP.SkyLine_Parser_Expressions_Parse_Expression(PREC)
	return EXPRESSION
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Prefix is a function to parse infix expressions
//::
//:: that are found by the parsers main function and type finder.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_Prefix() SLAST.SL_Expression {
	EXPRESION := &SLAST.SL_EN_Prefix{
		TokenConstruct: SLP.SL_CurrentToken,
		Operator:       SLP.SL_CurrentToken.Literal,
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	EXPRESION.Right = SLP.SkyLine_Parser_Expressions_Parse_Expression(PREFIX)
	return EXPRESION
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_GroupedExpr is a that parses grouped expressions
//::
//:: that are found by the parsers main function and type finder.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_GroupedExpr() SLAST.SL_Expression {
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	EXPRESSION := SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_RPAREN) {
		return nil
	}
	return EXPRESSION
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Expression will parse a given expression and
//::
//:: return the operation
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_Expression(Prec int) SLAST.SL_Expression {
	Postfix := SLP.SL_PostFix_Parser_Functions[SLP.SL_CurrentToken.Token_Type]
	if Postfix != nil {
		return (Postfix())
	}
	Prefix := SLP.SL_Prefix_Parser_Functions[SLP.SL_CurrentToken.Token_Type]
	if Prefix == nil {
		SLP.SkyLine_Parser_Helper_Log_Errors(true, SLP.SL_CurrentToken.Token_Type)
		return nil
	}
	Left := Prefix()
	for !SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) && Prec < SLP.SkyLine_Parser_Helper_Peek_Precedence() {
		Infix := SLP.SL_Infix_Parser_Functions[SLP.SL_PeekToken.Token_Type]
		if Infix == nil {
			return Left
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		Left = Infix(Left)
	}
	return Left
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Conditional is a function to parse conditonal if
//::
//:: else, else if statements that are picked up by the scanner
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_Conditional() SLAST.SL_Expression {
	Conditional := &SLAST.SL_EN_Conditional_IfElse{
		TokenConstruct: SLP.SL_CurrentToken,
	}

	if Conditional == nil {
		msg := fmt.Sprintf(SkyErr.ErrorMap[SkyErr.SkyLine_Parser_Unexpected_Null_Conditional_Unit], " < ' "+SLP.SL_CurrentToken.Literal+" ' >")
		tree := SkyErr.CreateStandardErrorTree(
			Error,
			Technology,
			fmt.Sprint(SkyErr.SkyLine_Parser_Unexpected_Null_Conditional_Unit),
			msg,
			fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
			"",
			SLP.GenErrorLine(),
		)
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		return nil
	}

	Conditional.Condition = SLP.SkyLine_Parser_Token_Parse_Brackets()
	if Conditional.Condition == nil {
		return nil
	}

	// Parse unit
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
		msg := fmt.Sprintf(SkyErr.ErrorMap[SkyErr.SkyLine_Parser_Missing_RightFacing_Curely_Brace_Token], " < ' "+SLP.SL_CurrentToken.Literal+" ' >")
		tree := SkyErr.CreateStandardErrorTree(
			Error,
			Technology,
			fmt.Sprint(SkyErr.SkyLine_Parser_Missing_RightFacing_Curely_Brace_Token),
			msg,
			fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
			"",
			SLP.GenErrorLine(),
		)
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		return nil
	}

	// The consequence
	Conditional.Consequence_Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
	if Conditional.Consequence_Unit == nil {
		msg := SkyErr.ErrorMap[SkyErr.SkyLine_Parser_Unexpected_Null_Conditional_Unit]
		tree := SkyErr.CreateStandardErrorTree(
			Error,
			Technology,
			fmt.Sprint(SkyErr.SkyLine_Parser_Unexpected_Null_Conditional_Unit),
			msg,
			fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
			"",
			SLP.GenErrorLine(),
		)
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		return nil
	}

	// Else?

	if SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_ELSE) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()

		if SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_IF) {
			SLP.SkyLine_Parser_Helper_LoadNextToken()
			Conditional.Alternative_Unit = &SLAST.SL_UnitBlockStatement{
				Statements: []SLAST.SL_Statement{
					&SLAST.Expression_Statement{
						Expression: SLP.SkyLine_Parser_Expressions_Parse_Conditional(),
					},
				},
			}
			return Conditional
		}

		if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
			msg := fmt.Sprintf(SkyErr.ErrorMap[SkyErr.SkyLine_Parser_Missing_RightFacing_Curely_Brace_Token], " < ' "+SLP.SL_CurrentToken.Literal+" ' >")
			tree := SkyErr.CreateStandardErrorTree(
				Error,
				Technology,
				fmt.Sprint(SkyErr.SkyLine_Parser_Missing_RightFacing_Curely_Brace_Token),
				msg,
				fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
				"",
				SLP.GenErrorLine(),
			)
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
			return nil
		}

		Conditional.Alternative_Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
		if Conditional.Alternative_Unit == nil {
			msg := SkyErr.ErrorMap[SkyErr.SkyLine_Parser_Unexpected_Null_Conditional_Unit]
			tree := SkyErr.CreateStandardErrorTree(
				Error,
				Technology,
				fmt.Sprint(SkyErr.SkyLine_Parser_Unexpected_Null_Conditional_Unit),
				msg,
				fmt.Sprint(SLP.SL_Scanner.ReadLineNum()+1),
				"",
				SLP.GenErrorLine(),
			)
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, tree)
		}
	}
	return Conditional
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_ConditionalSwitch is a function to parse conditonal
//::
//:: switch case statements. In SkyLine, switch and case have a default. Switch and case
//::
//:: can also be defined with sw (switch), df (default), cs (case)
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_ConditionalSwitch() SLAST.SL_Expression {
	SwitchCaseExpr := &SLAST.SL_EN_Switch_ExpressionStatement{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	SwitchCaseExpr.Value = SLP.SkyLine_Parser_Token_Parse_Brackets()
	if SwitchCaseExpr.Value == nil {
		return nil
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	for !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_RBRACE) {
		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_EOF) {
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, "unterminated switch statement")
			return nil
		}

		tmp := &SLAST.SL_EN_Case_ExpressionStatement{
			TokenConstruct: SLP.SL_CurrentToken,
		}

		if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_DEFAULT) {
			tmp.Default = true
		} else if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_CASE) {
			SLP.SkyLine_Parser_Helper_LoadNextToken()
			if SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_DEFAULT) {
				tmp.Default = true
			} else {
				tmp.Expression = append(tmp.Expression, SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST))
				for SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_COMMA) {
					SLP.SkyLine_Parser_Helper_LoadNextToken()
					SLP.SkyLine_Parser_Helper_LoadNextToken()
					tmp.Expression = append(tmp.Expression, SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST))
				}
			}
		} else {
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, fmt.Sprintf("Expected case|default got %s", SLP.SL_CurrentToken.Literal))
			return nil
		}
		if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
			msg := fmt.Sprintf("Expected nect token to be a left face brace ( '{' ) vyt got %s", SLP.SL_CurrentToken.Literal)
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, msg)
			fmt.Printf("error\n")
			return nil
		}
		tmp.Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
		if !SLP.SkyLine_Parser_Helper_CurrentTokenIs(SkyLine_Backend_Tokens.TOKEN_RBRACE) {
			msg := fmt.Sprintf("Expected next token to be right face brace ( '}' ) but you gave %s instead", SLP.SL_CurrentToken.Literal)
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, msg)
			fmt.Printf("error\n")
			return nil
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		SwitchCaseExpr.Conditions = append(SwitchCaseExpr.Conditions, tmp)
	}
	count := 0
	for _, c := range SwitchCaseExpr.Conditions {
		if c.Default {
			count++
		}
	}
	if count > 1 {
		msg := "A switch-statement should only have one default block"
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, msg)
		return nil

	}
	return SwitchCaseExpr
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_ConditionalLoop is a function to parse conditonal
//::
//:: for loop statements. This function will basically parse a for loop with a condition
//::
//:: and since for loops are based on truthy values then we call it a conditional loopc
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_ConditionalLoop() SLAST.SL_Expression {
	ConditionalLoop := &SLAST.SL_EN_Conditional_Loop{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LPAREN) {
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	ConditionalLoop.Condition = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_RPAREN) {
		return nil
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_LBRACE) {
		return nil
	}
	ConditionalLoop.Consequence = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
	return ConditionalLoop
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_WithinRangeLoop is a function to parse standard
//::
//:: range based foreach loops. This is basically saying for each value within range of
//::
//:: x then do n which can be quite simple to parse and implement.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_WithinRangeLoop() SLAST.SL_Expression {
	Range := &SLAST.SL_EN_For_Each_Loop{
		TokenConstruct: SLP.SL_CurrentToken,
	}

	SLP.SkyLine_Parser_Helper_LoadNextToken()

	Range.Identifier = SLP.SL_CurrentToken.Literal

	if SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_COMMA) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		if !SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_IDENT) {
			SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, fmt.Sprintf("Second argument to foreach must be an identifier, got %v", SLP.SL_PeekToken))
			return nil
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		Range.Index = Range.Identifier
		Range.Identifier = SLP.SL_CurrentToken.Literal
	}

	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_INSIDE) {
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Range.Value = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	if Range.Value == nil {
		return nil
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	Range.Unit = SLP.SkyLine_Backend_Module_Parser_ParseUnitBlockExpressions()
	return Range
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword is a function designed to parse
//::
//:: basic variable assignment. This allows variables to be changed once assigned
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword(VarName SLAST.SL_Expression) SLAST.SL_Expression {
	Statement := &SLAST.SL_EN_VariableAssignmentStatement{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	if Note, ok := VarName.(*SLAST.SL_Identifier); ok {
		Statement.Name = Note
	} else {
		msg := fmt.Sprintf("Expected assign token to be IDENT, got %s instead", VarName.SkyLine_NodeInterface_Token_Literal())
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, msg)
	}
	OPERATION := SLP.SL_CurrentToken
	SLP.SkyLine_Parser_Helper_LoadNextToken()

	switch OPERATION.Token_Type {
	case SkyLine_Backend_Tokens.TOKEN_PLUS_EQUALS:
		Statement.Operator = "+="
	case SkyLine_Backend_Tokens.TOKEN_MINUS_EQUALS:
		Statement.Operator = "-="
	case SkyLine_Backend_Tokens.TOKEN_DIVEQ:
		Statement.Operator = "/="
	case SkyLine_Backend_Tokens.TOKEN_ASTERISK_EQUALS:
		Statement.Operator = "*="
	default:
		Statement.Operator = "="
	}
	Statement.Value = SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
	return Statement
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Importing is a function to parse standard
//::
//:: import expression to import specific files or modules within the code.
//::
//:: import(x, y, x) -> foreach _ x { eval_module }
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Expressions_Parse_Importing() SLAST.SL_Expression {
	expression := &SLAST.SL_ImportExpression{TokenConstruct: SLP.SL_CurrentToken}
	if !SLP.SkyLine_Parser_Helper_PeekTokenCmp("(") {
		fmt.Println("Missing ( in import statement")
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	expression.Name = SLP.SkyLine_Parser_Functions_Parse_ExpressionList(")")
	for SLP.SkyLine_Parser_Helper_PeekTokenCmp(SkyLine_Backend_Tokens.TOKEN_SEMICOLON) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()
	}
	return expression
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//::
//:: SkyLine_Parser_Expressions_Parse_Selector is a function to parse standard
//::
//:: module calls which are results from importing files or modules.
//::
//:: x = import(...) -> x::Object...()=>[]=>{}.
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (p *SkyLine_Parser) SkyLine_SelectorExpression(exp SLAST.SL_Expression) SLAST.SL_Expression {
	p.SkyLine_Parser_Helper_ExpectPeek(SkyLine_Backend_Tokens.TOKEN_IDENT)
	index := &SLAST.SL_String{
		TokenConstruct: p.SL_CurrentToken,
		Value:          p.SL_CurrentToken.Literal,
	}
	return &SLAST.SL_EN_Index_Expression{
		Left:  exp,
		Index: index,
	}
}
