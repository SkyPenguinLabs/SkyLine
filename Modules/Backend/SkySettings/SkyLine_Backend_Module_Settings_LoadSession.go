///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Settings_LoadSession
// Extension         | .go ( golang source code file )
// Purpose           | Will load all data into the session information (OSNAME, OSUSER, OSVERSION, KERNEL ETC)
// Directory         | Modules/Backend/SkyConsole
// Modular Directory | SkyLine/Modules/Backend/SkyConsole
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

import "runtime"

func init() {
	SkySession.OS_Name = runtime.GOOS
	SkySession.OS_Arch = runtime.GOARCH
	SkySession.Go_Version = runtime.Version()
}
