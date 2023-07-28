//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _              _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___ |   __|___ ___ ___ ___ ___|_|___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___ |   __| . |  _| -_|   |_ -| |  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|    |__|  |___|_| |___|_|_|___|_|___|___|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
// This library defines all forensics related functions and libraries which allow you to dissect, inject or mess around with images or files in some shape or form while also inspecting
//
// specific file formats such as PNG, JPEG, GIF, BMP, ELF, PE, EXE and etc files.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the functions that will be called by registry
//
//
//
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"log"
	"os"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkySTDLibHelp "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

// Alias for getting chunks
var Metadata PNG_MetaData

func NewHash(pairs map[SkyEnv.HashKey]SkyEnv.HashPair) *SkyEnv.SL_HashMap {
	return &SkyEnv.SL_HashMap{Pairs: pairs}
}

func NewArray(elements []SkyEnv.SL_Object) *SkyEnv.SL_Array {
	return &SkyEnv.SL_Array{Elements: elements}
}

func ConvertToSL_HASH(data map[string][]string) SkyEnv.SL_Object {
	hashPairs := make(map[SkyEnv.HashKey]SkyEnv.HashPair)

	for key, values := range data {
		slKey := &SkyEnv.SL_String{Value: key}
		slValues := make([]SkyEnv.SL_Object, len(values))
		for i, value := range values {
			slValues[i] = &SkyEnv.SL_String{Value: value}
		}

		hashPair := SkyEnv.HashPair{
			Key:   slKey,
			Value: NewArray(slValues),
		}
		hashPairs[slKey.SL_HashKeyType()] = hashPair
	}

	return NewHash(hashPairs)
}

func SkyLine_Forensics_Image_CallLocateArchive(Argumnts ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"image.FindArchive",
		Argumnts,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	file := Argumnts[0].(*SkyEnv.SL_String).Value
	hex, res := Image_Controller_Function_External_1_Verify_Archive(file)
	if res {
		return &SkyEnv.SL_String{Value: hex}
	} else {
		return &SkyEnv.SL_String{Value: "false"}
	}
}

func SkyLine_Forensics_Image_CallInjectImage(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"image.InjectFiles",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(2, 3),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	// Out : Output file
	// In  : Art
	var In, FileInj, Out string
	In = Arguments[0].(*SkyEnv.SL_String).Value
	if len(Arguments) != 3 {
		Out = In
	} else {
		Out = Arguments[2].(*SkyEnv.SL_String).Value
	}
	FileInj = Arguments[1].(*SkyEnv.SL_String).Value
	return &SkyEnv.SL_Boolean{
		Value: Image_Controller_Function_External_2_Inject_Into_File(In, Out, FileInj),
	}
}

func SkyLine_Forensics_Image_CallGetMetaChunkData(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"image.DumpMeta",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	data, x := os.Open(Arguments[0].(*SkyEnv.SL_String).Value)
	if x != nil {
		log.Fatal(x)
	}
	defer data.Close()
	Reader, x := SkyLine_Forensics_Controllers_ProcessImage(data)
	if x != nil {
		log.Fatal(x)
	}
	m := Metadata.PNG_MetaData_GetAll(Reader)
	return ConvertToSL_HASH(m)
}

func SkyLine_Forensics_Image_CallGGetChunkName(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"forensics.ChunkType",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(2),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		log.Fatal(x)
	}
	data, x := os.Open(Arguments[0].(*SkyEnv.SL_String).Value)
	if x != nil {
		log.Fatal(x)
	}
	defer data.Close()
	Reader, x := SkyLine_Forensics_Controllers_ProcessImage(data)
	if x != nil {
		log.Fatal(x)
	}
	return &SkyEnv.SL_String{
		Value: Metadata.PNG_MetaData_GetChunkType(Reader, Arguments[1].(*SkyEnv.SL_String).Value),
	}
}

func SkyLine_Forensics_Image_CallGetNumChunks(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"forensics.ChunkCount",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		log.Fatal(x)
	}
	data, x := os.Open(Arguments[0].(*SkyEnv.SL_String).Value)
	if x != nil {
		log.Fatal(x)
	}
	defer data.Close()
	Reader, x := SkyLine_Forensics_Controllers_ProcessImage(data)
	if x != nil {
		log.Fatal(x)
	}
	return &SkyEnv.SL_String{
		Value: Metadata.PNG_MetaData_GetChunkCount(Reader),
	}
}

func SkyLine_Forensics_Image_CallGetOffsets(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"foresnics.LoadOffsets",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		log.Fatal(x)
	}
	image := Arguments[0].(*SkyEnv.SL_String).Value
	data, x := os.Open(image)
	if x != nil {
		log.Fatal(x)
	}
	defer data.Close()
	Reader, x := SkyLine_Forensics_Controllers_ProcessImage(data)
	if x != nil {
		log.Fatal(x)
	}
	Offsets := Metadata.PNG_MetaData_Offsets(Reader)
	arr := make([]SkyEnv.SL_Object, 0)
	for i := 0; i < len(Offsets); i++ {
		arr = append(arr, &SkyEnv.SL_String{Value: Offsets[i]})
	}
	return &SkyEnv.SL_Array{Elements: arr}
}

func SkyLine_Forensics_Image_CallCreateImage(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"forensics.Create",
		Arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(4),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT,
			SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		log.Fatal(x)
	}
	typeof := Arguments[3].(*SkyEnv.SL_String).Value
	out := Arguments[0].(*SkyEnv.SL_String).Value
	pw := Arguments[1].(*SkyEnv.SL_String).Value
	ph := Arguments[2].(*SkyEnv.SL_String).Value
	if out == "" && pw == "" && ph == "" {
		out = "OutputFileSkyline_Environment_File_Name_Null"
		pw = "600"
		ph = "1000"
	}
	CreateImageSettings.Output = out
	CreateImageSettings.PixelHeight = ph
	CreateImageSettings.PixelWidth = pw
	return &SkyEnv.SL_Boolean{
		Value: SkyLine_Forensics_Module_Image_GenerateImage(typeof),
	}
}
