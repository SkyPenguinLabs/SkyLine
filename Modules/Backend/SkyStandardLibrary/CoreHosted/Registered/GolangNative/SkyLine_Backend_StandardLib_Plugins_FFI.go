//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  _  | |_ _ ___|_|___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|   __| | | | . | |   |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|  |_|___|_  |_|_|_|
//	primary units that define the functions to register the sub func's  //            |___|                               |___|
//////////////////////////////////////////////////////////////////////////
//
// This section of files contain mathemaical functions that can be registered into the SkyLine programming language. Most of these automate the backend of golang's basic
//
// interfaces for math and automate most of the backend. However, some other functions may be algorithmic implementations, specific sets or tweaks that can also be added.
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//				^...^
//			   <_* *_>
//				 \_/
//
// This file defines all of the functions dedicated to foreign function interfaces such as lua or go plugins
//
package SkyLine_Backend_Native_Plugins

import (
	"fmt"
	"plugin"
	"reflect"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	Helpers "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"

	"log"
)

func LoadPlugin(Environ *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"LoadPlugin", args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(2),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		log.Fatal(x.Error())
	}
	symbol := args[1].(*SkyEnv.SL_String).Value
	name := args[0].(*SkyEnv.SL_String).Value
	p, err := plugin.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	v, err := p.Lookup(symbol)
	if err != nil {
		log.Fatalf("error finding symbol: %s", err)
	}
	fnValue := reflect.ValueOf(v)
	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		log.Fatalf("symbol is not a function")
	}
	expectedFnType := reflect.TypeOf((*func(*SkyEnv.SkyLineEnvironment, ...SkyEnv.SL_Object) SkyEnv.SL_Object)(nil)).Elem()
	if fnType != expectedFnType {
		fmt.Println("Function type -> ", fnType)
		log.Fatalf("symbol has unexpected function signature")
	}
	argValues := make([]reflect.Value, 0, len(args)+1)
	argValues = append(argValues, reflect.ValueOf(Environ))
	for _, arg := range args {
		argValues = append(argValues, reflect.ValueOf(arg))
	}
	resultValues := fnValue.Call(argValues)
	return resultValues[0].Interface().(SkyEnv.SL_Object)
}

/*

Further development notes: Try to get the plugin loader (FFI) implementation to just export and assign those functions proper values
within the environment.

*/
