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
// This file holds all supported trigonometric functions within mathematics that again SkyLine can currently support
//
package SkyLine_Standard_Library_Math

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	"math"
)

func Math_Trig_Tan(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x1 := AutomateError_ExactArgCheckAndType("math.Tan", 1, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, args...); x1 == nil {
		return &SkyEnv.SL_Float{
			Value: math.Tan(args[0].(*SkyEnv.SL_Float).Value),
		}
	} else {
		return &SkyEnv.SL_Error{Message: x1.Error()}
	}
}

func Math_Trig_Cos(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x1 := AutomateError_ExactArgCheckAndType("math.Cos", 1, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, args...); x1 == nil {
		return &SkyEnv.SL_Float{
			Value: math.Cos(args[0].(*SkyEnv.SL_Float).Value),
		}
	} else {
		return &SkyEnv.SL_Error{Message: x1.Error()}
	}
}

func Math_Trig_Sin(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x1 := AutomateError_ExactArgCheckAndType("math.Sin", 1, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, args...); x1 == nil {
		return &SkyEnv.SL_Float{
			Value: math.Sin(args[0].(*SkyEnv.SL_Float).Value),
		}
	} else {
		return &SkyEnv.SL_Error{Message: x1.Error()}
	}
}
