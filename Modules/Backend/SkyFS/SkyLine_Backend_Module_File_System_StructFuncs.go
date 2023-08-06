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

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (Current *SL_CurrentFile) FileSystem_Modify_Name(filename string) {
	if st := CheckStat(filename); st != nil {
		if !st.IsDir() {
			Current.SL_FileName = filename
		} else {
			log.Fatal("Error: File was a directory, currently directories are not supported")
		}
	}
}

func (Current *SL_CurrentFile) FileSystem_ModifyDir() {
	if st := CheckStat(Current.SL_FileName); st != nil {
		fp, x := filepath.Abs(Current.SL_FileName)
		if x != nil {
			log.Fatal(x)
		}
		Current.SL_FileRealPath = fp
	}
}

func (Current *SL_CurrentFile) FileSystem_ModifyInfo() {
	f, x := os.Open(Current.SL_FileName)
	if x != nil {
		log.Fatal(x)
	}
	scanner := bufio.NewScanner(f)
	Current.SL_FileStorage = []string{}
	Current.SL_LineCount = 0
	for scanner.Scan() {
		Current.SL_LineCount++
		Current.SL_FileStorage = append(Current.SL_FileStorage, scanner.Text())
	}
}

func (Current *SL_CurrentFile) FileSystem_ModifyExtension() {} // Do nothing rn

func (Current *SL_CurrentFile) FileSystem_IndexLine(line int) string {
	var str string
	line = line - 1
	if line >= 0 && line < len(Current.SL_FileStorage) {
		str = Current.SL_FileStorage[line]
	}
	return str
}

func DrawUtilsBox(variable, line2 string) string {
	var box string
	BL := BoxFormat{
		TL: "┏",
		TR: "┓",
		BL: "┗",
		BR: "┛",
		HZ: "━",
		VT: "┃",
	}
	l := strings.Split(variable, "\n")
	var mlen int
	for _, lin := range l {
		if len(lin) > mlen {
			mlen = len(lin)
		}
	}
	var LINUX_GREY = "\033[38;5;249m" // Color Scheme | Linux unicode set ( GREY   )
	box += LINUX_GREY + BL.TL + strings.Repeat(BL.HZ, mlen) + BL.TR + "\n"
	for _, line := range l {
		if line2 == line {
			box += BL.VT + "\033[31m" + line + LINUX_GREY + strings.Repeat(" ", mlen-len(line)) + BL.VT + "\n"
		} else {
			box += BL.VT + line + strings.Repeat(" ", mlen-len(line)) + BL.VT + "\n"
		}
	}
	box += BL.BL + strings.Repeat(BL.HZ, mlen) + BL.BR + "\n"
	return box
}

func (Current *SL_CurrentFile) DrawBoxWithinLineRange(lineCode string) string {
	var lines []string
	lineIndex := -1
	for i, line := range Current.SL_FileStorage {
		if line == lineCode {
			lineIndex = i
			break
		}
	}
	if lineIndex != -1 {
		start := lineIndex - 5
		if start < 0 {
			start = 0
		}
		end := lineIndex + 5
		if end > len(Current.SL_FileStorage) {
			end = len(Current.SL_FileStorage)
		}

		lines = Current.SL_FileStorage[start:end]
	}
	var str string
	for i := 0; i < len(lines); i++ {
		str += lines[i] + "\n"
	}
	return DrawUtilsBox(str, lineCode)
}
