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
// Filename      |  SkyLine_Script_Language_Backend_Models.go
// Project       |  SkyLine programming language
// Line Count    |  800+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the data structures, type structures, type aliases, constant lists, sub functiosn for data types, SLC_Object types, object methods
//                 flag types, settings, configuration files, token types, token keywords, regex patterns, signal patterns, byte/string/bin arrays, results etc for the SkyLine
//                 prime systems. These systems include but are not limited to the Parser, Scanner, Evaluator, Executor, Engine, Scripter, Writer, Reader, Filler, environment, etc.
//
//
package SkyLine_Backend

import (
	"strings"
	"unicode/utf8"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This next section breaks down all of the tokenizer/scanner based data types or data structures which help the scanner categorize specific tokens or keywords
//
//
const (
	TOKEN_DEFAULT         = "DEFAULT"                                // Default keyword    | Implemented
	TOKEN_FOR             = "FOR"                                    // For loop token     | Implemented
	TOKEN_IMPORT          = "IMPORT"                                 // Import             | Implemented
	TOKEN_MODULE          = "module"                                 // Module             | Implemented
	TOKEN_KEYWORD_ENGINE  = "ENGINE"                                 // ENGINE             | Implemented
	TOKEN_ENGINE_TYPE     = "ENGINE::ENVIRONMENT_MODIFIER->CALL:::>" // ENGINE ENV MODIFY  | Implemented
	TOKEN_FOREACH         = "FOREACH"                                // For every element  | Implemented
	TOKEN_INSIDE          = "IN"                                     // Within range       | Implemented
	TOKEN_REGISTER        = "REGISTER"                               // STD LIB Registry   | Implemented
	TOKEN_ILLEGAL         = "ILLEGAL"                                // Illegal character  | Implemented
	TOKEN_EOF             = "EOF"                                    // End Of File        | Implemented
	TOKEN_IDENT           = "TOKEN_IDENT"                            // Identifier         | Implimented
	TOKEN_INT             = "INT"                                    // TYPE integer       | Implemented
	TOKEN_FLOAT           = "FLOAT"                                  // TYPE float         | Implemented
	TOKEN_STRING          = "STRING"                                 // TYPE string        | Implemented
	TOKEN_CONSTANT        = "CONST"                                  // Constant           | Implemented
	TOKEN_FUNCTION        = "FUNCTION"                               // Function           | Implemented
	TOKEN_LET             = "LET"                                    // let statement      | Implemented
	TOKEN_TRUE            = "TOKEN_TRUE"                             // boolean type true  | Implemented
	TOKEN_FALSE           = "FALSE"                                  // boolean type false | Implemented
	TOKEN_IF              = "IF"                                     // If statement       | Implemented
	TOKEN_ELSE            = "ELSE"                                   // Else statement     | Implemented
	TOKEN_RETURN          = "RETURN"                                 // return statement   | Implemented
	TOKEN_SWITCH          = "SWITCH"                                 // Switch statement   | Implemented
	TOKEN_CASE            = "CASE"                                   // Case statement 	 | Implemented
	TOKEN_REGEXP          = "REGEXP"                                 // Regex Type         | Not implemented
	TOKEN_LTEQ            = "<="                                     // LT or equal to     | Implemented
	TOKEN_GTEQ            = ">="                                     // GT or equal to     | Implemented
	TOKEN_ASTERISK_EQUALS = "*="                                     // Multiply equals    | Implemented
	TOKEN_BANG            = "!"                                      // Boolean operator   | Implemented
	TOKEN_ASSIGN          = "="                                      // General assignment | Implemented
	TOKEN_PLUS            = "+"                                      // General operator   | Implemented
	TOKEN_MINUS           = "-"                                      // General operator   | Implemented
	TOKEN_ASTARISK        = "*"                                      // General operator   | Implemented
	TOKEN_SLASH           = "/"                                      // General operator   | Implemented
	TOKEN_LT              = "<"                                      // Boolean operator   | Implemented
	TOKEN_GT              = ">"                                      // Boolean operator   | Implemented
	TOKEN_EQ              = "=="                                     // Boolean operator   | Implemented
	TOKEN_MINUS_EQUALS    = "-="                                     // Minus equals       | Implemented
	TOKEN_NEQ             = "!="                                     // Boolean operator   | Implemented
	TOKEN_DIVEQ           = "/="                                     // Division operator  | Implemented
	TOKEN_PERIOD          = "."                                      // Method Call        | Implemented
	TOKEN_PLUS_EQUALS     = "+="                                     // Plus equals        | Implemented
	TOKEN_COMMA           = ","                                      // Seperation         | Implemented
	TOKEN_SEMICOLON       = ";"                                      // SemiColon          | Implemented
	TOKEN_COLON           = ":"                                      // Colon              | Implemented
	TOKEN_LPAREN          = "("                                      // Args start         | Implemented
	TOKEN_RPAREN          = ")"                                      // Args end           | Implemented
	TOKEN_LINE            = "|"                                      // Line con           | Implemented
	TOKEN_LBRACE          = "{"                                      // Open  f            | Implemented
	TOKEN_RBRACE          = "}"                                      // Close f            | Implemented
	TOKEN_LBRACKET        = "["                                      // Open               | Implemented
	TOKEN_RBRACKET        = "]"                                      // Close              | Implemented
	TOKEN_OROR            = "||"                                     // Condition or or    | Implemented
	TOKEN_ANDAND          = "&&"                                     // Boolean operator   | Implemented
	TOKEN_BACKTICK        = "`"                                      // Backtick           | Implemented
	TOKEN_POWEROF         = "**"                                     // General operator   | Implemented
	TOKEN_MODULO          = "%"                                      // General operator   | Implemented
	TOKEN_NEWLINE         = '\n'                                     // COND               | Implemented
	TOKEN_PLUS_PLUS       = "++"                                     // Plus Plus          | Not implemented
	TOKEN_QUESTION        = "?"                                      // Question que       | Not implemented
	TOKEN_DOTDOT          = ".."                                     // Range              | Not implemented
	TOKEN_CONTAINS        = "~="                                     // Contains           | Not implemented
	TOKEN_NOTCONTAIN      = "!~"                                     // Boolean operator   | Not implemented
	TOKEN_MINUS_MINUS     = "--"                                     // Minus minus        | Not Implemented
)

var keywords = map[string]Token_Type{
	"Func":     TOKEN_FUNCTION,       // Function
	"function": TOKEN_FUNCTION,       // Function
	"let":      TOKEN_LET,            // Variable declaration let
	"set":      TOKEN_LET,            // Variable declaration set
	"cause":    TOKEN_LET,            // Variable declaration cause
	"allow":    TOKEN_LET,            // Variable declaration allow
	"true":     TOKEN_TRUE,           // Boolean true
	"false":    TOKEN_FALSE,          // Boolean false
	"if":       TOKEN_IF,             // Conditional start
	"else":     TOKEN_ELSE,           // Conditional alternative
	"return":   TOKEN_RETURN,         // Return decl
	"ret":      TOKEN_RETURN,         // Return decl
	"const":    TOKEN_CONSTANT,       // Constant type
	"constant": TOKEN_CONSTANT,       // Constant type
	"switch":   TOKEN_SWITCH,         // Switch statement
	"sw":       TOKEN_SWITCH,         // Switch statement
	"case":     TOKEN_CASE,           // Case statement
	"cs":       TOKEN_CASE,           // Case statement
	"default":  TOKEN_DEFAULT,        // Switch alternative
	"df":       TOKEN_DEFAULT,        // Switch alternative
	"register": TOKEN_REGISTER,       // Register
	"ENGINE":   TOKEN_KEYWORD_ENGINE, // Engine caller
	"import":   TOKEN_IMPORT,         // Import data
	"for":      TOKEN_FOR,            // For loop
	"STRING":   TOKEN_STRING,         // STRING data type
	"BOOLEANT": TOKEN_TRUE,           // Boolean
	"BOOLEANF": TOKEN_FALSE,          // Boolean
	"foreach":  TOKEN_FOREACH,        // Foreach
	"in":       TOKEN_INSIDE,         // in

}

type Token_Type string

type Token struct {
	Token_Type Token_Type
	Literal    string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section is seperated to be specific for the SkyLine scanner or lexer, these structures also help implement and keep track of current tokens being parsed
//
//

type ScannerInterface interface {
	NT() Token
}

type ScannerStructure struct {
	CharInput string
	POS       int
	RPOS      int
	Char      byte
	Chars     []rune
	PrevTok   Token
	Prevch    byte
	CurLine   int // Current line
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section of the models file is segmented to be specific for the skyline objects definitions. These structures and constant values help define the types and objects
//
// given that skyline is completely object orritented. This also defines all of the data type and keyword structures for the Objects such as switch, case, if, else, elif, func etc
//
type Type_Object string

// Data types
const (
	ImportingType   = "Import"
	IntegerType     = "Integer"     // integger
	FloatType       = "Float"       // float integer
	BooleanType     = "Boolean"     // bool true or bool false
	NilType         = "Nil"         // null
	ReturnValueType = "ReturnValue" // return
	ErrorType       = "Error"       // error
	FunctionType    = "Function"    // function
	StringType      = "String"      // string
	BuiltinType     = "Builtin"     // built in
	ArrayType       = "Array"       // array
	HashType        = "Hash"        // hash
	RegisterType    = "Registry"    // Register backend
)

type Switch struct {
	Token   Token
	Value   Expression
	Choices []*Case
}

type Case struct {
	Token Token
	Def   bool
	Expr  []Expression
	Block *BlockStatement
}

type SLC_Object interface {
	SL_RetrieveDataType() Type_Object
	SL_InspectObject() string
	InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object
	ToInterface() interface{}
}

type HashKey struct {
	Type_Object Type_Object
	Value       uint64
}

type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value int64
}

type Float struct {
	Value float64
}

type Boolean_Object struct {
	Value bool
}

type String struct {
	Value  string
	Offset int
}

// Offset for string reset
func (str *String) Reset() {
	str.Offset = 0
}

type Nil struct{}

type ReturnValue struct {
	Value SLC_Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*Ident
	Body       *BlockStatement
	Env        *Environment_of_environment
}

type BuiltinFunction func(env *Environment_of_environment, args ...SLC_Object) SLC_Object

type Constant struct {
	Token Token
	Name  *Ident
	Value Expression
}

type ForeachStatement struct {
	Token Token
	Index string
	Ident string
	Value Expression
	Body  *BlockStatement
}

type Builtin struct {
	Fn BuiltinFunction
}

type Array struct {
	Elements []SLC_Object
	Offset   int
}

// Offset reset for array structure
func (array *Array) Reset() {
	array.Offset = 0
}

type HashPair struct {
	Key   SLC_Object
	Value SLC_Object
}

type Hash struct {
	Pairs  map[HashKey]HashPair
	Offset int
}

// Offset for hash reset
func (hash *Hash) Reset() {
	hash.Offset = 0
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines the AST generat functions and interfaces, this allows us to create nodes, statements, expressions and other various types for the AST. This also defines
//
// functions and data type structures such as statement and expression type structures which the AST uses.
//
var U UserInterpretData

type Node interface {
	SL_ExtractNodeValue() string
	SL_ExtractStringValue() string
}

type Statement interface {
	Node
	SN()
}

type Expression interface {
	Node
	EN()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token Token
	Name  *Ident
	Value Expression
}

type AssignmentStatement struct {
	Token    Token
	Name     *Ident
	Value    Expression
	Operator string
}

type Ident struct {
	Token Token
	Value string
}

type ReturnStatement struct {
	Token       Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

type IntegerLiteral struct {
	Token Token
	Value int64
}

type FloatLiteral struct {
	Token Token
	Value float64
}

type PrefixExpression struct {
	Token    Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean_AST struct {
	Token Token
	Value bool
}

type ConditionalExpression struct {
	Token       Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type BlockStatement struct {
	Token      Token
	Statements []Statement
}

type FunctionLiteral struct {
	Token      Token
	Parameters []*Ident
	Body       *BlockStatement
}

type CallExpression struct {
	Token     Token      // the '(' token
	Function  Expression // Ident or FunctionLiteral
	Arguments []Expression
}

type StringLiteral struct {
	Token Token
	Value string
}

type ArrayLiteral struct {
	Token    Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token Token // the '[' token
	Left  Expression
	Index Expression
}

type PostfixExpression struct {
	Token    Token
	Operator string
}

type HashLiteral struct {
	Token Token // the '{' token
	Pairs map[Expression]Expression
}

type RegisterValue struct {
	Value SLC_Object
}

type Register struct {
	Token         Token
	RegistryValue Expression
}

type ENGINE_Value struct {
	Value SLC_Object
}

type ENGINE struct {
	Token       Token
	EngineValue Expression
}

type EngineCallValues struct {
	Name        string
	Version     string
	Require     []string
	Languages   string
	Description string
	SOS         string
	Prepped     bool // Engine has parsed data
}

type Module struct {
	Name  string
	Attrs SLC_Object
}

type ImportExpression struct {
	Token Token
	Name  Expression
}

type ForLoopExpression struct {
	Token       Token
	Condition   Expression
	Consequence *BlockStatement
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines the environment structures for the skyline environment
//
//
//
type Environment_of_environment struct {
	Store  map[string]SLC_Object
	Outer  *Environment_of_environment
	ROM    map[string]bool // Read Only mode for constants
	permit []string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines the most basic evaluation models which are used
//
//
//
var (
	NilValue   = &Nil{}
	TrueValue  = &Boolean_Object{Value: true}
	FalseValue = &Boolean_Object{Value: false}
)

type ObjectCallExpression struct {
	Token      Token
	SLC_Object Expression
	Call       Expression
}

var FileCurrent FileCurrentWithinParserEnvironment

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines all of the parser functions, types, constants, variables and maps which are used to communicate with other functions within Skyline. This also allows it to
//
// work with the tokens, types, prefix's, infix's and other various statements and expressions.
//
const (
	_ int = iota
	LOWEST
	ASSIGNMENT
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	INDEX
	GTEQS
	LTEQS
	POWER
	MOD
	DOT_DOT
	REGEXP_MATCH
	TERNARY
	COND
)

var Precedences = map[Token_Type]int{
	TOKEN_ANDAND:          COND,
	TOKEN_OROR:            COND,
	TOKEN_EQ:              EQUALS,
	TOKEN_GTEQ:            GTEQS,
	TOKEN_LTEQ:            LTEQS,
	TOKEN_NEQ:             EQUALS,
	TOKEN_LT:              LESSGREATER,
	TOKEN_GT:              LESSGREATER,
	TOKEN_PLUS:            SUM,
	TOKEN_PLUS_EQUALS:     SUM,
	TOKEN_MINUS:           SUM,
	TOKEN_MINUS_EQUALS:    SUM,
	TOKEN_SLASH:           PRODUCT,
	TOKEN_ASSIGN:          ASSIGNMENT,
	TOKEN_POWEROF:         POWER,
	TOKEN_QUESTION:        TERNARY,
	TOKEN_ASTARISK:        PRODUCT,
	TOKEN_ASTERISK_EQUALS: PRODUCT,
	TOKEN_DIVEQ:           PRODUCT,
	TOKEN_LPAREN:          CALL,
	TOKEN_PERIOD:          CALL,
	TOKEN_NOTCONTAIN:      REGEXP_MATCH,
	TOKEN_CONTAINS:        REGEXP_MATCH,
	TOKEN_DOTDOT:          DOT_DOT,
	TOKEN_LBRACKET:        INDEX,
	MODULECALL:            INDEX,
}

type (
	PrefixParseFn  func() Expression
	InfixParseFn   func(Expression) Expression
	PostfixParseFn func() Expression
)

type Parser struct {
	Lex             *ScannerStructure
	Errors          []string
	PreviousToken   Token
	CurrentToken    Token
	PeekToken       Token
	PrefixParseFns  map[Token_Type]PrefixParseFn
	InfixParseFns   map[Token_Type]InfixParseFn
	PostfixParseFns map[Token_Type]PostfixParseFn
}

//////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////// ENVIRONMENT MODIFIER
//////////////////////////////////////////////////////////////////////////////////////////////////

type Iterable interface {
	Reset()
	Next() (SLC_Object, SLC_Object, bool)
}

func (env *Environment_of_environment) Names(prefix string) []string {
	var ret []string
	for key := range env.Store {
		if strings.HasPrefix(key, prefix) {
			ret = append(ret, key)
		}
		if strings.HasPrefix(key, "object.") {
			ret = append(ret, key)
		}
	}
	return ret
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines TOINTERFACE functions which are used to take a given key or type structure and convert it to an interface type.
//
//
//

func (ENGINE_VAL *ENGINE_Value) ToInterface() interface{} {
	return "<ENGINE_MODIFICATION : SkyLine External Environment Modifier>"
}

func (Arr *Array) ToInterface() interface{} {
	return "<ARRAY>"
}

func (hash *Hash) ToInterface() interface{} {
	return "<HASH>"
}

func (null *Nil) ToInterface() interface{} {
	return "<NULL>"
}

func (Func *Function) ToInterface() interface{} {
	return "<FUNCTION>"
}

func (Err *Error) ToInterface() interface{} {
	return "<ERROR>"
}

func (Mod *Module) ToInterface() interface{} {
	return "<Module>"
}

func (BuiltIn *Builtin) ToInterface() interface{} {
	return "<BUILT-IN-FUNCTION>"
}

func (Ret *ReturnValue) ToInterface() interface{} {
	return "<RETURN_VALUE>"
}

func (float *Float) ToInterface() interface{} {
	return float.Value
}

func (Int *Integer) ToInterface() interface{} {
	return Int.Value
}

func (Str *String) ToInterface() interface{} {
	return Str.Value
}

func (Boolean *Boolean_Object) ToInterface() interface{} {
	return Boolean.Value
}

func (RegisterValue *RegisterValue) ToInterface() interface{} {
	return "<REGISTRY>"
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// The NEXT functions which is the next brick of code which defines functions to check and itterate over specfic elements or values of a specific data type that can be itterated
//
// over such as String, Hashes or arrays.
//
//

func (Arr *Array) Next() (SLC_Object, SLC_Object, bool) {
	if Arr.Offset < len(Arr.Elements) {
		Arr.Offset++
		element := Arr.Elements[Arr.Offset-1]
		return element, &Integer{Value: int64(Arr.Offset - 1)}, true
	}
	return nil, &Integer{Value: 0}, false
}

func (hash *Hash) Next() (SLC_Object, SLC_Object, bool) {
	if hash.Offset < len(hash.Pairs) {
		idx := 0
		for _, pair := range hash.Pairs {
			if hash.Offset == idx {
				hash.Offset++
				return pair.Key, pair.Value, true
			}
			idx++
		}
	}
	return nil, &Integer{Value: 0}, false
}

func (str *String) Next() (SLC_Object, SLC_Object, bool) {
	if str.Offset < utf8.RuneCountInString(str.Value) {
		str.Offset++
		chars := []rune(str.Value)
		value := &String{Value: string(chars[str.Offset-1])}
		return value, &Integer{Value: int64(str.Offset - 1)}, true
	}
	return nil, &Integer{Value: 0}, false
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// The following code unerneath each of these sections defines all of the identifiers and libraries currently avalible for the SkyLine programming language. These maps and variables
//
// are used in the context of register which allows you to `register` all of the functions from one standard library into the environment of the programming language or the environment
//
// which the interpreter or SL engine has started. These also verify that they are able to be imported and are not currently unavalible.
//

var ConstantIdents = map[string]bool{
	"math.abs":                          true,
	"math.cos":                          true,
	"math.sin":                          true,
	"math.sqrt":                         true,
	"math.tan":                          true,
	"math.cbrt":                         true,
	"math.rand":                         true,
	"math.out":                          true,
	"crypt.hash":                        true,
	"io.input":                          true,
	"io.clear":                          true,
	"io.box":                            true,
	"io.restore":                        true,
	"io.listen":                         true,
	"fpng.new":                          true,
	"ForensicsPng.PngChunkCount":        true, // Chunk count of image
	"ForensicsPng.PngSettingsNew":       true, // Forensics PNG New() method to create new settings
	"ForensicsPng.PngInject":            true, // Forensics PNG PNG FORMAT injection
	"ForensicsPng.InjectRegular":        true, // Forensics PNG regular file injection
	"ForensicsUtils.FindSigUnknownFile": true, // Forensics Utils Find filetype by signature
	"ForensicsUtils.CheckZIPSig":        true, // Forensics Utils check ZIP signature
	"ImageUtils.CreateImage":            true, // Forensics utils create image
	"ImageUtils.CreationNew":            true, // Image Utils Creation new
	"ImageUtils.InjectImage":            true, // Image Utils inject image
	"http.Get":                          true,
	"http.MethodGet":                    true,
	"http.MethodDelete":                 true,
	"http.MethodPut":                    true,
	"http.MethodPatch":                  true,
	"http.MethodOptions":                true,
	"http.MethodHead":                   true,
	"http.MethodPost":                   true,
	"http.MethodTrace":                  true,
	"env.Getenv":                        true,
	"env.Setenv":                        true,
	"env.Environment":                   true,
	"Google.CastDeviceInfo":             true,
	"Google.CastDeviceReboot":           true,
	"Google.CastDeviceDescription":      true,
	"Google.CastDeviceWiFiForget":       true,
	"Google.CastDeviceWiFiScan":         true,
	"Google.CastDeviceWiFiScanResults":  true,
	"Google.CastDeviceWiFiConfigured":   true,
	"Google.CastDeviceAlarms":           true,
	"Google.CastDeviceTimezones":        true,
	"Google.CastDeviceLegacyConf":       true,
	"Google.CastDeviceBleStat":          true,
	"Google.CastDeviceBlePaired":        true,
	"Google.CastDeviceSetName":          true,
	"Google.CastDeviceApplications":     true,
	"Roku.KeyPressHome":                 true,
	"Roku.KeyPressPlay":                 true,
	"Roku.KeyPressUp":                   true,
	"Roku.KeyPressDown":                 true,
	"Roku.KeyPressLeft":                 true,
	"Roku.KeyPressRight":                true,
	"Roku.KeyPressSelect":               true,
	"Roku.KeyPressRewind":               true,
	"Roku.KeyPressFFW":                  true,
	"Roku.KeyPressOptions":              true,
	"Roku.KeyPressPause":                true,
	"Roku.KeyPressBack":                 true,
	"Roku.KeyPressPoweroff":             true,
	"Roku.KeyPressVolumeUp":             true,
	"Roku.KeyPressVolumeDown":           true,
	"Roku.KeyPressVolumeMute":           true,
	"Roku.KeyPressDeviceDown":           true,
	"Roku.KeyPressDeviceUp":             true,
	"Roku.KeyPressDevAppLaunch":         true,
	"Roku.KeyPressDevAppInstall":        true,
	"Roku.KeyPressDevDisableSGR":        true,
	"Roku.KeyPressDevEnableSGR":         true,
	"Roku.KeyPressDevTV":                true,
	"Roku.KeyPressDevSGNODES":           true,
	"Roku.KeyPressDevActiveTVS":         true,
	"Roku.KeyPressDevDial":              true,
	"Roku.DeviceBrowse":                 true,
	"Roku.DeviceInformation":            true,
	"Roku.DeviceApplications":           true,
	"Roku.DeviceActiveApplications":     true,
	"Apple.AirPlayServerInfo":           true,
	"Apple.AirPlayPlayBackInfo":         true,
	"Apple.AirPlayScrubInfo":            true,
	"Apple.AirPlayStreamInfo":           true,
	"Apple.AirPlayInfo":                 true,
	"Apple.DAAPMain":                    true,
	"Apple.DAAPLogin":                   true,
	"Apple.DAAPDatabase":                true,
	"Amazon.TvDeviceInformation":        true,
	"Amazon.TvDeviceDescription":        true,
	"file.Open":                         true,
	"file.New":                          true,
	"file.Overwrite":                    true,
	"file.Write":                        true,
}

var StandardLibNames = map[string]bool{
	"math":                true,
	"io":                  true,
	"forensics":           true,
	"os":                  true,
	"http":                true,
	"crypt":               true,
	"IoT/Apple/Database":  true,
	"IoT/Roku/Database":   true,
	"IoT/Google/Database": true,
	"IoT/Amazon/Database": true,
	"forensics/Utils":     true,
	"forensics/Apple":     true,
	"forensics/Sub":       true,
	"json":                true,
	"net":                 true,
	"xml":                 true,
	"env":                 true,
	"file":                true,
	"forensics/PNG":       true,
}

/*

 */

var RegisterStandard = map[string]func(){
	"io":                  RegisterIO,
	"math":                RegisterMath,
	"crypt":               RegisterCrypt,
	"http":                RegisterHTTP,
	"env":                 RegisterEnvironment,
	"file":                RegisterFile,
	"forensics/PNG":       RegisterForensicsPNG,
	"forensics/Apple":     RegisterForensicsApple,
	"forensics/Utils":     RegisterImageUtilitiesForensics,
	"forensics/Sub":       RegisterForensicsSubFunctions,
	"IoT/Apple/Database":  RegisterAppleIoTDatabase,
	"IoT/Roku/Database":   RegisterRokuIoTDatabase,
	"IoT/Google/Database": RegisterGoogleIoTDatabase,
	"IoT/Amazon/Database": RegisterAmazonIoTDatabase,
}

var Datatypes = []string{
	"string.",
	"float.",
	"object.",
	"hash.",
	"array.",
	"boolean.",
}

var Static_String_Methods = []string{
	"Length",         // Length of string
	"Methods",        // Methods
	"Ord",            // ord
	"Integer",        // to integer
	"Float",          // to float
	"Boolean",        // To byte
	"Upper",          // To uppercase
	"Lower",          // To lowercase
	"Title",          // To title
	"Split",          // Split
	"Trim",           // Trim
	"UnlinkRegistry", // Unlink a libraries registration | This is a string because we unlink a library name
	"View",
}

var Static_Integer_Methods = []string{
	"chr",
	"Methods", // Grab all methods
	"View",    // View values
}

var Static_Float_Methods = []string{
	"Methods", // Grab all methods
	"View",
}

var Static_Array_Methods = []string{
	"Reverse", // Reverse
	"Append",  // Append
	"Copy",    // Copy
	"Swap",    // Swap
	"Less",    // Less
	"Compare", // Compare
	"PopR",    // Pop right
	"PopL",    // Pop left
	"Length",  // Length of the array
	"Methods", // Grab all methods
	"Typeof",  // Typeof
	"View",    // View
}

var Static_Hash_Methods = []string{
	"Keys",    // Dump all the hash's keys
	"Methods", // Grab all methods
	"View",
}

var Static_Boolean_Methods = []string{
	"Methods", // Grab all methods
	"View",
}

var Static_BuiltInFunction_Methods = []string{
	"Methods", // Grab all methods
	"View",
}

// All functions for each library
var MathLibFunctionsRegistered = []string{
	"math.tan",
	"math.cos",
	"math.sin",
	"math.rand",
	"math.abs",
	"math.cbrt",
	"math.sqrt",
}

var HTTPLibFunctionsRegistered = []string{
	"http.Get",
}

var IOLibFunctionsRegistered = []string{
	"io.box",
	"io.listen",
	"io.restore",
	"io.clear",
	"io.input",
}

var CryptLibFunctionsRegistered = []string{
	"crypt.hash",
}

var ForensicsFunctionsRegistered = []string{
	"fpngnew",
	"fpngmeta",
	"fpngPngSettingsNew",
	"fpngInjectPNG",
	"fpngInjectImage",
	"fpngFindSigUnknownFile",
	"fpngCheckZIPSig",
	"fpngCreationNew",
	"fpngInjectRegular",
	"fpngCreateImage",
}

var EnvironmentFunctionsRegistered = []string{
	"env.setenv",
	"env.getenv",
	"env.environment",
}

type UserInterpretData struct {
	OperatingSystem             string
	OperatingSystemArchitecture string
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines all of the script information which works with the variables, the flags, command line details, frontend work, art work etc etc. These are all really
//
// variables to store flags, version numbers or art.
//

var Stars = `
				*      *
			   * 		*
					*
				*				*
			*	     *
						*

				*			*
						*
`

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
// This code section defines all models for the file system within SkyLine. Of course when using file system in this context we are not talking about direct FS (/, and \\) rather
//
// a system to manage files, allowed file types, file scanning, signature scanning, file verification, integrity checks and much more.
//

type FileCurrentWithinParserEnvironment struct {
	Filename      string
	FileLocation  string
	FileExtension string
	FileBasename  string
	IsDir         bool
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
// This section defines the colors and the specific hex code for the frontend, this is also in a seperate section due to it being colors and a constant list.
//
//
//

const (
	SUNRISE_HIGH_DEFINITION         = "\033[38;5;214m"
	SUNRISE_LIGHT_DEFINITION        = "\033[38;5;215m"
	SKYLINE_HIGH_DEFPURPLE          = "\033[38;5;57m"
	SKYLINE_HIGH_DEFAQUA            = "\033[38;5;51m"
	SKYLINE_HIGH_DEFRED             = "\033[38;5;196m"
	SKYLINE_SUNRISE_HIGH_DEF_YELLOW = "\033[38;5;190m"
	SKYLINE_HIGH_DEFWARN            = "\033[38;5;213m"
	SKYLINE_HIGH_FIXBLUE            = "\033[38;5;121m"
	SKYLINE_SICK_BLUE               = "\033[38;5;81m"
	SKYLINE_HIGH_RES_VIS_RED        = "\033[38;5;196m"
	SKYLINE_HIGH_RES_VIS_BLUE       = "\033[38;5;122m"
	SKYLINE_HIGH_RES_VIS_SUNSET     = "\033[38;5;217m"
	SKYLINE_WHITE                   = "\033[38;5;249m"
	SKYLINE_RESTORE                 = "\033[39m"
)
