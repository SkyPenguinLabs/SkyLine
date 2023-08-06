//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _                __ _____ _____ _____
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___   |  |   __|     |   | |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|  |  |__   |  |  | | | |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_____|_____|_|___|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This part of the standard library contains any and all functions for JSON data sets and functions. This includes dumping relative information, dumping file information,
//
// gnerating golang structures for plugins, generating infromation, generating strings, dumping to a hash map and so on from there.
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the initation or direct function calls that will be named as registered return functions
//
package SkyLine_Standard_Library_JSON

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkySTDLibHelp "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

func IniateCallToStructure(hash bool, arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"json.ToGo",
		arguments, SkySTDLibHelp.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(1, 2),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x == nil {
		var in []byte
		datachan := arguments[0].(*SkyEnv.SL_String).Value
		if CheckifFile(datachan) {
			in, _ = os.ReadFile(datachan)
		} else {
			reader := strings.NewReader(datachan)
			buffer := new(bytes.Buffer)
			_, x := buffer.ReadFrom(reader)
			if x != nil {
				return &SkyEnv.SL_Error{
					Message: fmt.Sprint(x),
				}
			}
			in = buffer.Bytes()
		}
		var data interface{}
		_ = json.Unmarshal(in, &data)
		if hash {
			return MapToHash(data)
		} else {
			if len(arguments) != 2 {
				return &SkyEnv.SL_Error{Message: "You are trying to convert json to go, give a primary structure tag as the second positional argument"}
			}
			result := arguments[1].(*SkyEnv.SL_String).Value
			return &SkyEnv.SL_String{
				Value: GenerateGolangStructure_JSON(data, result),
			}
		}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func IniateMapToHash(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"JSON.ToMap",
		args,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x == nil {
		return IniateCallToStructure(true, args...)
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}
