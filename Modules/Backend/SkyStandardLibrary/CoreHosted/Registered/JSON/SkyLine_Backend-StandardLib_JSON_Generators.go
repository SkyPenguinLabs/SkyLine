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
// This file defines all of the generation or field generation functions for the language
//
package SkyLine_Standard_Library_JSON

import (
	"bytes"
	"fmt"
	"strings"
)

func GenerateGolangStructure_JSON(data interface{}, StructureName string) (strux string) {
	var buffer bytes.Buffer
	GenerateStructureFromJSON(&buffer, data, StructureName)
	return buffer.String()
}

func GenerateStructureFromJSON(buffer *bytes.Buffer, data interface{}, StructureName string) {
	switch value := data.(type) {
	case map[string]interface{}:
		GenerateStructureFromMapJSON(buffer, value, StructureName)
	case []interface{}:
		GenerateStructureFromArrayJSON(buffer, value, StructureName)
	}
}

func GenerateStructureFromMapJSON(buffer *bytes.Buffer, value map[string]interface{}, StructureName string) {
	buffer.WriteString(
		fmt.Sprintf(
			"type %s struct {\n", StructureName,
		),
	)
	for key, value := range value {
		FieldType, FieldName := DetermineFieldTypeNameJSON(value, key)
		UtilitiesGenerateField(buffer, FieldType, FieldName)
	}
	buffer.WriteString("}\n\n")
	for key, value := range value {
		if subMap, ok := value.(map[string]interface{}); ok {
			SubStructName := UtilitiesCaptialize(key)
			GenerateStructureFromJSON(buffer, subMap, SubStructName)
		}
	}
}

func GenerateStructureFromArrayJSON(buffer *bytes.Buffer, data []interface{}, structName string) {
	FieldType := fmt.Sprintf("[]%s", structName)
	FieldName := "Items"
	UtilitiesGenerateField(buffer, FieldType, FieldName)
	for _, element := range data {
		if SubsMap, ok := element.(map[string]interface{}); ok {
			SubstructureName := UtilitiesCaptialize(structName)
			GenerateStructureFromJSON(buffer, SubsMap, SubstructureName)
		}
	}
}
func DetermineFieldTypeNameJSON(value interface{}, Key string) (FieldType, FieldName string) {
	switch value.(type) {
	case string:
		FieldType = "string"
	case float64:
		FieldType = "float64"
	case bool:
		FieldType = "bool"
	case map[string]interface{}:
		FieldType = UtilitiesCaptialize(Key)
	case []interface{}:
		FieldType = fmt.Sprintf("[]%s", UtilitiesCaptialize(Key))
	default:
		FieldType = "interface{}"
	}
	FieldName = UtilitiesCaptialize(Key)
	return FieldType, FieldName
}

func UtilitiesGenerateField(buffer *bytes.Buffer, FieldType, FieldName string) {
	buffer.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", FieldName, FieldType, FieldName))
}

func UtilitiesCaptialize(value string) string {
	return strings.ToUpper(string(value[0])) + value[1:]
}
