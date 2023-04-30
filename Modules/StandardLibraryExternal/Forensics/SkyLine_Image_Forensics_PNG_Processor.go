package SkyLine_Standard_External_Forensics

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// This path of (n) contains all digital forensics based functions for images, files, memory, binary files and much more among that list.
//
// n = StandardLibraryExternal/Forensics
//
// This type of module contains much more files and information for general digital forensics and steganography. This includes functions and
// types as well as parsers that can look up signatures in images, verify unifentified files, pull data from files and images, extract signatures
// and uses the syscall table to read and write data to images with an option that allows you to encrypt and then inject data into an image. Currently
// this library only works with mainly PNG images but also has support for JPG/JPEG images and as far as general injection goes it works with most images
// which are common that includes BMP, JPG, JPEG, PNG, GIF, WEBP, WebM and other various file formats.
//
//
// - Date of start : Wed 01 Mar 2023 02:18:32 PM EST
//
// - File name     : SkyLine_Image_Forensics_PNG_Processor.go
//
// - File contains : Contains functions to grab information and to grab or injectd ata into the image

// PNG_Meta_GetChunkCount will grab the metadata of the image and print it out
// Represented as forensics.png.meta()
func (META *PNG_Meta) PNG_Meta_GetChunkCount(byter *bytes.Reader) string {
	META.Verify_Image_Is_PNG(byter)
	StartByte := 0
	var ct string
	for ct != SIG_END {
		//inmap := make(map[string]interface{})
		//META.R_O(byter) // Get the offset
		//META.R_C(byter) // Get the chunk or read it
		//inmap["Offset"] = fmt.Sprintf("%#02x", META.Offset)
		//inmap["ChunkLength"] = strconv.Itoa(int(META.Chunk.Size))
		//inmap["ChunkType"] = META.R_TO_STR()
		//inmap["Chunk_Critical"] = META.Verify_CRITICAL_Chunk()
		//inmap["Chunk_CRC"] = fmt.Sprintf("%x", META.Chunk.CRC)
		//fmt.Println("Chunk | " + strconv.Itoa(StartByte))
		//fmt.Printf("       Offset   : %#02x\n", META.Offset)
		//fmt.Printf("       C-Length : %s bytes \n", strconv.Itoa(int(META.Chunk.Size)))
		//fmt.Printf("       C-Type   : %s\n", META.R_TO_STR())
		//fmt.Printf("       C-Verif  : %s\n", META.Verify_CRITICAL_Chunk())
		//fmt.Printf("       C-CRC    : %x\n", META.Chunk.CRC)
		//metamap[fmt.Sprintf("Chunk_%s", fmt.Sprint(strconv.Itoa(StartByte)))] = inmap
		ct = META.R_TO_STR()
		StartByte++
	}
	return strconv.Itoa(StartByte)
}

// PNG_Meta_GetOffsets will grab the offsets of an image
func (META *PNG_Meta) PNG_Meta_GetOffsets(byter *bytes.Reader) []string {
	var results []string
	META.Verify_Image_Is_PNG(byter)
	var ct string
	for ct != SIG_END {
		results = append(results, fmt.Sprintf("%#02x", META.Offset))
	}
	return results
}

// Settings will load the general settings for the other programs, this is required before any injection can continue

type Settings struct {
	Key             string // Encryption key if any
	Decode          bool   // Decode payload before injection
	Encode          bool   // Encode payload before injection
	OutputFile      string // Output file resulting in the input file being injected
	InputFile       string // Input file to inject
	FileMode        int64  // File mode of created output file from resulting injection
	ImageOffset     int64  // Offset to inject at
	Payload         string // Payload to inject
	InjectableChunk string // Chunk to inject at
}

// We pass all settings as string to avoid useless conversions for the moment from type arguments
func (SessionSets *Settings) Settings_Inject_New(key, ofile, infile, filemode, decode, encode, offset, payload, chunktoinject string) {
	SessionSets.Key = key
	SessionSets.OutputFile = ofile
	SessionSets.InputFile = infile
	fm, x := strconv.ParseInt(filemode, 0, 64)
	if x != nil {
		fmt.Println("E | (F)->New() ")
		fmt.Println("           | Code    : Unknown")
		fmt.Println("           | Message : Could not parse the filemode as an integer, overlflowed integer 64 data type?")
		return
	} else {
		SessionSets.FileMode = fm
	}
	decode = strings.ToLower(decode)
	dec, x := strconv.ParseBool(decode)
	if x != nil {
		fmt.Println("E | (F)->New()")
		if decode != "true" && decode != "false" {
			fmt.Println("           | Code     : Unknown 0000_ffff  ")
			fmt.Println("           | Argument : encode=" + encode)
			fmt.Println("           | Message  : Boolean parse error, the value you have provided as a boolean value is not true and not false, you have -> ", decode)
		} else {
			fmt.Println("           | Code     : Unknown 0000_ffff")
			fmt.Println("           | Message  : Boolean parse error -> ", x)
		}
		fmt.Println("SkyLine self healing WARN: Fixing this value to be set to false")
		SessionSets.Decode = false
	} else {
		SessionSets.Decode = dec
	}
	encode = strings.ToLower(decode)
	enc, x := strconv.ParseBool(encode)
	if x != nil {
		fmt.Println("E | (F)->New()")
		if encode != "true" && encode != "false" {
			fmt.Println("           | Code     : Unknown 0000_ffff  ")
			fmt.Println("           | Argument : encode=" + encode)
			fmt.Println("           | Message  : Boolean parse error, the value you have provided as a boolean value is not true and not false, you have -> " + encode)
		} else {
			fmt.Println("           | Code     : Unknown 0000_ffff")
			fmt.Println("           | Message  : Boolean parse error -> ", x)
		}
		fmt.Println("SkyLine self healing WARN: Fixing this (encode=) value to be set to false")
		SessionSets.Encode = false
	} else {
		SessionSets.Encode = enc
	}
	off, x := strconv.ParseInt(offset, 0, 64)
	if x != nil {
		fmt.Println("E | (F)->New()")
		fmt.Println("           | Code     : Unknown 0000_ffff")
		fmt.Println("           | Message  : Integer parse error, could not parse given offset as a integer64 signed. Overflows int64 data type???---> ", offset)
		return
	} else {
		SessionSets.ImageOffset = off
	}
	SessionSets.InjectableChunk = chunktoinject
	SessionSets.Payload = payload
}

// Injection and other methods for PNG images, this requires the forensics.new() function to be called before hand to iniate the settings
func (SessionSets *Settings) Injection_Standard_Payload(reader *bytes.Reader) bool {
	var META PNG_Meta
	META.Verify_Image_Is_PNG(reader)
	if SessionSets.ImageOffset != 0x00 &&
		SessionSets.ImageOffset != 0 &&
		SessionSets.Payload != "" &&
		SessionSets.OutputFile != "" &&
		SessionSets.FileMode != 0 {
		var metac PNG_Meta
		metac.Chunk.FD = []byte(SessionSets.Payload)
		metac.Chunk.Type = metac.STR_TO_I(SessionSets.InjectableChunk)
		metac.Chunk.Size = metac.Create_Sz()
		metac.Chunk.CRC = metac.Create_CRC()
		binmarshal := metac.Write_Marshal()
		binmarshalbytes := binmarshal.Bytes()
		return Injection_Writer(SessionSets.OutputFile, SessionSets.ImageOffset, false, int(SessionSets.FileMode), reader, binmarshalbytes)
	} else {
		fmt.Println("W | SkyLine Forensics")
		fmt.Println("         | Code     : 0x01 ")
		fmt.Println("         | Message  : Payload or image offset was missing from forensics.PngSettingsNew()")
		fmt.Println("SL Debug -> Offset       ? ", (SessionSets.ImageOffset == 0x00))
		fmt.Println("SL Debug -> Payload      ? ", (SessionSets.Payload == ""))
		fmt.Println("SL Debug -> Output File  ? ", (SessionSets.OutputFile == ""))
		fmt.Println("SL Debug -> File Mode    ? ", (SessionSets.FileMode))
		return false
	}
}
