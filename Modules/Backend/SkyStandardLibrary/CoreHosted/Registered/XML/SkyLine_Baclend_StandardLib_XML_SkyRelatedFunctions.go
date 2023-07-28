//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             __ __     _____     __
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |___|     |___|  |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|-   -|___| | | |___|  |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__|   |_|_|_|   |_____|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
// This section of the standard library is dedicated to dumping, mapping, matching, adding, generating, parsing, running or loading XML files which can also be parsed
//
// as PLIST files. However, given that PLIST files are technically XML, we put them under their own library as BPLIST ( Binary Property List ) works with regular PLIST
//
// parsers and programs.
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all functions that will be registered as standard library functions. This means that these are the initated functions before calling the primary functions
//
package SkyLine_Standard_Library_XML

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	SkySTDLibHelp "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

func XML_CallToConveryJsonToXML(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"xml.FromJson",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		file := Arguments[0].(*SkyEnv.SL_String).Value
		if file != "" {
			f, x := os.Open(file)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			defer f.Close()
			content, x := ioutil.ReadAll(f)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			var Data map[string]interface{}
			x = json.Unmarshal(content, &Data)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			root := ConstructPrimeXMLTree(Data)
			bytes, x := xml.MarshalIndent(root, "", " ")
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			return &SkyEnv.SL_String{Value: string(bytes)}
		} else {
			return &SkyEnv.SL_Error{Message: "In call to xml.FromJson -> Requires a real file name, got Null"}
		}
	}
}

// This function will take a input XML file and convert it to JSON
func XML_CallToConvertXMLtoJson(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"xml.ToJson",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		file := Arguments[0].(*SkyEnv.SL_String).Value
		if file != "" {
			f, x := os.Open(file)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			defer f.Close()
			Content, x := ioutil.ReadAll(f)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			var RootNode XML_NodeConverterXMLtoJson
			x = xml.Unmarshal(Content, &RootNode)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			jsondata, x := ConvertPrimeXMLNodeToJsonNode(RootNode)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			return &SkyEnv.SL_String{Value: string(jsondata)}
		} else {
			return &SkyEnv.SL_Error{Message: "In call to xml.ToJson() -> Requires a real file name, got NULL"}
		}
	}
}

func XML_CallToParse(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"xml.Parse",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		file := Arguments[0].(*SkyEnv.SL_String).Value
		if file != "" {
			f, x := os.Open(file)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			defer f.Close()
			content, x := ioutil.ReadAll(f)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			var rootnode XML_Node
			x = xml.Unmarshal(content, &rootnode)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			hash := XML_Lib_DumpIntoHashMap(rootnode) // assign hash values
			hash2 := ConvertXMLDataToHashMap(hash)
			return hash2
		} else {
			return &SkyEnv.SL_Error{Message: "In call to xml.Parse() -> Requires a real file name, got NULL"}
		}
	}
}
