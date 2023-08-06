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
// This file defines converter functions to convert and return specific data types such as converting strings to real integers and data types
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

// Automates string conversion to integer
func SkyLine_Forensics_Converters_AtoI(data string) int {
	Data, x := strconv.Atoi(data)
	if x != nil {
		log.Fatal(x)
	}
	return Data
}

// Automates string conversion to unsigned integer binary
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Converters_AtoB(data string) uint32 {
	return binary.BigEndian.Uint32([]byte(data))
}

// Automates writing meta chunk types to a string
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Converters_RToStr() string {
	encoded_hex := fmt.Sprintf("%x", PngMeta.Chunk.Type)
	decoded_hex, _ := hex.DecodeString(encoded_hex)
	return string(decoded_hex)
}
