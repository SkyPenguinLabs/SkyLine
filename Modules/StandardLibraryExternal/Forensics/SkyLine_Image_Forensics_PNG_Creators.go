package SkyLine_Standard_External_Forensics

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
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
// - File name     : SkyLine_Image_Forensics_PNG_Creators.go
//
// - File contains : Contains image creations like creating the chunks and more for the image files

// Creates CRC chunk
func (META *PNG_Meta) Create_CRC() uint32 {
	var Out = new(bytes.Buffer)
	META.W_Bin(Out, binary.BigEndian, META.Chunk.Type)
	META.W_Bin(Out, binary.BigEndian, META.Chunk.FD)
	return crc32.ChecksumIEEE(Out.Bytes())
}

// Create image chunk size by returning the length of the data
func (META *PNG_Meta) Create_Sz() uint32 {
	return uint32(len(META.Chunk.FD))
}
