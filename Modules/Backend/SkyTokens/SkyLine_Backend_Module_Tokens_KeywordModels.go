///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Token_KeywordModels
// Extension         | .go ( golang source code file )
// Purpose           | Defines a map to map specific identifiers to keywords within the language
// Directory         | Modules/Backend/SkyTokens
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens
// Package Name      | SkyLine_Backend_Tokens
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file helps better tell the scanner to categorize a specific identifier as a keyword if it matches. Well, instead of telling it this file rather contains a map that maps
//
// specific identifiers or to their respected token values. This helps the scanner know what to parse if it is not a single character.
//
//
// Tokens in SkyLine are the following
//
// Token Named Value And Map Result | Token Description  | Tokens assigned
// -------------------------------- | ------------------ | ----------------
// TOKEN_FUNCTION_DEFINE_LITERAL    | defines function   | define, func
// TOKEN_FUNCTION                   | allows functions   | Func, function
// TOKEN_ALLOW                      | variable decl      | allow, set, cause, let
// TOKEN_TRUE                       | boolean true       | true, BOOLEANT
// TOKEN_FALSE                      | boolean false      | false, BOOLEANF
// TOKEN_RETURN                     | return values      | ret, return
// TOKEN_CONSTANT                   | constant variable  | constant, const
// TOKEN_SWITCH                     | switch expression  | switch, sw
// TOKEN_CASE                       | case expression    | case, cs
// TOKEN_DEFAULT                    | default expression | default, df
// TOKEN_REGISTER                   | register library   | register
// TOKEN_KEYWORD_ENGINE             | SLC call           | ENGINE
// TOKEN_IMPORT                     | import files       | import
// TOKEN_FOR                        | for expression     | for
// TOKEN_STRING                     | string             | STRING
// TOKEN_INSIDE                     | within expression  | in
// TOKEN_NULL                       | empty expression   | null
//
package SkyLine_Backend_Tokens

var SkyLine_Keywords = map[string]SL_TokenDataType{
	"define":   TOKEN_FUNCTION_DEFINE_LITERAL, // Define function literal
	"func":     TOKEN_FUNCTION_DEFINE_LITERAL, // Function definition
	"Func":     TOKEN_FUNCTION,                // Function
	"function": TOKEN_FUNCTION,                // Function
	"let":      TOKEN_FUNCTION,                // Variable declaration let
	"set":      TOKEN_ALLOW,                   // Variable declaration set
	"cause":    TOKEN_ALLOW,                   // Variable declaration cause
	"allow":    TOKEN_ALLOW,                   // Variable declaration allow
	"true":     TOKEN_TRUE,                    // Boolean true
	"false":    TOKEN_FALSE,                   // Boolean false
	"if":       TOKEN_IF,                      // Conditional start
	"else":     TOKEN_ELSE,                    // Conditional alternative
	"return":   TOKEN_RETURN,                  // Return decl
	"ret":      TOKEN_RETURN,                  // Return decl
	"const":    TOKEN_CONSTANT,                // Constant type
	"constant": TOKEN_CONSTANT,                // Constant type
	"switch":   TOKEN_SWITCH,                  // Switch statement
	"sw":       TOKEN_SWITCH,                  // Switch statement
	"case":     TOKEN_CASE,                    // Case statement
	"cs":       TOKEN_CASE,                    // Case statement
	"default":  TOKEN_DEFAULT,                 // Switch alternative
	"df":       TOKEN_DEFAULT,                 // Switch alternative
	"register": TOKEN_REGISTER,                // Register
	"ENGINE":   TOKEN_KEYWORD_ENGINE,          // Engine caller
	"import":   TOKEN_IMPORT,                  // Import data
	"for":      TOKEN_FOR,                     // For loop
	"STRING":   TOKEN_STRING,                  // STRING data type
	"BOOLEANT": TOKEN_TRUE,                    // Boolean
	"BOOLEANF": TOKEN_FALSE,                   // Boolean
	"foreach":  TOKEN_FOREACH,                 // Foreach
	"in":       TOKEN_INSIDE,                  // in
	"null":     TOKEN_NULL,                    // Null
	"jmp":      TOKEN_JUMP,                    // Jump to current lable
	"jump":     TOKEN_JUMP,                    // Jump to current lable
}
