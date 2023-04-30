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
// Filename      |  SkyLine_Standard_PreRegistered_Library_Mathematics.go
// Project       |  SkyLine programming language
// Line Count    |  2 active lines
// Status        |  Inactive
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines a registry function to register built in functions
//
package SkyLine_Backend

import (
	"fmt"
	"log"
	"strconv"
)

var Builtins = map[string]*Builtin{}
var BuiltInVariables_String = map[string]*String{}
var BuiltInVariables_Float = map[string]*Float{}
var BuiltInVariables_Integer = map[string]*Integer{}
var BuiltInVariables_Boolean = map[string]*Boolean_Object{}

func RegisterBuiltin(name string, fun BuiltinFunction) {
	Builtins[name] = &Builtin{Fn: fun}
}

func UnlikeRegistryFunctions(name string) {
	delete(Builtins, name)
}

func RegisterVariable(name string, object SLC_Object) {
	switch object.(type) {
	case *String:
		BuiltInVariables_String[name] = &String{Value: object.SL_InspectObject()}
	case *Float:
		conv, x := strconv.ParseFloat(object.SL_InspectObject(), 64)
		if x != nil {
			log.Fatalf("SKYLINE DEVELOPER ERROR (Register variable): Could not register variable %s due to %s", name, fmt.Sprint(x))
		}
		BuiltInVariables_Float[name] = &Float{Value: conv}
	case *Integer:
		conv, x := strconv.ParseInt(object.SL_InspectObject(), 10, 64)
		if x != nil {
			log.Fatalf("SKYLINE DEVELOPER ERROR (Register variable): Could not register variable %s due to %s", name, fmt.Sprint(x))
		}
		BuiltInVariables_Integer[name] = &Integer{Value: conv}
	case *Boolean_Object:
		conv, x := strconv.ParseBool(object.SL_InspectObject())
		if x != nil {
			log.Fatalf("SKYLINE DEVELOPER ERROR (Register variable): Could not register variable %s due to %s", name, fmt.Sprint(x))
		}
		BuiltInVariables_Boolean[name] = &Boolean_Object{Value: conv}
	}
}
