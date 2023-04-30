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
// Filename      |  SkyLine_Standard_PreRegistered_Library_Mathematics.go
// Project       |  SkyLine programming language
// Line Count    |  0+ active lines
// Status        |  Inactive
// Package       |  SkyLine_Backend
//
//
// Defines       | functions to grab process informaton, just delete this file eventually once tested. It does not work the way it needs to
//
package SkyLine_Backend

import (
	"log"
	"os/exec"
	"strings"
)

func VerErr(x error) {
	if x != nil {
		log.Fatal("ERROR >> ", x)
	}
}

type ProcessInformation struct {
	ProcessName string // Name of the given process             | csc
	ProcessID   string // Process ID of the selected process    | 7886
	ProcessPath string // Path of the currently running process | /proc/7886/fd/0
}

// For linux we use the command as it is much easier to leverage
func (PI *ProcessInformation) PIDbyProgramName(Progname string) {
	cout, x := exec.Command("pidof", Progname).Output()
	VerErr(x)
	o := string(cout)
	oline := strings.Split(o, "\n")
	for _, l := range oline {
		if l != "" {
			PI.ProcessID = l
			PI.ProcessName = Progname
		}
	}
}
