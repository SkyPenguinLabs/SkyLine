//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _              _____               _         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___ |_   _|___ ___ _____|_|___ ___| |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|  | | -_|  _|     | |   | .'| | |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|      | |___|_| |_|_|_|_|_|_|__,| |_|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Defines -> This section of the standard library contains information for the console based functions that include frontend based functions such as tables, organizations,
//
// data analytics, informational organization, color, output etc.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the functions that are related to the console based functions in a way that these functions call and prepare then return resulting arguments
//
package SkyLine_StandardLib_Console

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	Helpers "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
	"fmt"
)

func Console_Frontend_GenerateAnsiEscapeSequence(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"console.HtmlToAnsi",
		Arguments,
		Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(2, 3),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		var returnonlystraightcode bool
		if len(Arguments) == 3 {
			if v, ok := Arguments[2].(*SkyEnv.SL_Boolean); ok {
				returnonlystraightcode = v.Value
			} else {
				return &SkyEnv.SL_Error{Message: "Sorry, third positional argument in console.HtmlToAnsi needs to be a boolean data type"}
			}
		} else {
			returnonlystraightcode = false
		}
		code := Arguments[0].(*SkyEnv.SL_String).Value
		Message := Arguments[1].(*SkyEnv.SL_String).Value
		r, g, b, x := ParseRGB(code)
		if x != nil {
			return &SkyEnv.SL_Error{Message: x.Error()}
		}
		resultingString := fmt.Sprintf("%s%s", ForegroundANSI(r, g, b, returnonlystraightcode), Message)
		return &SkyEnv.SL_String{Value: resultingString}
	}
}

func AssignValues(Arguments []string) {
	defaults := []*string{
		&T.ColumnTitleColor,
		&T.HeaderCrossOriginL_Color,
		&T.HeaderCrossOriginR_Color,
		&T.CrossLineX_Color,
		&T.CrossLineY_Color,
		&T.CrossLineX,
		&T.CrossLineY,
		&T.HeaderCrossOriginR,
		&T.HeaderCrossOriginL,
	}

	for i, arg := range Arguments {
		if i < len(defaults) {
			*defaults[i] = arg
		} else {
			break
		}
	}
}

func Console_Frontend_TableNewCreation(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"console.TableNew",
		Arguments,
		Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(0, 9),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	valuemap := make([]string, 0)
	for _, arg := range Arguments {
		valuemap = append(valuemap, arg.(*SkyEnv.SL_String).Value)
	}
	AssignValues(valuemap)
	return &SkyEnv.SL_NULL{}
}

func Console_Frontend_TableCreation(Arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"console.Table",
		Arguments,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(2),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	// Now we need to convert and find every single array within the argument list
	TwoDimensionalArray := make([][]string, 0)
	SLArrayTableRows := Arguments[0].(*SkyEnv.SL_Array).Elements

	for _, slRow := range SLArrayTableRows {
		if slRowArray, ok := slRow.(*SkyEnv.SL_Array); ok {
			row := make([]string, len(slRowArray.Elements))
			for j, slElement := range slRowArray.Elements {
				if slString, ok := slElement.(*SkyEnv.SL_String); ok {
					row[j] = slString.Value
				} else {
					return &SkyEnv.SL_Error{Message: "Array value contains non-string element"}
				}
			}
			TwoDimensionalArray = append(TwoDimensionalArray, row)
		} else {
			return &SkyEnv.SL_Error{Message: "Array value was not two-dimensional"}
		}
	}

	SLArrayColumns := make([]string, 0)
	columns := Arguments[1].(*SkyEnv.SL_Array).Elements
	for i := 0; i < len(columns); i++ {
		if ColumnValue, ok := columns[i].(*SkyEnv.SL_String); ok {
			SLArrayColumns = append(SLArrayColumns, ColumnValue.Value)
		}
	}
	return &SkyEnv.SL_String{Value: Console_Lib_DrawTableSepColumnBased(TwoDimensionalArray, SLArrayColumns)}
}
