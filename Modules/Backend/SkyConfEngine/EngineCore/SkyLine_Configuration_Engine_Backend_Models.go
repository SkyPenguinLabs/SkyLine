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
//
//
// Contains -> This file contains most of the models or data type structures for the language, each section is split into its own reason.
package SkyLine_Configuration_Engine_Backend_Source

import "regexp"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x01]
//
// This next section defines all constants used within the engine and used within the main code block. These lists enclude token types, token integers, int reps,
//
// keywords, values and any other form of variable that can be worked with but does not include VAR lists
//

const (
	ASSIGN_Token             = "="            // Assignment
	MODIFY_Token             = "->"           // Modifier
	SEMICOLON_Token          = ";"            // Break
	COMMA_Token              = ","            // List def
	LPAREN_Token             = "("            // Start compression
	RPAREN_Token             = ")"            // End compression
	LBRACE_Token             = "{"            // Start block
	RBRACE_Token             = "}"            // End block
	LBRACKET_Token           = "["            // Start array
	RBRACKET_Token           = "]"            // End array
	SET_Token                = "SET"          // Assignment
	CONSTANT_Token           = "CONSTANT"     // Constant def
	BOOL_TRUE_Token          = "TRUE"         // Boolean -> true
	BOOL_FALSE_Token         = "FALSE"        // Boolean -> false
	INTEGER_Token            = "INTEGER"      // Integer token
	STRING_Token             = "STRING"       // String token
	ENGINE_Token             = "ENGINE"       // Engine block definition
	ENGINE_INITATE_Token     = "INIT"         // Engine initate parse definition
	END_OF_FILE              = "EOF"          // End of the file has been reached
	ALIENATED_UNVERSED       = "ILLEGAL_TYPE" // Illegal or unsupported data type, this calls alienated due to it not being in bounds of the engine
	IDENTIFIER               = "IDENT"        // Identifier
	OBJECT_NULL              = "NULL"         // Object | Null
	OBJECT_INTEGER           = "INTEGER"      // Object | Integer
	OBJECT_STRING            = "STRING"       // Object | String
	OBJECT_BOOLEAN           = "BOOLEAN"      // Object | Boolean
	OBJECT_ARRAY             = "ARRAY"        // Object | Array
	OBJECT_BUILT_IN_FUNCTION = "Builtin"      // Object | Built-in function
	OBJECT_ERROR             = "ERROR"        // Object | Error
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x02]
//
// This next section contains token data for the scanner as well as constant values for the evaluator and parser which are used to better categorize
//
// and work with individual tokens that work within the configuration language. Note that this is a regex engine so it works a bit differently and is not a standard
//
// lexer or scanner

type TokenDataType string

type TokenRegistry struct {
	TokenDataType TokenDataType
	Literal       string
}

type ScannnerInterface interface {
	NT() TokenRegistry
}

type ScannerStructureRegister struct {
	CharInput        string
	POS              int
	RPOS             int
	Char             byte
	Chars            []rune
	PreviousRegistry TokenRegistry
	PrevCh           byte
}

var ScannerTokenizationRegularExpressions = map[*regexp.Regexp]TokenDataType{
	regexp.MustCompile(`^=$`):  ASSIGN_Token,
	regexp.MustCompile(`^;$`):  SEMICOLON_Token,
	regexp.MustCompile(`^\($`): LPAREN_Token,
	regexp.MustCompile(`^\)$`): RPAREN_Token,
	regexp.MustCompile(`^,$`):  COMMA_Token,
	regexp.MustCompile(`^{`):   LBRACE_Token,
	regexp.MustCompile(`^}`):   RBRACE_Token,
	regexp.MustCompile(`^\[$`): LBRACKET_Token,
	regexp.MustCompile(`^\]$`): RBRACKET_Token,
}

var KeyMap = map[string]TokenDataType{
	"set":      SET_Token,
	"true":     BOOL_TRUE_Token,
	"false":    BOOL_FALSE_Token,
	"ENGINE":   ENGINE_Token,
	"INIT":     ENGINE_INITATE_Token,
	"constant": CONSTANT_Token,
	"const":    CONSTANT_Token,
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x3]
//
// The code below this code brick will define all AST based data structures and representations for the data structures and AST nodes. This will include tree reps
//
// leaf node representations, statements and structures used and so on from here. These type's are also going to be used when working with the evaluator

type AbstractSyntaxTree_Node interface {
	TokenConstructLiteral() string
	TokenConstructToString() string
}

type AbstractSyntaxTree_Statement interface {
	AbstractSyntaxTree_Node
	SatatementLeafNode()
}

type AbstractSyntaxTree_Expression interface {
	AbstractSyntaxTree_Node
	ExpressionLeafNode()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x04]->[0x03]
//
// This brick ties into section 0x3 with the definitions or data structures here and representations being apart of the AST's development phase where it defines
//
// statement and expression representations as well as the proper leaf nodes to do so. These type structures will only be Expressions NOT statements

type ENGINE_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Value         AbstractSyntaxTree_Expression
	SubUnits      []*INIT_Expression_AbstractSyntaxTree
} //ENGINE
type INIT_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Default       bool
	Expression    []AbstractSyntaxTree_Expression
	Sub_UNIT      *BlockStatement_Statement_AbstractSyntaxTree
} //INIT
type Identifier_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Value         string
} // IDENT::
type BooleanDataType_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Value         bool
} // Boolean::TRUE->FALSE
type StringDataType_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Value         string
} // String::""
type IntegerDataType_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Value         int64
} // Integer::1-9

type PrefixExpression_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Operator      string
	Right         AbstractSyntaxTree_Expression
} // Prefix::
type InfixExpression_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Left          AbstractSyntaxTree_Expression
	Operator      string
	Right         AbstractSyntaxTree_Expression
} // Inffix::
type CallFunction_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Function      AbstractSyntaxTree_Expression
	Arguments     []AbstractSyntaxTree_Expression
} // CallFunction::A-Z..a-z()
type ArrayLiteral_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Elements      []AbstractSyntaxTree_Expression
} // Array::[]
type IndexLit_Expression_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Left          AbstractSyntaxTree_Expression
	Index         AbstractSyntaxTree_Expression
} // Index::[1..INF]

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x5]::[0x04]->[0x03]
//
//
// This next brick ties into 0x04 and 0x03 due to this being the second part of 0x03 where the type definitions are all about statement nodes rather than
//
// expressions like the previous brick was about

type Constant_Statement_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Name          *Identifier_Expression_AbstractSyntaxTree
	Value         AbstractSyntaxTree_Expression
} // Constant

type Assignment_Statement_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Name          *Identifier_Expression_AbstractSyntaxTree
	Value         AbstractSyntaxTree_Expression
} // set

type BlockStatement_Statement_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Statements    []AbstractSyntaxTree_Statement
} // {}

type Expression_Statement_AbstractSyntaxTree struct {
	TokenRegister TokenRegistry
	Expression    AbstractSyntaxTree_Expression
} // expression statement

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x6]:::[0x5]::[0x04]->[0x03]
//
//
// This brick ties into the AST bricks and will define the program function and structure to parse the programs statements and

type Engine_Prog struct {
	Statements []AbstractSyntaxTree_Statement
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x7]
//
//
// This brick will contain all the type structures and data types required to define an object. SkyLine as a language is an Object Orriented programming language
//
// Everything inside of SkyLine is an object which means that the engine should also be using objects. Which is why this section will be dedicated to interfaces,
//
// definitions and other forms for the Object side of the language. The code under this brick will contain base functions for the objects types

var (
	NULL  = &ObjectNULL{}
	TRUE  = &ObjectBoolean{Value: true}
	FALSE = &ObjectBoolean{Value: false}
)

const (
	NULL_OBJ    = "NULL"
	ERROR_OBJ   = "ERROR"
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	STRING_OBJ  = "STRING"
	BUILTIN_OBJ = "BUILTIN"
	ARRAY_OBJ   = "ARRAY"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x8]->[0x7]
//
//
// This brick will contain all the type structures and data types required to define an object. SkyLine as a language is an Object Orriented programming language
//
// Everything inside of SkyLine is an object which means that the engine should also be using objects. Which is why this section will be dedicated to interfaces,
//
// definitions and other forms for the Object side of the language. The code under this brick will contain base functions for the objects types

type ObjectsDataType string

type SLC_Object interface {
	ObjectDataType() ObjectsDataType
	ObjectInspectFunc() string
}

type HashingKey struct {
	Type  ObjectsDataType
	Value uint64
}

type HashAbleInterface interface {
	GrabHashKey() HashingKey
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x9]::[0x8]->[0x7]
//
//
// This brick will contain all the type structures and data types required to define an object. SkyLine as a language is an Object Orriented programming language
//
// Everything inside of SkyLine is an object which means that the engine should also be using objects. Which is why this section will be dedicated to interfaces,
//
// definitions and other forms for the Object side of the language. The code under this brick will contain base functions for all data types in SkylineConfig

type BuiltinFunction func(args ...SLC_Object) SLC_Object

type ObjectString struct {
	Value string
}

type ObjectInteger struct {
	Value int64
}

type ObjectBoolean struct {
	Value bool
}

type ObjectArray struct {
	Elements []SLC_Object
}

type ObjectNULL struct{}

type ObjectERROR struct {
	Message string
}

type ObjectBUILTINFUNCTION struct {
	Function BuiltinFunction
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x10]
//
// The next section of code will define all functions and types for the parser that will allow for functions such as prefix and infix expressions to be properly
//
//  parsed and called when the parsers main function is called. This allows us to work with the actual sections of the parser based on specific indexes.
//
//

// PRECS
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
	ARROW
)

var ParserPrecedences = map[TokenDataType]int{
	LPAREN_Token: CALL,
	LBRACE_Token: INDEX,
	MODIFY_Token: ARROW,
}

type (
	PrefixParserFunctions func() AbstractSyntaxTree_Expression
	InfixParserFunctions  func(AbstractSyntaxTree_Expression) AbstractSyntaxTree_Expression
)

type SLC_Parser_Structure struct {
	Scanner              ScannnerInterface
	EngineErrors         []string
	CurrentToken         TokenRegistry
	PeekToken            TokenRegistry
	PrefixParseFunctions map[TokenDataType]PrefixParserFunctions
	InfixParseFunctions  map[TokenDataType]InfixParserFunctions
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x10A]
//
// This code section contains all models, variables and constants related to the evaluator and system step to finalize and execute the code
//
//

var Builtins = map[string]*ObjectBUILTINFUNCTION{}

func RegisterBuiltin(name string, fun BuiltinFunction) {
	Builtins[name] = &ObjectBUILTINFUNCTION{Function: fun}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x10B]
//
// This code section defines the environment related structures
//
//

type Engine_Environment_Of_Environment struct {
	StoreObj    map[string]SLC_Object
	ReadOnly    map[string]bool
	EngineOuter *Engine_Environment_Of_Environment
	PermitMod   []string
}
