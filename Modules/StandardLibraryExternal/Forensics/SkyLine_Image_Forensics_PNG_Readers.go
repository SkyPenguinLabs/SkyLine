package SkyLine_Standard_External_Forensics

import (
	"bufio"
	"bytes"
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
// - File name     : SkyLine_Image_Forensics_PNG_Readers.go
//
// - File contains : All reader for the PNG image forensics library

func Process_Given_Image(file *os.File) (reader *bytes.Reader, x error) {
	st, x := file.Stat()
	PrepareErrorAndLog(
		x,
		file.Name(),
		fmt.Sprint(ERROR_CODE_FILE_COULD_NOT_STAT),
		"File stat FAIL", "Could not stat the file because of a given error"+fmt.Sprint(x),
		"Make sure that the supplied input file exists...",
	)
	var sizeof = st.Size()
	byter := make([]byte, sizeof)
	buffer := bufio.NewReader(file)
	if _, x = buffer.Read(byter); x != nil {
		return nil, x
	}
	reader = bytes.NewReader(byter)
	return reader, nil
}
