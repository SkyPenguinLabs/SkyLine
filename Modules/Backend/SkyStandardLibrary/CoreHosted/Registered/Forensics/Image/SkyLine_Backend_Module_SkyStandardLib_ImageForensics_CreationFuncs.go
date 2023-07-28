//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _              _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___ |   __|___ ___ ___ ___ ___|_|___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___ |   __| . |  _| -_|   |_ -| |  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|    |__|  |___|_| |___|_|_|___|_|___|___|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
// This library defines all forensics related functions and libraries which allow you to dissect, inject or mess around with images or files in some shape or form while also inspecting
//
// specific file formats such as PNG, JPEG, GIF, BMP, ELF, PE, EXE and etc files.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file dsefines all of the creation functions for images, this current system allows you to create basic images
//
// such as GIF, JPG, PNG, and BMP
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func SkyLine_Forensics_Module_Image_GenerateImage(typeof string) bool {
	PixelWidth, PixelHeight := SkyLine_Forensics_Converters_AtoI(CreateImageSettings.PixelWidth), SkyLine_Forensics_Converters_AtoI(CreateImageSettings.PixelHeight)

	//:::::::::::::
	//:: New image:
	//:::::::::::::
	Image := image.NewRGBA(
		image.Rect(0, 0, PixelWidth, PixelHeight),
	)

	Color := color.RGBA{0, 0, 0, 255}
	Image.Set(0, 0, Color)

	//::::::::::::::::
	//:: Create File :
	//::::::::::::::::

	f, x := os.Create(CreateImageSettings.Output)
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()

	//:::::::::::::::::::::
	//:: Encode Base Type :
	//:::::::::::::::::::::
	switch strings.ToLower(typeof) {
	case "gif":
		gif.Encode(f, Image, nil)
	case "png":
		png.Encode(f, Image)
	case "bmp":
		f.Write(BMP_Base_Payload)
	case "jpg":
		jpeg.Encode(f, Image, nil)
	default:
		gif.Encode(f, Image, nil)
	}
	return true
}
