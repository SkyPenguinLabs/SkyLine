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
// Defines       | This file defines a function that takes a object and converts it into a native boolean object
//
package SkyLine_Backend

func OBJECT_CONV_TO_NATIVE_BOOL(o SLC_Object) bool {
	if r, ok := o.(*ReturnValue); ok {
		o = r.Value
	}
	switch obj := o.(type) {
	case *Boolean_Object:
		return obj.Value
	case *String:
		return obj.Value != ""
	case *Nil:
		return false
	case *Integer:
		if obj.Value == 0 {
			return false
		}
		return true
	case *Float:
		if obj.Value == 0.0 {
			return false
		}
		return true
	case *Array:
		if len(obj.Elements) == 0 {
			return false
		}
		return true
	case *Hash:
		if len(obj.Pairs) == 0 {
			return false
		}
		return true
	default:
		return true
	}
}
