///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_DataType_Array
// Extension         | .go ( golang source code file )
// Purpose           | Defines all models for the array data type parsing
// Directory         | Modules/Backend/SkyEvaluator
// Modular Directory | SkyLine/Modules/Backend/SkyEvaluator
// Package Name      | SkyLine_Backend_Module_Evaluation
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// The final part in standard interpretation inside of a programming language is to evaluate and execute the data or keys. In the case of SkyLine, it relies on different
//
// forms of engines which can use the byte code compiler or use the evaluator. The byte code compiler is a whole different story but the evaluator will take advantage of the
//
// AST and then check and execute conditions, statements, values, or modifications accordingly. The evaluator can also sometimes be complex to use but it still manages to stay
//
// one of the fastest ones to write.
//
package SkyLine_Backend_Evaluation

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_Eval_Array_Index_Expression(Array, Index SkyEnv.SL_Object) SkyEnv.SL_Object {
	ArrayObject := Array.(*SkyEnv.SL_Array)
	var Idx int
	switch v := Index.(type) {
	case *SkyEnv.SL_Integer:
		Idx = v.Value
	case *SkyEnv.SL_Integer32:
		Idx = int(v.Value)
	case *SkyEnv.SL_Integer64:
		Idx = int(v.Value)
	case *SkyEnv.SL_Integer8:
		Idx = int(v.Value)
	}
	MaxIndexAllowed := int64(len(ArrayObject.Elements) - 1)

	if Idx < 0 || Idx > int(MaxIndexAllowed) {
		return SkyLine_Null_ALLIAS
	}
	return ArrayObject.Elements[Idx]

}
