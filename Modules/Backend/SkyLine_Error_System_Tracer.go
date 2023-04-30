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
// Filename      |  SkyLine_Error_System_Tracer.go
// Project       |  SkyLine programming language
// Line Count    |  0 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the output based functions for the error system, well not all
//
// STATE         | Needs to be organized and worked on
// Resolution    | Functions can all be automated better, half of them are not used, constants can be thrown into the Models file
//
//
package SkyLine_Backend

// Func Prepare and return error line for base error system
func (parser *Parser) Prepare_Base_Error_Message(
	error_message string,
	Code_Line string,
	Fix bool,
	FixStmt string,
) string {
	var Out string
	Out += SKYLINE_HIGH_RES_VIS_RED + "E | " +
		SKYLINE_HIGH_RES_VIS_BLUE +
		ParserErrorSystem_GetFileName() + ":" + parser.GetLineCound()
	Out += SKYLINE_RESTORE + "\n\n"
	Out += SKYLINE_HIGH_RES_VIS_SUNSET + parser.GetLineCound() + " | " + SKYLINE_RESTORE + Code_Line + "\n"
	if Fix {
		Out += SKYLINE_HIGH_RES_VIS_SUNSET + "F | " + SKYLINE_RESTORE + FixStmt + "\n"
	}
	Out += "\n" + SKYLINE_HIGH_RES_VIS_RED + "Error: " + SKYLINE_RESTORE + SKYLINE_HIGH_RES_VIS_BLUE + error_message + "\n"
	Out += "\n" + SKYLINE_RESTORE
	return Out
}
