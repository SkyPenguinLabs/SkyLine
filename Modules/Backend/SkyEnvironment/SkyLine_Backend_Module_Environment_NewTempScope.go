///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_NewTempScop
// Extension         | .go ( golang source code file )
// Purpose           | Defines environmental new temporary scope
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file and module basically contains a definition to a function that helps create a new temporary environmental scope which is basically a better way to
//
// create new scopes within the programming language.
//
//
package SkyLine_Backend_Modules_Objects

func SL_NewEnvironmentTemporaryScope(OuterEnv *SkyLineEnvironment, Keys []string) *SkyLineEnvironment {
	Environment := SL_NewEnvironment()
	Environment.SkyLine_Outer = OuterEnv
	Environment.SkyLine_Permit = Keys
	return Environment
}
