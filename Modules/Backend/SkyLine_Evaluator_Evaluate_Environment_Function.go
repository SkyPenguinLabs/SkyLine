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
// Filename      |  SkyLine_Evaluator_Evaluate_Environment_Function.go
// Project       |  SkyLine programming language
// Line Count    |  20 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines a function to extend a functions environment
//
// STATE         | Needs to be organized and worked on
// Resolution    | Functions can all be automated better, half of them are not used, constants can be thrown into the Models file
//
package SkyLine_Backend

import (
	"fmt"
	"os"
)

func extendFunctionEnv(fn *Function, args []SLC_Object) *Environment_of_environment {
	Env := NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		Env.Set(param.Value, args[i])
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(Map_Eval[ERROR_EXTEND_FUNCTION_NOT_ENOUGH_ARGUMENTS](fmt.Sprint(fn.SL_InspectObject()), fmt.Sprint(len(args)), fmt.Sprint(len(fn.Parameters))))
				os.Exit(0)
			}
		}()
	}
	return Env
}
