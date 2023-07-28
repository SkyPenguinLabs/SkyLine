///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Exporter
// Extension         | .go ( golang source code file )
// Purpose           | Defines an exportation function to export the hash of a environment. This is used for modules
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

import "unicode"

func (e *SkyLineEnvironment) ExportedHash() *SL_HashMap {
	pairs := make(map[HashKey]HashPair)
	for k, v := range e.SkyLine_Storage {
		if unicode.IsUpper(rune(k[0])) {
			s := &SL_String{Value: k}
			pairs[s.SL_HashKeyType()] = HashPair{Key: s, Value: v}
		}
	}
	return &SL_HashMap{Pairs: pairs}
}
