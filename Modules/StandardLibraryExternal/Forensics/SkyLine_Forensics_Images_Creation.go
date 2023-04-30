package SkyLine_Standard_External_Forensics

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
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
// - File name     : SkyLine_Forensics_Images_Creation.go
//
// - File contains : Image generation programs for BMP, JPG, PNG, WEBP, GIF
type Create_Image_Settings_Pre struct {
	PixelHeight string // Height of the pixel
	PixelWidth  string // Width of the pixel
	Output      string // Output image
}

var CreationSettings Create_Image_Settings_Pre

func SkyLine_Forensics_Image_Creation_Utility(typeofimage string) bool {
	pw, x := strconv.Atoi(CreationSettings.PixelWidth)
	if x != nil {
		log.Fatal(x)
	}
	ph, x := strconv.Atoi(CreationSettings.PixelHeight)
	if x != nil {
		log.Fatal(x)
	}
	img := image.NewRGBA(
		image.Rect(0, 0, pw, ph),
	)
	col := color.RGBA{0, 0, 0, 255}
	img.Set(0, 0, col)
	f, x := os.Create(CreationSettings.Output)
	if x != nil {
		return false
	}
	defer f.Close()
	switch strings.ToLower(typeofimage) {
	case "gif":
		gif.Encode(f, img, nil)
	case "jpg":
		jpeg.Encode(f, img, nil)
	case "png":
		png.Encode(f, img)
	case "bmp":
		f.Write(BMP_PAT_V)
	default:
		gif.Encode(f, img, nil)
	}
	return true
}
