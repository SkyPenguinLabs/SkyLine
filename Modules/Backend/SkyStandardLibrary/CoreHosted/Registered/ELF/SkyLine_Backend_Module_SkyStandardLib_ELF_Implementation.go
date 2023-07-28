package SkyLine_ELF_Parsing

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	SkyHelpers "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

func InitateNewSession(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyHelpers.SkyLine_Standard_Library_Helper_CheckArguments("elf.New", args, SkyHelpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT), SkyHelpers.SkyLine_Standard_Library_Helper_ExactArguments(1)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	arg := args[0].(*SkyEnv.SL_String).Value
	if arg != "" {
		ELF_Session_Initation(arg)
	} else {
		return &SkyEnv.SL_Error{Message: "Argument in call to `elf.New` can not be empty"}
	}
	return &SkyEnv.SL_NULL{}
}

func CallToReturnHeaderInfo(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	LoadElfHeader(ELF_CTX.Header)
	Hash := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
	for k, v := range ReturnStorage.HeaderParsed {
		key := &SkyEnv.SL_String{Value: k}
		val := &SkyEnv.SL_String{Value: v}
		Hash[key.SL_HashKeyType()] = SkyEnv.HashPair{Key: key, Value: val}
	}
	return &SkyEnv.SL_HashMap{Pairs: Hash}
}
