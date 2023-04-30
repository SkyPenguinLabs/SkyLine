package main

import (
	"fmt"
	Mod "main/Modules/Backend"
	"os"
	"runtime"
)

func init() {
	Mod.ParseFlags()
}

func Help() {
	Mod.Banner()
	fmt.Printf("CSC (Cyber Security Core) | Golang version      ( %s ) \n", runtime.Version())
	fmt.Printf("CSC (Cyber Security Core) | OS Picked Up        ( %s ) \n", Mod.U.OperatingSystem)
	fmt.Printf("CSC (Cyber Security Core) | Architecture        ( %s ) ", Mod.U.OperatingSystemArchitecture)
	fmt.Print("\n\n")
	keys := `
	(bool)   --help             | Load the SkyLine help menu ( this menu )
	(bool)   --engine           | Enable engine code checkin 
	(bool)   --source           | Tell skyline to enable source code input
	(bool)   --dout             | Disable/Enable variable value output when variable is placed 
	(bool)   --repl             | Load the Read Eval Print loop ( console )
	(string) --eval             | Load code through the evaluator as a single line
	(string) --EF               | SkyLine Configuration Engine source code file
	(string) --i                | Input source code file 
	`
	fmt.Println(keys)
	os.Exit(0)
}

func main() {

	// Call to help
	if *Mod.ProgramConfig.RunHelp {
		Help()
	}

	// Call to run source code file
	if *Mod.ProgramConfig.RunSource {
		if *Mod.ProgramConfig.SourceFile != "" {
			if err := Mod.RunAndParseFiles(*Mod.ProgramConfig.SourceFile, *Mod.ProgramConfig.DisableOut); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(0)
			}
		} else {
			fmt.Println("[-] Error: Please give a input file with the file extension .csc")
		}
	}
	// Call to run evaluator
	if *Mod.ProgramConfig.RunEvalCode != "" {
		code := *Mod.ProgramConfig.RunEvalCode
		Mod.RunAndParseNoFile(code, *Mod.ProgramConfig.DisableOut)
		os.Exit(0)
	}
	// Call to run engine
	if *Mod.ProgramConfig.RunEngine {
		engineFile := *Mod.ProgramConfig.RunEngineFile
		Mod.RunSyntaxCheckEngine(engineFile)
		os.Exit(0)
	}
	// Run REPL
	if *Mod.ProgramConfig.RunREPL {
		fmt.Println("\x1b[H\x1b[2J\x1b[3J")
		Mod.Banner()
		Mod.Start(os.Stdin, os.Stdout, *Mod.ProgramConfig.DisableOut)
		return
	}
}
