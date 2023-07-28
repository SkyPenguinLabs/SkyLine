///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Grabbers
// Extension         | .go ( golang source code file )
// Purpose           | Defines retriever functions
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file and module contains functions known as Grabbers or retrievers. We call these functions retrievers or grabbers because they can locate and check for variables
//
// within the environment. For example, if wewanted to get the value of a given variable we would call SL_ENVIRONMENT.Get()
//
package SkyLine_Backend_Modules_Objects

import "strings"

func (SL_ENV *SkyLineEnvironment) Get(VarName string) (SL_Object, bool) {
	Objectmem, ok := SL_ENV.SkyLine_Storage[VarName]
	if !ok && SL_ENV.SkyLine_Outer != nil {
		Objectmem, ok = SL_ENV.SkyLine_Outer.Get(VarName)
	}
	return Objectmem, ok
}

func (SL_ENV *SkyLineEnvironment) LocateNames(SL_PREFIX string) []string {
	var VALS []string
	for key := range SL_ENV.SkyLine_Storage {
		// this might be an issue later on
		if strings.HasPrefix(key, SL_PREFIX) {
			VALS = append(VALS, key)
		}
		if strings.HasPrefix(key, "object.") {
			VALS = append(VALS, key)
		}
	}
	return VALS
}
