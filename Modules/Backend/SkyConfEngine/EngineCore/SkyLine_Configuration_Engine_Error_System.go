////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
//
// The SkyLine configuration language is a language and engine designed to act as a modification language to the SkyLine programming language. This language is
// very minimal and contains a regex base lexer, a very basic parser, a few keywords, a base interpreter and that is all as well as some backend engine code. This
// language is purely modified to be an extension to the SkyLine programming language, something that can be a pre processor language post processing for the main
// SkyLine script. Below is more technical information for the language
//
// Lexer       : Regex based lexer with minimal tokens and base syntax
// Parser      : Base parser with minimal tokens and base syntax with simple error systems
// REPL        : Does not exist
// Environment : Extremely minimal
// Types       : String, Boolean, Integer
// Statements  : set, import, use, errors, output, system, constant/const
// Filename    : SkyLine_Configuration_Engine_Error_System.go
//
//
// Contains -> This file defines the entire error syetem and every error for the SkyLine Configuration Language. THis error system is defined below in terms of understanding
// as well as codes, type systems, type checking and much more among that list.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// The SkyLine configuration language/engine is a very unique engine of its kind, it does not use standard error systems as it makes its own, anytime there is an error
//
// there are unique codes put in place to do this. Now, working with SLC given how small it is means there is only a small amount of codes and errors that can happen or
//
// that can exist within this engine. We can give our standard range of 1-100 for standard errors, 100-200 for parser errors, 200-300 for scanner errors, 300-400 for evaluator
//
// errors and anything above is a developer error code. Below the chart is shown for categorization
//
//
//
// ---- CHART ( error codes ) -----
//
// 1-100      | Standard errors, files, ignoring, flags, config etc
// 100-200    | parser errors, failed symbols, non prefix parsing functions etc
// 200-300    | scanner errors, false token, illegal token, illegal syntax, error during parse etc
// 300-400    | evaluator token, expecting, missing semicolon, missing syntax, missing symbol, broken tag etc
// 400-10,000 | developer errors such as missing code in map, missing argument call, panic or fatal error etc etc
//
//
// The below tags all define the design for the error system for the configuration engine. This will go by color codes or definitions etc etc.
//
//---- DESIGN ----
//
// The design for the error system is quite simple, it will show the error, the code, the filename and more. The example is shown below
//
// - [E]:(Engine)->[1]
//          |
//      	|[E]: missing semicolon
//          |
//          |[S]: set x := 10
//          |
//          |[F]:
//
// The top row will be red with grey where grey will highlight the engine, the number at the end represents the code
//
// further down you have E which is the error message, below the statement and the final statement which will be a possible fix
//
//
package SkyLine_Configuration_Engine_Backend_Source

const (
	SLC_FileSystem_NULL_FIELDS                     = "10"  // File does not have anything in it NULL
	SLC_FileSystem_MissingEngineCode               = "20"  // before iniation you must define one engine(true) block and one init(true) block
	SLC_FileSystem_ErrorWhenOpeningOrLoadingFile   = "30"  // Could not load the file due to
	SLC_Parser_ExpectRightBracket                  = "200" // Missing '}
	SLC_Parser_ExpectLeftBracket                   = "205" // Missing '{
	SLC_Parser_ExpectLeftParen                     = "210" // Missing '(
	SLC_Parser_ExpectRightParen                    = "215" // Missing ')
	SLC_Parser_ExpectSemicolom                     = "220" // Missing ';
	SLC_Parser_OnlyOneDefaultParseInEngineINIT     = "225" // The ENGINE (INIT) statement should only have one DEF call to it
	SLC_Parser_Unterminated_IniationStatement      = "230" // Unterminated INIT statement
	SLC_Parser_IntegerOverflowParsingError         = "235" // Integer parsing overflow
	SLC_Parser_NoPrefixParsingFunctionFound        = "240" // No prefix parser function found
	SLC_Parser_UnterminatedConstantStatement       = "245" // Unterminated constant statement
	SLC_Parser_ExpectedTokenButGotSoOnStatement    = "250" // Expected next token to be ''... but got ''...
	SLC_Evaluator_Unknown_Token_Operator           = "300" // Unknown token / operator
	SLC_Evaluator_DataType_Mismatch                = "310" // Non compatible data type
	SLC_Evaluator_Identifier_Not_Found             = "320" // Ident not found
	SLC_Evaluator_Array_Index_Operator_Unsupported = "330" // Array index expression false
	SLC_Evaluator_ValueCall_Not_A_Function         = "340" // Not a function
	SLC_Evaluator_ATTEMPT_TO_MOD_CONSTANT          = "350" // Attempting to modify a constant value
	SLC_Evaluator_FAULT_WHEN_PARSING               = "360" // The engine encountered a segment fault when parsing the left or right values of ->.
)

func CallErrorStr(Code, Message, Statement string) string {
	var cerr string
	cerr += "\033[38;5;242m- \033[38;5;196m[E]:\033[38;5;242m(\033[38;5;209mEngine\033[38;5;242m)\033[39m \033[38;5;242m| " + Code
	cerr += "\n          |"
	cerr += "\n          |\033[38;5;196m[E]: " + Message + "\033[38;5;242m"
	cerr += "\n          |     "
	cerr += "\n          |\033[38;5;86m[S]: " + Statement + "\033[38;5;242m"
	return cerr
}
