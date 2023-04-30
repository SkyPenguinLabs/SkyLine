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
// Filename      |  SkyLine_Language_Envionment_TemporaryScope.go
// Project       |  SkyLine programming language
// Line Count    |  10 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Creates a new temprorary environment or scope rather for the specified for loops and range loops
//
//
package SkyLine_Backend

func NewTempScop(outer *Environment_of_environment, keys []string) *Environment_of_environment {
	env := NewEnvironment()
	env.Outer = outer
	env.permit = keys
	return env
}
