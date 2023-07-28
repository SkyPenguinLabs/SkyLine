///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_ParseMainFunction
// Extension         | .go ( golang source code file )
// Purpose           | Defines the main registry function and new function for the parser
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
	SLScanner "SkyLine/Modules/Backend/SkyScanner"
	SLTK "SkyLine/Modules/Backend/SkyTokens"
)

func SkyLineNewParser(Scanner *SLScanner.SL_ScannerStructure) *SkyLine_Parser {
	SLP := &SkyLine_Parser{
		SL_Scanner:       Scanner,
		SL_Parser_Errors: []string{},
	}

	SLP.SkyLine_Parser_Helper_LoadNextToken()
	SLP.SkyLine_Parser_Helper_LoadNextToken()

	//:::::::::::::::::::::::::::::::::
	//:: Map of prefix parser functions
	//:::::::::::::::::::::::::::::::::
	SLP.SL_Prefix_Parser_Functions = map[SLTK.SL_TokenDataType]SkyLine_Parser_PrefixParser_Function{
		SLTK.TOKEN_BANG:                    SLP.SkyLine_Parser_Expressions_Parse_Prefix,                    // Parser Call Function (Prefix) | Parse (!)
		SLTK.TOKEN_FUNCTION_DEFINE_LITERAL: SLP.SkyLine_Parser_Functions_Parse_Function_Definition,         // Parser Call Function (Prefix) | Parse (define, func)
		SLTK.TOKEN_FUNCTION:                SLP.SkyLine_Parser_Functions_Parse_Function_LiteralToken,       // Parser Call Function (Prefix) | Parse (Func, function)
		SLTK.TOKEN_Lable:                   SLP.SkyLine_Parser_Lables_Parse_Lable_Definition,               // Parser Call Function (Prefix) | Parse (@lable)
		SLTK.TOKEN_IDENT:                   SLP.SkyLine_Parser_Function_DataType_Identifier,                // Parser Call Function (Prefix) | Parse (Identifier)
		SLTK.TOKEN_EOF:                     SLP.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_Broken, // Parser Call Function (Prefix) | Parse (EOF)
		SLTK.TOKEN_ILLEGAL:                 SLP.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_Broken, // Parser Call Function (Prefix) | Parse (Illegal)
		SLTK.TOKEN_BYTESTART:               SLP.SkyLine_Parser_Function_DataType_Byte,                      // Parser Call Function (Prefix) | Parse (Byte)
		SLTK.TOKEN_STRING:                  SLP.SkyLine_Parser_Function_DataType_String,                    // Parser Call Function (Prefix) | Parse (STRING)
		SLTK.TOKEN_TRUE:                    SLP.SkyLine_Parser_Function_DataType_Boolean,                   // Parser Call Function (Prefix) | Parse (true, BOOLEANT)
		SLTK.TOKEN_FALSE:                   SLP.SkyLine_Parser_Function_DataType_Boolean,                   // Parser Call Function (Prefix) | Parse (false, BOOLEANF)
		SLTK.TOKEN_INT:                     SLP.SkyLine_Parser_Function_DataType_Integer,                   // Parser Call Function (Prefix) | Parse (INTEGER)
		SLTK.TOKEN_INTEGER8:                SLP.SkyLine_Parser_Function_DataType_Integer,                   // Parser Call Function (Prefix) | Parse (INTEGER 8)
		SLTK.TOKEN_INTEGER16:               SLP.SkyLine_Parser_Function_DataType_Integer,                   // Parser Call Function (Prefix) | Parse (INTEGER 16)
		SLTK.TOKEN_INTEGER32:               SLP.SkyLine_Parser_Function_DataType_Integer,                   // Parser Call Function (Prefix) | Parse (INTEGER 32)
		SLTK.TOKEN_INTEGER64:               SLP.SkyLine_Parser_Function_DataType_Integer,                   // Parser Call Function (Prefix) | Parse (INTEGER 64)
		SLTK.TOKEN_FLOAT:                   SLP.SkyLine_Parser_Function_DataType_Float64,                   // Parser Call Function (Prefix) | Parse (FLOAT)
		SLTK.TOKEN_NULL:                    SLP.SkyLine_Parser_Function_DataType_Null,                      // Parser Call Function (Prefix) | Parse (NULL)
		SLTK.TOKEN_IF:                      SLP.SkyLine_Parser_Expressions_Parse_Conditional,               // Parser Call Function (Prefix) | Parse (if, else, else if)
		SLTK.TOKEN_SWITCH:                  SLP.SkyLine_Parser_Expressions_Parse_ConditionalSwitch,         // Parser Call Function (Prefix) | Parse (switch, sw, case, cs, default, df)
		SLTK.TOKEN_FOR:                     SLP.SkyLine_Parser_Expressions_Parse_ConditionalLoop,           // Parser Call Function (Prefix) | Parse (for)
		SLTK.TOKEN_FOREACH:                 SLP.SkyLine_Parser_Expressions_Parse_WithinRangeLoop,           // Parser Call Function (Prefix) | Parse (foreach & in)
		SLTK.TOKEN_LBRACE:                  SLP.SkyLine_Parser_Function_DataType_HashMap,                   // Parser Call Function (Prefix) | Parse ( '{' )
		SLTK.TOKEN_LBRACKET:                SLP.SkyLine_Parser_Function_DataType_Array,                     // Parser Call Function (Prefix) | Parse ( '[' )
		SLTK.TOKEN_LPAREN:                  SLP.SkyLine_Parser_Expressions_Parse_GroupedExpr,               // Parser Call Function (Prefix) | Parse ( '(' )
		SLTK.TOKEN_MINUS:                   SLP.SkyLine_Parser_Expressions_Parse_Prefix,                    // Parser Call Function (Prefix) | Parse ( '-' )
		SLTK.TOKEN_IMPORT:                  SLP.SkyLine_Parser_Expressions_Parse_Importing,                 // Parser Call Function (Prefix) | Parse ( ' import ' )
	}

	//::::::::::::::::::::::::::::::::::
	//:: Map of infix parser functions
	//::::::::::::::::::::::::::::::::::

	SLP.SL_Infix_Parser_Functions = map[SLTK.SL_TokenDataType]SkyLine_Parser_InfixParser_Function{
		SLTK.TOKEN_PLUS:              SLP.SkyLine_Parser_Expressions_Parse_Infix,
		SLTK.TOKEN_ANDAND:            SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '&&' )
		SLTK.TOKEN_OROR:              SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '||' )
		SLTK.TOKEN_EQ:                SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '==' )
		SLTK.TOKEN_GT:                SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '>' )
		SLTK.TOKEN_LT:                SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '<' )
		SLTK.TOKEN_LTEQ:              SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '<=' )
		SLTK.TOKEN_GTEQ:              SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '>=' )
		SLTK.TOKEN_NEQ:               SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '!=' )
		SLTK.TOKEN_SLASH:             SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '/' )
		SLTK.TOKEN_MINUS:             SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '-' )
		SLTK.TOKEN_ASTARISK:          SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '*' )
		SLTK.TOKEN_MODULO:            SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '%' )
		SLTK.TOKEN_POWEROF:           SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '**' )
		SLTK.TOKEN_BITWISE_OP_AND:    SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '&' )
		SLTK.TOKEN_BITWISE_OP_OR:     SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '|' )
		SLTK.TOKEN_BITWISE_OP_LSHIFT: SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '<<' )
		SLTK.TOKEN_BITWISE_OP_RSHIFT: SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '>>' )
		SLTK.TOKEN_BITWISE_OP_XOR:    SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '^' )
		SLTK.TOKEN_ASTERISK_EQUALS:   SLP.SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword,                // Parser Call Function (Infix) | Parse symbol ( '*=' )
		SLTK.TOKEN_MINUS_EQUALS:      SLP.SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword,                // Parser Call Function (Infix) | Parse symbol ( '-=' )
		SLTK.TOKEN_DIVEQ:             SLP.SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword,                // Parser Call Function (Infix) | Parse symbol ( '/=' )
		SLTK.TOKEN_PLUS_EQUALS:       SLP.SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword,                // Parser Call Function (Infix) | Parse symbol ( '+=' )
		SLTK.TOKEN_ASSIGN:            SLP.SkyLine_Parser_Expressions_Parse_AssignmentNoKeyword,                // Parser Call Function (Infix) | Parse symbol ( '=' )
		SLTK.TOKEN_PERIOD:            SLP.SkyLine_Parser_Functions_Parse_FunctionObjectCall,                   // Parser Call Function (Infix) | Parse symbol ( '.' )
		SLTK.TOKEN_LPAREN:            SLP.SkyLine_Parser_Functions_Parse_FunctionCall,                         // Parser Call Function (Infix) | Parse symbol ( '(' )
		SLTK.TOKEN_LBRACKET:          SLP.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_IndexExpression, // Parser Call Function (Infix) | Parse symbol ( '[' )
		SLTK.TOKEN_DOTDOT:            SLP.SkyLine_Parser_Expressions_Parse_Infix,                              // Parser Call Function (Infix) | Parse symbol ( '..' )
		SLTK.TOKEN_MODULECALL:        SLP.SkyLine_SelectorExpression,                                          // Parser Call Function (Infix) | Parse symbol ( '::' )
	}

	//::::::::::::::::::::::::::::::::::
	//:: Map of postfix parser functions
	//::::::::::::::::::::::::::::::::::
	SLP.SL_PostFix_Parser_Functions = map[SLTK.SL_TokenDataType]SkyLine_Parser_PostFixParser_Function{
		SLTK.TOKEN_MINUS_MINUS: SLP.SkyLine_Parser_Expressions_Parse_Postfix, // Parser Call Function (Postfix) | Parse (--)
		SLTK.TOKEN_PLUS_PLUS:   SLP.SkyLine_Parser_Expressions_Parse_Postfix, // Parser Call Function (Postfix) | Parse (++)
	}

	return SLP
}
