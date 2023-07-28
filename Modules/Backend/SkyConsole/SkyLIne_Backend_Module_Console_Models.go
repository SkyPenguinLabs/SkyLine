///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Console_Models
// Extension         | .go ( golang source code file )
// Purpose           | Define all models and types for the console application
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
	SkyFS "SkyLine/Modules/Backend/SkyFS"
	SkyParser "SkyLine/Modules/Backend/SkyParser"
	SkyScanner "SkyLine/Modules/Backend/SkyScanner"

	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type (
	//:::::::::::::::::::::::::::::::
	//:: Format for the box output
	//:::::::::::::::::::::::::::::::
	BoxFormat struct {
		TL string
		TR string
		BL string
		BR string
		HZ string
		VT string
	}

	//::::::::::::::::::::::::::::::::
	//:: REPL settings / current data
	//:::::::::::::::::::::::::::::::
	SL_REPL struct {
		LineCount    int
		Threads      int
		Environment  []string
		EnvironmentB []byte
		Out          io.Writer
		In           io.Reader
		Callout      bool
		CurrentEnv   *SkyEnvironment.SkyLineEnvironment
	}
)

// Note: These are not color schemes but are defined that way for simplicity, they are just sets of ANSI codes.
var (
	REPL         SL_REPL
	LINUX_CLS    = "\x1b[H\x1b[2J\x1b[3J" // Color Scheme | Linux unicode set ( Clear  )
	LINUX_RET    = "\x1b[0m"              // Color Scheme | Linux unicode set ( Reset  )
	LINUX_AQUA   = "\033[38;5;51m"        // Color Scheme | Linux unicode set ( AQUA   )
	LINUX_PURPLE = "\033[38;5;56m"        // Color Scheme | Linux unicode set ( VIOLET )
	LINUX_GREY   = "\033[38;5;249m"       // Color Scheme | Linux unicode set ( GREY   )
	LINUX_WHITE  = "\033[38;5;255m"       // Color Scheme | Linux unicode set ( WHITE   )
	REPL_OUT     = "SkyLine|%s(%s)>> "

	REPL_Commands = map[string]func(){
		"exit": func() {
			os.Exit(0)
		},
		"REPL_SETTINGS": func() {
			fmt.Println("Line Count| ", REPL.LineCount)
			fmt.Println("Threads   | ", REPL.Threads)
		},
		"REPL_ENVIRONMENT": func() {
			SkyLine_Draw_Around_REPL_ENV()
			// draw box for environment
		},
		"REPL_SAVE": func() {
			dir := "SkyLine_Temp_Directory"
			filename := "SkyLine_REPL_Save.sl"
			x := os.Mkdir(dir, 0755)
			if x != nil {
				if x.Error() != "mkdir SkyLine_Temp_Directory: file exists" {
					fmt.Println("[-] Error when saving file to directory: ", x)
				}
			}
			joined := strings.Join(REPL.Environment, "\n")
			x = ioutil.WriteFile(fmt.Sprintf("%s/%s", dir, filename), []byte(joined), 0644)
			if x != nil {
				fmt.Println("[-] Error when trying to write the environment -> ", x)
			}
			DrawUtilsBox(fmt.Sprintf("Environment saved in -> %s/%s", dir, filename))
		},
		"REPL_REPLAY_LAST": func() {
			if len(REPL.Environment) > 0 {
				ExecuteSingleLine(REPL.Environment[len(REPL.Environment)-1], REPL.Out, REPL.Callout, REPL.CurrentEnv)
			} else {
				fmt.Println("[-] Nothing todo ( REPLAY_FAIL -> Reason[No environment data to load]) ")
				fmt.Println("[F] Check   : REPL_ENVIRONMENT")
			}
		},
		"REPL_REPLAY_ALL": func() {
			if len(REPL.Environment) > 0 {
				for i := 0; i < len(REPL.Environment); i++ {
					ExecuteSingleLine(REPL.Environment[i], REPL.Out, REPL.Callout, REPL.CurrentEnv)
				}
			} else {
				fmt.Println("[-] Nothing todo ( REPLAY_FAIL -> Reason[No environment data to load]) ")
				fmt.Println("[F] Check   : REPL_ENVIRONMENT")
			}
		},
		"REPL_LOAD_SAVE": func() {
			f, x := os.Open("SkyLine_Temp_Directory/SkyLine_REPL_Save.sl")
			if x != nil {
				fmt.Println("[-] Nothing todo ( Session Replay ) -> Could not load workspace : ", x.Error())
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				REPL.Environment = append(REPL.Environment, scanner.Text())
			}
			for i := 0; i < len(REPL.Environment); i++ {
				fmt.Print(LINUX_GREY)
				REPL.LineCount++
				fmt.Println(SkyLine_Code_Box(REPL.Environment[i]))
			}
			fmt.Println("\n --- Executed Save ---")
			input, err := ioutil.ReadFile("SkyLine_Temp_Directory/SkyLine_REPL_Save.sl")
			if err != nil {
				log.Fatal(err)
			}
			SkyFS.Current.FileSystem_Modify_Name("SkyLine_Temp_Directory/SkyLine_REPL_Save.sl")
			SkyFS.Current.FileSystem_ModifyInfo()
			SkyFS.Current.FileSystem_ModifyExtension()
			SkyFS.Current.FileSystem_ModifyDir()
			Execute(string(input), REPL.CurrentEnv)

		},
	}
)

func Execute(input string, Environment *SkyEnvironment.SkyLineEnvironment) int {
	Scan := SkyScanner.New(input)
	Parse := SkyParser.SkyLineNewParser(Scan)
	Prog := Parse.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
	if len(Parse.SkyLine_Parser_Helper_Ret_Errors()) != 0 {
		for _, msg := range Parse.SkyLine_Parser_Helper_Ret_Errors() {
			fmt.Printf("\t%s\n", msg)
		}
		os.Exit(0)
	}
	InitateScanner := SkyScanner.New("")
	InitateParser := SkyParser.SkyLineNewParser(InitateScanner)
	InitateProg := InitateParser.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
	SkyEval.SkyLine_Call_Eval(InitateProg, Environment)
	SkyEval.SkyLine_Call_Eval(Prog, Environment)
	return 0
}
