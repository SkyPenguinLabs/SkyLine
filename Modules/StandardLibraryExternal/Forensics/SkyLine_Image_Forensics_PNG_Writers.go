package SkyLine_Standard_External_Forensics

import (
	"bytes"
	"io"
	"io/fs"
	"log"
	"os"
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
// - File name     : SkyLine_Image_Forensics_PNG_Writers.go
//
// - File contains : This file contains a writer to write and inject the data into the image

// STAT CODES

const (
	SUCCESS = 9000
	FAIL    = -9000
)

func Injection_Writer(output_file string, Offset int64, decode bool, permission int, reader *bytes.Reader, b []byte) bool {
	// output_file defines the output image to write to                           | Exmaple : out.png
	// offset will define the offset the data will start being injected at        | Example : 0x8999
	// Decode will give you the option to overwrise and decode the encoded chunk  | Example : true
	// Permission will be a set of codes or permissions when creating the file    | Example : 0671
	w, x := os.OpenFile(output_file, os.O_RDWR|os.O_CREATE, fs.FileMode(permission))
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
