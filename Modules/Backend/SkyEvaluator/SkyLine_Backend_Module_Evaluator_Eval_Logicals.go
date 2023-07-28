///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_Logicals
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions for logical statements or expressions such as for loops, while loops, if else, switch case etc
// Directory         | Modules/Backend/SkyEvaluator
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEvaluator
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
	SkyAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_Eval_Conditional_ForLoop(ForLoop *SkyAST.SL_EN_Conditional_Loop, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	Ret := &SkyEnv.SL_Boolean{
		Value: true,
	}
	for {
		Condition := SkyLine_Call_Eval(ForLoop.Condition, Env)

		if SkyLine_Evaluator_CheckError(Condition) {
			return Condition
		}

		if SkyLine_Evaluator_IsTruthy(Condition) {
			Ret2 := SkyLine_Call_Eval(ForLoop.Consequence, Env)
			if !SkyLine_Evaluator_CheckError(Ret2) && (Ret2.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_RETURN_OBJECT || Ret2.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_ERROR_OBJECT) {
				return Ret2
			}
		} else {
			break
		}
	}
	return Ret
}

func SkyLine_Evaluator_Eval_Conditional_Switch(SwitchStatement *SkyAST.SL_EN_Switch_ExpressionStatement, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	Object := SkyLine_Call_Eval(SwitchStatement.Value, Env)

	for _, Condition := range SwitchStatement.Conditions {
		if Condition.Default {
			continue
		}

		for _, value := range Condition.Expression {
			Output := SkyLine_Call_Eval(value, Env)

			if Object.SkyLine_ObjectFunction_GetDataType() == Output.SkyLine_ObjectFunction_GetDataType() && (Object.SkyLine_ObjectFunction_GetTrueValue() == Output.SkyLine_ObjectFunction_GetTrueValue()) {
				BlockOut := SkyLine_Evaluator_Eval_UnitAndBlock_Statement(Condition.Unit, Env)
				return BlockOut
			}
		}
	}
	for _, Condition := range SwitchStatement.Conditions {
		if Condition.Default {
			Out := SkyLine_Evaluator_Eval_UnitAndBlock_Statement(Condition.Unit, Env)
			return Out
		}
	}
	return nil
}

func SkyLine_Evaluator_Eval_Conditional_Statement(Conditional *SkyAST.SL_EN_Conditional_IfElse, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	Condition := SkyLine_Call_Eval(Conditional.Condition, Env)
	if SkyLine_Evaluator_CheckError(Condition) {
		return Condition
	}
	if SkyLine_Evaluator_IsTruthy(Condition) {
		SkyLine_Call_Eval(Conditional.Consequence_Unit, Env)
	} else if Conditional.Alternative_Unit != nil {
		return SkyLine_Call_Eval(Conditional.Alternative_Unit, Env)
	}
	return SkyLine_Null_ALLIAS
}
