///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Set
// Extension         | .go ( golang source code file )
// Purpose           | Defines functions to set variables or load variables into the environment
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This module defines all of the environmental functions that exist within the language to set or load variables into the VENV ( Virtual Environment ). These functions
//
// allow SkyLine to set and store dynamic objects into the environment that have a name and a value. For example
//
// (10 + 20) / 90 * 20
//
// will not be stored, it has a value but does not have a name
//
// but
//
// set x := 10;
//
// will be stored as `x:10` because x is the variable name and 10 is the value of x. We store variable names and values only to avoid collison and memory corruption when loading
//
// or executing code in the evaluator or virtual machine
//
package SkyLine_Backend_Modules_Objects

import (
	"fmt"
	"os"
)

func (SL_ENV *SkyLineEnvironment) Load(VarName string, Value SL_Object) SL_Object {
	Current := SL_ENV.SkyLine_Storage[VarName]
	if Current != nil && SL_ENV.SkyLine_RO[VarName] {
		fmt.Println("[+] Should not be trying to modify a constant value")
		os.Exit(0)
	}
	if len(SL_ENV.SkyLine_Permit) > 0 {
		for _, v := range SL_ENV.SkyLine_Permit {
			if v == VarName {
				SL_ENV.SkyLine_Storage[VarName] = Value
				return Value
			}
		}
		if SL_ENV.SkyLine_Outer != nil {
			return SL_ENV.SkyLine_Outer.Load(VarName, Value)
		}
		fmt.Printf("scoping weirdness; please report a bug\n")
		os.Exit(5)
	}
	SL_ENV.SkyLine_Storage[VarName] = Value
	return Value
}

func (SL_ENV *SkyLineEnvironment) LoadConstant(CName string, Value SL_Object) SL_Object {
	SL_ENV.SkyLine_Storage[CName] = Value
	SL_ENV.SkyLine_RO[CName] = true
	return Value
}
