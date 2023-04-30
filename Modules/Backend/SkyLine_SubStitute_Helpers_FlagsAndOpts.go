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
// Filename      |  SkyLine_SubStitute_Helpers_FlagsAndOpts.go
// Project       |  SkyLine programming language
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the options, parsing functions, run functions and environment settings that will be defined by flags. All of these flags come from
//                 running the binary interpreter in `/usr/bin/skyline` where each flag is defined with - or --. These flags will dictate the behavior of the interpreter and
//                 the skyline configuration engine. We work with the engine as a third party package to provide a much more deeper check when loading SLC configuration files.
//
//
package SkyLine_Backend

import (
	"flag"
	"fmt"

	ConfigurationEngine "github.com/SkyPenguin-Solutions/SkyLineConfigurationEngine/Engine/Backend"
)

type ProgramConfiguration struct {
	RunEngine     *bool   // Run SLC to check configuration
	RunSource     *bool   // Run a source code file
	RunREPL       *bool   // Run Read Eval Print Loop
	RunServer     *bool   // Run documentation server
	RunEvalCode   *string // Will just run code regularly
	RunEngineFile *string // The engine source code file to check
	RunHelp       *bool   // Run help menu
	DisableOut    *bool   // Disable resulting evaluation output
	LoadConfig    *bool   // Configuration file
	SourceFile    *string // Source code file to run
}

var (
	ProgramConfig ProgramConfiguration
	Version       = "0.0.5"
)

func ParseFlags() {
	ProgramConfig.RunHelp = flag.Bool("help", false, "")     // Call help menu
	ProgramConfig.RunEngine = flag.Bool("engine", false, "") // Call engine
	ProgramConfig.DisableOut = flag.Bool("dout", false, "")  // Call to disable out on repl or files
	ProgramConfig.RunREPL = flag.Bool("repl", false, "")     // Call REPL
	ProgramConfig.RunServer = flag.Bool("server", false, "") // Call server
	ProgramConfig.RunSource = flag.Bool("source", false, "") // Call source code conv
	ProgramConfig.LoadConfig = flag.Bool("conf", false, "")  // Call configuration loader
	ProgramConfig.SourceFile = flag.String("i", "", "")      // Call to input source code
	ProgramConfig.RunEvalCode = flag.String("eval", "", "")  // Call to eval a single line of code
	ProgramConfig.RunEngineFile = flag.String("EF", "", "")  // Call EF ( Engine Source Code Input File)
	// Now continue to parse
	flag.Parse()
	// Finall check the results
}

func PrepareOutConfig(name, value string) {
	str := fmt.Sprintf("\033[38;5;122m[*]\033[38;5;242m Environment Configured \033[38;5;48m%s\033[49m\033[38;5;242m\033[59m  %s\033[39m", name, value)
	fmt.Println(str)
}

func RunSyntaxCheckEngine(filename string) {
	ConfigurationEngine.OutputBanner()
	fmt.Print("\n\n\n")
	ConfigurationEngine.OutputBoxOCode(filename)
	errorList := ConfigurationEngine.Start(filename)
	fmt.Print("\n\n")
	if !errorList {
		PrepareOutConfig("| Name         |", ConfigurationEngine.Exportable_data.ProjectData.Name)
		PrepareOutConfig("| Description  |", ConfigurationEngine.Exportable_data.ProjectData.Description)
		PrepareOutConfig("| Version      |", ConfigurationEngine.Exportable_data.ProjectData.Version)
		PrepareOutConfig("| Languages    |", ConfigurationEngine.Exportable_data.ProjectData.Languages)
		PrepareOutConfig("| Supported OS |", ConfigurationEngine.Exportable_data.ProjectData.SuportedOS)
		for idx := 0; idx < len(ConfigurationEngine.Exportable_data.ProjectData.Require); idx++ {
			PrepareOutConfig("| Package      |", ConfigurationEngine.Exportable_data.ProjectData.Require[idx])
		}
	}
}
