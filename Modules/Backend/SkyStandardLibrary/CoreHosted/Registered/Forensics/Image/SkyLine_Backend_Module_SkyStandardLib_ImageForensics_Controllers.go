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
// This file defines all automated controllers which are used to write binary data, execute specific tasks etc. Typically this is used during steganography.
//
// Controllers also contain type specific functions, such as chunk creation, image creation, image setup, auto binary readers, auto binary writers and other various functions
//
// that may be specific to the image library and forensics library.
//
package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section defines all seek and sub functions for file automation
//
//
//

//::::::::::::::::::::::
//:: File seek        ::
//::::::::::::::::::::::
//
//
//
//SkyLine_Image_Forensics_Controller_M2_F_S:
//    mov edx, eax      ; Copy WHENCE to EDX
//    mov eax, 0x8      ; SYS_LSEEK = 8
//    int 0x80          ; Perform system call
//    ret				; Return information
//
//
func SkyLine_Forensics_Controllers_FileSeek(f *os.File, OFFSET int64, WHENCE int) {
	if _, x := f.Seek(OFFSET, WHENCE); x != nil {
		log.Fatal(x)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section of the file contains all reader and writer functions and controllers. This means that the data is reading or writing some form of data no matter the data type
//
//
//

//::::::::::::::::::::::
//:: File Read        ::
//::::::::::::::::::::::
func SkyLine_Forensics_Controllers_FileRead(f *os.File, data []byte) {
	if _, x := f.Read(data); x != nil {
		log.Fatal(x)
	}
}

//::::::::::::::::::::::
//:: Binary Reader    ::
//::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_BinaryR(Reader io.Reader, order binary.ByteOrder, information interface{}) {
	if x := binary.Read(Reader, order, information); x != nil {
		if x == io.EOF {
			fmt.Println("Warning: IO reader failed to read data any further due to IO EOF")
			os.Exit(0)
		}
		log.Fatal(x)
	}
}

//::::::::::::::::::::::
//:: File Write       ::
//::::::::::::::::::::::
func SkyLine_Forensics_Controllers_FileWrite(f *os.File, data []byte) {
	if _, x := f.Write(data); x != nil {
		log.Fatal(x)
	}
}

//::::::::::::::::::::::
//:: Binary Writer    ::
//::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_BinaryW(Writer io.Writer, order binary.ByteOrder, information interface{}) {
	if x := binary.Write(Writer, order, information); x != nil {
		log.Fatal(x)
	}
}

//::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Read Chunks Data ::
//::::::::::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensic_Controllers_ChunkBytes_R(reader *bytes.Reader, ChunkLength uint32) {
	PngMeta.Chunk.FieldData = make([]byte, ChunkLength)
	if x := binary.Read(reader, binary.BigEndian, &PngMeta.Chunk.FieldData); x != nil {
		log.Fatal(x)
	}
}

//:::::::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Read Specified Offset ::
//:::::::::::::::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Fornesics_Controllers_Offset_R(reader *bytes.Reader) {
	PngMeta.Offset, _ = reader.Seek(0, 1)
}

//::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Read Chunk Type  ::
//::::::::::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_ChunkType_R(reader *bytes.Reader) {
	PngMeta.SkyLine_Forensics_Controllers_BinaryR(reader, binary.BigEndian, &PngMeta.Chunk.Type)
}

//::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Read Chunk Size  ::
//::::::::::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_ChunkSize_R(reader *bytes.Reader) {
	PngMeta.SkyLine_Forensics_Controllers_BinaryR(reader, binary.BigEndian, &PngMeta.Chunk.Size)
}

//::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Read Chunk CRC   ::
//::::::::::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Fornesics_Controllers_ChunkCRC_R(reader *bytes.Reader) {
	PngMeta.SkyLine_Forensics_Controllers_BinaryR(reader, binary.BigEndian, &PngMeta.Chunk.CRC)
}

//:::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Read All MetaData ::
//:::::::::::::::::::::::::::::::::::::::::::s
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_Chunk_R(reader *bytes.Reader) {
	PngMeta.SkyLine_Forensics_Controllers_ChunkSize_R(reader)                     // Read chunks size
	PngMeta.SkyLine_Forensics_Controllers_ChunkType_R(reader)                     // Read chunks type
	PngMeta.SkyLine_Forensic_Controllers_ChunkBytes_R(reader, PngMeta.Chunk.Size) // Read chunks bytes
	PngMeta.SkyLine_Fornesics_Controllers_ChunkCRC_R(reader)                      // Read chunk CRC
}

//::::::::::::::::::::::::::::::::::::::::::
//:: PNG Format Reader | Process Image    ::
//::::::::::::::::::::::::::::::::::::::::::

func SkyLine_Forensics_Controllers_ProcessImage(file *os.File) (reader *bytes.Reader, x error) {
	stat, x := file.Stat()
	if x != nil {
		log.Fatal(x)
	}
	var sizeof = stat.Size()
	byter := make([]byte, sizeof)
	buffer := bufio.NewReader(file)
	if _, x = buffer.Read(byter); x != nil {
		return nil, x
	}
	reader = bytes.NewReader(byter)
	return reader, nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This file defines all junk creation controllers which will generate chunks of specific file formats like PNG images
//
//
//

//:::::::::::::::::::::::::::::::::::
//:: Controller Create Junk CRC32  ::
//:::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_CreateChunk_CRC32() uint32 {
	Out := new(bytes.Buffer)
	PngMeta.SkyLine_Forensics_Controllers_BinaryW(Out, binary.BigEndian, PngMeta.Chunk.Type)
	PngMeta.SkyLine_Forensics_Controllers_BinaryW(Out, binary.BigEndian, PngMeta.Chunk.FieldData)
	return crc32.ChecksumIEEE(Out.Bytes())
}

//:::::::::::::::::::::::::::::::::::
//:: Controller Create Junk Size   ::
//:::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_CreateChunk_Size() uint32 {
	return uint32(len(PngMeta.Chunk.FieldData))
}

//:::::::::::::::::::::::::::::::::::
//:: Controller Create Metadata    ::
//:::::::::::::::::::::::::::::::::::
func (PngMeta *PNG_MetaData) SkyLine_Forensics_Controllers_CreateMarshaler() *bytes.Buffer {
	byter := new(bytes.Buffer)
	PngMeta.SkyLine_Forensics_Controllers_BinaryW(byter, binary.BigEndian, PngMeta.Chunk.Size)
	PngMeta.SkyLine_Forensics_Controllers_BinaryW(byter, binary.BigEndian, PngMeta.Chunk.Type)
	PngMeta.SkyLine_Forensics_Controllers_BinaryW(byter, binary.BigEndian, PngMeta.Chunk.FieldData)
	PngMeta.SkyLine_Forensics_Controllers_BinaryW(byter, binary.BigEndian, PngMeta.Chunk.CRC)
	return byter
}
