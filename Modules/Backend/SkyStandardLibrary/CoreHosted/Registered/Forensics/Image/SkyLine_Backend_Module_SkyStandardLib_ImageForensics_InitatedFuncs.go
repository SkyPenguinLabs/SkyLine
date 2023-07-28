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
// This file contains functions that are called by SkyLine related controllers and exporters. These are the prime functions to use for the file within this directory
//
// that is SkyRelated_Functions.go
//
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"bytes"
	"fmt"
	"strconv"
)

func (PngMeta *PNG_MetaData) PNG_MetaData_GetChunkType(reader *bytes.Reader, ChunkAddress string) string {
	PngMeta.Verify_PngHeader(reader)
	var ct string
	for ct != SIG_END {
		PngMeta.SkyLine_Fornesics_Controllers_Offset_R(reader)
		PngMeta.SkyLine_Forensics_Controllers_Chunk_R(reader)
		addr := fmt.Sprintf("%#02x", PngMeta.Offset)
		if addr == ChunkAddress {
			ChunkType := PngMeta.SkyLine_Forensics_Converters_RToStr()
			return ChunkType
		}
	}
	return "Unfound"
}
func (PNG_MetaData *PNG_MetaData) PNG_MetaData_GetAll(reader *bytes.Reader) map[string][]string {
	res := make(map[string][]string)
	PNG_MetaData.Verify_PngHeader(reader)
	count := 0
	var CurrentType string
	for CurrentType != SIG_END {
		PNG_MetaData.SkyLine_Fornesics_Controllers_Offset_R(reader)
		PNG_MetaData.SkyLine_Forensics_Controllers_Chunk_R(reader)
		TempMap := make([]string, 0)
		TempMap = append(TempMap,
			fmt.Sprintf("%#02x", PNG_MetaData.Offset),
			strconv.Itoa(int(PNG_MetaData.Chunk.Size)),
			PNG_MetaData.SkyLine_Forensics_Converters_RToStr(),
			PNG_MetaData.Verify_CriticalChunks(), fmt.Sprintf("%x", PNG_MetaData.Chunk.CRC))
		CurrentType = Metadata.SkyLine_Forensics_Converters_RToStr()
		res[strconv.Itoa(count)] = TempMap
		count++
	}
	return res
}

func (PngMeta *PNG_MetaData) PNG_MetaData_Offsets(reader *bytes.Reader) []string {
	var res []string
	PngMeta.Verify_PngHeader(reader)
	var ct string
	for ct != SIG_END {
		PngMeta.SkyLine_Fornesics_Controllers_Offset_R(reader)
		PngMeta.SkyLine_Forensics_Controllers_Chunk_R(reader)
		res = append(res, fmt.Sprintf("%#02x", PngMeta.Offset))
		ct = PngMeta.SkyLine_Forensics_Converters_RToStr()
	}
	return res
}

func (PngMeta *PNG_MetaData) PNG_MetaData_GetChunkCount(reader *bytes.Reader) string {
	PngMeta.Verify_PngHeader(reader)
	Start := 0
	var ct string
	for ct != SIG_END {
		Start++
		PngMeta.SkyLine_Fornesics_Controllers_Offset_R(reader)
		PngMeta.SkyLine_Forensics_Controllers_Chunk_R(reader)
		ct = PngMeta.SkyLine_Forensics_Converters_RToStr()
	}
	return strconv.Itoa(Start)
}
