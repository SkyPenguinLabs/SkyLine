//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             ______  _ _____
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     | / |     |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|-   -|/ /|  |  |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_/ |_____|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This code section defines input and output functions and controllers. These controllers allow you to do specific things with the systems IO such as reading, writing
//
// converting and controlling input while also being able to mess around with specific terminal formats.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines the types and models this standard library may rely on
//
//
package SkyLine_Standard_Library_IO

import (
	"strconv"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

const (
	Linux_Clear   = "\x1b[H\x1b[2J\x1b[3J" // clear screen
	Linux_FReturn = "\033[39m"             // Return forground color
)

type (
	InputFunction func(value string) SkyEnv.SL_Object
)

var (
	InputMapType = map[SkyEnv.SL_Object]InputFunction{
		&SkyEnv.SL_String{}: func(value string) SkyEnv.SL_Object {
			return &SkyEnv.SL_String{Value: value}
		},
		&SkyEnv.SL_Integer{}: func(value string) SkyEnv.SL_Object {
			c, x := strconv.ParseInt(value, 0, 64)
			if x != nil {
				return &SkyEnv.SL_Error{Message: "Could not return this value, it was not able to be parsed as a integer which means it was either a float or character but this function does not support that as input"}
			}
			return &SkyEnv.SL_Integer{Value: int(c)}
		},
		&SkyEnv.SL_Boolean{}: func(value string) SkyEnv.SL_Object {
			c, x := strconv.ParseBool(value)
			if x != nil {
				return &SkyEnv.SL_Error{Message: "Could not return this value, could not be parsed as boolean as per requested"}
			}
			return &SkyEnv.SL_Boolean{Value: c}
		},
		&SkyEnv.SL_Float{}: func(value string) SkyEnv.SL_Object {
			c, x := strconv.ParseFloat(value, 64)
			if x != nil {
				return &SkyEnv.SL_Error{Message: "Could not return this value, it was not able to parse as a float value which means it was either a character or an integer, this function input does not accept anything but a float value"}
			}
			return &SkyEnv.SL_Float{Value: c}
		},
	}
)
