///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_Models
// Extension         | .go ( golang source code file )
// Purpose           | Define constant definitions for string values of Tokens
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file helps define the proper interfaces and models that the lexer/scanner will rely on to keep track and store information. These structures help the scanner when categorizing
//
// cutting or moving onto other tokens as well as keeping track of the lines currently being parsed or scanned in the file.
//
//
package SkyLine_Backend_Scanner

import (
	SkyErrors "SkyLine/Modules/Backend/SkyErrorSystem"
	SLTK "SkyLine/Modules/Backend/SkyTokens"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Type list
type (
	SL_ScannerInterface interface {
		NT() SLTK.SL_TokenConstruct
	}

	SL_ScannerStructure struct {
		CharacterInput            string                 // Scanners current character based input
		Scanner_POS               int                    // Scanners current position
		Scanner_RPOS              int                    // Scanners READ position
		Scanner_Character         rune                   // Scanners current character
		Scanner_Characters        []rune                 // Scanners list of current characters
		Scanner_PreviousToken     SLTK.SL_TokenConstruct // Scanners previous token construct
		Scanner_PreviousCharacter byte                   // Scanners previous character
		Scanner_CurrentLine       int                    // Scanners current line
	}

	// Functions
	// These functions are used for map comparisons later on when the implementation of regex comes into play

	TokenizerFunction func(*SL_ScannerStructure) SLTK.SL_TokenConstruct
	TokenizerPeeked   func(*SL_ScannerStructure) SLTK.SL_TokenConstruct
)

var Datatypes = []string{
	"string.",
	"float.",
	"object.",
	"hash.",
	"array.",
	"boolean.",
}

var ConstantIdents = make(map[string]bool)

type Controller struct {
	ConstantIdentifiers []struct {
		Name   string `json:"Name"`
		Module string `json:"Module"`
	} `json:"ConstantIdentifiers"`
}

func init() {
	f, x := ioutil.ReadFile("/usr/share/Modules/Backend/SkyDB/ProgramaticFiles/ConstantIdentifiersStandard.json")
	if x != nil {
		log.Fatal(x)
	}
	var controller Controller
	x = json.Unmarshal(f, &controller)
	if x != nil {
		log.Fatal(x)
	}
	var newarr []string
	var secarr []string
	for i := 0; i < len(controller.ConstantIdentifiers); i++ {
		ConstantIdents[controller.ConstantIdentifiers[i].Name] = true
		newarr = append(newarr, controller.ConstantIdentifiers[i].Name)
		secarr = append(secarr, controller.ConstantIdentifiers[i].Module)
	}
	SkyErrors.ErrorConstEnv.Update_LibraryFunctions(newarr, secarr)
}
