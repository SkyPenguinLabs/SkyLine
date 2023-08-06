package SkyLine_Network

import (
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkySTDLibHelp "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

func SkyLine_Network_Module_ImplementInterfaces(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	// now check and return interfaces
	if inters, x := RetrieveInterfaceNames(); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		RetArr := make([]SkyEnv.SL_Object, 0)
		for i := 0; i < len(inters); i++ {
			RetArr = append(RetArr, &SkyEnv.SL_String{Value: inters[i]})
		}
		return &SkyEnv.SL_Array{Elements: RetArr}
	}
}

func SkyLine_Network_Module_Implement_NetAddrByName(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkySTDLibHelp.SkyLine_Standard_Library_Helper_CheckArguments(
		"network.NetByName",
		arguments,
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ExactArguments(1),
		SkySTDLibHelp.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	// Now get the interfaces
	interfacename := arguments[0].(*SkyEnv.SL_String).Value
	network, x := GetInterfaceIPByName(interfacename)
	if x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	} else {
		return &SkyEnv.SL_String{Value: network}
	}
}
