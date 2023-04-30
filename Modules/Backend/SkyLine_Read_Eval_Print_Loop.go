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
// Filename      |  SkyLine_Read_Eval_Print_Loop.go
// Project       |  SkyLine programming language
// Line Count    |  30 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines a REPL starter function
//
// State         | Working but can be worked on
// Changes?      | Can be thrown into another sub function file or file dedicated to REPL
//
//
package SkyLine_Backend

import (
	"bufio"
	"fmt"
	"io"
	"runtime"
	"strings"
)

const prompt = "SkyLine|%s(%s)>> "

var counter int

func CallBox(line string) string {
	longest := 0
	longest = len(line)
	counter++
	return fmt.Sprintf("[%d]  │ %s%s", counter, line, strings.Repeat(" ", longest-len(line)))
}

func Start(in io.Reader, Out io.Writer, callout bool) {
	scanner := bufio.NewScanner(in)
	Env := NewEnvironment()

	for {
		fmt.Printf(prompt, runtime.GOOS, runtime.GOARCH)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		fmt.Print("\033[1A\033[K")
		fmt.Printf("%s\n", SKYLINE_WHITE+CallBox(line)+SKYLINE_RESTORE)
		l := LexNew(line)
		parser := New_Parser(l)
		program := parser.ParseProgram()
		if len(parser.ParserErrors()) != 0 {
			printParserErrors(line, Out, parser.ParserErrors())
			continue
		}
		evaluated := Eval(program, Env)
		if evaluated == nil {
			continue
		}
		if callout {
			io.WriteString(Out, evaluated.SL_InspectObject()+"\n")
		}
	}
}

func printParserErrors(line string, Out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(Out, msg)
		io.WriteString(Out, "\n")
	}
}
