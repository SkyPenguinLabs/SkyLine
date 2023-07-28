package SkyLine_ELF_Parsing

import (
	"bytes"
	"debug/elf"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"strings"
)

func ELF_CheckContextError() {
	if ELF_CTX.HandleError != nil {
		log.Fatal(ELF_CTX.HandleError)
	}
}

func ELF_ReadMagicUnit() {
	ELF_CTX.File.Read(ELF_CTX.MagicUnit[:16])
	expected := []byte{'\x7f', 'E', 'L', 'F'}
	if !bytes.Equal(ELF_CTX.MagicUnit[:4], expected) {
		log.Fatal("Error: File is not a valid ELF file")
	}
}

func ELF_InitateHeaderMap() {
	if v, ok := ELF_MapDataActions_Funcs[elf.Data(ELF_CTX.MagicUnit[elf.EI_DATA])]; ok {
		v()
	} else {
		log.Fatal("Corruption in ELF header, endianness unknown")
	}
	ELF_CTX.File.Seek(0, io.SeekStart)
	ELF_CTX.HandleError = binary.Read(ELF_CTX.File, ELF_CTX.FileHeader.Endian, ELF_CTX.Header)
	ELF_CheckContextError()
	switch HT := ELF_CTX.Header.(type) {
	case *elf.Header32:
		ELF_CTX.FileHeader.MachineName = elf.Machine(HT.Machine)
	case *elf.Header64:
		ELF_CTX.FileHeader.MachineName = elf.Machine(HT.Machine)
	}
}

func ELF_IniateArchitecture() {
	if v, ok := ELF_MapArchitecture_Funcs[elf.Class(ELF_CTX.MagicUnit[elf.EI_CLASS])]; ok {
		v()
	} else {
		log.Fatal("Invalid ELF class when parsing or initating architecture")
	}
}

func LoadElfHeader(Header interface{}) {
	ReturnStorage.HeaderParsed = make(map[string]string)
	if hdr, ok := Header.(*elf.Header64); ok {
		ReturnStorage.HeaderParsed["Magic"] = fmt.Sprintf("% x", hdr.Ident)
		if elf.Class(hdr.Ident[elf.EI_CLASS]) == elf.ELFCLASS64 {
			ReturnStorage.HeaderParsed["Class"] = "ELF64"
		} else if elf.Class(hdr.Ident[elf.EI_CLASS]) == elf.ELFCLASS32 {
			ReturnStorage.HeaderParsed["Class"] = "ELF32"
		}
		if v, x := ELF_DataActions_Funcs[elf.Data(hdr.Ident[elf.EI_DATA])]; x {
			ReturnStorage.HeaderParsed["Data"] = strings.TrimSpace(v())
		} else {
			log.Fatal("Error: Could not get Data action, should be either ELFDATA2LSB or ELFDATA2MSB  |Only two supported")
		}
		ReturnStorage.HeaderParsed["Version"] = strings.TrimSpace(fmt.Sprint(elf.Version(hdr.Version)))
		ReturnStorage.HeaderParsed["ABI"] = strings.TrimSpace(fmt.Sprint(elf.OSABI(hdr.Ident[elf.EI_OSABI])))
		ReturnStorage.HeaderParsed["ABIV"] = fmt.Sprintf("%d", hdr.Ident[elf.EI_ABIVERSION])
		ReturnStorage.HeaderParsed["Type"] = strings.TrimSpace(fmt.Sprint(elf.Type(hdr.Type)))
		ReturnStorage.HeaderParsed["Machine"] = strings.TrimSpace(fmt.Sprint(elf.Machine(hdr.Machine)))
		ReturnStorage.HeaderParsed["Entry"] = fmt.Sprintf("0x%x", hdr.Entry)
		ReturnStorage.HeaderParsed["ProgramHdrOffset"] = fmt.Sprintf("0x%x", hdr.Phoff)
		ReturnStorage.HeaderParsed["SectionHdrOffset"] = fmt.Sprintf("0x%x", hdr.Shoff)
		ReturnStorage.HeaderParsed["Flags"] = fmt.Sprintf("0x%x", hdr.Flags)
		ReturnStorage.HeaderParsed["ElfHdrSz"] = fmt.Sprintf("%d", hdr.Ehsize)
		ReturnStorage.HeaderParsed["ProgHdrEntrySz"] = fmt.Sprintf("%d", hdr.Phentsize)
		ReturnStorage.HeaderParsed["NumProgHdrEnt"] = fmt.Sprintf("%d", hdr.Phnum)
		ReturnStorage.HeaderParsed["SzSectionHdrEntry"] = fmt.Sprintf("%d", hdr.Shentsize)
		ReturnStorage.HeaderParsed["NumHrdEntry"] = fmt.Sprintf("%d", hdr.Shnum)
		ReturnStorage.HeaderParsed["IdxSecHdrStrTable"] = fmt.Sprintf("%d", hdr.Shstrndx)
	}
	if hdr, ok := Header.(*elf.Header32); ok {
		ReturnStorage.HeaderParsed["Magic"] = fmt.Sprintf("%x", hdr.Ident)
		ReturnStorage.HeaderParsed["Class"] = strings.TrimSpace(fmt.Sprint(elf.Class(hdr.Ident[elf.EI_CLASS])))
		ReturnStorage.HeaderParsed["Data"] = strings.TrimSpace(fmt.Sprint(elf.Data(hdr.Ident[elf.EI_DATA])))
		ReturnStorage.HeaderParsed["Version"] = strings.TrimSpace(fmt.Sprint(elf.Version(hdr.Version)))
		ReturnStorage.HeaderParsed["ABI"] = strings.TrimSpace(fmt.Sprint(elf.OSABI(hdr.Ident[elf.EI_OSABI])))
		ReturnStorage.HeaderParsed["ABIV"] = fmt.Sprintf("%d", hdr.Ident[elf.EI_ABIVERSION])
		ReturnStorage.HeaderParsed["Type"] = strings.TrimSpace(fmt.Sprint(elf.Type(hdr.Type)))
		ReturnStorage.HeaderParsed["Machine"] = strings.TrimSpace(fmt.Sprint(elf.Machine(hdr.Machine)))
		ReturnStorage.HeaderParsed["Entry"] = fmt.Sprintf("0x%x", hdr.Entry)
		ReturnStorage.HeaderParsed["ProgramHdrOffset"] = fmt.Sprintf("0x%x", hdr.Phoff)
		ReturnStorage.HeaderParsed["SectionHdrOffset"] = fmt.Sprintf("0x%x", hdr.Shoff)
		ReturnStorage.HeaderParsed["Flags"] = fmt.Sprintf("0x%x", hdr.Flags)
		ReturnStorage.HeaderParsed["ElfHdrSz"] = fmt.Sprintf("%d", hdr.Ehsize)
		ReturnStorage.HeaderParsed["ProgHdrEntrySz"] = fmt.Sprintf("%d", hdr.Phentsize)
		ReturnStorage.HeaderParsed["NumProgHdrEnt"] = fmt.Sprintf("%d", hdr.Phnum)
		ReturnStorage.HeaderParsed["SzSectionHdrEntry"] = fmt.Sprintf("%d", hdr.Shentsize)
		ReturnStorage.HeaderParsed["NumHrdEntry"] = fmt.Sprintf("%d", hdr.Shnum)
		ReturnStorage.HeaderParsed["IdxSecHdrStrTable"] = fmt.Sprintf("%d", hdr.Shstrndx)
	}
}
