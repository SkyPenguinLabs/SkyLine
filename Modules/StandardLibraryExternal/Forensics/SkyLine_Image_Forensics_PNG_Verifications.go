package SkyLine_Standard_External_Forensics

import (
	"bytes"
	"encoding/binary"
	"log"
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
// - File name     : SkyLine_Image_Forensics_PNG_Verifications.go
//
// - File contains : All functions to verify certain signatures or to verify certain steps and required data prior to injection or forensics
//
//

// Categorization
func (META *PNG_Meta) Verify_CRITICAL_Chunk() (chunk_t string) {
	farchar := string([]rune(META.R_TO_STR())[0])
	switch farchar {
	case strings.ToUpper(farchar):
		chunk_t = "Critical"
	default:
		chunk_t = "Ancillary"
	}
	return chunk_t
}

// Verify that the image is a PNG image
func (META *PNG_Meta) Verify_Image_Is_PNG(byter *bytes.Reader) {
	var header PNG_Header
	META.R_BIN(byter, binary.BigEndian, &header.HEAD)
	arr := make([]byte, 8)
	binary.BigEndian.PutUint64(arr, header.HEAD)
	if string(arr[1:4]) != "PNG" {
		log.Fatal("Provided file is not a valid PNG format")
	}
}
