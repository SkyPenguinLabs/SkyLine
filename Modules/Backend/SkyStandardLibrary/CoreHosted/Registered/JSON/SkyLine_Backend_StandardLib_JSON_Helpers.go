//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _                __ _____ _____ _____
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___   |  |   __|     |   | |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|  |  |__   |  |  | | | |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_____|_____|_|___|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This part of the standard library contains any and all functions for JSON data sets and functions. This includes dumping relative information, dumping file information,
//
// gnerating golang structures for plugins, generating infromation, generating strings, dumping to a hash map and so on from there.
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all help related functions for the libraries
//
package SkyLine_Standard_Library_JSON

import "os"

func CheckifFile(data string) bool {
	var x error
	_, x = os.Stat(data)
	return x == nil
	// if nil then data is a file
	// else then it is not a file
}
