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
// This file defines all of the base functions for color related functions
//
package SkyLine_StandardLib_Console

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Works only for the current standard emulators and linux systems
func ForegroundANSI(r, g, b uint8, onlyreturn bool) string {
	if onlyreturn {
		return fmt.Sprintf(`\033[38;2;%d;%d;%dm`, r, g, b)
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// Parses a HTML RGB color code
// without #
func ParseRGB(code string) (r, g, b uint8, x error) {
	code = strings.Trim(code, "#")
	if len(code) != 6 {
		return 0x00, 0x00, 0x00, errors.New("Error parsing code: The RGB color code must be 6 characters long (2 R, 2 G, 2 B) and must not include # inside of the string. Please try again")
	}
	decR, _ := strconv.ParseUint(code[0:2], 16, 8)
	decG, _ := strconv.ParseUint(code[2:4], 16, 8)
	decB, _ := strconv.ParseUint(code[4:6], 16, 8)
	r = uint8(decR)
	g = uint8(decG)
	b = uint8(decB)
	return r, g, b, x
}
