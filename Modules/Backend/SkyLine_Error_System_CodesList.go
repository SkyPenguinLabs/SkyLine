/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Error_System_CodesList.go
// Project       |  SkyLine programming language
// Line Count    |  500+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines ALOT of data and information. In basic terms all this file defines is a code and error code system for the frontend of SkyLine which allows for
//
//                 more in depth parser errors, scanner errors, tokenization errors, evaluation errors, file errors, engine error, system errors and much more. This file REQUIRES ALOT
//
//                 of actual work and needs to be replaced heavily, for right now the state works
//
// STATE         | Needs to be organized
// Resolution    | Make extra states and functions for the language as well as ensure that the messages are pretty good.
//
//
package SkyLine_Backend

import (
	"fmt"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// | The offical SkyLine error system, this file includes code systems, notes as big as this and messages as well as loggers and other statements which allow you to make a
///////  | much more verbose and advanced error system. This system uses a custom code and output system to define its errors. Do note that this is still in beta and needs to be
///////  | tested much further before deployment and that it is important to note that not all codes need to be used.
//////   |
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Date  : Mon 27 Feb 2023 10:28:49 PM EST
// Codes :
// Types : Files, Parser errors, argument error, type errors, unknown operators, operational errors, mathematical errors, byte errors, binary errors, flow errors, systematic errors, os errors, support errors, network errors and more.
//
//
// The code system for SkyLine will differ based on the type of code and the type of system you are trying to troubleshoot, for example
// certain modules, functions, nodes and starters have reserved sets of codes which all mean different things and which have their own error messages.
// With every error based on the error and type of code if it is supported or required will allow for recomendations such as missing imports, semicolons etc
// below you will find a list of all the error codes and their reserved spots
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
//
//
//
//
//
// Startup errors
const (
	ERROR_SOURCE_FLAG_DEFINED_FILE_DOES_NOT_EXIST = 5  // File does not exist
	ERROR_SOURCE_FLAG_DEFINED_FILE_IS_A_DIRECTORY = 10 // File is a directory
	ERROR_SOURCE_FLAG_DEFINED_FILE_IS_EMPTY       = 15 // File is empty, this is bad to run in some cases
	ERROR_SOURCE_FLAG_DEFINED_FILE_IS_INVALID_SRC = 20 // File is a invalid source code file ( does not end in .csc, .skyline, .sl, .srcsl, .datasl)
)

// Parser Errors
const (
	ERROR_DURING_PEEK_IN_PARSER                     = 100 // Peek error, missing token after so and so character
	ERROR_FILE_INTEGRITY_BAD_FILE                   = 105 // File from import or require was not a CSC/SkyLine source file
	ERROR_FILE_INTEGRITY_IS_DIRECTORY               = 110 // File from import or require was a directory
	ERROR_FILE_INTEGRITY_IS_EMPTY                   = 115 // File from import or require was empty
	ERROR_FILE_INTEGRITY_DOES_NOT_EXIST             = 120 // File from import or require does not exist or was not found
	ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED           = 125 // File from import or require was not able to be read by the reader
	ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR        = 130 // Parser failed to parse an integer of the supported types
	ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT         = 135 // Parser has failed to parse due to an unterminated string, constant, function or type
	ERROR_EXPECTED_ASSIGNMENT_OPERATOR_BEFORE_TOKEN = 140 // Parser could not find an assignment operator before the check
	ERROR_UNTERMINATED_CONSTANT_VALUE               = 145 // Parser found a constant value that was not terminated
	ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV = 150 // Parser could not locate the prefix parse function for the following token
	ERROR_COULD_NOT_PARSE_FLOAT_VALUE               = 155 // Parser could not parse the provided input as a float64
	ERROR_PARSER_FOUND_NIL_EXPRESSION_UNEXPECTED    = 160 // Parser found NIL expression in block statement for conditionals
	ERROR_PARSER_EXPECTED_LBRACE_BUT_GOT_SOMETHING  = 165 // Parser was expecting a left brace { but rather got another unexpected token
	ERROR_PARSER_EXPECTED_RBRACE_BUT_GOT_SOMETHING  = 170 // Parser was expecting a right closing brace but got another unexpected token
)

// Evaluator Errors
const (
	ERROR_NOT_ENOUGH_ARGUMENTS_IN_CALL_TO_FUNCTION             = 300 // Arguments in call to function were not correct, function requires ... but unexpected to get ...
	ERROR_INVALID_OPERATOR_DURING_EVALUATION                   = 310 // Invalid prefix operator when supported includes (! and -)
	ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE       = 320 // Function or prefix parse function did not exist nor did its respective type exist
	ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX        = 330 // Index operator unsupported or non existent
	ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY            = 340 // Key within hash could not be hashable in hash literal
	ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY             = 350 // Key within hash in evaluation to hash literal could not be parsed
	ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN = 360 // Identifier was not found or located within the current environment
	ERROR_INVALID_DATATYPE_INIFX_OPERATION_LR                  = 370 // the type of the left value does not equal the type of the right value
)

// Environment Errors
const (
	ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS = 800 // Not enough arguments in call to function
	ERROR_CONSTANT_VALUE_ATTEMPT_TO_CHANGE     = 830 // Cannot change the value of constant x
)

// Standard Library Errors
const (
	ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS         = 1000 // In call to standard library function ... does not have enough arguments, was given ... but requires ...
	ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED   = 1010 // In call to standard library function ... the given argument at (n) was provided to be the wrong data type
	ERROR_STANDARD_PRESET_CODE_ARGUMENT_WAS_NOT_FOUND_OR_SUPPORTED = 1020 // In call to standard library function ... the given argument which is a default argument was not found or expected
	ERROR_STANDARD_FUNCTION_GOT__ARGUMENTS                         = 1030 // In call to standard library function ... the amount of arguments supplied was either too little or too large
	ERROR_STANDARD_FUNCTION_COULD_NOT_SET_ARGS_VAL_AFTER_SEP_NIL   = 1040 // In call to standard library function ... when splitting arguments usually with system:arg where arg = n, n was left empty which is bad
	ERROR_STANDARD_FUNCTION_COULD_NOT_LOCATE_ARGUMENT_SEPERATOR    = 1050 // In call to standard library function ... when trying to split the arguments, the seperator was not found
)

// Developer Errors
const (
	ERROR_TOO_MANY_ARGUMENTS_IN_CALL_TO_ERROR_MAP = 19000 // In call to error map within source code, there was way too many arguments
)

// Code list errors
const (
	ERROR_SUPPORT_IN_MODIFY_ENVIRONMENT_MODIFY_NOT_SUPPORTED = 10500 // In call to modify(), the first argument modif(arg:setting) where arg = first, the environment modifier was not found
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// | This next section here will define the maps and function calls necessary to operate with the error codes. These variables are maps
///////  | which result in a function that returns a new message and error setting called by and called from sprintf. These messages can be
///////  | of SLC_Object error type and will need to return a new error message and setting or recomendation.
///////  |
///////  | - Date: Tue 28 Feb 2023 04:37:17 PM EST
///////  |
///////  |
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// C_1 -> Map_Eval   : Eval        error specific map
// C_2 -> Map_Parser : Parser      error specific map
// C_3 -> Map_Std    : Stdlib      error specific map
// C_4 -> Map_Envi   : Environment error specific map
//

// type

type ErrorSystemReturn struct {
	Message    string // Message
	Code       int    // Error code
	Suggestion string // Suggested fix
}

func GenerateErrorTooMuchArgs() ErrorSystemReturn {
	return ErrorSystemReturn{
		Message:    "[DE] (Dev Error) : Too many arguments used in call to map",
		Code:       ERROR_TOO_MANY_ARGUMENTS_IN_CALL_TO_ERROR_MAP,
		Suggestion: "Constact developers if this issue appears",
	}
}

func GenerateErrorDeveloper(err string) ErrorSystemReturn {
	return ErrorSystemReturn{
		Message:    "Error system function call was not found, developer error ( not enough arguments in call to error map) - " + err,
		Code:       ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS,
		Suggestion: "Contact developers if this issue appears",
	}
}

// C4 | Environment errors during token identification
var Map_Envi = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_CONSTANT_VALUE_ATTEMPT_TO_CHANGE: func(Arguments ...string) ErrorSystemReturn {
		return ErrorSystemReturn{
			Message:    "Attempting to change constant value which is ILLEGAL",
			Code:       ERROR_CONSTANT_VALUE_ATTEMPT_TO_CHANGE,
			Suggestion: "Do not change constants",
		}
	},
	ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS: func(Arguments ...string) ErrorSystemReturn {
		return ErrorSystemReturn{
			Message:    "A certain function was called but was not given the proper argument counts.",
			Code:       ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS,
			Suggestion: "Meet the function argument requirements",
		}
	},
}

// C3 | Standard Library errors
var Map_STDLIB = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf("Standard function call to (%s) requires at least %s argument(s) but you have (%s) argument(s)",
					Arguments[0],
					Arguments[1],
					Arguments[2],
				),
				Code: ERROR_STANDARD_FUNCTION_DOES_NOT_HAVE_PROPER_ARGUMENTS,
				Suggestion: fmt.Sprintf("Suggestion : When calling function (%s) use only %s argument(s) to prevent this error", Arguments[0],
					Arguments[1]),
			}
		} else {
			return GenerateErrorDeveloper("Standard function call to (%s) requires at least %s argument(s) but you have (%s) argument(s)")
		}
	},
	ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 4 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf("Standard library call to function (%s) was given the wrong data type at POS (%s) in the argument list, the data type of the variable given was %s but the function expects you to at least have %s",
					Arguments[0],
					Arguments[1],
					Arguments[2],
					Arguments[3],
				),
				Code:       ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED,
				Suggestion: "Use the correct data type the function requires",
			}
		} else {
			return GenerateErrorDeveloper("Standard library call to function (%s) was given the wrong data type at POS (%s) in the argument list")
		}
	},
	ERROR_STANDARD_PRESET_CODE_ARGUMENT_WAS_NOT_FOUND_OR_SUPPORTED: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"In call to standard library function %s, this function has a param that is pre set on the backend or static which requires a specific title. You gave this function the following argument %s in POS %s but this function requires a different argument which may not be supported",
					Arguments[0],
					Arguments[1],
					Arguments[2],
				),
				Code:       ERROR_STANDARD_PRESET_CODE_ARGUMENT_WAS_NOT_FOUND_OR_SUPPORTED,
				Suggestion: "Use the std.lookup function to lookup standard library functions and fill out the required details to lookup the functions details and requirements or suggestions",
			}
		} else {
			return GenerateErrorDeveloper(
				"In call to standard library function NULL, this function has a param that is pre set on the backend or static which requires a specific title. You gave this function the following argument %s in POS %s but this function requires a different argument which may not be supported",
			)
		}
	},
}

// C2 | Parser error codes, this map will return error codes for the parser itself.
var Map_Parser = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 0 {
			return ErrorSystemReturn{
				Code:       ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV,
				Message:    "RULE(Prefix_Parse_Function_Not_Found) : Could not locate the prefix parsing function for the token",
				Suggestion: "Consider adding a new token that exists within the environment and can be parsed",
			}
		} else {
			return GenerateErrorTooMuchArgs()
		}
	},
	ERROR_UNTERMINATED_CONSTANT_VALUE: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			if strings.Contains(Arguments[0], ":") {
				Arguments[0] = strings.Trim(Arguments[0], ";")
			}
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"RULE(UNTERMINATED) : Missing semicolon in statement -> [ %s %s %s ] ",
					Arguments[1],
					Arguments[0],
					Arguments[2],
				),
				Code:       ERROR_UNTERMINATED_CONSTANT_VALUE,
				Suggestion: "Consider adding a (';') <semicolon> at the end of the statement",
			}
		} else {
			return GenerateErrorDeveloper("RULE(UNTERMINATED) : Missing semicolon in statement ->")
		}
	},
	ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			if strings.Contains(Arguments[0], ";") {
				Arguments[0] = strings.Trim(Arguments[0], ";")
			}

			return ErrorSystemReturn{
				Message:    fmt.Sprintf("RULE(UNTERMINATED) : Missing semicolon in statement -> [ %s %s %s ] ", Arguments[1], Arguments[0], Arguments[2]),
				Code:       ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT,
				Suggestion: "Make sure you add a (';') at the end of the statement",
			}
		} else {
			return GenerateErrorDeveloper("RULE(UNTERMINATED) : Missing semicolon in statement ->")
		}
	},
	ERROR_DURING_PEEK_IN_PARSER: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 2 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When peeking or looking for the next token, the parser could not find the required token. Expected next token to be %s but rather found %s",
					Arguments[0],
					Arguments[1],
				),
				Code:       ERROR_DURING_PEEK_IN_PARSER,
				Suggestion: "Check to make sure your token is correct",
			}
		} else {
			return GenerateErrorDeveloper("When peeking or looking for the next token, the parser could not find the required token")
		}
	},
	ERROR_FILE_INTEGRITY_BAD_FILE: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing the file PARSE(%s), the file fails to become parsed or called due to the file being bad. ",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_BAD_FILE,
				Suggestion: "File extension INVALID -> File types are .csc, .sl, .skyline, .SkyLine, .core, .seccore",
			}
		} else {
			return GenerateErrorDeveloper("When parsing the file PARSE(NULL[DEVERROR]), the file fails to become parsed or called due to the file being bad.")
		}
	},
	ERROR_FILE_INTEGRITY_DOES_NOT_EXIST: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing for file (%s) the standard checks detected that this file did not exist within the given directory",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_DOES_NOT_EXIST,
				Suggestion: fmt.Sprintf("Make sure file %s is correct ", Arguments[0]),
			}
		} else {
			return GenerateErrorDeveloper("When parsing for file (NULL[DEV_ERROR]) the standard checks detected that this file did not exist within the given directory")
		}
	},
	ERROR_FILE_INTEGRITY_IS_DIRECTORY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing for file (%s) the standard checks detected that this was not a file but rather a directory. Execution of CSC files within a directory is not supported in the current version...",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_IS_DIRECTORY,
				Suggestion: "Provide a source code file, not a directory",
			}
		} else {
			return GenerateErrorDeveloper("When parsing for file (NULL[DEV_ERROR]) the standard checks detected that this was not a file but rather a directory.")
		}
	},
	ERROR_FILE_INTEGRITY_IS_EMPTY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing for file (%s) the standard checks detected that this was a empty file, this is a security issue and the interpreter will not accept any file that is empty.",
					Arguments[0],
				),
				Code:       ERROR_FILE_INTEGRITY_IS_EMPTY,
				Suggestion: "Provide a source code file that has something inside of it and is not NULL",
			}
		} else {
			return GenerateErrorDeveloper("When parsing for file (NULL[DEV_ERROR]) the standard checks detected that this was a empty file")
		}
	},
	ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing the file (%s) the Input and Output buffer when reading the file has failed for some reason as it gave this error (%s)",
					Arguments[0],
					Arguments[0],
				),
				Code:       ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED,
				Suggestion: "Provide a source code file that may not be corrupted, damaged, injected or infected with data or is being used by another process",
			}
		} else {
			return GenerateErrorDeveloper("When parsing the file (NULL[DEV_ERROR]) the Input and Output buffer when reading the file has failed for some reason as it gave this error")
		}

	},
	ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"When parsing integer (%s) the parser came across an error that made it bad, unsafe or made the parser error out when parsing the provided integer",
					Arguments[0],
				),
				Code:       ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR,
				Suggestion: "Provide a number variable that does not overflow the Integer64 data type indicating that it can be a real integer and can be parsed by SkyLine",
			}
		} else {
			return GenerateErrorDeveloper("When parsing integer (NULL[DEV_ERROR]) the parser came across an error that made it bad, unsafe or made the parser error out when parsing the provided integer")
		}
	},
}

// C1 | Evaluation step errors

var Map_Eval = map[int]func(Arguments ...string) ErrorSystemReturn{
	ERROR_NOT_ENOUGH_ARGUMENTS_IN_CALL_TO_FUNCTION: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 3 {
			return ErrorSystemReturn{
				Message: fmt.Sprintf(
					"In call to function %s, there were not enough arguments that were needed as this function requires %s argumen(s) but you gave it %s argument(s)",
					Arguments[0],
					Arguments[1],
					Arguments[2],
				),
				Code:       ERROR_NOT_ENOUGH_ARGUMENTS_IN_CALL_TO_FUNCTION,
				Suggestion: "Use the correct number of arguments",
			}
		} else {
			return GenerateErrorDeveloper("In call to function [NULL_ERROR_DEV], there were not enough arguments that were needed as this function requires %s argumen(s) but you gave it %s argument(s)")

		}
	},
	ERROR_INVALID_OPERATOR_DURING_EVALUATION: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_OPERATOR_DURING_EVALUATION,
				Suggestion: "Use the supported operators",
				Message: fmt.Sprintf(
					"When evalutating a given operator (%s) the evaluator found that it was not supported, please check the laws of the evaluator",
					Arguments[0],
				),
			}
		} else {
			return GenerateErrorDeveloper("USE SUPPORTED OPERATORS")
		}
	},
	ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE,
				Suggestion: "Make sure evaluator has this valid",
				Message: fmt.Sprintf(
					"Function type (%s) was not found, this is bad please make sure the function is the supported type or this feature was implemented",
					Arguments[0],
				),
			}

		} else {
			return GenerateErrorDeveloper("Make sure evaluator has this valid | INVALID DATA TYPE [ dev error ] ")

		}
	},
	ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code: ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX,
				Message: fmt.Sprintf(
					"Index operator (%s) unsupported or non existent",
					Arguments[0],
				),
				Suggestion: "Make sure the index operator is supported ",
			}

		} else {
			return GenerateErrorDeveloper("Index operator (NULL[DEV_ERROR]) unsupported or non existent")

		}
	},
	ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY,
				Message:    fmt.Sprintf("Key (%s) was not able to be used as a valid hash key", Arguments[0]),
				Suggestion: "make sure the hash key you gave is correct and stable",
			}
		} else {
			return GenerateErrorDeveloper("Key (%s) was not able to be used as a valid hash key")
		}
	},
	ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 1 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY,
				Message:    fmt.Sprintf("Could not parse hash key (%s)", Arguments[0]),
				Suggestion: "Make sure this key can be used within the hash",
			}
		} else {
			return GenerateErrorDeveloper("Couuld not parse hash key")
		}
	},
	ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN: func(Arguments ...string) ErrorSystemReturn {
		if len(Arguments) == 2 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN,
				Suggestion: "The identifier you tried to parse was not assigned or registered to the environment, check if it exists",
				Message:    fmt.Sprintf("The identifier (%s) of type (%s) could not be found or was not known to the interpreters environment", Arguments[0], Arguments[1]),
			}

		} else {
			return GenerateErrorDeveloper("The identifier you tried to parse was not assigned or registered to the environment, check if it exists")
		}
	},
	ERROR_INVALID_DATATYPE_INIFX_OPERATION_LR: func(Arguments ...string) ErrorSystemReturn {
		// needs left data type and right data type
		if len(Arguments) == 2 {
			return ErrorSystemReturn{
				Code:       ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN,
				Suggestion: "Make sure the type of the left value is equal to the type of the right value",
				Message: fmt.Sprintf(
					"During evaluation of the infix operation, the data type of the left value which yeilds (%s) must be equal to the type of the right value (%s) which in this case is not true, fix it.",
					Arguments[0],
					Arguments[1],
				),
			}

		} else {
			return GenerateErrorDeveloper("Make sure the type of the left value is equal to the type of the right value")
		}
	},
}
