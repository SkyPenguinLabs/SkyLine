///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Settings_Models
// Extension         | .go ( golang source code file )
// Purpose           | Define all models for script settings
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

type (
	//::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: Script settings is the settings for the script
	//::::::::::::::::::::::::::::::::::::::::::::::::::
	SL_Settings struct {
		SkyLine_Run_SLC                *bool   // Tell SkyLine to only run and parse SLC files
		SkyLine_Run_File               *bool   // Tell SkyLine to run a source code file
		SkyLine_Run_Eval               *bool   // Tell SkyLine to evaluate a single line of code
		SkyLine_Use_VM                 *bool   // Tell SkyLine to use SkyVM
		SkyLine_Compile                *bool   // Tell SkyLine to compile to bytecode
		SkyLine_AllowOut               *bool   // Tell SkyLine to allow IO access from the program ( when a variable for example `x` is defined then placed it will output the value)
		SkyLine_InputFile              *string // Tell SkyLine to load the given input file
		SkyLine_EngineFile             *string // Tell SkyLine to load an SLC engine file
		SkyLine_Console                *bool   // Tell SkyLine to run the REPL
		SkyLine_Help                   *bool   // Tell SkyLine to print the help menu
		SkyLine_ReplExecEnd            *bool   // This will tell SkyLine to draw a box around the code it is executing, then print the results under neath with the SL banner
		SkyLine_Tooling_SLC_AutoGen    *bool   // This will tell SkyLine to automatically generate templates for code projects and call the SLC engine on demand
		SkyLine_Tooling_SLC_AutoGenSrc *string // This will tell SkyLine to pass the data to the SLC engine for building the project.
	}

	//::::::::::::::::::::::::::::::::::
	// User / Session settings
	//::::::::::::::::::::::::::::::::::
	SL_Session struct {
		OS_Name    string
		OS_User    string
		OS_Arch    string
		Go_Version string
	}
)

var (
	SkySession  SL_Session
	SkySettings SL_Settings

	HelpMenu = `
	
	CSC (Cyber Security Core) | Golang version      ( %s ) 
	CSC (Cyber Security Core) | OS Picked Up        ( %s ) 
	CSC (Cyber Security Core) | Architecture        ( %s ) 


		(bool)   --help             | Load the SkyLine help menu ( this menu )
		(bool)   --SLC              | Enable syntactic checking for SkyLineConfig
		(bool)   --SL               | Enable execution of SkyLine source code files
		(bool)   --Output           | Disable/Enable variable value output when variable is placed 
		(bool)   --Repl             | Load the Read Eval Print loop ( console )
		(string) --Eval             | Load code through the evaluator as a single line
		(string) --source           | Tell skyline to load a given source code file
		(string) --Esource          | Tell skyline to load a given source code file for SkyLineConfiguration

	`
)
