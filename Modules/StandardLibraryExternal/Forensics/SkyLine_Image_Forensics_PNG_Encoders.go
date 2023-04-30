package SkyLine_Standard_External_Forensics

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
// - File name     : SkyLine_Image_Forensics_PNG_Encoders.go
//
// - File contains : Payload encoders before injection was made or gets prepared

func Mult_Encode_OR_Decode(enc_key string, in []byte) []byte {
	byter := make([]byte, len(in))
	for k := 0; k < len(in); k++ {
		byter[k] += in[k] ^ enc_key[k%len(enc_key)]
	}
	return byter
}

func Decode_Array(encoded_array []byte, key string) []byte {
	return Mult_Encode_OR_Decode(key, encoded_array)
}

func Encode_Array(decoded_array []byte, key string) []byte {
	return Mult_Encode_OR_Decode(key, decoded_array)
}
