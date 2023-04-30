/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Parser_SubRoutine_FileSys_And_Trace_Management.go
// Project       |  SkyLine programming language
// Line Count    |  0 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the File methods which allow us to check the data and the file names within the language.
//
//
//
package SkyLine_Backend

import (
	"log"
	"path/filepath"
)

func (File *FileCurrentWithinParserEnvironment) New(filename string)   { File.Filename = filename }  // Method | Assigns new file
func (File *FileCurrentWithinParserEnvironment) Get_Name() string      { return File.Filename }      // Method | Returns file name
func (File *FileCurrentWithinParserEnvironment) Get_Extension() string { return File.FileExtension } // Method | Returns file extension
func (File *FileCurrentWithinParserEnvironment) Get_Basename() string  { return File.FileBasename }  // Method | Returns file basename
func (File *FileCurrentWithinParserEnvironment) GetAbsolute() string {
	abs, x := filepath.Abs(File.Filename)
	if x != nil {
		log.Fatal(x)
	}
	return abs
} // Method | Returns
