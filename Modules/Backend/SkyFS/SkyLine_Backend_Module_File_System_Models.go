///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This is the file system of the language which helps to define file line tracing systems, error tracing, system and file location, getting current locations from so
//
// and so provided file, etc. In our case, we need to be able to load, unload, trace etc on files during the process of errors.
//
//
// This file defines all of the types and models for this module
package SkyLine_File_System

type (
	//:::::::::::::::::::::::::::::::::::::::::::::::
	//:: This defines the data in the Current File ::
	//:::::::::::::::::::::::::::::::::::::::::::::::

	SL_CurrentFile struct {
		SL_FileName     string   // This is the files direct name
		SL_SourceType   string   // This is the files source type (SL, SLC, skyline, sl, etc...)
		SL_LineCount    int      // This is the number of lines the file has
		SL_FileRealPath string   // This is the files real
		SL_FileStorage  []string // This is the file storage or the current line direct to access for error systems
	}

	//:::::::::::::::::::::::::::::::
	//:: Format for the box output
	//:::::::::::::::::::::::::::::::
	BoxFormat struct {
		TL string
		TR string
		BL string
		BR string
		HZ string
		VT string
	}
)

var Current SL_CurrentFile
