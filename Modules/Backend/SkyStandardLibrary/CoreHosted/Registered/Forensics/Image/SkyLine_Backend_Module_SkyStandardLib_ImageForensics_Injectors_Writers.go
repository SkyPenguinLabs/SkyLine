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
// This file defines all injection utilities that are dedicated to injecting images with files, writing to files, editing files, cutting chunks out and more.
//
// These can be used to inject regular image formats or inject and hide specific secrets within the image.
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func SkyLine_Forensics_InjectRegularImage_Payload(filename, payload string) bool {
	f, x := os.OpenFile(filename, os.O_WRONLY, 0644)
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	/*
		mov eax, 0x0    ; SEEK_SET = 0
		mov ebx, 0x6    ; OFFSET   = 6
		mov ecx, [f]    ; load file pointer
		mox edx, 0x0    ; WHENCE = 0
		call FileSeek	; Call
	*/
	SkyLine_Forensics_Controllers_FileSeek(f, 6, 0)
	/*
		mov eax, [START]    ; Start addres
		mov ebx, [f]        ; file pointer
		call FileWrite		; Calls
	*/
	SkyLine_Forensics_Controllers_FileWrite(f, Payload_Image_1_Start)
	/*
	   mov eax, 0x2      ; SEEK_CUR = 2
	   mov ebx, 0x0      ; OFFSET = 0
	   mov ecx, [f]      ; Load the file pointer
	   mov edx, 0x0      ; WHENCE = 0
	   call FileSeek     ; Call
	*/
	SkyLine_Forensics_Controllers_FileSeek(f, 0, 2)
	SkyLine_Forensics_Controllers_FileWrite(f, Payload_Image_1_Middle)
	SkyLine_Forensics_Controllers_FileWrite(f, []byte(payload))
	SkyLine_Forensics_Controllers_FileWrite(f, Payload_Image_1_End)
	return true
}

func SkyLine_Forensics_InjectPng_Payload(outfile string, Offset int64, decode bool, permission int, reader *bytes.Reader, b []byte) bool {
	// output_file defines the output image to write to                           | Exmaple : out.png
	// offset will define the offset the data will start being injected at        | Example : 0x8999
	// Decode will give you the option to overwrise and decode the encoded chunk  | Example : true
	// Permission will be a set of codes or permissions when creating the file    | Example : 0671
	w, x := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, fs.FileMode(permission))
	if x != nil {
		log.Fatal("Fatal: Problem writing to the output file!")
	}
	reader.Seek(0, 0)
	var buff = make([]byte, Offset)
	reader.Read(buff)
	w.Write(buff)
	w.Write(b)
	//if c.Decode {
	//	r.Seek(int64(len(b)), 1) // right bitshift to overwrite encode chunk
	//}
	_, x = io.Copy(w, reader)
	return x == nil
}

// Controller function | Injects a file into another file
func Image_Controller_Function_External_2_Inject_Into_File(input, output, file string) bool {
	in, x := os.Open(input)
	if x != nil {
		log.Fatal(x)
		return false
	}
	defer in.Close()
	f, x := os.Open(file)
	if x != nil {
		log.Fatal(x)
		return false
	}
	defer f.Close()
	out, x := os.Create(output)
	if x != nil {
		log.Fatal(x)
		return false
	}
	defer out.Close()
	_, x = io.Copy(out, in)
	if x != nil {
		log.Fatal(x)
		return false
	}
	_, x = io.Copy(out, f)
	if x != nil {
		log.Fatal(x)
		return false
	}
	return true
}

func (SessionSets *Png_Module_Settings) Injection_Standard_Payload(reader *bytes.Reader) bool {
	var META PNG_MetaData
	META.Verify_PngHeader(reader)
	if SessionSets.ImageOffset != 0x00 &&
		SessionSets.ImageOffset != 0 &&
		SessionSets.Payload != "" &&
		SessionSets.OutputFile != "" &&
		SessionSets.FileMode != 0 {
		var metac PNG_MetaData
		metac.Chunk.FieldData = []byte(SessionSets.Payload)
		metac.Chunk.Type = metac.SkyLine_Forensics_Converters_AtoB(SessionSets.InjectableChunk)
		metac.Chunk.Size = metac.SkyLine_Forensics_Controllers_CreateChunk_Size()
		metac.Chunk.CRC = metac.SkyLine_Forensics_Controllers_CreateChunk_CRC32()
		binmarshal := metac.SkyLine_Forensics_Controllers_CreateMarshaler()
		binmarshalbytes := binmarshal.Bytes()
		return SkyLine_Forensics_InjectPng_Payload(SessionSets.OutputFile, SessionSets.ImageOffset, false, int(SessionSets.FileMode), reader, binmarshalbytes)
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
