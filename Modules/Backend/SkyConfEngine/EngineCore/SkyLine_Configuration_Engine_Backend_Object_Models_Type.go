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
// File contains -> This file contains all necessary functions for the objects within SLC allowing us to view the objects type

package SkyLine_Configuration_Engine_Backend_Source

func (ObjInteger *ObjectInteger) ObjectDataType() ObjectsDataType  { return OBJECT_INTEGER } // Return integer object
func (ObjBoolean *ObjectBoolean) ObjectDataType() ObjectsDataType  { return OBJECT_BOOLEAN } // Return boolean object
func (ObjString *ObjectString) ObjectDataType() ObjectsDataType    { return OBJECT_STRING }  // Return string object
func (ObjArray *ObjectArray) ObjectDataType() ObjectsDataType      { return OBJECT_ARRAY }   // Return array object
func (ObjNullField *ObjectNULL) ObjectDataType() ObjectsDataType   { return OBJECT_NULL }    // Return null object
func (ObjErrorField *ObjectERROR) ObjectDataType() ObjectsDataType { return OBJECT_ERROR }   // Return Error object
func (ObjBuiltIn *ObjectBUILTINFUNCTION) ObjectDataType() ObjectsDataType {
	return OBJECT_BUILT_IN_FUNCTION
} // Return built in function object
