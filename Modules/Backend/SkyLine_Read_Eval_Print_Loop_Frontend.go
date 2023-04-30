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
// Filename      |  SkyLine_Parser_Subroutine_FileSys_ParserRunner.go
// Project       |  SkyLine programming language
// Line Count    |  80 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines all frontend art or backend settings functions or flags for this programming language.
//
// State         | Working but can be worked on
// Changes?      | This can all be organized and thrown into respected files, types and constants and vars should go into models files and the functions can get thrown into frontend
//                 based files such as Art or Box or IO files
//
//
//
package SkyLine_Backend

import (
	"fmt"
	"runtime"
)

func Banner() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	U.OperatingSystem = runtime.GOOS
	U.OperatingSystemArchitecture = runtime.GOARCH
	switch runtime.GOOS {
	case "linux":
		fmt.Println("\t\t\t	 \033[38;5;51m┏━┓\x1b[0m")
		fmt.Println("\t\t\t	\033[38;5;56m┃\033[38;5;51m┃ ┃\x1b[0m")
		fmt.Println("\t\t\t    \033[38;5;56m━━━━┛\x1b[0m")
		fmt.Println("\t\t\t	\033[38;5;249m   Sky Line Interpreter| V 0.0.5")
		fmt.Print("\n\n\n\n\033[39m")
	default:
		fmt.Println("\t\t\t\t	 \u001b[38;5;51m┏━┓\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;56m┃\u001b[38;5;51m┃ ┃\u001b[0m")
		fmt.Println("\t\t\t\t    \u001b[38;5;56m━━━━┛\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;249mSky Line Interpreter| V 0.0.5")

	}
}
