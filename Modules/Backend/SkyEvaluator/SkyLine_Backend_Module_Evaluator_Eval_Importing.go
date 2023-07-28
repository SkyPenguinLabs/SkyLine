///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_Importing
// Extension         | .go ( golang source code file )
// Purpose           | Defines all of the import and evaluation functions for modules and imports
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
	SkyParser "SkyLine/Modules/Backend/SkyParser"
	SkyScanner "SkyLine/Modules/Backend/SkyScanner"
	"fmt"
	"io/ioutil"
	"os"
)

func Eval_Final_Module(pathname string) SkyEnv.SL_Object {
	byter, x := ioutil.ReadFile(pathname)
	if x != nil {
		return SkyLine_Evaluator_CreateError("IO Error: Could not read module '%s' -> %s", pathname, x)
	}
	lex := SkyScanner.New(string(byter))
	parser := SkyParser.SkyLineNewParser(lex)
	mdod := parser.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
	if len(parser.SL_Parser_Errors) != 0 {
		return SkyLine_Evaluator_CreateError("Parser error: %s", parser.SL_Parser_Errors)
	}
	env := SkyEnv.SL_NewEnvironment()
	SkyLine_Call_Eval(mdod, env)
	return env.ExportedHash()
}

func EvalModule(pathname string) SkyEnv.SL_Object {
	_, x := os.Stat(pathname)
	if x != nil {
		return SkyLine_Evaluator_CreateError("SkyLine Statistics: Could not get the file name to be parsed -> %s", fmt.Sprint(x))
	}
	return Eval_Final_Module(pathname)
}

func evalModuleIndexExpression(module, index SkyEnv.SL_Object) SkyEnv.SL_Object {
	moduleObject := module.(*SkyEnv.SL_Module)
	return SkyLine_Evaluator_Eval_Hash_Map_IndexExpression(moduleObject.SL_Attributes, index)
}
