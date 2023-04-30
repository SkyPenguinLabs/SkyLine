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
// Filename      |  SkyLine_Parser_NextLoadFault.go
// Project       |  SkyLine programming language
// Line Count    |  5 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines a function to load the next token during parsing
//
// State         | Working but needs work
// Resolution    | Can just be moved into a "helpers" file
//
//
package SkyLine_Backend

func (parser *Parser) NextLoadFaultToken() {
	parser.PreviousToken = parser.CurrentToken
	parser.CurrentToken = parser.PeekToken
	parser.PeekToken = parser.Lex.Scan_NT()
}
