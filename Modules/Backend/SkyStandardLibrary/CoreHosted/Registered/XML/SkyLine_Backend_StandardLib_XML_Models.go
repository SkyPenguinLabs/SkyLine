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
// This file holds all models for the XML library or package that can be imported into SkyLine. This includes node and hash xml relations
//
package SkyLine_Standard_Library_XML

import "encoding/xml"

type (
	//:::::::::::::::::::::::::::::::::::::::::::::::
	//:: Defines a base node for XML visualization ::
	//:::::::::::::::::::::::::::::::::::::::::::::::
	Node struct {
		Name  xml.Name
		Text  string `xml:",chardata"`
		Nodes []Node `xml:",any"`
	}

	//:::::::::::::::::::::::::::::::::::::::::::
	//:: Defines a base node for XML dumping   ::
	//:::::::::::::::::::::::::::::::::::::::::::
	XML_Node struct {
		XMLName xml.Name
		Attrs   []xml.Attr `xml:",any,attr"`
		Nodes   []XML_Node `xml:",any"`
		Text    string     `xml:",chardata"`
	}

	//:::::::::::::::::::::::::::::::::::::::::::::::
	//:: Defines a base node for XML Conversions   ::
	//:::::::::::::::::::::::::::::::::::::::::::::::

	// This is for the function that will take a JSON file and convert it to JSON
	XML_NodeConverterJsonToXML struct {
		XMLName xml.Name
		Attrs   []xml.Attr                    `xml:",any,attr"`
		Nodes   []*XML_NodeConverterJsonToXML `xml:",any"`
		Text    string                        `xml:",chardata"`
	}

	// This is for the function thart will take an XML file and convert it to JSON
	XML_NodeConverterXMLtoJson struct {
		XMLName xml.Name
		Text    string `xml:",chardata"`
		Attrs   []xml.Attr
		Nodes   []XML_NodeConverterXMLtoJson `xml:",any"`
	}
)
