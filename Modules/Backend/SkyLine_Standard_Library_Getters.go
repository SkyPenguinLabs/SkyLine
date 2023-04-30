////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// This file contains all the 'getters' for specific settings within the script, this will return the data type of a value, check the value
// check arguments, verify script data and much more which is required before being called.
//
//
//
package SkyLine_Backend

import "strings"

// Get the data type of the argument
func Get_ArgType(arg SLC_Object) Type_Object {
	return arg.SL_RetrieveDataType()
}

// Get the value of the argument
func Get_ArgValue(arg SLC_Object) string {
	return arg.SL_InspectObject()
}

// Check if the argument is of the proper type
func Get_ArgSupportCompare(arg SLC_Object, list_supported []string) bool {
	for k := 0; k < len(list_supported); k++ {
		if list_supported[k] == strings.ToLower(arg.SL_InspectObject()) {
			return true
		}
	}
	return false
}
