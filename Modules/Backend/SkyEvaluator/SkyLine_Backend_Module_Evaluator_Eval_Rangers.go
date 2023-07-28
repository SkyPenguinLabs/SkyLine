///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_Rangers
// Extension         | .go ( golang source code file )
// Purpose           | Defines all range based evaluation functions which means it evaluates for each or for x in range loops
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
	SkyAST "SkyLine/Modules/Backend/SkyAST"
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_Eval_ForEach_Range_Loop(Ranger *SkyAST.SL_EN_For_Each_Loop, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	Value := SkyLine_Call_Eval(Ranger.Value, Env)
	helper, ok := Value.(SkyEnv.SL_Iterable)
	if !ok {
		return SkyLine_Evaluator_CreateError("%s object does not implement the iterabvle interface", Value.SkyLine_ObjectFunction_GetDataType())
	}
	var permission []string
	permission = append(permission, Ranger.Identifier)
	if Ranger.Index != "" {
		permission = append(permission, Ranger.Index)
	}
	child := SkyEnv.SL_NewEnvironmentTemporaryScope(Env, permission)
	helper.SkyLine_Reset_Offset()
	ret, idx, ok := helper.SkyLine_Next_Itteration()
	for ok {
		child.Load(Ranger.Identifier, ret)
		idxName := Ranger.Index
		if idxName != "" {
			child.Load(Ranger.Index, idx)
		}
		rt := SkyLine_Call_Eval(Ranger.Unit, child)
		if !SkyLine_Evaluator_CheckError(rt) && (rt.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_RETURN_OBJECT || rt.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_ERROR_OBJECT) {
			return rt
		}
		ret, idx, ok = helper.SkyLine_Next_Itteration()
	}
	return &SkyEnv.SL_NULL{}
}
