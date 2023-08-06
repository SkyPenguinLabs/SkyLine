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
// This file defines all of the helping or assistant functions
package SkyLine_File_System

import (
	"io/fs"
	"log"
	"os"
)

func CheckStat(filename string) fs.FileInfo {
	f, x := os.Stat(filename)
	if x != nil {
		log.Fatal(x)
	}
	return f
}
