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
// This file is for functions that are basic calculations or basic functions used in all fields of mathematics which includes the following
//
// - calculus
// - trigenometry
// - geomtry
// - algebra
// - linear algebra
// - ...
//
// in general, these functions are used in most fields if not all fields of mathematics
//
package SkyLine_Standard_Library_Math

import SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"

func Math_Generic_PowerOf(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x1 := AutomateError_ExactArgCheckAndType("math.Pow", 2, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x1 == nil {
		x := args[0].(*SkyEnv.SL_Integer).Value
		y := args[1].(*SkyEnv.SL_Integer).Value
		p := int64(1)
		for y > 0 {
			if y&1 != 0 {
				p *= int64(x)
			}
			y >>= 1
			x *= x
		}
		return &SkyEnv.SL_Integer{Value: int(p)}
	} else {
		return &SkyEnv.SL_Error{Message: x1.Error()}
	}
}
