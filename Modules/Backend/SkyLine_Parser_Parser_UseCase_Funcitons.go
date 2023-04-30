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
// Line Count    |  40 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines all functions for the parser such as parser errors, getting current lines and other various states
//
// State         | Working but needs work
// Resolution    | Can all be organized into other files such as parser helper functions and other files alike
//
//
package SkyLine_Backend

import "fmt"

func (parser *Parser) ParserErrors() []string          { return parser.Errors }
func (parser *Parser) PeekTokenIs(typ Token_Type) bool { return parser.PeekToken.Token_Type == typ }
func (parser *Parser) GetLineCound() string            { return fmt.Sprint(parser.Lex.CurLine) }

func (parser *Parser) CurrentTokenIs(typ Token_Type) bool {
	return parser.CurrentToken.Token_Type == typ
}

func (parser *Parser) NT() {
	parser.PreviousToken = parser.CurrentToken
	parser.CurrentToken = parser.PeekToken
	parser.PeekToken = parser.Lex.Scan_NT()
}

func (parser *Parser) PeekError(typ Token_Type) {
	msg := Map_Parser[ERROR_DURING_PEEK_IN_PARSER](fmt.Sprint(typ), string(parser.PeekToken.Token_Type)).Message
	parser.Errors = append(parser.Errors, msg)
}

func (parser *Parser) ExpectPeek(typ Token_Type) bool {
	if parser.PeekTokenIs(typ) {
		parser.NT()
		return true
	}
	parser.PeekError(typ)
	return false
}

func (parser *Parser) ParseProgram() *Program {
	program := &Program{
		Statements: []Statement{},
	}

	for !parser.CurrentTokenIs(TOKEN_EOF) {
		stmt := parser.SkyLine_Statement()
		program.Statements = append(program.Statements, stmt)
		parser.NT()
	}

	return program
}
