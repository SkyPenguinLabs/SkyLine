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
// Defines       | This file defines all invoke initation before the environment is fully started or filled with developer defined functions
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

import SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"

func RegisterInvokes() {
	//::::::::::::::::::::::::::::::::::::::::::::::
	//:: String invoke registry
	//::::::::::::::::::::::::::::::::::::::::::::::
	SkyEnv.RegisterInvokeString("Length", "string", String_Length)
	SkyEnv.RegisterInvokeString("Split", "string", String_Split)
	SkyEnv.RegisterInvokeString("Upper", "string", String_Upper)
	SkyEnv.RegisterInvokeString("Lower", "string", String_Lower)
	SkyEnv.RegisterInvokeString("Ord", "string", String_Ord)
	SkyEnv.RegisterInvokeString("ToInt", "string", String_ToInt)
	SkyEnv.RegisterInvokeString("ToBool", "string", String_ToBoolean)
	SkyEnv.RegisterInvokeString("ToFloat", "string", String_ToFloat)
	SkyEnv.RegisterInvokeString("Methods", "string", String_Methods)
	SkyEnv.RegisterInvokeString("Outln", "string", Integer_Outln)
	SkyEnv.RegisterInvokeString("ToHex", "String", String_ToHex)
	SkyEnv.RegisterInvokeString("Format", "String", String_Format)
	SkyEnv.RegisterInvokeString("Substr", "String", String_SubString)
	SkyEnv.RegisterInvokeString("Byte", "String", String_BArr)
	SkyEnv.RegisterInvokeString("Hex", "String", String_Hex)
	//::::::::::::::::::::::::::::::::::::::::::::::
	//:: Array invoke registry
	//::::::::::::::::::::::::::::::::::::::::::::::
	SkyEnv.RegisterInvokeString("View", "array", Array_View)
	SkyEnv.RegisterInvokeString("Typeof", "array", Array_Typeof)
	SkyEnv.RegisterInvokeString("Length", "array", Array_Length)
	SkyEnv.RegisterInvokeString("Prepend", "array", Array_Prepend)
	SkyEnv.RegisterInvokeString("Append", "array", Array_Append)
	SkyEnv.RegisterInvokeString("PopR", "array", Array_PopRight)
	SkyEnv.RegisterInvokeString("PopL", "array", Array_PopLeft)
	SkyEnv.RegisterInvokeString("Reverse", "array", Array_Reverse)
	SkyEnv.RegisterInvokeString("Copy", "array", Array_Copy)
	SkyEnv.RegisterInvokeString("Swap", "array", Array_Swap)
	SkyEnv.RegisterInvokeString("Ascii", "array", Array_Ascii)
	SkyEnv.RegisterInvokeString("Hex", "array", Array_Hex)
	//::::::::::::::::::::::::::::::::::::::::::::::
	//:: Hash map invoke registry
	//::::::::::::::::::::::::::::::::::::::::::::::
	SkyEnv.RegisterInvokeString("Keys", "hash", Hash_DisplayKeys)
	SkyEnv.RegisterInvokeString("Length", "hash", GetHashLength)
	//::::::::::::::::::::::::::::::::::::::::::::::
	//:: Integer map invoke registry
	//::::::::::::::::::::::::::::::::::::::::::::::
	SkyEnv.RegisterInvokeString("Outln", "integer", Integer_Outln)
	SkyEnv.RegisterInvokeString("Out", "integer", Integer_Out)
	SkyEnv.RegisterInvokeString("Bool", "integer", Integer_Bool)
	SkyEnv.RegisterInvokeString("String", "integer", Integer_Str)
	//::::::::::::::::::::::::::::::::::::::::::::::
	//:: Boolean  map invoke registry
	//::::::::::::::::::::::::::::::::::::::::::::::
	SkyEnv.RegisterInvokeString("T", "boolean", BoolT)
	SkyEnv.RegisterInvokeString("F", "boolean", BoolF)
	SkyEnv.RegisterInvokeString("String", "boolean", BoolStr)
}
