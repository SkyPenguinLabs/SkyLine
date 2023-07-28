///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_EvaluationModels
// Extension         | .go ( golang source code file )
// Purpose           | Defines all models for the Evaluator
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
	"context"
)

var (
	SkyLine_Null_ALLIAS  = &SkyEnv.SL_NULL{}
	SkyLine_True_ALLIAS  = &SkyEnv.SL_Boolean{Value: true}
	SkyLine_False_ALLIAS = &SkyEnv.SL_Boolean{Value: false}
	SkyLine_Node_Context = context.Background()

	// Built in functions
	BuiltIn                  = map[string]*SkyEnv.SL_Builtin{}
	BuiltInVariables_String  = map[string]*SkyEnv.SL_String{}
	BuiltInVariables_Float   = map[string]*SkyEnv.SL_Float{}
	BuiltInVariables_Integer = map[string]*SkyEnv.SL_Integer{}
	BuiltInVariables_Boolean = map[string]*SkyEnv.SL_Boolean{}

	// Standard library names that exist
	StandardLibNames = map[string]bool{
		"math":                true,
		"io":                  true,
		"File":                true,
		"http":                true,
		"Google/Smart_API_DB": true,
		"Amazon/Smart_API_DB": true,
		"Roku/Smart_API_DB":   true,
		"Apple/Smart_API_DB":  true,
		"Other/Smart_API_DB":  true,
		"xml":                 true,
		"json":                true,
		"env":                 true,
		"console":             true,
		"forensics/image":     true,
	}

	// Standard library to functions | register(...)
	RegisterStandard = map[string]func(){
		"io":                  RegisterIO,
		"math":                RegisterMath,
		"File":                RegisterFile2,
		"http":                RegisterHTTP,
		"Google/Smart_API_DB": RegisterGoogleIoTDatabase,
		"Amazon/Smart_API_DB": RegisterAmazonIoTDatabase,
		"Roku/Smart_API_DB":   RegisterRokuIoTDatabase,
		"Apple/Smart_API_DB":  RegisterAppleIoTDatabase,
		"Other/Smart_API_DB":  RegisterOtherIoTDatabases_ARRIS,
		"xml":                 RegisterXML,
		"json":                RegisterJson,
		"env":                 RegisterEnv,
		"console":             RegisterConsole,
		"forensics/image":     RegisterForensicsImagePath,
	}
)

func SkyLine_Call_Eval(SkyNode SkyAST.SL_Node, SkyEnvironment *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	return SkyLineEvaluateContext(context.Background(), SkyNode, SkyEnvironment)
}
