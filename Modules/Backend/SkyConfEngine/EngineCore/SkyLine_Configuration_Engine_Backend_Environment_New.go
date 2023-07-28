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
// Contains -> This file contains all of the newwer environment creation functions. These functions "spawn" or rather cretae a new environment before the engine can parse the data

package SkyLine_Configuration_Engine_Backend_Source

func Start_Engine_Environment_Create() *Engine_Environment_Of_Environment {
	Store, Reader := make(map[string]SLC_Object), make(map[string]bool)
	return &Engine_Environment_Of_Environment{
		StoreObj:    Store,
		ReadOnly:    Reader,
		EngineOuter: nil,
	}
}

func Start_Engine_Enclosed_Environment(OutterShell *Engine_Environment_Of_Environment) *Engine_Environment_Of_Environment {
	Environ := Start_Engine_Environment_Create()
	Environ.EngineOuter = OutterShell
	return Environ
}
