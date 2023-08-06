///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_Cause_Prog
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions for evaluating the statements inside of a program
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

func SkyLine_Evaluator_Eval_Program(Program *SkyAST.SL_Prog, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	var Results SkyEnv.SL_Object

	for _, statement := range Program.ProgramStatements {
		Results = SkyLine_Call_Eval(statement, Env)

		switch Results := Results.(type) {
		case *SkyEnv.SL_Return:
			return Results.Value
		case *SkyEnv.SL_Error:
			return Results
		}
	}
	return Results
}
