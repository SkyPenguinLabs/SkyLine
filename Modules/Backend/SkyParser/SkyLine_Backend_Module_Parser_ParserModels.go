///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParserModels
// Extension         | .go ( golang source code file )
// Purpose           | Defines all the models for the parser such as structures, types and more
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
//s
package SkyLine_Backend_Module_Parser

import (
	SLAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SLSCANNER "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner"
	SLTK "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens"
)

const (
	_ int = iota
	LOWEST
	ASSIGNMENT
	EQUALS
	LESSGREATER
	BitwiseOR    // |
	BitwiseXOR   // ^
	BitwiseAND   // &
	BitwiseShift // << or >>
	SUM          // + or -
	PRODUCT      // * / or %
	PREFIX       // -Y or !Y
	CALL         // F(x...)
	INDEX        // ARR OF TYPE ARRAY[...]
	GTEQS        //
	LTEQS
	POWER
	MOD
	DOT_DOT
	REGEXP_MATCH
	TERNARY
	COND
)

var Precedences = map[SLTK.SL_TokenDataType]int{
	SLTK.TOKEN_ANDAND:            COND,
	SLTK.TOKEN_OROR:              COND,
	SLTK.TOKEN_EQ:                EQUALS,
	SLTK.TOKEN_GTEQ:              GTEQS,
	SLTK.TOKEN_LTEQ:              LTEQS,
	SLTK.TOKEN_NEQ:               EQUALS,
	SLTK.TOKEN_LT:                LESSGREATER,
	SLTK.TOKEN_GT:                LESSGREATER,
	SLTK.TOKEN_PLUS:              SUM,
	SLTK.TOKEN_PLUS_EQUALS:       SUM,
	SLTK.TOKEN_MINUS:             SUM,
	SLTK.TOKEN_MINUS_EQUALS:      SUM,
	SLTK.TOKEN_SLASH:             PRODUCT,
	SLTK.TOKEN_ASSIGN:            ASSIGNMENT,
	SLTK.TOKEN_POWEROF:           POWER,
	SLTK.TOKEN_QUESTION:          TERNARY,
	SLTK.TOKEN_ASTARISK:          PRODUCT,
	SLTK.TOKEN_ASTERISK_EQUALS:   PRODUCT,
	SLTK.TOKEN_DIVEQ:             PRODUCT,
	SLTK.TOKEN_LPAREN:            CALL,
	SLTK.TOKEN_PERIOD:            CALL,
	SLTK.TOKEN_NOTCONTAIN:        REGEXP_MATCH,
	SLTK.TOKEN_CONTAINS:          REGEXP_MATCH,
	SLTK.TOKEN_DOTDOT:            DOT_DOT,
	SLTK.TOKEN_LBRACKET:          INDEX,
	SLTK.TOKEN_MODULECALL:        INDEX,
	SLTK.TOKEN_MODULO:            MOD,
	SLTK.TOKEN_BITWISE_OP_AND:    BitwiseAND,
	SLTK.TOKEN_BITWISE_OP_OR:     BitwiseOR,
	SLTK.TOKEN_BITWISE_OP_LSHIFT: BitwiseShift,
	SLTK.TOKEN_BITWISE_OP_RSHIFT: BitwiseShift,
	SLTK.TOKEN_BITWISE_OP_XOR:    BitwiseXOR,
}

type (
	//::::::::::::::::::::::::::::::::
	//:: Function list
	//::::::::::::::::::::::::::::::::
	SkyLine_Parser_PrefixParser_Function  func() SLAST.SL_Expression
	SkyLine_Parser_InfixParser_Function   func(SLAST.SL_Expression) SLAST.SL_Expression
	SkyLine_Parser_PostFixParser_Function func() SLAST.SL_Expression

	//:::::::::::::::::::::::::::::::::
	//:: Modules dependant structures
	//:::::::::::::::::::::::::::::::::

	SkyLine_Parser struct {
		SL_Scanner                  *SLSCANNER.SL_ScannerStructure                                  // Scanner / lexical analysis structure
		SL_PreviousToken            SLTK.SL_TokenConstruct                                          // Parsers previously parsed token
		SL_CurrentToken             SLTK.SL_TokenConstruct                                          // Parsers current token being parsed
		SL_PeekToken                SLTK.SL_TokenConstruct                                          // Parsers next token to parse or the peeked token
		SL_Prefix_Parser_Functions  map[SLTK.SL_TokenDataType]SkyLine_Parser_PrefixParser_Function  // Parsers prefix parsing functions
		SL_Infix_Parser_Functions   map[SLTK.SL_TokenDataType]SkyLine_Parser_InfixParser_Function   // Parsers infix parsing functions
		SL_PostFix_Parser_Functions map[SLTK.SL_TokenDataType]SkyLine_Parser_PostFixParser_Function // Parsers postfix parsing functions list
		SL_Parser_Errors            []string                                                        // Parsers list of errors
	}
)
