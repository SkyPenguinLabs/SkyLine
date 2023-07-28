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
// This file defines grabbers and generation functions. These functions help for dumping XML into specific formats such as Arrays, Maps, converting it to json and much more
//
// other various functions. This can help when working with larger sets of files and instead of parsing everything manually and using for loops, we can dump it into a map.
//
package SkyLine_Standard_Library_XML

// Dump and convert values into a well required hash map
func XML_Lib_DumpIntoHashMap(node XML_Node) map[string]interface{} {
	maped := make(map[string]interface{})
	if len(node.Nodes) == 0 {
		maped[node.XMLName.Local] = node.Text
	} else {
		cmap := make(map[string]interface{})
		for _, cn := range node.Nodes {
			cmap[cn.XMLName.Local] = XML_Lib_DumpIntoHashMap(cn)
		}
		maped[node.XMLName.Local] = cmap
	}
	return maped
}
