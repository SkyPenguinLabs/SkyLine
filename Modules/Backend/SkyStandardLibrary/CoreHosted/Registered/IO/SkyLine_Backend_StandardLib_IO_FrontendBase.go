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
// This file defines all frontend related functions for the IO pack
//
//
package SkyLine_Standard_Library_IO

import (
	"bufio"
	"os"
	"strings"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkySTDLibHelp "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"

	"fmt"
)

func IO_Clear() SkyEnv.SL_Object {
	fmt.Println(Linux_Clear)
	return &SkyEnv.SL_NULL{}
}

func IO_ReturnColor() SkyEnv.SL_Object {
	fmt.Println(Linux_FReturn)
	return &SkyEnv.SL_NULL{}
}

func IO_Input(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"io.input",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(2),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		// Construct arguments
		prompt := Arguments[0].(*SkyEnv.SL_String).Value
		ExpectedType := Arguments[1].(*SkyEnv.SL_String).Value
		var expected SkyEnv.SL_Object
		switch strings.ToLower(ExpectedType) {
		case "string":
			expected = &SkyEnv.SL_String{}
		case "integer":
			expected = &SkyEnv.SL_Integer{}
		case "boolean":
			expected = &SkyEnv.SL_Boolean{}
		case "float":
			expected = &SkyEnv.SL_Float{}
		default:
			expected = &SkyEnv.SL_String{}
		}
		fmt.Print(prompt)
		retret := bufio.NewReader(os.Stdin)
		var out string
		for {
			out, _ = retret.ReadString('\n')
			out = strings.Replace(out, "\n", "", -1)
			if v, ok := InputMapType[expected.SkyLine_ObjectFunction_GetDataType()]; ok {
				return v(out)
			} else {
				fmt.Println("Error with output reading params")
				fmt.Println(" | out=" + out)
				fmt.Println(" | func=", InputMapType[expected.SkyLine_ObjectFunction_GetDataType()])
				fmt.Println(" | Expected=" + expected.SkyLine_ObjectFunction_GetDataType())
			}
		}
	}
}
