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
// File contains -> This file hashes the SkyLine configuration language data types and allows for valid parsing

package SkyLine_Configuration_Engine_Backend_Source

import "hash/fnv"

func (ObjInteger *ObjectInteger) GrabHashKey() HashingKey {
	return HashingKey{
		Type:  ObjInteger.ObjectDataType(),
		Value: uint64(ObjInteger.Value),
	}
}

func (ObjBoolean *ObjectBoolean) GrabHashKey() HashingKey {
	var val uint64
	if ObjBoolean.Value {
		val = 1 // true:BINARY
	} else {
		val = 0 // false:BINARY
	}
	return HashingKey{
		Type:  ObjBoolean.ObjectDataType(),
		Value: val,
	}
}

func (ObjString *ObjectString) GrabHashKey() HashingKey {
	STR_HASH := fnv.New64a()
	STR_HASH.Write([]byte(ObjString.Value))
	return HashingKey{
		Value: STR_HASH.Sum64(),
		Type:  ObjString.ObjectDataType(),
	}
}
