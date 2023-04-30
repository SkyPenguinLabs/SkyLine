/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Environment_Modification_Via_Modify.go
// Project       |  SkyLine programming language
// Line Count    |  200+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all proper functions for the error system and environmental modiy
//
// STATE         | Useless and needs to be organized
// Resolution    | The data in this file should be thrown into models or a file that is dedicated towards the error system of SkyLine
//
//
package SkyLine_Backend

import (
	"fmt"
	"strings"
)

type ErrorSystemSettings struct {
	Tree bool // Output the entire tree
	Box  bool // Output the box of code where the error occurred
	Line bool // Output a simple message
}

type ImportSystemSettings struct {
	ExpectSingleFile bool // This tells the import system to just take single files
	ExpectDirectory  bool // This tells the import system to take directories not single files
	ExpectAce        bool // This tells the import system to expect all forms and import no matter what, if directory then call directory import
}

func (ESys *ErrorSystemSettings) TreeValid() bool {
	return ESys.Tree
}

func (ESys *ErrorSystemSettings) BoxValid() bool {
	return ESys.Box
}

func (ESys *ErrorSystemSettings) LineValid() bool {
	return ESys.Line
}

// Ace a setting in the program will return the value if to grab both directories and files
func (ISys *ImportSystemSettings) GetAll() bool {
	return ISys.ExpectAce
}

func (ISys *ImportSystemSettings) GetExpectDir() bool {
	return ISys.ExpectDirectory
}

func (ISys *ImportSystemSettings) GetExpectSingleFile() bool {
	return ISys.ExpectSingleFile
}

func ActiveImporterSystemSettings(settings ...bool) {
	ImporterSys.ExpectAce = settings[0]
	ImporterSys.ExpectDirectory = settings[1]
	ImporterSys.ExpectSingleFile = settings[2]
}

func (ESys *ErrorSystemSettings) CallError(tree *TreeNode, BoxOCode string, Line string) {
	if ESys.Box {
		RetTreeSys(tree, "", false)
	}
	if ESys.Line {
		fmt.Println(Line)
	}
	if ESys.Box {
		fmt.Println(BoxOCode)
	}
}

var (
	// System settings
	ErrorSys    ErrorSystemSettings
	ImporterSys ImportSystemSettings
	// Maps
	ModifyEnvironmentTypes = map[string]bool{
		"errors":   true,
		"importer": true,
	}
	ModifyEnvironmentSupportLists = map[string]func(string) bool{
		"errors": func(setting string) bool {
			switch setting {
			case "verbose":
				ErrorSys.Tree = true
			case "basic":
				ErrorSys.Line = true
				ErrorSys.Box = false
				ErrorSys.Tree = false
			case "masters":
				ErrorSys.Box = true
				ErrorSys.Line = true
				ErrorSys.Tree = true
			default:
				ErrorSys.Box = false
				ErrorSys.Line = false
				ErrorSys.Tree = true
			}
			return true
		},
		"importer": func(setting string) bool {
			switch strings.ToLower(setting) {
			case "expect_directory":
				ActiveImporterSystemSettings(false, true, false)
			case "expect_ace":
				ActiveImporterSystemSettings(true, false, false)
			case "expect_file":
				ActiveImporterSystemSettings(false, false, true)
			default:
				ActiveImporterSystemSettings(false, false, true)
			}
			return true

		},
	}
)

func VerifyModification(modifier, setting string) bool {
	if ok := ModifyEnvironmentTypes[strings.ToLower(modifier)]; ok {
		return ModifyEnvironmentSupportLists[strings.ToLower(modifier)](setting)
	} else {
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_SUPPORT_IN_MODIFY_ENVIRONMENT_MODIFY_NOT_SUPPORTED) + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error -> Environment Modify ( Verify Modification ) " + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + " When parsing values of modify() the first argument was not a valid environment system",
									Children: []*TreeNode{
										{
											Type: SKYLINE_HIGH_DEFAQUA + "[Sub Branch] " + SKYLINE_RESTORE,
											Children: []*TreeNode{
												{
													Type: SKYLINE_HIGH_DEFAQUA + "[SB-I] Modify what system     ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + modifier + SKYLINE_RESTORE,
												},
												{
													Type: SKYLINE_HIGH_DEFAQUA + "[SB-I] System Supported       ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(ok) + SKYLINE_RESTORE,
												},
												{
													Type: SKYLINE_HIGH_DEFAQUA + "[SB-I] Modify by what setting ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + setting + SKYLINE_RESTORE,
												},
											},
										},
									},
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "Suggestion" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_SICK_BLUE + `[S] Consider adding a supported system? For example modify("errors:verbose")` + SKYLINE_RESTORE,
								},
							},
						},
					},
				},
			},
		}
		RetTreeSys(root, "", false)
		return false
	}
}
