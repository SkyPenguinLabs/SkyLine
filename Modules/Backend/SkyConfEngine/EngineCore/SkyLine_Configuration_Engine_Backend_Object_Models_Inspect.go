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
// File contains -> This file contains information on the objects and functions that allow the interpreter or engine to access the current values of the object

package SkyLine_Configuration_Engine_Backend_Source

import (
	"bytes"
	"fmt"
	"strings"
)

func (ObjInteger *ObjectInteger) ObjectInspectFunc() string {
	return fmt.Sprintf("%d", ObjInteger.Value)
} // Grab the integer value

func (ObjBoolean *ObjectBoolean) ObjectInspectFunc() string {
	return fmt.Sprintf("%t", ObjBoolean.Value)
} // Grab the boolean value

func (ObjString *ObjectString) ObjectInspectFunc() string {
	return ObjString.Value
} // Grab the string value

func (ObjNULL *ObjectNULL) ObjectInspectFunc() string {
	return "NULL"
} // Grab the null field value

func (ObjERROR *ObjectERROR) ObjectInspectFunc() string {
	return ObjERROR.Message
} // Grab the error field and message

func (ObjBuiltIn *ObjectBUILTINFUNCTION) ObjectInspectFunc() string {
	return "Object<BUILT-IN-FUNCTION>"
} // Grab the function value

func (ObjArray *ObjectArray) ObjectInspectFunc() string {
	var Out bytes.Buffer
	ArrayElems := []string{}
	for _, element := range ObjArray.Elements {
		ArrayElems = append(ArrayElems, element.ObjectInspectFunc())
	}
	Out.WriteString("[")
	Out.WriteString(strings.Join(ArrayElems, ", "))
	Out.WriteString("];")
	return Out.String()
}
