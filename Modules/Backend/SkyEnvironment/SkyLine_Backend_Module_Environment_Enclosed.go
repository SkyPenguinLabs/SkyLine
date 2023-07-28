///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Enclosed
// Extension         | .go ( golang source code file )
// Purpose           | Defines functions for setting a new enclosed environment
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines a function that can set a new enclosed environment. This will create a new environment which is spawned or called
//
// by an outer environmental parameter.
//
package SkyLine_Backend_Modules_Objects

func SL_NewEnclosedEnvironment(SL_Outer *SkyLineEnvironment) *SkyLineEnvironment {
	Environment := SL_NewEnvironment()
	Environment.SkyLine_Outer = SL_Outer
	return Environment
}
