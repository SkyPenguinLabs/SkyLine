///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_EnvironAndObject_Models
// Extension         | .go ( golang source code file )
// Purpose           | Define structures, types, interfaces, alliases, constants and values for the SkyLine environment
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file is like any other file, it defines the objects within the file, defines the type structures and alliases for each object
//
// as well as defines constants and object interfaces. Currently SkyLine only supports basic types which are shown in the table below which also correspond to
//
// how they are parsed on the backend using golang's type system.
//
// Data Type (SL)   | Data Type (Golang)
// -----------------|-------------------
// Integer          | int
// String           | string
// boolean          | boolean
// Float            | float32
// NULL             | nil
//
package SkyLine_Backend_Modules_Objects

import (
	"strings"

	SkyAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
)

//
// Define object data types
//
const (
	SKYLINE_DATATYPE_INTEGER_OBJECT   = "Integer"    // Defines | Integer objects (1)
	SKYLINE_DATATYPE_INTEGER8_OBJECT  = "Integer8"   // Defines | Integer objects (1)
	SKYLINE_DATATYPE_INTEGER16_OBJECT = "Integer16"  // Defines | Integer objects (1)
	SKYLINE_DATATYPE_INTEGER32_OBJECT = "Integer32"  // Defines | Integer objects (1)
	SKYLINE_DATATYPE_INTEGER64_OBJECT = "Integer64"  // Defines | Integer objects (1)
	SKYLINE_DATATYPE_BOOLEAN_OBJECT   = "BOOLEAN"    // Defines | Boolean objects (true, false)
	SKYLINE_DATATYPE_STRING_OBJECT    = "STRING"     // Defines | String objects ("")
	SKYLINE_DATATYPE_FLOAT_OBJECT     = "FLOAT"      // Defines | Float objects (INTEGER.INTEGER)
	SKYLINE_DATATYPE_NULL_OBJECT      = "NULL"       // Defines | Null objects (NULL)
	SKYLINE_DATATYPE_ERROR_OBJECT     = "ERROR"      // Defines | Error objects (_error_)
	SKYLINE_DATATYPE_BUILTIN_OBJECT   = "BUILTIN"    // Defines | Builtin objects (_builtin_)
	SKYLINE_DATATYPE_FUNCTION_OBJECT  = "FUNCTION"   // Defines | Function objects (Func, function)
	SKYLINE_DATATYPE_FILE_OBJECT      = "FILE"       // Defines | File objects (FILE_RETURN)
	SKYLINE_DATATYPE_HASH_OBJECT      = "HASH"       // Defines | Hash objects ( {INTERFACE: INTERFACE} )
	SKYLINE_DATATYPE_ARRAY_OBJECT     = "ARRAY"      // Defines | Array objects ( [...] )
	SKYLINE_DATATYPE_RETURN_OBJECT    = "RETURN_VAR" // Defines | Return objects ( ret, return, <-TYPE )
	SKYLINE_DATATYPE_MODULE_OBJECT    = "MODULE"     // Defines | Module objects ( module 'module' )
	SKYLINE_DATATYPE_REGISTER_OBJECT  = "REGISTER"   // Defines | Register objects ( object )
	SKYLINE_DATATYPE_BYTE_OBJECT      = "BYTE"       // Defines | Byte objects ( '' )
)

type (
	// Module dependants
	ObjectDataType string

	SkyLine_BuiltinFunction func(Environ *SkyLineEnvironment, InvokeArgs ...SL_Object) SL_Object

	// Environment types
	SkyLineEnvironment struct {
		SkyLine_Storage map[string]SL_Object // Map STRING(Object) to the Objects structure
		SkyLine_RO      map[string]bool      // Map an object in storage to read only if its a constant
		SkyLine_Permit  []string             // Stores thre name of an object in storage
		SkyLine_Outer   *SkyLineEnvironment  // Parent environment storage to allow nesting scopes
	}

	// Object interfaces
	SL_Object interface {
		SkyLine_ObjectFunction_GetDataType() ObjectDataType                                                             // Get the data type of the Object
		SkyLine_ObjectFunction_GetTrueValue() string                                                                    // Get the value that the interpreter or byte code compiler sees
		SkyLine_ObjectFunction_GetInterface() interface{}                                                               // Convert to interface
		SkyLine_ObjectFunction_InvokeObject(Call string, Environ SkyLineEnvironment, InvokeArgs ...SL_Object) SL_Object // Object call functions (TYPE.CALL)
	}

	SL_Hashable interface {
		SL_HashKeyType() HashKey
	}

	SL_Iterable interface {
		SkyLine_Reset_Offset()
		SkyLine_Next_Itteration() (SL_Object, SL_Object, bool)
	}

	// Object strutures

	SL_ENGINE_VALUE struct {
		Value SL_Object
	}

	SL_Functions struct {
		Sky_Function_Arguments []*SkyAST.SL_Identifier
		Unit                   *SkyAST.SL_UnitBlockStatement
		Defaults               map[string]SkyAST.SL_Expression
		Env                    *SkyLineEnvironment
	}

	SL_Module struct {
		SL_ModuleNames string
		SL_Attributes  SL_Object
	}

	SL_Builtin struct {
		Function SkyLine_BuiltinFunction
	}

	SL_String struct {
		Value  string
		Offset int
	}

	SL_Byte struct {
		Value byte
	}

	SL_Boolean struct {
		Value bool
	}

	SL_Integer8 struct {
		Value int8
	}

	SL_Integer16 struct {
		Value int16
	}

	SL_Integer32 struct {
		Value int32
	}

	SL_Integer64 struct {
		Value int64
	}

	SL_Integer struct {
		Value int
	}

	SL_Float struct {
		Value float64
	}

	SL_NULL struct{}

	SL_Error struct {
		Message string
	}

	SL_Return struct {
		Value SL_Object
	}

	SL_Array struct {
		Elements []SL_Object
		Offset   int
	}

	SL_RegisterType struct {
		Value SL_Object
	}

	// Hash implementation
	HashKey struct {
		ObjectDataType ObjectDataType
		Value          uint64
	}

	HashPair struct {
		Key   SL_Object
		Value SL_Object
	}

	SL_HashMap struct {
		Pairs  map[HashKey]HashPair
		Offset int
	}

	SL_InvokeMethod_String  func(Value string, Env SkyLineEnvironment, args ...SL_Object) SL_Object
	SL_InvokeMethod_Integer func(Value int, Env SkyLineEnvironment, args ...SL_Object) SL_Object
	SL_IvokeMethod_Array    func(Values []SL_Object, ArrayObject *SL_Array, Env SkyLineEnvironment, args ...SL_Object) SL_Object
	SL_InvokeMethod_Map     func(Value string, Env SkyLineEnvironment, HashObject *SL_HashMap, args ...SL_Object) SL_Object
)

var (
	SL_InvokeStringMap  = map[string]SL_InvokeMethod_String{}
	SL_InvokeIntegerMap = map[string]SL_InvokeMethod_Integer{}
	SL_InvokeArrayMap   = map[string]SL_IvokeMethod_Array{}
	SL_InvokeHashMapMap = map[string]SL_InvokeMethod_Map{}

	RegisterTypeMap = map[string]func(Name string, data interface{}){
		"string": func(Name string, data interface{}) {
			if conv, ok := data.(func(Value string, Env SkyLineEnvironment, args ...SL_Object) SL_Object); ok {
				//fmt.Println("[+] Register string")
				SL_InvokeStringMap[Name] = conv
			}
		},
		"integer": func(Name string, data interface{}) {
			if conv, ok := data.(func(Value int, Env SkyLineEnvironment, args ...SL_Object) SL_Object); ok {
				//fmt.Println("[+] Register integer")
				SL_InvokeIntegerMap[Name] = conv
			}
		},
		"array": func(Name string, data interface{}) {
			if conv, ok := data.(func(Values []SL_Object, ArrayObject *SL_Array, Env SkyLineEnvironment, args ...SL_Object) SL_Object); ok {
				SL_InvokeArrayMap[Name] = conv
				//fmt.Println("[+] Register array")
			}
		},
		"hash": func(Name string, data interface{}) {
			if conv, ok := data.(func(Value string, Env SkyLineEnvironment, HashObject *SL_HashMap, args ...SL_Object) SL_Object); ok {
				SL_InvokeHashMapMap[Name] = conv
				//fmt.Println("[+] Register hash")
			}
		},
	}
)

func RegisterInvokeString(Name, DataType string, data interface{}) {
	if v, ok := RegisterTypeMap[strings.ToLower(DataType)]; ok {
		v(Name, data)
	}
}
