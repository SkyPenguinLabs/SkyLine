//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____             _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_ _ ___ ___| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | |   |  _|  _| | . |   |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |___|_|_|___|_| |_|___|_|_|___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This sub unit defines all of the environment related functions that will tell or inform  the user about the environment. This typically is its own package and depends
//
// on whether or not the user used `register("env")`. This sub unit follows the same ideas as the sub unit above
//
//
package SkyLine_Environment

import (
	"fmt"
	"os"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func Environment_GetEnvironment(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	environmemt := os.Environ()
	Hashp := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
	for idx := 1; idx < len(environmemt); idx++ {
		key := &SkyEnv.SL_String{Value: environmemt[idx]}
		val := &SkyEnv.SL_String{Value: os.Getenv(environmemt[idx])}
		NewHashCreate := SkyEnv.HashPair{Key: key, Value: val}
		Hashp[key.SL_HashKeyType()] = NewHashCreate
	}
	return &SkyEnv.SL_HashMap{Pairs: Hashp}
}

func Environment_GetEnvironmentPath(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if len(arguments) != 1 {
		return &SkyEnv.SL_Error{Message: fmt.Sprintf("Argument error (Standard Library): wwrong number of arguments! this function requires %d and you gave `%d` argument(s)", 1, len(arguments))}
	}
	if arguments[0].SkyLine_ObjectFunction_GetDataType() != SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT {
		return &SkyEnv.SL_Error{
			Message: fmt.Sprintf("Argument error (Standard Library): Data type error! This function requires a data type of type STRING but you gave %s", arguments[0].SkyLine_ObjectFunction_GetDataType()),
		}
	}
	input := arguments[0].(*SkyEnv.SL_String).Value
	return &SkyEnv.SL_String{Value: os.Getenv(input)}
}

func Environment_SetEnvironment(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if len(arguments) != 2 {
		return &SkyEnv.SL_Error{
			Message: fmt.Sprintf("Argument error (Standard Library): Wrong number of arguments! This function requires %d and you gave `%d` argument(s)", 2, len(arguments)),
		}
	}
	if arguments[0].SkyLine_ObjectFunction_GetDataType() != SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT {
		return &SkyEnv.SL_Error{
			Message: fmt.Sprintf("argument must be a string, got=%s", arguments[0].SkyLine_ObjectFunction_GetDataType()),
		}
	}
	if arguments[1].SkyLine_ObjectFunction_GetDataType() != SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT {
		return &SkyEnv.SL_Error{
			Message: fmt.Sprintf("argument must be a string, got=%s", arguments[1].SkyLine_ObjectFunction_GetDataType()),
		}
	}
	name := arguments[0].(*SkyEnv.SL_String).Value
	value := arguments[1].(*SkyEnv.SL_String).Value
	os.Setenv(name, value)
	return &SkyEnv.SL_NULL{}
}
