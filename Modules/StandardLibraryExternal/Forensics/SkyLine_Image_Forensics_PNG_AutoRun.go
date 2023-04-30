package SkyLine_Standard_External_Forensics

import (
	"encoding/binary"
	"fmt"
	"io"
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
// - File name     : SkyLine_Image_Forensics_PNG_AutoRun.go
//
// - File contains : Contains the binary readers and writer functions

func (META *PNG_Meta) W_Bin(w io.Writer, order binary.ByteOrder, data interface{}) {
	if x := binary.Write(w, order, data); x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_CODE_COULD_NOT_WRITE_BINARY_DATA_IN_WBPD),
			"Could not write the binary data because "+fmt.Sprint(x),
			"Data could have been corrupted, make sure data you are trying to write is valid",
		)
	}
}

func (META *PNG_Meta) R_BIN(r io.Reader, order binary.ByteOrder, data interface{}) {
	if x := binary.Read(r, order, data); x != nil {
		if x == io.EOF {
			PrepareErrorAndLog(
				x,
				fmt.Sprint(ERROR_CODE_COULD_NOT_READ_BINARY_DATA_IN_RBPD),
				"IO ( Reader ) failed to read the binary data any further due to EOF",
				"Data might have been corrupted or we just found the end of the file.",
			)
		} else {
			PrepareErrorAndLog(
				x,
				fmt.Sprint(ERROR_CODE_COULD_NOT_READ_BINARY_DATA_IN_RBPD),
				"Could not read the binary data because "+fmt.Sprint(x),
				"Data could have been corrupted, make sure the data being read in the image or file is not corrupted or damaged",
			)
		}
	}
}
