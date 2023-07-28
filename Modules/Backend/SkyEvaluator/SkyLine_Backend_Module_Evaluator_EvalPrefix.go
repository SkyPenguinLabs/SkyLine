///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_EvalPrefix
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions for prefix expression evaluation
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

import SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"

func SkyLine_Evaluator_Eval_Prefix_BANG_Operator_ExpressionBoolean(RightValue SkyEnv.SL_Object) SkyEnv.SL_Object {
	switch RightValue {
	case SkyLine_True_ALLIAS:
		return SkyLine_False_ALLIAS
	case SkyLine_False_ALLIAS:
		return SkyLine_True_ALLIAS
	case SkyLine_Null_ALLIAS:
		return SkyLine_True_ALLIAS
	default:
		return SkyLine_False_ALLIAS
	}
}

func SkyLine_Evaluator_Eval_Prefix_Minux_Operator_Expression(Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	switch ObjectT := Right.(type) {
	case *SkyEnv.SL_Integer:
		return &SkyEnv.SL_Integer{
			Value: -ObjectT.Value,
		}
	case *SkyEnv.SL_Float:
		return &SkyEnv.SL_Float{
			Value: -ObjectT.Value,
		}
	default:
		return SkyLine_Evaluator_CreateError("Unknown error for prefix minus operator expression: -%s", Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_Prefix_Expression(Op string, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	switch Op {
	case "!":
		return SkyLine_Evaluator_Eval_Prefix_BANG_Operator_ExpressionBoolean(Right)
	case "-":
		return SkyLine_Evaluator_Eval_Prefix_Minux_Operator_Expression(Right)
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for prefix expression, supported(!, -) but got %s", Right.SkyLine_ObjectFunction_GetDataType())
	}
}
