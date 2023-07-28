//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____ _ _     _____         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_| |___|   __|_ _ ___| |_ ___ ______
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | | -_|__   | | |_ -|  _| -_|     |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |_|_|___|_____|_  |___|_| |___|_|_|_|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|                   |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This file defines all models and types used by the library
//
//
package SkyLine_Standard_Library_File

type FileInit struct {
	Filename string
	Mode     int
	Lines    []string
}

var File FileInit
