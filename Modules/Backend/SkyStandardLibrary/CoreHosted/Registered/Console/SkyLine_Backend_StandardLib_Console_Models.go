//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _              _____               _         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___ |_   _|___ ___ _____|_|___ ___| |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|  | | -_|  _|     | |   | .'| | |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|      | |___|_| |_|_|_|_|_|_|__,| |_|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Defines -> This section of the standard library contains information for the console based functions that include frontend based functions such as tables, organizations,
//
// data analytics, informational organization, color, output etc.
//
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines and declares all of the required constants and variables used within this library
//
package SkyLine_StandardLib_Console

type (
	Table struct {
		HeaderCrossOriginL       string // This will be the far left cross origin header frame meet
		HeaderCrossOriginL_Color string // This will be the far left cross origin header frame meet
		HeaderCrossOriginR       string // This will be the far right cross origin header frame meet
		HeaderCrossOriginR_Color string // This will be the far right cross origin header frame meet
		CrossLineY               string // This will be the far y cross line meet
		CrossLineY_Color         string // This will be the far y cross line meet color
		CrossLineX               string // This will be the far x cross line meet
		CrossLineX_Color         string // This will be the far x cross line meet color
		ColumnTitleColor         string // This will be the column title color
		RowDataColor             string // This will be the row data color
	}
)

var (
	// Table assignment
	T Table
)
