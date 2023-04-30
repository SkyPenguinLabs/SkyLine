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
// - File name     : SkyLine_Forensics_Library_Codes.go
//
// - File contains : All codes and error systems used within this library

// ERROR CODES
const (
	ERROR_CODE_FILE_NOT_FOUND                      = 10 // FIle or image not found
	ERROR_CODE_FILE_IS_DIRECTORY                   = 11 // supplied data is a directory
	ERROR_CODE_FILE_COULD_NOT_STAT                 = 12 // File supplied as image could not be stated
	ERROR_CODE_COULD_NOT_PARSE_OFFSET              = 20 // Offset supplied is not a real integer.
	ERROR_CODE_COULD_NOT_CREATE_OUTPUT_FILE        = 21 // File supplied as an output file could not be created due to ...
	ERROR_CODE_COULD_NOT_WRITE_BINARY_DATA_IN_WBPD = 30 // Could not write the data using binary.write()  FROM ->*PNG_Meta.W_BIN()->UINT32_x
	ERROR_CODE_COULD_NOT_READ_BINARY_DATA_IN_RBPD  = 35 // Could not read the data using binary.read() FROM->*PNG_Meta.R_BIN()->UINT32_X
	ERROR_READ_BYTE_BUFFER_COULD_NOT_READ_BYTE     = 40 // Buffer for file could not read the next byte within the image
	ERROR_READ_BYTE_BUFFER_COULD_NOT_PEAK_AT       = 41 // Buffer for file could not peek the next byte at ...
	ERROR_OS_OPERATION_SET_FAILED_CREATE           = 50 // Operating System Operation has failed at create
	ERROR_IO_OPERATION_SET_FAILED_COPY             = 60 // IO operation has failed at copy
	ERROR_IO_OPERATION_SET_FAILED_READ_FILE        = 61 // IO Operation has failed to open the file
)

func PrepareErrorAndLog(x error, arguments ...string) {
	if x != nil {
		fmt.Println("     | Code    -> " + arguments[0]) // 1 represents the code we are running
		fmt.Println("     | Type    -> " + arguments[1]) // 2 represents the type of the error or the message
		fmt.Println("     | Message -> " + arguments[2]) // 3 represents the error message we want to display
		fmt.Println("|----|")
		os.Exit(0)
	}
}

func PrepareErrorAndWait(message string) {
	fmt.Println("E | " + message)
}
