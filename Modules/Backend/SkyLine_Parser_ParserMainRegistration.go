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
// Filename      |  SkyLine_Parser_ParserMainRegistration.go
// Project       |  SkyLine programming language
// Line Count    |  80+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines -> This file defines all of the parser prefix and infix parsing maps which is the way to parse specific tokens. In this file it helps us output and run specific functions
//            based on a token's type
//
//
package SkyLine_Backend

func New_Parser(l *ScannerStructure) *Parser {
	parser := &Parser{
		Lex:    l,
		Errors: []string{},
	}
	parser.NextLoadFaultToken()
	parser.NextLoadFaultToken()

	parser.PrefixParseFns = map[Token_Type]PrefixParseFn{
		TOKEN_IDENT:          parser.SkyLine_Identifier,
		TOKEN_INT:            parser.SkyLine_IntegerLiteral,
		TOKEN_FLOAT:          parser.SkyLine_FloatLiteral,
		TOKEN_BANG:           parser.SkyLine_PrefixExpression,
		TOKEN_MINUS:          parser.SkyLine_PrefixExpression,
		TOKEN_TRUE:           parser.SkyLine_Boolean,
		TOKEN_FALSE:          parser.SkyLine_Boolean,
		TOKEN_LPAREN:         parser.SkyLine_GroupedExpression,
		TOKEN_IF:             parser.SkyLine_ConditionalExpression,
		TOKEN_FUNCTION:       parser.SkyLine_FunctionLiteral,
		TOKEN_STRING:         parser.parseStringLiteral,
		TOKEN_LBRACKET:       parser.parseArrayLiteral,
		TOKEN_LBRACE:         parser.parseHashLiteral,
		TOKEN_LINE:           parser.SkyLine_GroupImportExpression,
		TOKEN_SWITCH:         parser.SkyLine_SwitchCase_Expression,
		TOKEN_REGISTER:       parser.SkyLine_GroupedExpression,
		TOKEN_KEYWORD_ENGINE: parser.SkyLine_GroupedExpression,
		TOKEN_IMPORT:         parser.SkyLine_ImportExpression,
		TOKEN_FOR:            parser.SkyLine_ForLoop,
		TOKEN_FOREACH:        parser.SkyLine_ForEach,
	}

	parser.InfixParseFns = map[Token_Type]InfixParseFn{
		TOKEN_MODULO:          parser.SkyLine_InfixExpression,
		TOKEN_PLUS:            parser.SkyLine_InfixExpression,
		TOKEN_MINUS:           parser.SkyLine_InfixExpression,
		TOKEN_ASTARISK:        parser.SkyLine_InfixExpression,
		TOKEN_SLASH:           parser.SkyLine_InfixExpression,
		TOKEN_EQ:              parser.SkyLine_InfixExpression,
		TOKEN_NEQ:             parser.SkyLine_InfixExpression,
		TOKEN_LT:              parser.SkyLine_InfixExpression,
		TOKEN_GT:              parser.SkyLine_InfixExpression,
		TOKEN_LPAREN:          parser.parseCallExpression,
		TOKEN_LBRACKET:        parser.parseIndexExpression,
		TOKEN_GTEQ:            parser.SkyLine_InfixExpression,
		TOKEN_LTEQ:            parser.SkyLine_InfixExpression,
		TOKEN_PLUS_EQUALS:     parser.SkyLine_Assignment,
		TOKEN_DIVEQ:           parser.SkyLine_Assignment,
		TOKEN_MINUS_EQUALS:    parser.SkyLine_Assignment,
		TOKEN_ASTERISK_EQUALS: parser.SkyLine_Assignment,
		TOKEN_ASSIGN:          parser.SkyLine_Assignment,
		TOKEN_ANDAND:          parser.SkyLine_InfixExpression,
		TOKEN_OROR:            parser.SkyLine_InfixExpression,
		TOKEN_POWEROF:         parser.SkyLine_InfixExpression,
		TOKEN_PERIOD:          parser.parseMethodCallExpression,
		MODULECALL:            parser.SkyLine_SelectorExpression,
	}

	parser.PostfixParseFns = map[Token_Type]PostfixParseFn{
		TOKEN_MINUS_MINUS: parser.ParsePostfixExpression,
		TOKEN_PLUS_PLUS:   parser.ParsePostfixExpression,
	}
	return parser
}
