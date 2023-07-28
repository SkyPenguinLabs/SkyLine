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
// Defines       | This file defines the iniation function which will register everything into the environment by default that is allowed to be there.
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
package SkyLine_Module_SkyStandardLib_Direct

import (
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkyEval "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEvaluator"
	SkyModeELF "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/ELF"
	SkyPlugin "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/GolangNative"
	SkyModePwn "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/Utilities"
)

func InitateCallRegisterPwnFuncs() {
	SkyEval.SkyLine_Register_Builtin("Pack32", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.Pwn_Pack32(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("Pack64", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.Pwn_Pack64(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("NewBuf", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.BufferImpl_New(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("BufGet", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.BufferImpl_Get(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("BufSet", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.BufferImpl_Set(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("BufUnget", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.BufferImpl_Unget(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("BufDestroy", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyModePwn.BufferImpl_Destroy(InvokeArgs...))
	})
	SkyModeELF.InitateELFLibs()
}

func InitateCallRegisterStandardCall() {
	SkyEval.SkyLine_Register_Builtin("mode", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SwitchMode(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("print", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (Print(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("println", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (Println(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("sprint", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (Sprint(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("args", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (Arguments(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("version", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (Version(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("LoadPlugin", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyPlugin.LoadPlugin(Environ, InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("Typeof", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (GetDataType(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("sprintf", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (Sprintf(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("int", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (ConvertInt8_16_32_64_to_int(InvokeArgs...))
	})
}
