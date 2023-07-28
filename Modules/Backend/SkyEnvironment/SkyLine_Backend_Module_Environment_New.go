///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_New
// Extension         | .go ( golang source code file )
// Purpose           | Defines a function to create a new environment
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines a function that will spin up a SkyLine environment. This environment will actually help with the use cases of working with
//
// variables as the virtual environment is what helps the language store variables in memory. Since we can not just call a bunch of variables
//
// unorganized we should create a shell or environment that can hold those values, variables or in the VM's case operations.
//
package SkyLine_Backend_Modules_Objects

func SL_NewEnvironment() *SkyLineEnvironment {
	Storage := make(map[string]SL_Object) // Storage
	ROO := make(map[string]bool)          // Read Only Object
	return &SkyLineEnvironment{
		SkyLine_Storage: Storage,
		SkyLine_RO:      ROO,
		SkyLine_Outer:   nil,
	}
}
