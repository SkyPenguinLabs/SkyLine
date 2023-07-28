//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____ _ _     _____         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_| |___|   __|_ _ ___| |_ ___ ______
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | | -_|__   | | |_ -|  _| -_|     |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |_|_|___|_____|_  |___|_| |___|_|_|_|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|                   |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This file defines all primary functions for the library
//
//
package SkyLine_Standard_Library_File

import (
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	Helpers "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"

	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
)

func FileLib_Carve(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"File.Carve",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(4),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
		),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	Input := args[0].(*SkyEnv.SL_String).Value
	Output := args[3].(*SkyEnv.SL_String).Value
	// Byte arrays
	StartByteArr := args[1].(*SkyEnv.SL_Array).Elements
	EndByteArr := args[2].(*SkyEnv.SL_Array).Elements
	// Converted arrays | Start Bytes
	StartBytes := make([]byte, 0)
	for i := 0; i < len(StartByteArr); i++ {
		if v, ok := StartByteArr[i].(*SkyEnv.SL_Integer); ok {
			StartBytes = append(StartBytes, byte(v.Value))
		}
	}
	// Converted Arrays | End bytes
	EndBytes := make([]byte, 0)
	for idx := 0; idx < len(EndByteArr); idx++ {
		if v, ok := EndByteArr[idx].(*SkyEnv.SL_Integer); ok {
			EndBytes = append(EndBytes, byte(v.Value))
		}
	}
	// Call function
	SkyLine_File_Lib_CarveSector(Input, StartBytes, EndBytes, Output)
	return &SkyEnv.SL_NULL{}
}

func FileLib_WriteFile(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.Write", args, Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x == nil {
		payload := args[0].SkyLine_ObjectFunction_GetTrueValue()
		if File.Filename == "" {
			return &SkyEnv.SL_Error{
				Message: "File sys error: Sorry, you provided 1 argument in call to File.Write which means SkyLine assumed you ran File.New() but it seems as if you did not. Please place the following instruction above this call -> INSTRUCT (`%s`)" + `File.New("NAME_OF_FILE")`,
			}
		} else {
			f, x := os.OpenFile(File.Filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if x != nil {
				return &SkyEnv.SL_Error{Message: "File Sys Error: Failed to open file due to -> " + fmt.Sprint(x)}
			} else {
				defer f.Close()
				_, x = f.WriteString(payload)
				if x != nil {
					return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
				} else {
					return &SkyEnv.SL_Boolean{Value: true}
				}
			}
		}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_ExtractLables_FileOut(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.Extractlables", args, Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x == nil {
		lable := args[0].(*SkyEnv.SL_String).Value
		labledata := FileLib_Extract_Lable(File.Lines, lable)
		if labledata != nil {
			var arr []SkyEnv.SL_Object
			for i := 0; i < len(labledata); i++ {
				arr = append(arr, &SkyEnv.SL_String{Value: labledata[i]})
			}
			return &SkyEnv.SL_Array{Elements: arr}
		} else {
			return &SkyEnv.SL_Error{Message: "Sorry but you must use File.New() to initate the file loading process before files can be properly read from."}
		}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_OverWrite_WriteFile(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.Overwrite", args, Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x == nil {
		payload := []byte(args[0].SkyLine_ObjectFunction_GetTrueValue())
		x = ioutil.WriteFile(File.Filename, payload, fs.FileMode(File.Mode))
		if x != nil {
			return &SkyEnv.SL_Error{Message: "File Sys error: Could not write to file -> " + fmt.Sprint(x)}
		} else {
			return &SkyEnv.SL_Boolean{Value: true}
		}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_IniateNewFunction(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.New", args, Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(1, 2), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x == nil {
		if len(args) == 2 {
			var x error
			File.Mode, x = strconv.Atoi(args[1].SkyLine_ObjectFunction_GetTrueValue())
			if x != nil {
				return &SkyEnv.SL_Error{Message: "Conversion Error: Could not conver `" + fmt.Sprint(args[0].SkyLine_ObjectFunction_GetTrueValue()) + "` to integer -> " + fmt.Sprint(x)}
			}
		}
		if args[0].SkyLine_ObjectFunction_GetTrueValue() != "" {
			_, x := os.Stat(args[0].SkyLine_ObjectFunction_GetTrueValue())
			if x != nil {
				return &SkyEnv.SL_Error{Message: "File Sys error: Could not stat file -> " + args[0].SkyLine_ObjectFunction_GetTrueValue() + " : because ( " + fmt.Sprint(x) + ")"}
			}
			File.Filename = args[0].SkyLine_ObjectFunction_GetTrueValue()
			File.Lines = FileLib_ReadLines_InitateData(File.Filename)
			return &SkyEnv.SL_Boolean{Value: true}
		} else {
			return &SkyEnv.SL_Error{Message: "File sys error: Refused to try opening file, data was NULL"}
		}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_OpenAndOutFile(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.Open", args, Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(0, 1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x == nil {
		var file string
		if len(args) == 1 {
			file = args[0].SkyLine_ObjectFunction_GetTrueValue()
			File.Filename = file
		} else {
			if File.Filename != "" {
				file = File.Filename
			} else {
				return &SkyEnv.SL_Error{
					Message: "File sys error: Sorry, you provided 0 arguments which means SkyLine assumed you ran File.New() but it seems as if you did not. Please place the following instruction above this call -> INSTRUCT (`%s`)" + `File.New("NAME_OF_FILE")`}
			}
		}
		dt, x := ioutil.ReadFile(file)
		if x != nil {
			return &SkyEnv.SL_Error{Message: "File sys error: Sorry, could not read file due to -> " + fmt.Sprint(x)}
		} else {
			return &SkyEnv.SL_String{Value: string(dt)}
		}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

// This function will check through adatabase structured of your choice as long as the mime type
// comes first or rather the byte order.
func FileLib_CheckFileMimeThroughDB(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.Mime", args, Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(1, 2), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		var dbload string
		if len(args) > 1 {
			dbload = args[1].(*SkyEnv.SL_String).Value
		} else {
			dbload = "Modules/Backend/SkyDB/ProgramaticFiles/FileSignatures.json"
		}
		file := args[0].(*SkyEnv.SL_String).Value
		resarr := make([]SkyEnv.SL_Object, 0)
		files := FileLib_Attempt_Mime_Type_Through_ShortDB(dbload, file)
		for i := 0; i < len(files); i++ {
			resarr = append(resarr, &SkyEnv.SL_String{Value: files[i]})
		}
		return &SkyEnv.SL_Array{Elements: resarr}
	}
}

func FileLib_GetFileHeader(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("File.Head", args, Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(1, 2), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		Range := args[1].(*SkyEnv.SL_Integer).Value
		File := args[0].(*SkyEnv.SL_String).Value
		header, ascii := FileLib_GrabHeader_OfSpecificSize(File, Range)
		concentrate := header + "@" + ascii
		return &SkyEnv.SL_String{Value: concentrate}
	}
}
