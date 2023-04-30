package SkyLine_Standard_External_Forensics

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
// - File name     : SkyLine_Forensics_Images_Controller.go
//
// - File contains : File contains any other functions which can be run on sub images. For example, ZIP signature checking, Base64 image checking, all of it is a sub
//                   function that can be put inside of the main controller for the forensics library.

const (
	GEN_Signature = '\x50'
)

var (
	ZIP_Signature = []byte{'\x4b', '\x03', '\x04'}
)

// Controller function | Verifies the factor of a ZIP signature existing within the image data
func Image_Controller_Function_External_1_Verify_Archive(imagefile string) bool {
	f, x := os.Open(imagefile)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_CODE_FILE_NOT_FOUND),
			"File (Err when trying to open)",
			"Make sure the file that you gave as an input is a valid input file, check the name, path, location etc in order to verify that it can be opened by OS operations",
		)
		return false
	}
	defer f.Close()
	buffer := bufio.NewReader(f)
	status, _ := f.Stat()
	for lbuf := int64(0); lbuf < status.Size(); lbuf++ {
		byter, x := buffer.ReadByte()
		if x != nil {
			PrepareErrorAndLog(
				x,
				fmt.Sprint(ERROR_READ_BYTE_BUFFER_COULD_NOT_READ_BYTE),
				"Buffer ^IO^ (Err when trying to read the next byte)",
				"Could not read the next byte or read any byte within the buffer -> "+fmt.Sprint(x),
			)
			return false
		}
		if byter == GEN_Signature {
			B_Mark := make([]byte, 3)
			B_Mark, x = buffer.Peek(3)
			if x != nil {
				PrepareErrorAndLog(
					x,
					fmt.Sprint(ERROR_READ_BYTE_BUFFER_COULD_NOT_PEAK_AT),
					"Buffer ^IO^ (Err when trying to peek at 3)",
					"GOt error from buffer -> "+fmt.Sprint(x),
				)
			}
			if bytes.Equal(B_Mark, ZIP_Signature) {
				return true
			}
		}
	}
	return false
}

// Controller function | Injects a file into another file
func Image_Controller_Function_External_2_Inject_Into_File(input, output, file string) bool {
	// open the input file, this is the file we will be taking text or data from and moving it into the output
	in, x := os.Open(input)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_CODE_FILE_NOT_FOUND),
			"File (Err when trying to open the file)",
			"Make sure that the file you gave as an input file exists, SkyLine got -> "+fmt.Sprint(x),
		)
		return false
	}
	defer in.Close()
	f, x := os.Open(file)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_CODE_FILE_NOT_FOUND),
			"File (Err when trying to open the file)",
			"Make sure that the file you gave as an input file exists, SkyLine got -> "+fmt.Sprint(x),
		)
		return false
	}
	defer f.Close()
	// create the output file
	out, x := os.Create(output)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_OS_OPERATION_SET_FAILED_CREATE),
			"OS ^Operations^ (Err when trying to create the output file [ "+fmt.Sprint(output)+" ])",
			"SkyLine got -> "+fmt.Sprint(x),
		)
		return false
	}
	defer out.Close()
	_, x = io.Copy(out, in)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_IO_OPERATION_SET_FAILED_COPY),
			"IO ^Operation^ (Err when trying to copy data from input file to output file)",
			"SkyLine got -> "+fmt.Sprint(x),
		)
		return false
	}
	_, x = io.Copy(out, f)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_IO_OPERATION_SET_FAILED_COPY),
			"IO ^Operation^ (Err when trying to copy data from file to inject to output file)",
			"SkyLine got -> "+fmt.Sprint(x),
		)
		return false
	}
	return true
}

// Controller function | Gets the name of a file using a small leightweight database
func Image_Controller_Function_External_3_IdentifyUnknownFile(file string) (FileInf string) {
	f, x := ioutil.ReadFile(file)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_IO_OPERATION_SET_FAILED_READ_FILE),
			"IO ^Operation^ (Err when trying to read the file)",
			"IO Utilities has failed to read the file, SkyLine got -> "+fmt.Sprint(x),
		)
		return ""
	} else {
		for _, v := range Sig {
			if strings.HasSuffix(file, v.Sufix) || bytes.Contains(f, []byte(v.Sign)) {
				FileInf = v.Format
			}
		}
	}
	return FileInf
}
