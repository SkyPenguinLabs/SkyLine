///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Object_StringType
// Extension         | .go ( golang source code file )
// Purpose           | Defines required states for the String data type
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This sector of files also known as the 'ObjectType' sector will be defining all of the functions for the current data types within SkyLine.
//
// WARN: This note exists for every single Object file in this directory
//
// These functions include
//						A - SkyLine_ObjectFunction_GetDataType
//						B - SkyLine_ObjectFunction_GetTrueValue
//						C - SkyLine_ObjectFunction_GetInterface
//                      D - SkyLine_ObjectFunction_InvokeObject
//
//
//
// A: Data Type return
//			- This function will tell the SkyLine interpreter to grab or return the data type of the current object.
//			  This helps during parsing and evaluation to ensure that the data type is what it is and to make sure that the object
//			  being executed has the proper type.
//
// B: Get True Value
//			- This function may be confusing at first, but it takes a value such as an array and makes the most syntactic appropriate
//			  representation of the data type and then returns it as a string. This function starts by first declaring a byte input and
//			  also writing those values of the data types proper specification and returning it as a string. For example, in SkyLine
//			  variables are dynamic, this means that everything is `interpreted` and variables have no constant type based values. If you
//			  declare a string variable like `set x := "name"; ` then the result of that from the function with lable `B(Get True Value)` will
//			  return `TYPE STRING VARIABLE (x) := STRING[name];` which can be helpful for optimization and syntactic purposes.`
//
// C: Get Interface:
//			- This function just returns a type interface representing the object
//
// D: Invoke Object:
//			- This allows objects to have their own specific functions such as ORD, LEN, SPLIT etc. Below is a table of common object
//			  call functions for different data types
//
//				| SkyLine Data Type | Object Call Function List  |
//   			------------------- | -------------------------- |
//				Integer             | [methods, ord]             |
//				String              | [methods, split, index ...]|
//              Boolean             | [methods]                  |
//              Float               | [methods]                  |
//               ...                | ...
//
// Object call functions are a ton easier and better to work with than requiring the user to require a data type such as `register(Array)`
//
// to get array object calls. In the further future this may actually be required due to the amount of object call functions that will exist
//
package SkyLine_Backend_Modules_Objects

import "unicode/utf8"

func (SL_ObjectString *SL_String) SkyLine_ObjectFunction_GetDataType() ObjectDataType {
	return SKYLINE_DATATYPE_STRING_OBJECT
}

func (SL_ObjectString *SL_String) SkyLine_ObjectFunction_GetInterface() interface{} {
	return SL_ObjectString.Value
}

func (SL_ObjectString *SL_String) SkyLine_Reset_Offset() {
	SL_ObjectString.Offset = 0
}

//
// Syntax
//

func (SL_ObjectString *SL_String) SkyLine_ObjectFunction_GetTrueValue() string {
	return SL_ObjectString.Value
}

//
// Invoke
//

func (SL_ObjectString *SL_String) SkyLine_ObjectFunction_InvokeObject(Call string, Environ SkyLineEnvironment, InvokeArgs ...SL_Object) SL_Object {
	if cx, ok := SL_InvokeStringMap[Call]; ok {
		return cx(SL_ObjectString.Value, Environ, InvokeArgs...)
	} else if Call == "puts" {
		println(SL_ObjectString.SkyLine_ObjectFunction_GetTrueValue())
	}
	return &SL_NULL{}
}

//
// Itteration
//

func (SL_ObjectString *SL_String) SkyLine_Next_Itteration() (SL_Object, SL_Object, bool) {
	if SL_ObjectString.Offset < utf8.RuneCountInString(SL_ObjectString.Value) {
		SL_ObjectString.Offset++
		chars := []rune(SL_ObjectString.Value)
		val := &SL_String{Value: string(chars[SL_ObjectString.Offset-1])}
		return val, &SL_Integer{Value: SL_ObjectString.Offset - 1}, true
	}
	return nil, &SL_Integer{Value: 0}, false
}
