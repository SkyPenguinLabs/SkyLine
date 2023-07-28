///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Token_Constants
// Extension         | .go ( golang source code file )
// Purpose           | Define constant definitions for string values of Tokens
// Directory         | Modules/Backend/SkyTokens
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens
// Package Name      | SkyLine_Backend_Tokens
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file helps define a tokens string value. For example, if our scanner comes across the byte of '=' ( which is 61 ) then the scanner will know to categorize this token
//
// as the TOKEN_ASSIGN. In a sense, this module helps define the token's string values, what the scanner should see, the patterns for the scanner as well as the type patterns
//
// and rule sets for tokens and keywords.
//
package SkyLine_Backend_Tokens

const (
	// New Data types
	TOKEN_BYTE      = "BYTE" // Byte value
	TOKEN_BYTESTART = "CHARACTER"
	//                                                                            ^ Call
	TOKEN_MODULECALL              = "::"                                     // imported object    | Implemented
	TOKEN_JUMP                    = "JUMP"                                   // Jump To @lable     | Not implemented
	TOKEN_DEFAULT                 = "DEFAULT"                                // Default keyword    | Implemented
	TOKEN_FOR                     = "FOR"                                    // For loop token     | Implemented
	TOKEN_IMPORT                  = "IMPORT"                                 // Import             | Implemented
	TOKEN_MODULE                  = "module"                                 // Module             | Implemented
	TOKEN_KEYWORD_ENGINE          = "ENGINE"                                 // ENGINE             | Implemented
	TOKEN_ENGINE_TYPE             = "ENGINE::ENVIRONMENT_MODIFIER->CALL:::>" // ENGINE ENV MODIFY  | Implemented
	TOKEN_FOREACH                 = "FOREACH"                                // For every element  | Implemented
	TOKEN_INSIDE                  = "IN"                                     // Within range       | Implemented
	TOKEN_REGISTER                = "REGISTER"                               // STD LIB Registry   | Implemented
	TOKEN_ILLEGAL                 = "ILLEGAL"                                // Illegal character  | Implemented
	TOKEN_EOF                     = "EOF"                                    // End Of File        | Implemented
	TOKEN_IDENT                   = "TOKEN_IDENT"                            // Identifier         | Implimented
	TOKEN_INT                     = "Integer"                                // TYPE integer       | Implemented
	TOKEN_INTEGER8                = "Integer8"                               // TYPE integer       | Implemented
	TOKEN_INTEGER16               = "Integer16"                              // TYPE integer       | Implemented
	TOKEN_INTEGER32               = "Integer32"                              // TYPE integer       | Implemented
	TOKEN_INTEGER64               = "Integer64"                              // TYPE integer       | Implemented
	TOKEN_FLOAT                   = "FLOAT"                                  // TYPE float         | Implemented
	TOKEN_STRING                  = "STRING"                                 // TYPE string        | Implemented
	TOKEN_NULL                    = "NULL"                                   // Type NULL          | Implemented
	TOKEN_CONSTANT                = "CONST"                                  // Constant           | Implemented
	TOKEN_FUNCTION                = "FUNCTION"                               // Function           | Implemented
	TOKEN_ALLOW                   = "SET"                                    // let statement      | Implemented
	TOKEN_TRUE                    = "TOKEN_TRUE"                             // boolean type true  | Implemented
	TOKEN_FALSE                   = "FALSE"                                  // boolean type false | Implemented
	TOKEN_IF                      = "IF"                                     // If statement       | Implemented
	TOKEN_ELSE                    = "ELSE"                                   // Else statement     | Implemented
	TOKEN_RETURN                  = "RETURN"                                 // return statement   | Implemented
	TOKEN_SWITCH                  = "SWITCH"                                 // Switch statement   | Implemented
	TOKEN_CASE                    = "CASE"                                   // Case statement 	   | Implemented
	TOKEN_REGEXP                  = "REGEXP"                                 // Regex Type         | Not implemented
	TOKEN_Lable                   = "@"                                      // Lables             | Not implemented
	TOKEN_LTEQ                    = "<="                                     // LT or equal to     | Implemented
	TOKEN_GTEQ                    = ">="                                     // GT or equal to     | Implemented
	TOKEN_ASTERISK_EQUALS         = "*="                                     // Multiply equals    | Implemented
	TOKEN_BANG                    = "!"                                      // Boolean operator   | Implemented
	TOKEN_ASSIGN                  = "="                                      // General assignment | Implemented
	TOKEN_PLUS                    = "+"                                      // General operator   | Implemented
	TOKEN_MINUS                   = "-"                                      // General operator   | Implemented
	TOKEN_ASTARISK                = "*"                                      // General operator   | Implemented
	TOKEN_SLASH                   = "/"                                      // General operator   | Implemented
	TOKEN_LT                      = "<"                                      // Boolean operator   | Implemented
	TOKEN_GT                      = ">"                                      // Boolean operator   | Implemented
	TOKEN_EQ                      = "=="                                     // Boolean operator   | Implemented
	TOKEN_MINUS_EQUALS            = "-="                                     // Minus equals       | Implemented
	TOKEN_NEQ                     = "!="                                     // Boolean operator   | Implemented
	TOKEN_DIVEQ                   = "/="                                     // Division operator  | Implemented
	TOKEN_PERIOD                  = "."                                      // Method Call        | Implemented
	TOKEN_PLUS_EQUALS             = "+="                                     // Plus equals        | Implemented
	TOKEN_COMMA                   = ","                                      // Seperation         | Implemented
	TOKEN_SEMICOLON               = ";"                                      // SemiColon          | Implemented
	TOKEN_COLON                   = ":"                                      // Colon              | Implemented
	TOKEN_LPAREN                  = "("                                      // Args start         | Implemented
	TOKEN_RPAREN                  = ")"                                      // Args end           | Implemented
	TOKEN_LINE                    = "|"                                      // Line con           | Implemented
	TOKEN_LBRACE                  = "{"                                      // Open  f            | Implemented
	TOKEN_RBRACE                  = "}"                                      // Close f            | Implemented
	TOKEN_LBRACKET                = "["                                      // Open               | Implemented
	TOKEN_RBRACKET                = "]"                                      // Close              | Implemented
	TOKEN_OROR                    = "||"                                     // Condition or or    | Implemented
	TOKEN_ANDAND                  = "&&"                                     // Boolean operator   | Implemented
	TOKEN_BACKTICK                = "`"                                      // Backtick           | Implemented
	TOKEN_POWEROF                 = "**"                                     // General operator   | Implemented
	TOKEN_MODULO                  = "%"                                      // General operator   | Implemented
	TOKEN_NEWLINE                 = '\n'                                     // COND               | Implemented
	TOKEN_PLUS_PLUS               = "++"                                     // Plus Plus          | Implemented
	TOKEN_QUESTION                = "?"                                      // Question que       | Not implemented
	TOKEN_DOTDOT                  = ".."                                     // Range              | Implemented
	TOKEN_CONTAINS                = "~="                                     // Contains           | Not implemented
	TOKEN_NOTCONTAIN              = "!~"                                     // Boolean operator   | Not implemented
	TOKEN_MINUS_MINUS             = "--"                                     // Minus minus        | Implemented
	TOKEN_BITWISE_OP_OR           = "|"                                      // Bitwise OR         | Implemented
	TOKEN_BITWISE_OP_XOR          = "^"                                      // Bitwise XOR        | Implemented
	TOKEN_BITWISE_OP_AND          = "&"                                      // Bitwise AND        | Implemented
	TOKEN_BITWISE_OP_LSHIFT       = "<<"                                     // Bitwise Shift L    | Implemented
	TOKEN_BITWISE_OP_RSHIFT       = ">>"                                     // Bitwise Shift R    | Implemented
	TOKEN_PEEKASSIGN              = ":="                                     // Peek assignment    | Implemented
	TOKEN_FUNCTION_DEFINE_LITERAL = "DEFINE_FUNCTION"                        // Define function    | Implemented
)
