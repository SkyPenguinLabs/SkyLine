package SkyLine_ELF_Parsing

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	SkyEval "SkyLine/Modules/Backend/SkyEvaluator"
)

func InitateELFLibs() {
	SkyEval.SkyLine_Register_Builtin("NewElf", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (InitateNewSession(InvokeArgs...))
	})
	SkyEval.SkyLine_Register_Builtin("ElfHead", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (CallToReturnHeaderInfo(InvokeArgs...))
	})
}
