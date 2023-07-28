//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ _____ _____     _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |_   _|_   _|  _  |___| __  |___ ___ _ _ ___ ___| |_ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|     | | |   | | |   __|___|    -| -_| . | | | -_|_ -|  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__| |_|   |_| |__|      |__|__|___|_  |___|___|___|_| |___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|               |_|
//////////////////////////////////////////////////////////////////////////
//
//
// This file defines all transport related functions or methods to create and auto use specific HTTP requests
//
//
package SkyLine_Standard_Library_HTTP

import (
	"fmt"
	"strings"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	Helpers "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

func Creator_CreateNewBody(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"http.TransportNewReader",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	TransportNewSL.ResponseReaderBod = strings.NewReader(args[0].(*SkyEnv.SL_String).Value)
	return &SkyEnv.SL_NULL{}
}

// Method fills the HttpRequest structure with data
func Creator_HttpNew(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"http.New",
		args,
		Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(0, 7),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_BOOLEAN_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_BOOLEAN_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT,
		),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	a := len(args)
	if a >= 1 {
		Httpstruct.Req_URL = args[0].(*SkyEnv.SL_String).Value
	}
	if a >= 2 {
		Httpstruct.Req_Method = args[1].(*SkyEnv.SL_String).Value
	}
	if a >= 3 {
		Httpstruct.Tor = args[2].(*SkyEnv.SL_Boolean).Value
	}
	if a >= 4 {
		Httpstruct.OutToFile = args[3].(*SkyEnv.SL_Boolean).Value
	}
	if a >= 5 {
		Httpstruct.Filename = args[4].(*SkyEnv.SL_String).Value
	}
	if a >= 6 {
		Httpstruct.TorProx = args[5].(*SkyEnv.SL_String).Value
	}
	if a >= 7 {
		if k, ok := args[6].(*SkyEnv.SL_Array); ok {
			for i := 0; i < len(k.Elements); i++ {
				if conv_value, ok := k.Elements[i].(*SkyEnv.SL_String); ok {
					if strings.Contains(conv_value.Value, ":") {
						parts := strings.SplitN(conv_value.Value, ":", 2)
						if len(parts) != 2 {
							return &SkyEnv.SL_Error{Message: "Header Parse error: Header format must be (key:value) or (header:value), the lenggth was not two for splitting parts"}
						}
						key := strings.TrimSpace(parts[0])
						value := strings.TrimSpace(parts[1])
						if key != "" || value != "" {
							Httpstruct.Headers = append(Httpstruct.Headers, conv_value.Value)
						} else {
							fmt.Println("[Warn]->HTTP : When parsing http headers ")
							fmt.Println("[Warn]->HTTP : Key value was empty    ? ", key == "")
							fmt.Println("[Warn]->HTTP : Value value was empty  ? ", value == "")
						}
					}
				}
			}
		} else {
			return &SkyEnv.SL_Error{Message: "Sorry but the data type for this argument must be a array"}
		}
	}
	return &SkyEnv.SL_NULL{}
}
