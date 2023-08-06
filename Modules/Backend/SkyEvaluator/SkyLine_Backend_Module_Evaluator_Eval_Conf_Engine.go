package SkyLine_Backend_Evaluation

import (
	SLC "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyConfEngine/EngineCore"
	SLE "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

type EngineCallValues struct {
	Name        string
	Version     string
	Require     []string
	Languages   string
	Description string
	SOS         string
	Prepped     bool // Engine has parsed data
}

var EngineCallVal EngineCallValues

func Assign() {
	EngineCallVal.Name = SLC.Exportable_data.ProjectData.Name
	EngineCallVal.SOS = SLC.Exportable_data.ProjectData.SuportedOS
	EngineCallVal.Description = SLC.Exportable_data.ProjectData.Description
	EngineCallVal.Languages = SLC.Exportable_data.ProjectData.Languages
	EngineCallVal.Require = SLC.Exportable_data.ProjectData.Require
	EngineCallVal.Version = SLC.Exportable_data.ProjectData.Version
	EngineCallVal.Prepped = true // Indicating this function was run
}

func EvalEngineCall(val SLE.SL_Object) SLE.SL_Object {
	switch Value := val.(type) {
	case *SLE.SL_String:
		SLC.StartEngine_RenderFile(Value.SkyLine_ObjectFunction_GetTrueValue(), false)
		Assign() // Call to assign
	default:
		return SkyLine_Evaluator_CreateError("Sorry but the data type provided %s is not `String` as the ENGINE keyword requires a string value which is the filename SLC can parse", Value)
	}
	return &SLE.SL_NULL{}
}
