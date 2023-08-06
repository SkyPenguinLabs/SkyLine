////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
//
// The SkyLine configuration language is a language and engine designed to act as a modification language to the SkyLine programming language. This language is
// very minimal and contains a regex base lexer, a very basic parser, a few keywords, a base interpreter and that is all as well as some backend engine code. This
// language is purely modified to be an extension to the SkyLine programming language, something that can be a pre processor language post processing for the main
// SkyLine script. Below is more technical information for the language
//
// Lexer       : Regex based lexer with minimal tokens and base syntax
// Parser      : Base parser with minimal tokens and base syntax with simple error systems
// REPL        : Does not exist
// Environment : Extremely minimal
// Types       : String, Boolean, Integer
// Statements  : set, import, use, errors, output, system, constant/const
//
// File Contains -> This file contains a function to start and continue to read the file for engine configuration

package SkyLine_Configuration_Engine_Backend_Source

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var IsRenderForCmdDbg bool

func StartEngine_RenderFile(Filename string, iscmdrender bool) (errors bool) {
	IsRenderForCmdDbg = iscmdrender
	fmt.Print("\n\n")
	f, x := os.Open(Filename)
	if x != nil {
		Message := CallErrorStr(
			fmt.Sprint(SLC_FileSystem_ErrorWhenOpeningOrLoadingFile),
			"Could not open file due to -> "+fmt.Sprint(x),
			Filename,
		)
		println(Message)
		os.Exit(0)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if line == nil {
		Message := CallErrorStr(
			fmt.Sprint(SLC_FileSystem_NULL_FIELDS),
			"Engine refused to parse file due to there being nothing there",
			Filename+" -> NULL",
		)
		println(Message)
		os.Exit(0)
	}
	data, x := ioutil.ReadFile(Filename)
	if x != nil {
		Message := CallErrorStr(
			fmt.Sprint(SLC_FileSystem_ErrorWhenOpeningOrLoadingFile),
			"Could not open file due to -> "+fmt.Sprint(x),
			Filename,
		)
		println(Message)
		os.Exit(0)
	}
	parser := NewParser(New(string(data)))
	program := parser.ParseProgram()
	if len(parser.Errors()) > 0 {
		log.Fatal(parser.Errors()[0])
		errors = true
	} else {
		errors = false
	}

	Env := Start_Engine_Environment_Create()
	result := Eval(program, Env)
	if _, ok := result.(*ObjectNULL); ok {
		return
	}
	_, x = io.WriteString(os.Stdout, result.ObjectInspectFunc()+"\n")
	if x != nil {
		log.Fatal(x)
	}
	return errors
}
