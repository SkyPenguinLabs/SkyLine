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
// Defines       | This file defines all of the proper models for the invokes which includes variables, maps, or types that this section or module requires
//
//
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//
// This is the registry module, this module apart of SkyLine_Backend allows us to register standard library based functions which are called with
//
// class.functionname
//
// this is pretty simple to understand however class and module keywords have not been implemented which means that you can not register your
// own custom standard module / library. Asides from that, we use the init() function because init functions will always run or be called before
// the main() function in go. Using registers under the init function we can ensure the environment has standard functions registered and placed
// into the environment before it is fully started and the input program is parsed. This eliminates the need to import("math") however in the
// further future import keywords will need to be added for standard library functions. This is becausethe bigger our standard library gets the
// more imports will need to be added and the harder the program will be to parse. Currently, due to the factor of how small the standard library
// is, it is not that bad to register the built in functions before a new environment for the input program is started which means it will not slow
// down runtime. However, as this again gets bigger we will need to eliminate registering before runtime unless they are standard functions such as
// .str, .int, integer, boolean, empty?, nil?, carries?, exported? etc which allow for a much heavier use case and do not require imports
// Using the import keyword will give the user the option to allow the program to import and register the standard library functions before
// runtime and parsing. This may cause collisions within the environment so we can actually cause another keyword to exist known as "register"
// followed by the library name. This keyword may be called like so register("math") pr register<<"math">> for a much more complex and parsed
// syntax. Allowing both register and import keywords allow the user to register the library functions before runtime and import files before
// runtime.
//
//
// - Mon 27 Feb 2023 10:23:16 PM EST
//
// as of the given date and time of writing this, SkyLine now will ask you to register the library before you use them if it is standard
// this includes crypt, math, net, http and much other built in libraries used within the SkyLine programming language
//
// ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This unit is the prime unit for any invoke or object call based functions that get called or pushed into the environment when the SkyEnvironment is loaded
//
// you do not need to register data types unless you are using specific functions that are invokes but not always going to be invoke calls but again are used as one.
//
//
package SkyLine_Module_Std_Lib_Invokes

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
)

var (
	Datatypes = []string{
		"string.",
		"float.",
		"object.",
		"hash.",
		"array.",
		"boolean.",
	}

	Static_String_Methods = []string{
		"Length",         // Length of string
		"Methods",        // Methods
		"Ord",            // ord
		"Integer",        // to integer
		"Float",          // to float
		"Boolean",        // To byte
		"Upper",          // To uppercase
		"Lower",          // To lowercase
		"Title",          // To title
		"Split",          // Split
		"Trim",           // Trim
		"UnlinkRegistry", // Unlink a libraries registration | This is a string because we unlink a library name
		"View",
		"ToHArr",
	}

	Static_Integer_Methods = []string{
		"chr",
		"Methods", // Grab all methods
		"View",    // View values
	}

	Static_Float_Methods = []string{
		"Methods", // Grab all methods
		"View",
	}

	Static_Array_Methods = []string{
		"Reverse", // Reverse
		"Append",  // Append
		"Copy",    // Copy
		"Swap",    // Swap
		"Less",    // Less
		"Compare", // Compare
		"PopR",    // Pop right
		"PopL",    // Pop left
		"Length",  // Length of the array
		"Methods", // Grab all methods
		"Typeof",  // Typeof
		"View",    // View
	}

	Static_Hash_Methods = []string{
		"Keys",    // Dump all the hash's keys
		"Methods", // Grab all methods
		"View",
		"Length", // Length of the hash keys
	}

	Static_Boolean_Methods = []string{
		"Methods", // Grab all methods
		"View",
	}

	Static_BuiltInFunction_Methods = []string{
		"Methods", // Grab all methods
		"View",
	}
)

type (
	InvokeMethFunctionArray func(Values []SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object
	InvokeMethFunctionStr   func(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object
	FunctionCall            func(FunctionName string, Arguments []SkyEnv.SL_Object) error
)
