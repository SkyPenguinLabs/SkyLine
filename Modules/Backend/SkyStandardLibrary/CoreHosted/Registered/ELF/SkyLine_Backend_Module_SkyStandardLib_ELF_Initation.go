package SkyLine_ELF_Parsing

import "os"

func ELF_Session_Initation(filename string) {
	ELF_CTX.File, ELF_CTX.HandleError = os.Open(filename)
	ELF_CheckContextError()
	defer ELF_CTX.File.Close() // Call to close once file is opened
	ELF_ReadMagicUnit()        // Read and verify magic header is E L F
	ELF_IniateArchitecture()   // Set architecture for further investigation (32 || 64)
	ELF_InitateHeaderMap()     // Set header map for further investigation to map endians and header types

}
