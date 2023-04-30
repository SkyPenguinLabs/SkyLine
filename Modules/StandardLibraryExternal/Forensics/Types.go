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
// - File name     : Types.go
//
// - File contains : All models, settings, codes and errors that can exist within or are used within this library.

const (
	PNG_SIG_1 = "PNG"
	PNG_SIG_2 = "89504E470D0A1A0A"
	PNG_PAT_3 = "89 50 4E 47 0D 0A 1A 0A"
	SIG_END   = "IEND"
)

var (
	// SIGNATURE FOR PNG IMAGE IN HEX
	PNG_PAT_V = []byte{
		0x89, 0x50, 0x4E,
		0x47, 0x0D, 0xA,
		0x1A, 0x0A,
	}

	BMP_PAT_V = []byte{
		0x42, 0x4d, 0x1e, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x1a,
		0x00, 0x00, 0x00, 0x0c, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x01,
		0x00, 0x01, 0x00, 0x18, 0x00,
		0x00, 0x00, 0xff, 0x00,
	}
)

type General struct {
	Input  string // Input image
	Output string // Output image
}

type PNG_Meta struct {
	Offset int64
	Chunk  PNG_Chunk
}

type PNG_Header struct {
	HEAD uint64
}

type PNG_Chunk struct {
	CRC  uint32 // CRC v
	FD   []byte // File Data v
	Type uint32 // Type v
	Size uint32 // Sizeof v
}

type Signatures struct {
	Sign   string
	Sufix  string
	Format string
}

var Sig = []Signatures{
	{`006E1EF0`, `*.ppt`, `PPT`},
	{`A0461DF0`, `*.ppt`, `PPT`},
	{`ECA5C100`, `*.doc`, `Doc file`},
	{`474946`, `*.gif`, `GIF`},
	{`GIF89a`, `*.gif`, `GIF`},
	{`FFD8FF`, `*.jpg`, `JPEG files`},
	{`JFIF`, `*.jpg`, `JPEG files`},
	{`504B03`, `*.zip`, `ZIP files`},
	{`LfLe`, `*.evt`, `Event file`},
	{`38425053`, `*.psd`, `Photoshop file`},
	{`8BPS`, `*.psd`, `Photoshop file`},
	{`4D5A`, `*.ocx`, `Active X`},
	{`415649204C495354`, `*.avi`, `AVI file`},
	{`AVI LIST`, `*.avi`, `AVI file`},
	{`57415645666D7420`, `*.wav`, `WAV file`},
	{`WAVEfmt`, `*.wav`, `WAV file`},
	{`25504446`, `*.pdf`, `PDF files`},
	{`%PDF`, `*.pdf`, `PDF files`},
	{`000100005374616E64617264204A6574204442`, `*.mdb`, `Microsoft database`},
	{`Standard Jet DB`, `*.mdb`, `Microsoft database`},
	{`2142444E`, `*.pst`, `PST file`},
	{`!BDN`, `*.pst`, `PST file`},
	{`4D6963726F736F66742056697375616C2053747564696F20536F6C7574696F6E2046696C65`, `*.sln`, `Microsft SLN file`},
	{`Microsoft Visual Studio Solution File`, `*.sln`, `Microsft SLN file`},
	{`504B030414000600`, `*.docx`, `Microsoft DOCX file`},
	{`504B030414000600`, `*.pptx`, `Microsoft PPTX file`},
	{`504B030414000600`, `*.xlsx`, `Microsoft XLSX file`},
	{`504B0304140008000800`, `*.xlsx`, `Java JAR file`},
	{`0908100000060500`, `*.xls`, `XLS file`},
	{`D0CF11E0A1B11AE1`, `*.msi`, `MSI file`},
	{`D0CF11E0A1B11AE1`, `*.doc`, `DOC`},
	{`D0CF11E0A1B11AE1`, `*.xls`, `Excel`},
	{`D0CF11E0A1B11AE1`, `*.vsd`, `Visio`},
	{`D0CF11E0A1B11AE1`, `*.ppt`, `PPT`},
	{`0A2525454F460A`, `*.pdf`, `PDF file`},
	{`.%%EOF.`, `*.pdf`, `PDF file`},
	{`4040402000004040`, `*.hlp`, `HLP file`},
	{`465753`, `*.swf`, `SWF file`},
	{`FWS`, `*.swf`, `SWF file`},
	{`CWS`, `*.swf`, `SWF file`},
	{`494433`, `*.mp3`, `MP3 file`},
	{`ID3`, `*.mp3`, `MP3 file`},
	{`MSCF`, `*.cab`, `Cab file`},
	{`0x4D534346`, `*.cab`, `Cab file`},
	{`ITSF`, `*.chm`, `Compressed Help`},
	{`49545346`, `*.chm`, `Compressed Help`},
	{`4C00000001140200`, `*.lnk`, `Link file`},
	{`4C01`, `*.obj`, `OBJ file`},
	{`4D4D002A`, `*.tif`, `TIF graphics`},
	{`MM`, `*.tif`, `TIF graphics`},
	{`000000186674797033677035`, `*.mp4`, `MP4 Video`},
	{`ftyp3gp5`, `*.mp4`, `MP4 Video`},
	{`0x00000100`, `*.ico`, `Icon file`},
	{`300000004C664C65`, `*.evt`, `Event file`},
	{`Rar!`, `*.rar`, `RAR file`},
	{`526172211A0700`, `*.rar`, `RAR file`},
	{`52657475726E2D506174683A20`, `*.eml`, `EML file`},
	{`Return-Path:`, `*.eml`, `EML file`},
	{`6D6F6F76`, `*.mov`, `MOV file`},
	{`moov`, `*.mov`, `MOV file`},
	{`7B5C72746631`, `*.rtf`, `RTF file`},
	{`{\rtf1`, `*.rtf`, `RTF file`},
	{`89504E470D0A1A0A`, `*.png`, `PNG file`},
	{`PNG`, `*.png`, `PNG file`},
	{`C5D0D3C6`, `*.eps`, `EPS file`},
	{`CAFEBABE`, `*.class`, `Java class file`},
	{`D7CDC69A`, `*.WMF`, `WMF file`},
}
