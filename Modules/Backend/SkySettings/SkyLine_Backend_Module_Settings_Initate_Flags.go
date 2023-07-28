///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Settings_Initate_Flags
// Extension         | .go ( golang source code file )
// Purpose           | Define all and loads all flags for the language
// Directory         | Modules/Backend/SkyConsole
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyConsole
// Package Name      | SkyLine_Backend_Module_Console
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file is also not directly needed however it helps when parsing flags or command line options that are provided by the user. Command line options allow users to be much more
//
// flexible during development and always have a reason to exist. In this case, command line flags can really tell SkyLine anything. Below are some examples of what flags allow users
//
// to do.
//
// run the SkyLine condfig engine, run the console, run a simple line of evaluation, use regex based lexers, use code for eval, parse files, allow files, allow specific syntax
//
// allow specific engines or optimization features, use configuration engines, generate projects, use projects, generate and load plugins, load backend plugins, specify import paths
//
// specify timers, disable locks, disable threads etc etc.
//
package SkyLine_Backend_Module_Settings

import "flag"

func init() {
	SkySettings.SkyLine_Run_File = flag.Bool("SL", false, "")
	SkySettings.SkyLine_Run_SLC = flag.Bool("SLC", false, "")
	SkySettings.SkyLine_AllowOut = flag.Bool("Output", false, "")
	SkySettings.SkyLine_Compile = flag.Bool("Compile", false, "")
	SkySettings.SkyLine_Run_Eval = flag.Bool("Eval", false, "")
	SkySettings.SkyLine_Use_VM = flag.Bool("Vm", false, "")
	SkySettings.SkyLine_Console = flag.Bool("Repl", false, "")
	SkySettings.SkyLine_Help = flag.Bool("help", false, "")
	SkySettings.SkyLine_InputFile = flag.String("source", "", "")
	SkySettings.SkyLine_EngineFile = flag.String("Esource", "", "")
	SkySettings.SkyLine_ReplExecEnd = flag.Bool("ReplExec", false, "")
	SkySettings.SkyLine_Tooling_SLC_AutoGen = flag.Bool("Generate", false, "")
	SkySettings.SkyLine_Tooling_SLC_AutoGenSrc = flag.String("Project", "Configuration/Project.json", "")
	flag.Parse()
}
