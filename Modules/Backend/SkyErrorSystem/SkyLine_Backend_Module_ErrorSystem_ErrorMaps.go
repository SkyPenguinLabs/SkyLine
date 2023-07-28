///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
//
// This module is designed to be used for construction and output of specific errors such as formatting information
//
package SkyLine_Error_System

// We want the tree to output the following
// Error : This defines the error message
// Type  : if it is a warning or a fatal error
// Tech  : This defines the technology it came from (Parser, Engine-Parser, Engine-Lexer, Regex-Lexer, Regex-Engine, Engine, Configuration, AST, Configuration, Initation etc)
// Line  : Line number or range
// BoxOCode : Box of range of code
func UseErrorMap(code int) string {
	return ErrorMap[code]
}

var ErrorMap = map[int]string{
	SkyLine_Parser_UUnterminated_Function_Parameter:              "Unterminated function parameter list, I need you to place a left facing parentheses ')' ",
	SkyLine_Parser_Unexpected_Null_Conditional_Unit:              "Unexpected NIL expression when parsing an alternative conditional unit",
	SkyLine_Parser_Unexpected_Token_Expected_Spec_Token:          "Unexpected token, expected \033[38;5;86m%s\033[31mbut found \033[38;5;86m%s",
	SkyLine_Initation_Flag_Defined_File_Does_Not_Exist:           "Could not locate or open file %s, file may not exist",
	SkyLine_Initation_Flag_Defined_File_Is_A_Directory:           "File that was attempted to be opened was a directory",
	SkyLine_Initation_Flag_Defined_File_Is_NULL:                  "File was not run, NULL symbols | UNSAFE",
	SkyLine_Initation_Flag_Defined_Unknown_Source:                "File does not have a direct source extension, invalid source code",
	SkyLine_Initation_Flag_Defined_Engine_But_No_Source_File:     "Engine called but you gave no source file",
	SkyLine_Initation_Flag_Defined_Engine_But_Unknown_Source:     "Engine called but you did not give a valid source code file",
	SkyLine_Initation_Flag_Defined_Evaluate_Line_But_Null:        "Evaluation could not find any symbols to execute",
	SkyLine_Parser_Unterminated_Constant_Statement:               "Unterminated constant statement, missing semicolon",
	SkyLine_Parser_Unterminated_Declaration_Statement:            "Unterminated variable declaration (let|set|cause|allow), missing semicolon",
	SkyLine_Parser_Unterminated_Code_Unit:                        "Unterminated code unit, missing semicolon",
	SkyLine_Parser_Unterminated_Switch_Block:                     "Unterminated switch case block, missing semicolon at the end of switch",
	SkyLine_Parser_Corrupted_Foreach_Expression:                  "Broken foreach expression",
	SkyLine_Parser_Corrupted_Switch_Satement:                     "Broken switch statement",
	SkyLine_Parser_Missing_RightFacing_Parenthesis_Token:         "Expecting right facing paren ' ( '  you gave me %s ",
	SkyLine_Parser_Missing_RightFacing_Curely_Brace_Token:        "Expecting right facing curly brace ' { ' you gave me %s ",
	SkyLine_Parser_Missing_LeftFacing_Curely_Brace_Token:         "Expecting left facing paren ' ) ' you gave me %s ",
	SkyLine_Parser_Missing_LeftFacing_Parenthesis_Token:          "Expecting left facing curly brace ' } ' you gave me %s ",
	SkyLine_Evaluator_UnknownOperator_When_FloatIntegerInfix:     "Unknown operator when evaluating Float64->Integer infix expression",
	SkyLine_Evaluator_UnknownOperator_When_IntegerInfix:          "Unknown operator when evaluating Integer->Integer infix expression",
	SkyLine_Evaluator_UnknownOperator_When_MinusPrefix:           "Unknown operator when evaluating postfix expression",
	SkyLine_Evaluator_UnknownOperator_When_PrefixExpression:      "Unknown operator when evaluating prefix expression",
	SkyLine_Evaluator_UnknownOperator_When_SpecialFloatInfix:     "Unknown operator when evaluating special float infix expression",
	SkyLine_Evaluator_UnknwonOperator_When_BooleanInfix:          "Unknown operator when evaluating boolean infix expression",
	SkyLine_Evaluator_UnknownIdentifier_When_IdentifierExecution: "Unknown identifier during evaluation",
	SkyLine_Evaluator_UnknownIndexOp_When_IndexExpression:        "Unknown Index operator when evaluating index expression",
	SkyLine_Evaluator_HashKeyFailure_When_HashLiteral:            "Hash key could not be used when evaluating hash literal",
	SkyLine_StandardLib_OutOfRange_Arguments:                     "Standard Library -? arguments are out of range",
	SkyLine_StandardLib_Mismatched_Arguments:                     "Standard Library -? mismatched positional argument types",
	SkyLine_StandardLib_Requiredms_Arguments:                     "Standard Library -? Arguments are out of range, needs specific positional argument count",
}
