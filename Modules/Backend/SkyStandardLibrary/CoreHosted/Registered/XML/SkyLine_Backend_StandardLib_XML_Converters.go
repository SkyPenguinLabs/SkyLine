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
// This file defines all of the code converters which can convert XML code to something like JSON or PLIST
//
package SkyLine_Standard_Library_XML

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	"encoding/json"
	"fmt"
)

func ConstructPrimeXMLTree(data interface{}) *XML_NodeConverterJsonToXML {
	node := &XML_NodeConverterJsonToXML{}
	switch value := data.(type) {
	case map[string]interface{}:
		for key, val := range value {
			childNode := ConstructPrimeXMLTree(val)
			childNode.XMLName.Local = key
			node.Nodes = append(node.Nodes, childNode)
		}
	case []interface{}:
		for _, val := range value {
			childNode := ConstructPrimeXMLTree(val)
			node.Nodes = append(node.Nodes, childNode)
		}
	default:
		node.Text = fmt.Sprintf("%v", value)
	}
	return node
}

func ConvertPrimeXMLNodeToJsonNode(node XML_NodeConverterXMLtoJson) ([]byte, error) {
	if len(node.Nodes) == 0 {
		return json.Marshal(node.Text)
	}

	jsonMap := make(map[string]interface{})
	if len(node.Attrs) > 0 {
		attrMap := make(map[string]string)
		for _, attr := range node.Attrs {
			attrMap[attr.Name.Local] = attr.Value
		}
		jsonMap["@attributes"] = attrMap
	}

	for _, childNode := range node.Nodes {
		childName := childNode.XMLName.Local
		childData, err := ConvertPrimeXMLNodeToJsonNode(childNode)
		if err != nil {
			return nil, err
		}

		if existingData, ok := jsonMap[childName]; ok {
			if existingSlice, isSlice := existingData.([]interface{}); isSlice {
				jsonMap[childName] = append(existingSlice, string(childData))
			} else {
				jsonMap[childName] = []interface{}{existingData, string(childData)}
			}
		} else {
			jsonMap[childName] = string(childData)
		}
	}

	return json.Marshal(XML_Helper_FindUnescapeJsonValues(jsonMap))
}

func ConvertValues(value interface{}) SkyEnv.SL_Object {
	switch v := value.(type) {
	case map[string]interface{}:
		convertedMap := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
		for k, v := range v {
			key := &SkyEnv.SL_String{Value: k}
			value := ConvertValues(v)
			hashPair := SkyEnv.HashPair{Key: key, Value: value}
			convertedMap[key.SL_HashKeyType()] = hashPair
		}
		return &SkyEnv.SL_HashMap{Pairs: convertedMap}
	case string:
		return &SkyEnv.SL_String{Value: v}
	case int:
		return &SkyEnv.SL_Integer{Value: v}
	case bool:
		return &SkyEnv.SL_Boolean{Value: v}
	case float32, float64:
		return &SkyEnv.SL_Float{Value: float64(v.(*SkyEnv.SL_Float).Value)}
	default:
		return nil
	}
}

func ConvertXMLDataToHashMap(data map[string]interface{}) *SkyEnv.SL_HashMap {
	resultingHash := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
	for k, v := range data {
		key := &SkyEnv.SL_String{Value: k}
		value := ConvertValues(v)
		hashPair := SkyEnv.HashPair{Key: key, Value: value}
		resultingHash[key.SL_HashKeyType()] = hashPair
	}
	return &SkyEnv.SL_HashMap{Pairs: resultingHash}
}
