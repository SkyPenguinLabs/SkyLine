//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             __ __     _____     __
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |___|     |___|  |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|-   -|___| | | |___|  |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__|   |_|_|_|   |_____|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
// This section of the standard library is dedicated to dumping, mapping, matching, adding, generating, parsing, running or loading XML files which can also be parsed
//
// as PLIST files. However, given that PLIST files are technically XML, we put them under their own library as BPLIST ( Binary Property List ) works with regular PLIST
//
// parsers and programs.
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all functions to help the XML library during specific operations such as dumping, parsing, loading or converting
//
package SkyLine_Standard_Library_XML

import "strings"

func XML_Helper_EraseJsonUnescapedString(str string) string {
	str = strings.ReplaceAll(str, "\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\", "\\")
	str = strings.ReplaceAll(str, "\\/", "/")
	str = strings.ReplaceAll(str, "\\b", "\b")
	str = strings.ReplaceAll(str, "\\f", "\f")
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "\\r", "\r")
	str = strings.ReplaceAll(str, "\\t", "\t")
	return str
}

func XML_Helper_FindUnescapeJsonValues(data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			data[key] = XML_Helper_FindUnescapeJsonValues(v)
		case []interface{}:
			for i, item := range v {
				if str, isString := item.(string); isString {
					v[i] = XML_Helper_EraseJsonUnescapedString(str)
				} else if nestedMap, isMap := item.(map[string]interface{}); isMap {
					v[i] = XML_Helper_FindUnescapeJsonValues(nestedMap)
				}
			}
		case string:
			data[key] = XML_Helper_EraseJsonUnescapedString(v)
		}
	}
	return data
}
