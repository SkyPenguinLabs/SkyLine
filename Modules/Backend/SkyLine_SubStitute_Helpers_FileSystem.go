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
// Filename      |  SkyLine_SubStitute_Helpers_FileSystem.go
// Project       |  SkyLine programming language
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | In a smaller sense this file defines the total file structure of the SkyLine programming language.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// DEFINES ( TYPE ) : This file defines the entire file system, configuration, requirements and much more to create or use a source code file. Currently, source code files
//
// work with different forms of text but are basically regular text files. In the case of skyline files we want to be able to do hundreds of things to verify those files, verify
//
// the integrity and much more. This file will define any help related functions to help along side of that or add extra features to the language such as allowing a much more verbose
//
// output log of everything that will be parsed or is assumed will be parsed. Checking directories will also allow for existence and will also be verified by the engine. In later versions
//
// the user will be able to specify what they want to allow for source code files to verify a file extension. Of course users will not be able to specify non SkyLine source code files which
//
// means they can only limit file extensions within the language. By defauly skyline will only accept files that end with '.csc' standing for CyberSecurityCore but we also will allow the following.
//
// [
//		-> '.sl' 		| (skyline)
//      -> '.skyline'   | (skyline)
//		-> '.SL' 		| (skyline)
//      -> '.slmod'     | (SkyLine Configuration Engine Source Code Files) / SLCESCF
//      -> '.SLE'       | (SkyLine Engine Code) / SLE
//      -> '.so'        | (SkyLine Plugin) / regular shared object files
// 	]
//
// This file will also store each constant, variable and type structure to interact with the file system. Sure we do have a seperate file to define code models
//
// however that is typically reserved for only the language's structures like tokens, types, functions, literals, AST, Object calls and much more. So this file will
//
// hold data structures for configuration files and much more that can be used to interact or help the file system store resulting information. YZou may be asking why
//
// do we call it "fs" well we call it this because while it is not a file system it is rather a file verification system to verify the general file structure of SkyLine.
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
package SkyLine_Backend

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var SearchPaths []string

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FindModule(name string) string {
	basename := fmt.Sprintf("%s.skyline", name)
	for _, p := range SearchPaths {
		filename := filepath.Join(p, basename)
		if Exists(filename) {
			fmt.Println("[=] Found module")
			return filename
		}
	}
	return ""
}

func RunAndParseNoFile(data string, outoncall bool) error {
	parser := New_Parser(LexNew(data))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}
	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return nil
	}
	var x error
	if outoncall {
		_, x = io.WriteString(os.Stdout, result.SL_InspectObject()+"\n")
		println(result.SL_InspectObject())
	}
	return x
}

func RunAndParseFiles(filename string, enable_out_on_eval bool) error {
	FileCurrent.New(filename)
	f, x := os.Open(filename)
	if x != nil {
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x)
			}
		}()
		fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_DOES_NOT_EXIST](filename))
		os.Exit(0)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	if line == nil {
		fmt.Println(Map_Parser[ERROR_FILE_INTEGRITY_IS_EMPTY](filename))
	}
	data, x := ioutil.ReadFile(filename)
	if x != nil {
		fmt.Println(Map_Parser[ERROR_FILE_INPUT_OUTPUT_BUFFER_FAILED](filename, fmt.Sprint(x)))
		os.Exit(1)
	}
	parser := New_Parser(LexNew(string(data)))
	program := parser.ParseProgram()
	if len(parser.ParserErrors()) > 0 {
		return errors.New(parser.ParserErrors()[0])
	}

	Env := NewEnvironment()
	result := Eval(program, Env)
	if _, ok := result.(*Nil); ok {
		return nil
	}
	if enable_out_on_eval {
		_, x = io.WriteString(os.Stdout, result.SL_InspectObject()+"\n")
	}
	return x
}
