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
// Filename      |  SkyLine_Error_System_TreeOut.go
// Project       |  SkyLine programming language
// Line Count    |  50+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the tree based error output and design functions. It is important to note that the error system of SKyLine should be shoved into one file
//
// STATE         | Needs to be organized and worked on
// Resolution    | Functions can all be automated better, half of them are not used, constants can be thrown into the Models file
//
//
package SkyLine_Backend

import "fmt"

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////// | The offical SkyLine error system, this file includes code systems, notes as big as this and messages as well as loggers and other statements which allow you to make a
///////  | much more verbose and advanced error system. This system uses a custom code and output system to define its errors. Do note that this is still in beta and needs to be
///////  | tested much further before deployment and that it is important to note that not all codes need to be used.
//////   |
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Date  : Mon 27 Feb 2023 10:28:49 PM EST
// Codes :
// Types : Files, Parser errors, argument error, type errors, unknown operators, operational errors, mathematical errors, byte errors, binary errors, flow errors, systematic errors, os errors, support errors, network errors and more.
//
//
// The code system for SkyLine will differ based on the type of code and the type of system you are trying to troubleshoot, for example
// certain modules, functions, nodes and starters have reserved sets of codes which all mean different things and which have their own error messages.
// With every error based on the error and type of code if it is supported or required will allow for recomendations such as missing imports, semicolons etc
// below you will find a list of all the error codes and their reserved spots
//
// Startup        | -> 0-100
// Parser         | -> 100-300
// Evaluator      | -> 300-500
// Operations     | -> 500-700
// Environment    | -> 800-1000
// Function Calls | -> 1000-2000
// Network calls  | -> 2000-3000
// Mathematics    | -> 4000-5000
// Data types     | -> 5000-6000
// Server errors  | -> 6000-6500
// Arguments      | -> 6500-7000
//
//
// Each message and each code depending on the type of code and the catgeory increments completely differently and is all formatted differently. For example below you will find a
// list of the error code reservations and their incremented positions
//
//
// Startup        | -> 0-100         | Increments by 5
// Parser         | -> 100-300       | Increments by 5
// Evaluator      | -> 300-500       | Increments by 10
// Operations     | -> 500-700       | Increments by 20
// Environment    | -> 800-1000      | Increments by 30
// STD library    | -> 1000-2000     | Increments by 10
// Network calls  | -> 2000-3000     | Increments by 20
// Mathematics    | -> 4000-5000     | Increments by 80
// Data types     | -> 5000-6000     | Increments by 90
// Server errors  | -> 6000-10,000   | Increments by 100
// Arguments      | -> 10,000-10,500 | Increments by 100s
//
// These arguments all have their own respected values and all have their own unique values and codes as well
// for example if you were to make a HTTP POST request to a given domain and the endpoint failes a code within the
// range of 2000-3000 will be displayed because this counts as a network error within the system. Anything past 10,500
// is just wild and much more verbose and a complicated error or traceback. Codes should not ever reach past that point
// as that typically means there is a panic within the program, a fatal log, a dev issue on the backend or some weird issue.
// It is important to note that within each system there is no line or column based traceback system right now as this is
// a main priority and I would rather much prefer to focus on the functionality of the error system first before the backend
//
//
// - Sub BRAMCH
//
// This file not only ties in with the codes above but is a direct link and implementation to colors
// as well as the formulated output from the skyline interpreter. This error system outputs as a tree
// and will branch off of errors and more unqiue subsets allowing for a much easier and better error
// handler.
//

type TreeNode struct {
	Type     string
	Children []*TreeNode
}

// The following constant list is a general setting of the output used within the error tree
const (
	T_V  = SKYLINE_HIGH_DEFRED + "│" // Tree vertical char
	T_H  = SKYLINE_HIGH_DEFRED + "─" // Tree horizontal char
	T_C  = SKYLINE_HIGH_DEFRED + "└" // Tree corner char
	T_T  = SKYLINE_HIGH_DEFRED + "├" // Tree tee char
	T_CR = SKYLINE_HIGH_DEFRED + "┼" // Tree Cross char
	T_SP = SKYLINE_HIGH_DEFRED + " " // Tree space char

)

func RetTreeSys(node *TreeNode, pref string, lastnode bool) {
	fmt.Print(pref)
	if lastnode {
		fmt.Print(T_C)
		pref += T_SP
	} else {
		fmt.Print(T_T)
		pref += T_V + T_SP
	}
	fmt.Println(T_H, node.Type)
	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		childPrefix := pref
		if lastnode {
			childPrefix += T_SP
		}
		if len(child.Children) > 0 {
			RetTreeSys(child, childPrefix, isLastChild)
			if !lastnode {
				fmt.Print(pref, T_T, T_H, T_C, T_H)
				fmt.Println(T_H, T_SP)
			}
		} else {
			fmt.Print(childPrefix)
			if isLastChild {
				fmt.Print(T_C)
				childPrefix += T_SP
			} else {
				fmt.Print(T_T)
				childPrefix += T_V + T_SP
			}
			fmt.Println(T_H, child.Type)
		}
	}
}
