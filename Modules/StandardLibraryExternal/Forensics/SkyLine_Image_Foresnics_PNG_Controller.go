package SkyLine_Standard_External_Forensics

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
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
// - File name     : SkyLine_Image_Foresnics_PNG_Controller.go
//
// - File contains : All the functions required for running the library, these are known as controller functions to the main definition

// Meta chunk type to a string
func (META *PNG_Meta) R_TO_STR() string {
	enc_hex := fmt.Sprintf("%x", META.Chunk.Type)
	dec_hex, _ := hex.DecodeString(enc_hex)
	return fmt.Sprintf("%s", dec_hex)
}

// Convert a string to a unsigned integer binary
func (META *PNG_Meta) STR_TO_I(in string) uint32 {
	return binary.BigEndian.Uint32([]byte(in))
}

// Preset readers
func (META *PNG_Meta) Write_Marshal() *bytes.Buffer {
	var byter = new(bytes.Buffer)
	META.W_Bin(byter, binary.BigEndian, META.Chunk.Size)
	META.W_Bin(byter, binary.BigEndian, META.Chunk.Type)
	META.W_Bin(byter, binary.BigEndian, META.Chunk.FD)
	META.W_Bin(byter, binary.BigEndian, META.Chunk.CRC)
	return byter
}
