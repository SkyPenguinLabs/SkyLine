///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Object_Module
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions for a modular definition
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

import (
	"fmt"
	"path"
)

func (SL_ObjectModule *SL_Module) SkyLine_ObjectFunction_GetDataType() ObjectDataType {
	return SKYLINE_DATATYPE_MODULE_OBJECT
}

func (SL_ObjectModule *SL_Module) SkyLine_ObjectFunction_GetTrueValue() string {
	return fmt.Sprintf("<SkyLine_Module '%s' > ", SL_ObjectModule.SL_ModuleNames)
}

func (SL_ObjectModule *SL_Module) GetModuleFile() string {
	fileName := path.Base(SL_ObjectModule.SL_ModuleNames) // Get the file name with extension
	return fileName[:len(fileName)-len(path.Ext(fileName))]
}

func (SL_ObjectModule *SL_Module) SkyLine_ObjectFunction_InvokeObject(Call string, Environ SkyLineEnvironment, InvokeArgs ...SL_Object) SL_Object {
	return nil
}

func (SL_ObjectModule *SL_Module) SkyLine_ObjectFunction_GetInterface() interface{} {
	return "<MODULE>"
}
