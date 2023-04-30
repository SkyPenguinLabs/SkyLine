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
// Filename      |  Skyline_Evaluator_SubBase_Functions.go
// Project       |  SkyLine programming language
// Line Count    |  30+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines some basic evaluator functions that can all be thrown into the error system file
//
// STATE         | Needs to be moved
// Resolution    | Functions should be collectively thrown into the error system file
//
package SkyLine_Backend

import "fmt"

func isTruthy(obj SLC_Object) bool {
	return obj != NilValue && obj != FalseValue
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj SLC_Object) bool {
	return obj != nil && obj.SL_RetrieveDataType() == ErrorType
}
