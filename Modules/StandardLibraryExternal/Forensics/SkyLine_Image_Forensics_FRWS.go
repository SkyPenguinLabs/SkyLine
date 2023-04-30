package SkyLine_Standard_External_Forensics

import (
	"fmt"
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
// - File name     : SkyLine_Image_Forensics_FRWS.go
//
// - File contains : All forms of readers and writers as well as seakers for automation
//

const (
	ERROR_FILE_OPERATION_FAILED_TO_WRITE = 70 // Failed to write data to the file
	ERROR_FILE_OPERATION_FAILED_TO_SEEK  = 71 // Failed to seek at or within the file
	ERROR_FILE_OPERATION_FAILED_TO_READ  = 72 // Failed to read data from the file
)

// File Write
func SkyLine_Image_Forensics_Controller_M2_F_W(f *os.File, data []byte) {
	if _, x := f.Write(data); x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_FILE_OPERATION_FAILED_TO_WRITE),
			"File ^Operation^ (Err when trying to write data to the file)",
			"SkyLine got error when writing data to the file -> "+fmt.Sprint(x),
		)
		return
	}
}

// File Read
func SkyLine_Image_Forensics_Controller_M2_F_R(f *os.File, data []byte) {
	if _, x := f.Read(data); x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_FILE_OPERATION_FAILED_TO_READ),
			"File ^Operation^ (Err when trying to read data from the file)",
			"SkyLine got the error when trying to read data from the file -> "+fmt.Sprint(x),
		)
		return
	}
}

// File Seek
func SkyLine_Image_Forensics_Controller_M2_F_S(f *os.File, OFFSET int64, E int) {
	if _, x := f.Seek(OFFSET, E); x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_FILE_OPERATION_FAILED_TO_SEEK),
			"File ^Operation^ (Err when trying to seek data from the file)",
			"SkyLine got error when seeking for offset ( "+fmt.Sprint(OFFSET)+" ) when using WHENCE ( "+fmt.Sprint(E)+" ) -> "+fmt.Sprint(x),
		)
		return
	}
}
