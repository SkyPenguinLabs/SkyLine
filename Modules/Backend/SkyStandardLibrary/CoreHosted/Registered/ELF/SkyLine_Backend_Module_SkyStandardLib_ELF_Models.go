package SkyLine_ELF_Parsing

import (
	"debug/elf"
	"encoding/binary"
	"os"
)

var ELF_DataActions_Funcs = map[elf.Data]func() string{
	elf.ELFDATA2LSB: func() string { return "Little Endian (ELF_DATA_2_LSB)" },
	elf.ELFDATA2MSB: func() string { return "Big Endian (ELF_DATA_2_MSB)" },
}

var ELF_MapDataActions_Funcs = map[elf.Data]func(){
	elf.ELFDATA2LSB: func() {
		ELF_CTX.FileHeader.Endian = binary.LittleEndian
	},
	elf.ELFDATA2MSB: func() {
		ELF_CTX.FileHeader.Endian = binary.BigEndian
	},
}

var ELF_MapArchitecture_Funcs = map[elf.Class]func(){
	elf.ELFCLASS32: func() {
		ELF_CTX.Header = new(elf.Header32)
		ELF_CTX.FileHeader.Architecture = elf.ELFCLASS32
	},
	elf.ELFCLASS64: func() {
		ELF_CTX.Header = new(elf.Header64)
		ELF_CTX.FileHeader.Architecture = elf.ELFCLASS64
	},
}

type SkyLine_ELF_Context struct {
	File        *os.File    // Save file handle
	MagicUnit   [16]byte    // Read 16 bytes into the file
	Header      interface{} // File header
	HandleError error       // Handle or loading error | CTX error
	FileSize    int64       // Total file size
	// Inline Structures
	FileHeader struct {
		Endian       binary.ByteOrder
		Architecture elf.Class
		MachineName  elf.Machine
	} // File head storage
}

type SkyLine_ReturnStore_ELF struct {
	HeaderParsed map[string]string
}

var (
	ELF_CTX       SkyLine_ELF_Context
	ReturnStorage SkyLine_ReturnStore_ELF
)
