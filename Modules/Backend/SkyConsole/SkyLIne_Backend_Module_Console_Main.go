///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Console_Main
// Extension         | .go ( golang source code file )
// Purpose           | Define the main function for the execution
// Directory         | Modules/Backend/SkyConsole
// Modular Directory | SkyLine/Modules/Backend/SkyConsole
// Package Name      | SkyLine_Backend_Module_Console
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This is known as the REPL otherwise and commonly known as a Read Eval Print Loop which goes straight to executing the next input and spins up a mock environment.
//
// This is not a necessary step within language development and is more of a feature. REPL helps users easily and quickly test environments and allows them to save workspaces
//
// and environments depending on the language. Languages like R-Script allow you to save the workspace while languages like ruby might not allow you to do so as easily or at all.
//
// SkyLine's REPL will be fully customized and much more modern than people may expect!
//
package SkyLine_Backend_Module_Console

import (
	SkyEnvironment "SkyLine/Modules/Backend/SkyEnvironment"
	SkyEval "SkyLine/Modules/Backend/SkyEvaluator"
	SkyParser "SkyLine/Modules/Backend/SkyParser"
	SkyScanner "SkyLine/Modules/Backend/SkyScanner"
	"bufio"
	"fmt"
	"io"
	"runtime"
)

func ExecuteSingleLine(line string, Out io.Writer, callout bool, Env *SkyEnvironment.SkyLineEnvironment) {
	l := SkyScanner.New(line)
	parser := SkyParser.SkyLineNewParser(l)
	program := parser.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
	if len(parser.SkyLine_Parser_Helper_Ret_Errors()) != 0 {
		ExecuteParserErrors(line, Out, parser.SkyLine_Parser_Helper_Ret_Errors())
	}
	evaluated := SkyEval.SkyLine_Call_Eval(program, Env)
	if evaluated == nil {
		return
	}
	if callout {
		io.WriteString(Out, evaluated.SkyLine_ObjectFunction_GetTrueValue()+"\n")
	}
}

func Start(in io.Reader, Out io.Writer, callout bool) {
	REPL.In = in
	REPL.Out = Out
	REPL.Callout = callout
	fmt.Println(LINUX_CLS)
	SkyLine_Console_Banner()
	scanner := bufio.NewScanner(in)
	Env := SkyEnvironment.SL_NewEnvironment()
	REPL.CurrentEnv = Env
	REPL.Threads++
	for {
		fmt.Printf(REPL_OUT, runtime.GOOS, runtime.GOARCH)
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		if m, ok := REPL_Commands[line]; ok {
			m()
			println(LINUX_RET)
		} else {

			REPL.LineCount++
			REPL.Environment = append(REPL.Environment, line)
			fmt.Print("\033[1A\033[K")
			fmt.Printf("%s\n", LINUX_WHITE+SkyLine_Code_Box(line)+LINUX_RET)
			l := SkyScanner.New(line)
			parser := SkyParser.SkyLineNewParser(l)
			program := parser.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
			if len(parser.SkyLine_Parser_Helper_Ret_Errors()) != 0 {
				ExecuteParserErrors(line, Out, parser.SkyLine_Parser_Helper_Ret_Errors())
				continue
			}
			evaluated := SkyEval.SkyLine_Call_Eval(program, Env)
			if evaluated == nil {
				continue
			}
			if callout {
				io.WriteString(Out, evaluated.SkyLine_ObjectFunction_GetTrueValue()+"\n")
			}
		}
	}
}
