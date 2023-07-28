///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Exporter
// Extension         | .go ( golang source code file )
// Purpose           | Defines a function to grab the names within the environment based on the prefix
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
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

import "strings"

func (env *SkyLineEnvironment) Names(prefix string) []string {
	var ret []string
	for key := range env.SkyLine_Storage {
		if strings.HasPrefix(key, prefix) {
			ret = append(ret, key)
		}
		if strings.HasPrefix(key, "object.") {
			ret = append(ret, key)
		}
	}
	return ret
}
