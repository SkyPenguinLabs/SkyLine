///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Evaluation
// Extension         | .go ( golang source code file )
// Purpose           | Defines all other sub and primary eval functions
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
	"fmt"
	"os"
	"strings"

	SkyAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_Eval_Index_Expression(Left, Index SkyEnv.SL_Object) SkyEnv.SL_Object {
	switch {
	case Left.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT &&
		Index.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT ||
		Index.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER8_OBJECT ||
		Index.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER16_OBJECT ||
		Index.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER32_OBJECT ||
		Index.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER64_OBJECT:

		return SkyLine_Evaluator_Eval_Array_Index_Expression(Left, Index)
	case Left.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_HASH_OBJECT:
		return SkyLine_Evaluator_Eval_Hash_Map_IndexExpression(Left, Index)
	case Left.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT:
		return SkyLine_Evaluator_Eval_String_Index_Expression(Left, Index)
	case Left.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_MODULE_OBJECT:
		return evalModuleIndexExpression(Left, Index)
	default:
		return SkyLine_Evaluator_CreateError("Index operator not supported: %s", Left.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_Expression(Expressions []SkyAST.SL_Expression, Env *SkyEnv.SkyLineEnvironment) []SkyEnv.SL_Object {
	var result []SkyEnv.SL_Object

	for _, expression := range Expressions {
		Evaled := SkyLine_Call_Eval(expression, Env)

		if SkyLine_Evaluator_CheckError(Evaled) {
			return []SkyEnv.SL_Object{Evaled}
		}
		result = append(result, Evaled)
	}
	return result
}

func SkyLine_Evaluator_Eval_Identifier(Node *SkyAST.SL_Identifier, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	if Value, ok := Env.Get(Node.Value); ok {
		return Value
	}
	if Builtin, ok := BuiltIn[Node.Value]; ok {
		return Builtin
	} else {
		if builtin_var, ok := BuiltInVariables_String[Node.Value]; ok {
			return builtin_var
		} else if builtin_var_i, ok := BuiltInVariables_Integer[Node.Value]; ok {
			return builtin_var_i
		} else if builtin_var_b, ok := BuiltInVariables_Boolean[Node.Value]; ok {
			return builtin_var_b
		} else if builtin_var_f, ok := BuiltInVariables_Float[Node.Value]; ok {
			return builtin_var_f
		} else if builtin_var_arr, ok := BuiltInVariables_Array[Node.Value]; ok {
			return builtin_var_arr
		}
	}
	fmt.Fprintf(os.Stderr, "Identifier not found %s \n", Node.Value)
	return SkyLine_Evaluator_CreateError("identifier not found " + Node.Value)
}

func SKyLine_Evaluator_Eval_ObjectCallingExpression(Called *SkyAST.SL_EN_Object_Call_Expression, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	Object := SkyLine_Call_Eval(Called.Object, Env)
	if call, ok := Called.Call.(*SkyAST.SL_EN_Call_Expression); ok {
		arguments := SkyLine_Evaluator_Eval_Expression(Called.Call.(*SkyAST.SL_EN_Call_Expression).Arguments, Env)
		ret := Object.SkyLine_ObjectFunction_InvokeObject(call.Function.SkyLine_NodeInterface_Token_Literal(), *Env, arguments...)
		if ret != nil {
			return ret
		}
		attempts := []string{}
		attempts = append(attempts, strings.ToLower(string(Object.SkyLine_ObjectFunction_GetDataType())))
		attempts = append(attempts, "object")
		for _, prefix := range attempts {
			name := prefix + "." + call.Function.SkyLine_NodeInterface_Get_Node_Value()
			if function, ok := Env.Get(name); ok {
				Extended := SkyLine_Eval_ExtendFunctionEnvironment(function.(*SkyEnv.SL_Functions), arguments)
				Extended.Load("self", Object)
				Evaled := SkyLine_Call_Eval(function.(*SkyEnv.SL_Functions).Unit, Extended)
				Object = SkyLine_Eval_Unwrap_Returns(Evaled)
				return Object
			}
		}
	}
	return SkyLine_Evaluator_CreateError("Failed to invoke method: %s ", Called.Call.(*SkyAST.SL_EN_Call_Expression).Function.SkyLine_NodeInterface_Get_Node_Value())
}
