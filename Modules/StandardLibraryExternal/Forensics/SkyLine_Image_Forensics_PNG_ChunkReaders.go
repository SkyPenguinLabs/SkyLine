package SkyLine_Standard_External_Forensics

import (
	"bytes"
	"encoding/binary"
	"log"
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
// - File name     : SkyLine_Image_Forensics_PNG_ChunkReaders.go
//
// - File contains : All the chunk and data readers for functions and calls
//
//
// Module organization -
//                  | C  -> Chunk
//                  | R  -> Read
//                  | S  -> Size
//                  | D  -> Data or bytes
//                  | C2 -> CRC
//                  | O  -> Offset
//                  | CR -> CRITICAL
//                  | T  -> Type
//
//
// Example : func R_O   = func read_offset
// Exmaple : func R_C2  = func read_CRC
// Example : func R_C_B = func read_chunk_bytes
//
//
// The reason we do this is for simplicity of call, while the name's may not be on point or human readable this code should not be touched fully in the future which
// means the code will not need to be read unless there are bugs, vulnerabilities etc which is why this note and pad is here
//

// Reads the offset
func (META *PNG_Meta) R_O(byter *bytes.Reader) { // O
	META.Offset, _ = byter.Seek(0, 1)
}

// Reads the chunks bytes of data
func (META *PNG_Meta) readChunkBytes(b *bytes.Reader, cLen uint32) {
	META.Chunk.FD = make([]byte, cLen)
	if err := binary.Read(b, binary.BigEndian, &META.Chunk.FD); err != nil {
		log.Fatal(err)
	}
}

// Reads the chunk type
func (META *PNG_Meta) R_C_T(byter *bytes.Reader) { // T
	META.R_BIN(byter, binary.BigEndian, &META.Chunk.Type)
}

// Reads the chunk size
func (META *PNG_Meta) R_C_S(byter *bytes.Reader) { // S
	META.R_BIN(byter, binary.BigEndian, &META.Chunk.Size)
}

// Reads the CRC chunk
func (META *PNG_Meta) R_C_C2(byter *bytes.Reader) { // C2
	META.R_BIN(byter, binary.BigEndian, &META.Chunk.CRC)
}

// Read the chunk
func (META *PNG_Meta) R_C(byter *bytes.Reader) { // C
	META.R_C_S(byter)
	META.R_C_T(byter)
	META.readChunkBytes(byter, META.Chunk.Size)
	META.R_C_C2(byter)
}
