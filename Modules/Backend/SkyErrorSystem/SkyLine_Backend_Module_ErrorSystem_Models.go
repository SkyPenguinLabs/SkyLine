/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Backend_Module_ErrorSystem_Models.go
// Project       |  SkyLine programming language
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//---------------------------------------------------------------------------------------------------------------------------------------------------------------------
//
// The SkyLine error system module defines errors for all possible systems such as SLC errors, engine parse errors, configuration errors, file errors, parser errors,
//
// evaluation errors, object errors, system errors, internal errors, external errors, load errors, Virtual Machine errors, compiler errors, code errors and environment
//
// errors. This system is MASSIVE and is defined for the majority of SkyLine but does not typically mix with SLC as SLC has its own unique error system and code system.
//
// This file defines all of the rule sets for the code systems or rather constants, variables, structures and more for the language that can be used when working with specific
//
// case's, evaluation, code parsing and more. Constants defined in this file are all relating to code while structures define error system structures.
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// | The offical SkyLine error system, this file includes code systems, notes as big as this and messages as well as loggers and other statements which allow you to make a
///////  | much more verbose and advanced error system. This system uses a custom code and output system to define its errors. Do note that this is still in beta and needs to be
///////  | tested much further before deployment and that it is important to note that not all codes need to be used.
//////   |
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// The code system for SkyLine will differ based on the type of code and the type of system you are trying to troubleshoot, for example
// certain modules, functions, nodes and starters have reserved sets of codes which all mean different things and which have their own error messages.
// With every error based on the error and type of code if it is supported or required will allow for recomendations such as missing imports, semicolons etc
// below you will find a list of all the error codes and their reserved spots
//
//
// Constant rules:
//			- Each constant must start with SkyLine
//          - Each constant must be of type integer
//          - Each constant must define the code type after skyline (SkyLine_ModuleOrSystem)
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
//
//
//
//
//				-!!!!!!!!!!!!!!!!!!!!![EXAMPLE CHART - LEGACY CODE BASE ]!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!-
//
// Startup        | -> 0-100
// Parser         | -> 100-300
// Evaluator      | -> 300-500
// Operations     | -> 500-700
// Environment    | -> 800-1000
// Function Calls | -> 1000-2000
// Network calls  | -> 2000-3000
// Mathematics    | -> 4000-5000
// Data types     | -> 5000-6000
// Server errors  | -> 6000-6500
// Arguments      | -> 6500-7000
//
//
// Each message and each code depending on the type of code and the catgeory increments completely differently and is all formatted differently. For example below you will find a
// list of the error code reservations and their incremented positions
//
//
// Startup        | -> 0-100         | Increments by 5
// Parser         | -> 100-300       | Increments by 5
// Evaluator      | -> 300-500       | Increments by 10
// Operations     | -> 500-700       | Increments by 20
// Environment    | -> 800-1000      | Increments by 30
// STD library    | -> 1000-2000     | Increments by 10
// Network calls  | -> 2000-3000     | Increments by 20
// Mathematics    | -> 4000-5000     | Increments by 80
// Data types     | -> 5000-6000     | Increments by 90
// Server errors  | -> 6000-10,000   | Increments by 100
// Arguments      | -> 10,000-10,500 | Increments by 100
// Map / Express  | -> 10,500-11,000 | Increments by 10
//
// These arguments all have their own respected values and all have their own unique values and codes as well
// for example if you were to make a HTTP POST request to a given domain and the endpoint failes a code within the
// range of 2000-3000 will be displayed because this counts as a network error within the system. Anything past 10,500
// is just wild and much more verbose and a complicated error or traceback. Codes should not ever reach past that point
// as that typically means there is a panic within the program, a fatal log, a dev issue on the backend or some weird issue.
// It is important to note that within each system there is no line or column based traceback system right now as this is
// a main priority and I would rather much prefer to focus on the functionality of the error system first before the backend
//
//
//
//
//			-!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!![ REAL CODE CHART - MODERN CODE BASE | MOST RECENT ]!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!-
//
package SkyLine_Error_System

const (
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Initation Errors are errors that happen before an env
	//::
	//:: is started.
	//::
	//:: RESERVED CODE SET -> 0-100
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	SkyLine_Initation_Flag_Defined_File_Does_Not_Exist       = 0  // Error Code ( SkyLine Initation ) | File that was meant to be loaded did not exist
	SkyLine_Initation_Flag_Defined_File_Is_A_Directory       = 5  // Error Code ( SkyLine Initation ) | File that was meant to be loaded is a directory
	SkyLine_Initation_Flag_Defined_File_Is_NULL              = 10 // Error Code ( SkyLine Initation ) | File that was meant to be loaded was NULL and had no code (DANGEROUS)
	SkyLine_Initation_Flag_Defined_Unknown_Source            = 15 // Error Code ( SkyLine Initation ) | File that was loaded did not have a valid extension
	SkyLine_Initation_Flag_Defined_Engine_But_No_Source_File = 20 // Error Code ( SkyLine Initation ) | Flag --SLC was defined but --Esource was not
	SkyLine_Initation_Flag_Defined_Engine_But_Unknown_Source = 25 // Error Code ( SkyLine Initation ) | Flag --SLC and --Esource were defined but the source code file did not have the correct extension
	SkyLine_Initation_Flag_Defined_Evaluate_Line_But_Null    = 30 // Error Code ( SkyLine Initation ) | Flag --Eval (-e) was called but no source code was found to load
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Parser Errors are errors that happen within the parser
	//::
	//:: and when code is being parsed
	//::
	//:: RESERVED CODE SET -> 100-200
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	SkyLine_Parser_Unterminated_Constant_Statement        = 100 // Error Code ( SkyLine Parsing ) | Unterminated constant statement, missing semicolon at end of statement
	SkyLine_Parser_Unterminated_Declaration_Statement     = 105 // Error Code ( SkyLine Parsing ) | Unterminated variable declaration, missing semicolon at end of statement
	SkyLine_Parser_Unterminated_Code_Unit                 = 110 // Error Code ( SkyLine Parsing ) | Unterminated unit / block statement, missing semicolon at the end of statement
	SkyLine_Parser_Unterminated_Switch_Block              = 115 // Error Code ( SkyLine Parsing ) | Unterminated switch|case|default unit, missing semicolon at the end of SWITCH DECL
	SkyLine_Parser_Corrupted_Switch_Satement              = 120 // Error Code ( SkyLine Parsing ) | Corrupted switch case expression, expected either case|default
	SkyLine_Parser_Corrupted_Foreach_Expression           = 125 // Error Code ( SkyLine Parsing ) | Corrupted foreach expression, second argument must be identifier not ...
	SkyLine_Parser_Missing_RightFacing_Parenthesis_Token  = 130 // Error Code ( SkyLine Parsing ) | Expected next token and position to spot '(' but got ...
	SkyLine_Parser_Missing_LeftFacing_Parenthesis_Token   = 135 // Error Code ( SkyLine Parsing ) | Expected next token and position to spot ')' but got ...
	SkyLine_Parser_Missing_RightFacing_Curely_Brace_Token = 140 // Error Code ( SkyLine Parsing ) | Expected next token and position to spot '{' but got ...
	SkyLine_Parser_Missing_LeftFacing_Curely_Brace_Token  = 145 // Error Code ( SkyLine Parsing ) | Expected next token and position to spot '}' but got ...
	SkyLine_Parser_Unexpected_Token_Expected_Spec_Token   = 150 // Error Code ( SkyLine Parsing ) | Expected next token and position to be ... but got ...
	SkyLine_Parser_Unexpected_Null_Conditional_Unit       = 155 // Error Code ( SkyLine Parsing ) | Unexpected NULL conditional unit.
	SkyLine_Parser_UUnterminated_Function_Parameter       = 160 // Error Code ( SkyLine Parsing ) | Unterminated function parameter list, missing ')'
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Evaluation errors are errors that happen during code
	//::
	//:: evaluation and execution.
	//::
	//:: RESERVED CODE SET -> 200-?
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	SkyLine_Evaluator_UnknownOperator_When_FloatIntegerInfix     = 200 // Error Code ( SkyLine Evaluation ) | When working with a float->integer infix expression, operator was unsupported
	SkyLine_Evaluator_UnknownOperator_When_IntegerInfix          = 205 // Error Code ( SkyLine Evaluation ) | When working with Integer->Integer infix expression, operator was unsupported
	SkyLine_Evaluator_UnknownOperator_When_SpecialFloatInfix     = 210 // Error Code ( SkyLine Evaluation ) | When working with a special float->integer infix expression, operator was unknown
	SkyLine_Evaluator_UnknwonOperator_When_BooleanInfix          = 215 // Error Code ( SkyLine Evaluation ) | When working wih a boolean infix expression, operator unsuppported
	SkyLine_Evaluator_UnknownOperator_When_MinusPrefix           = 220 // Error Code ( SkyLine Evaluation ) | When working with minux prefix expression, operator unknown
	SkyLine_Evaluator_UnknownOperator_When_PrefixExpression      = 225 // Error Code ( SkyLine Evaluation ) | When working with prefix expression, operator unknown
	SkyLine_Evaluator_MismatchType_When_InfixExpression          = 230 // Error Code ( SkyLine Evaluation ) | When working with infix expression, data type mismatch
	SkyLine_Evaluator_UnknownIdentifier_When_IdentifierExecution = 235 // Error Code ( SkyLine Evaluation ) | When evaluating identifier ..., seems to not exist within the environment
	SkyLine_Evaluator_UnknownIndexOp_When_IndexExpression        = 240 // Error Code ( SkyLine Evaluation ) | When evaluating index expression, unsupported index type or expression
	SkyLine_Evaluator_HashKeyFailure_When_HashLiteral            = 245 // Error Code ( SkyLine Evaluation ) | When evaluating hash expression, hash key failed to be used.
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Standard Library Based Errors
	//::
	//:: RESERVED CODE SET -> 0100-?
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	SkyLine_StandardLib_OutOfRange_Arguments = 0100 // Error Code ( SkyLine Std library ) | When parsing function arguments, out of range. Supposed to be at least ... and at most ...
	SkyLine_StandardLib_Mismatched_Arguments = 0130 // Error Code ( SkyLine Std library ) | When parsing function arguments, data type mismtch. Argument needed (type) but you gave (type)
	SkyLine_StandardLib_Requiredms_Arguments = 0160 // Error Code ( SkyLine Std library ) | When parsing function arguments, needed the exact argument count but you gave
)

type (
	ConstantEnvironment struct {
		LibraryIdentifiers []string // This will hold all STD library identifiers
		LibraryModules     []string // This will hold all module names
		FunctionCalls      []string // This will hold all function calls
	}
)

var (
	ErrorConstEnv ConstantEnvironment
)
