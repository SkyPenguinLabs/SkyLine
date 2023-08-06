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
// This file defines all of the functions to verify specific file chunks or image chunks
//
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"bytes"
	"encoding/binary"
	"log"
	"strings"
)

func (PngMeta *PNG_MetaData) Verify_CriticalChunks() (ChunkType string) {
	FarChar := string([]rune(PngMeta.SkyLine_Forensics_Converters_RToStr())[0])
	switch FarChar {
	case strings.ToUpper(FarChar):
		ChunkType = "Critical"
	default:
		ChunkType = "Ancillary"
	}
	return ChunkType
}

func (PngMeta *PNG_MetaData) Verify_PngHeader(reader *bytes.Reader) {
	var head PNG_Image_Header
	PngMeta.SkyLine_Forensics_Controllers_BinaryR(reader, binary.BigEndian, &head.HEADER)
	HeadArray := make([]byte, 8) // Header is 8 bytes long
	binary.BigEndian.PutUint64(HeadArray, head.HEADER)
	if string(HeadArray[1:4]) != "PNG" {
		log.Fatal("provided image file does not have a valid image header")
	}
}
