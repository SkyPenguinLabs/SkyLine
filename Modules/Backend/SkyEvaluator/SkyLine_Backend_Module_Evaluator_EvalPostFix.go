///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_EvalPostFix
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions for postfix parsing
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

func SkyLine_Evaluator_Eval_PostFix_PFExpression(Op string, Env *SkyEnv.SkyLineEnvironment, Node *SkyAST.SL_EN_Postfix) SkyEnv.SL_Object {
	switch Op {
	case "++":
		val, ok := Env.Get(Node.TokenConstruct.Literal)
		if !ok {
			return SkyLine_Evaluator_CreateError("%s is unknown", Node.TokenConstruct.Literal)
		}

		switch arg := val.(type) {
		case *SkyEnv.SL_Integer:
			v := arg.Value
			Env.Load(Node.TokenConstruct.Literal, &SkyEnv.SL_Integer{Value: v + 1})
			return arg
		default:
			return SkyLine_Evaluator_CreateError("%s is not an Integer", Node.TokenConstruct.Literal)
		}
	case "--":
		val, ok := Env.Get(Node.TokenConstruct.Literal)
		if !ok {
			return SkyLine_Evaluator_CreateError("%s is unknown", Node.TokenConstruct.Literal)
		}

		switch arg := val.(type) {
		case *SkyEnv.SL_Array:
			v := arg.Elements
			if len(v) > 1 {
				Env.Load(Node.TokenConstruct.Literal, &SkyEnv.SL_Array{Elements: v[:len(v)-1]})
				return arg
			} else {
				return arg
			}
		case *SkyEnv.SL_String:
			v := arg.Value
			if len(v) > 1 {
				Env.Load(Node.TokenConstruct.Literal, &SkyEnv.SL_String{Value: v[:len(v)-1]})
				return arg
			} else {
				return arg
			}
		case *SkyEnv.SL_Boolean:
			v := arg.Value
			if v {
				Env.Load(Node.TokenConstruct.Literal, &SkyEnv.SL_Boolean{Value: false})
			} else {
				Env.Load(Node.TokenConstruct.Literal, &SkyEnv.SL_Boolean{Value: true})
			}
			return arg
		case *SkyEnv.SL_Integer:
			v := arg.Value
			Env.Load(Node.TokenConstruct.Literal, &SkyEnv.SL_Integer{Value: v - 1})
			return arg
		default:
			return SkyLine_Evaluator_CreateError("%s is not an integer, string, boolean or array", Node.TokenConstruct.Literal)
		}

	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for a postfix expression? Requires either '++' or '--' but you gave %s", Op)
	}
}
