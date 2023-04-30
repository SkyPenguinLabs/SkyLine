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
// Filename      |  SkyLine_Language_Environment_ExportableHash.go
// Project       |  SkyLine programming language
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This defines the exportable hash function to export a environment and new parsers hash
//
//
package SkyLine_Backend

import "unicode"

func (e *Environment_of_environment) ExportedHash() *Hash {
	pairs := make(map[HashKey]HashPair)
	for k, v := range e.Store {
		if unicode.IsUpper(rune(k[0])) {
			s := &String{Value: k}
			pairs[s.HashKey()] = HashPair{Key: s, Value: v}
		}
	}
	return &Hash{Pairs: pairs}
}
