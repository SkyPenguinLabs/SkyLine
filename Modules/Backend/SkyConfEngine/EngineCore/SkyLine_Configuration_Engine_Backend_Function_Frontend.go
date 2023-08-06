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
//
// File    ->  SkyLine_Configuration_Language_Frontend_ArtWork.go
// Apart   ->  SLC/Modules/Backend
// Source  -> .go, .mod, .SL-Modify
//
// File contains    -> Artwork and frontend design based functions
//

package SkyLine_Configuration_Engine_Backend_Source

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func OutputBanner() {
	fmt.Println("\t\t\t  \033[38;5;51m       ┏━━━━")
	fmt.Println("\t\t\t	 \033[38;5;51m┃\033[38;5;51m┏━┓\x1b[0m")
	fmt.Println("\t\t\t	\033[38;5;56m┃\033[38;5;56m┃\033[38;5;51m┃ ┃\x1b[0m")
	fmt.Println("\t\t\t    \033[38;5;56m━━━━┛┃\x1b[0m")
	fmt.Println("\t\t\t    \033[38;5;56m     ┗━━━━")
	fmt.Println("\t\t\t	\033[38;5;249mSkyLine Configuration Engine 0.0.1 ")
}

func OutputBoxOCode(filename string) {
	input, x := ioutil.ReadFile(filename)
	if x != nil {
		println(x)
		os.Exit(0)
	}
	lines := strings.Split(string(input), "\n")
	longest := 0
	for _, line := range lines {
		if len(line) > longest {
			longest = len(line)
		}
	}
	lastLineNumber := len(lines)
	LENLEN := int(math.Log10(float64(lastLineNumber))) + 1
	width := longest + 6 + LENLEN
	fmt.Printf("[SOF] ┌%s┐\n", strings.Repeat("─", width-2))
	for i, line := range lines {
		fmt.Printf("[%*d]   │ %s%s \n", LENLEN, i+1, line, strings.Repeat(" ", longest-len(line)))
	}
	fmt.Printf("[EOF] └%s┘\n", strings.Repeat("─", width-2))
}

func Clear() {
	fmt.Print("\x1b[H\x1b[2J\x1b[3J")
}
