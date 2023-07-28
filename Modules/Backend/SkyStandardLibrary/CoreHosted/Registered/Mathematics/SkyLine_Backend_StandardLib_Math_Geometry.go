//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____     _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     |___| |_| |_
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___| | | | .'|  _|   |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_|_|_|__,|_| |_|_|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
// This section of files contain mathemaical functions that can be registered into the SkyLine programming language. Most of these automate the backend of golang's basic
//
// interfaces for math and automate most of the backend. However, some other functions may be algorithmic implementations, specific sets or tweaks that can also be added.
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file holds all supported geometric functions within mathematics that again SkyLine can currently support
//
package SkyLine_Standard_Library_Math

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	"math"
)

func Math_Geometry_SQRT(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Sqrt", 1, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, args...); x == nil {
		return &SkyEnv.SL_Float{Value: math.Sqrt(float64(args[0].(*SkyEnv.SL_Float).Value))}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Geoemtry_CBRT(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Cbrt", 1, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, args...); x == nil {
		x1 := args[0].(*SkyEnv.SL_Float).Value
		z := x1 / 3.0
		for i := 0; i < 10; i++ {
			z = z - ((z*z*z - x1) / (3 * z * z))
		}
		return &SkyEnv.SL_Float{Value: z}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}
