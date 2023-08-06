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
//
// Contains -> This file contains functions for project management. The SLC or SkyLine configuration language is a language and engine that was designed from the RPC project management
// system. Radical Processing Core's Engine (RPC-E) was an engine that would help manage and generate projects easier within RPC. This is because of how syntactically annoying RPC was.
// So for C++/C/Go wrapping and library calling we decided to also add the engine. This engine will help generate code templates using M4-Sugar ( For linux ) based on templates and help
// better set you up for project development! Current support for this system is LINUX
//
package SkyLine_Configuration_Engine_Backend_Source

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	M4_Sugar_LOC = "/usr/share/Modules/Backend/SkyConfEngine/EngineCore/SkyLine-Configuration-Engine-Templates/Linux"
)

var TemplatePaths = []string{
	"Backend",
	"Backend/Conf/EngineFiles",
	"Backend/Conf/ProjectData",
	"Backend/Conf/Requires",
	"Backend/Core/Scripts",
	"Backend/Core/Plugins",
}

var AllTemplates = []string{
	"slmod:Engine",
	"config:Project",
	"config:Requires",
	"core:Makefile",
	"core:Modules",
	"core:main",
	"core:FFI_IMPL_CPP",
	"core:FFI_IMPL_HPP",
	"core:FFI_FinalPlugin",
	"core:GoMod",
}

/*
With this code structure shown below, we are basically generating everything like so

- Backend
	| Conf
		| --> EngineFiles, ProjectData, Requires
	| Core
	    | -->  Installer, Scripts, Wrapper (CPP, Go)


*/

var FileMap = map[string]string{
	"core:main":            "main.sl",                                    // Main script file
	"slmod:Engine":         "Backend/Conf/EngineFiles/Engine.slmod",      // Engine modification scripts file for generation
	"config:Project":       "Backend/Conf/ProjectData/Project.json",      // Project data for internal project hooking
	"config:Requires":      "Backend/Conf/Requires/Requires.json",        // Requirements for project to be loaded by SLC
	"core:Modules":         "Backend/Core/Scripts/Module.sl",             // Example module for development
	"core:FFI_IMPL_CPP":    "Backend/Core/Plugins/CPPImplementation.cpp", // CPP Wrapper interfaces and files
	"core:FFI_IMPL_HPP":    "Backend/Core/Plugins/CPPImplementation.hpp", // CPP Wrapper interfaces and files
	"core:FFI_FinalPlugin": "Backend/Core/Plugins/PluginCompile.go",      // Shared file that includes C files and CPP files for C++->C code calling using CGO to generate plugin
	"core:Makefile":        "Backend/Core/Plugins/Makefile",              // Makefile / Install script
	"core:GoMod":           "Backend/Core/Plugins/go.mod",
}

func CreateDir(dir string) {
	x := os.MkdirAll(dir, os.ModePerm)
	if x != nil {
		log.Fatal(x)
	}
}

func Execute(command, arg1, write string) (string, error) {
	cmd := exec.Command(command, arg1)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func GenerateContentsOfFile(filename, startdir string) {
	var file_to_write string
	var fp string
	if resfile, ok := FileMap[filename]; ok {
		file_to_write = resfile
	} else {
		fmt.Println("Sorry filetype didnt exist, dev error (SLC|File-SkyLine_Configuration_Engine_Backend_ProjectManagement)|LINE=45")
		os.Exit(0)
	}
	if file_to_write != "" {
		file_to_write = startdir + file_to_write
	} else {
		log.Fatal("Issue when working with the file to write to, provided value was empty")
	}
	switch filename {
	case "core:GoMod":
		c := `
module main

go 1.17
				`
		file, x := os.Create(file_to_write)
		if x != nil {
			log.Fatal(x)
		}
		defer file.Close()
		file.Write([]byte(c))
	case "core:main":
		fp = M4_Sugar_LOC + "/ProjectMainScriptSL.m4"
	case "slmod:Engine":
		fp = M4_Sugar_LOC + "/ProjectEngineFileTemplate.m4"
	case "core:FFI_IMPL_CPP":
		fp = M4_Sugar_LOC + "/ProjectWrapperAutoConfCPP.m4"
	case "core:FFI_IMPL_HPP":
		fp = M4_Sugar_LOC + "/ProjectWrapperAutoConfHeaderPlusPlusFile.m4"
	case "core:FFI_FinalPlugin":
		fp = M4_Sugar_LOC + "/ProjectWrapperAutoConfGo.m4"
	case "config:Project":
		Project := Load_ProjectData{
			ProjectInformation: struct {
				Name        string `json:"Name"`
				Description string `json:"Description"`
				SupportedOS string `json:"Supported-OS"`
				Languages   string `json:"Languages"`
				Version     string `json:"Version"`
			}{
				Name:        "null",
				Description: "null",
				SupportedOS: "null",
				Languages:   "null",
				Version:     "null",
			},
		}
		Loaded, x := json.MarshalIndent(Project, "", "    ")
		if x != nil {
			log.Fatal(x)
			return
		}
		file, x := os.Create(file_to_write)
		if x != nil {
			log.Fatal(x)
		}
		defer file.Close()
		file.Write(Loaded)
	case "config:Requires":
		loadReq := Load_Requirements{
			Requirements: struct {
				Libraries       []string `json:"Libraries"`
				OperatingSystem string   `json:"Operating-System"`
				SLCVersion      string   `json:"SLC-Version"`
				SLVersion       string   `json:"SL-Version"`
			}{
				Libraries:       []string{"null1", "null2"},
				OperatingSystem: "Linux",
				SLCVersion:      "null",
				SLVersion:       "null",
			},
		}

		Loaded, x := json.MarshalIndent(loadReq, "", "    ")
		if x != nil {
			fmt.Println("Error:", x)
			return
		}
		file, x := os.Create(file_to_write)
		if x != nil {
			log.Fatal(x)
		}
		defer file.Close()
		file.Write(Loaded)
	case "core:Makefile":
		c := `
all:
	g++ -c -fPIC CPPImplementation.cpp -o CPPImplementation.o
	g++ -shared CPPImplementation.o -o CPPImplementation.so
	go build -buildmode=plugin -o SkyLinePlugin

clean:
	rm SkyLinePlugin CPPImplementation.o CPPImplementation.so
		`
		file, x := os.Create(file_to_write)
		if x != nil {
			log.Fatal(x)
		}
		defer file.Close()
		file.Write([]byte(c))
	case "core:Modules":
		c := `
set ExportedValue = [1,2,3,4];
// Exported
set x = 10;
// Not exported by import
		`
		file, x := os.Create(file_to_write)
		if x != nil {
			log.Fatal(x)
		}
		defer file.Close()
		file.Write([]byte(c))
	}
	if fp != "" {
		data, x := Execute("m4", fp, "")
		if x != nil {
			log.Fatal("Command Execution Error (SLC): ", x)
		}
		f, x := os.Create(file_to_write)
		if x != nil {
			log.Fatal(x)
		}
		defer f.Close()
		trimmed := strings.TrimSpace(string([]byte(data)))
		f.Write([]byte(trimmed))
	}
}

// Start directory is where SLC will generate or locate
// the primary start for generating the files and project data
// for example: say you provide "user/Desktop/Project" SLC will
// generate all paths and directories and files of the template
// project in that directory.
func SetupCall(StartDirectory string) {
	for _, k := range TemplatePaths {
		CreateDir(StartDirectory + k)
	}
}
