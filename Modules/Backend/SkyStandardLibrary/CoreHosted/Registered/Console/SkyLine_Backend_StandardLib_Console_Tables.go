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
// This file defines all relative setup and initation functions for table based output using specific characters
//
package SkyLine_StandardLib_Console

import (
	"strings"
)

func IfTableNilCheckDefaults() {
	defaults := map[*string]string{
		&T.ColumnTitleColor:         "\033[38;5;255m",
		&T.HeaderCrossOriginL_Color: "\033[38;5;57m",
		&T.HeaderCrossOriginR_Color: "\033[38;5;57m",
		&T.CrossLineX_Color:         "\033[38;5;57m",
		&T.CrossLineY_Color:         "\033[38;5;57m",
		&T.CrossLineX:               "━",
		&T.CrossLineY:               "┃",
		&T.HeaderCrossOriginR:       "┫",
		&T.HeaderCrossOriginL:       "┣",
	}

	for key, value := range defaults {
		if *key == "" {
			*key = value
		}
	}
}

func Console_Lib_DrawTableSepColumnBased(rows [][]string, columns []string) string {
	IfTableNilCheckDefaults()
	colwidth := make([]int, len(columns))
	for o, col := range columns {
		colwidth[o] = len(col)
		for _, rowdata := range rows {
			if len(rowdata[o]) > colwidth[o] {
				colwidth[o] = len(rowdata[o])
			}
		}
	}
	headsep := T.HeaderCrossOriginL_Color + T.HeaderCrossOriginL
	for _, w := range colwidth {
		headsep += strings.Repeat(T.CrossLineX_Color+T.CrossLineX, w+2) + T.HeaderCrossOriginR_Color + T.HeaderCrossOriginR
	}
	head := T.CrossLineY_Color + T.CrossLineY
	for i, col1 := range columns {
		head += " " + T.ColumnTitleColor + col1 + strings.Repeat(" ", colwidth[i]-len(col1)) + " " + T.CrossLineY_Color + T.CrossLineY
	}
	Rowdata := make([]string, len(rows))
	for k, row := range rows {
		RowT := T.CrossLineY_Color + T.CrossLineY
		for l, col := range row {
			RowT += " " + T.ColumnTitleColor + col + strings.Repeat(" ", colwidth[l]-len(col)) + " " + T.CrossLineY_Color + T.CrossLineY
		}
		Rowdata[k] = RowT
	}
	var table string
	table += headsep + "\n" + head + "\n" + headsep + "\n"
	for _, rt := range Rowdata {
		table += rt + "\n"
	}
	table += headsep + "\n"
	return table
}
