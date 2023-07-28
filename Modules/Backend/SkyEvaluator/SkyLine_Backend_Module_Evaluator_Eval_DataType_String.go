///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_DataType_String
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions to parse or evaluatte string data types
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

import SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"

func SkyLine_Evaluator_Eval_String_Index_Expression(String, Index SkyEnv.SL_Object) SkyEnv.SL_Object {
	Str := String.(*SkyEnv.SL_String).Value
	Idx := Index.(*SkyEnv.SL_Integer).Value

	MaxIndexAllowed := int64(len(Str))

	if Idx < 0 || int64(Idx) > MaxIndexAllowed {
		return SkyLine_Null_ALLIAS
	}

	chars := []rune(Str)

	Return := chars[Idx]

	return &SkyEnv.SL_String{Value: string(Return)}
}
