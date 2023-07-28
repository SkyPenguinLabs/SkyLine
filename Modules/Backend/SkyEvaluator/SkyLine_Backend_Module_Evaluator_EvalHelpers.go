///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_EvalHelpers
// Extension         | .go ( golang source code file )
// Purpose           | Defines all of the helper functions for the evaluator
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
	"context"
	"fmt"
	"log"
	"strconv"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_ObjectToNativeBoolean(Object SkyEnv.SL_Object) bool {
	if Ret, conv := Object.(*SkyEnv.SL_Return); conv {
		Object = Ret.Value
	}
	switch SL_Object := Object.(type) {
	case *SkyEnv.SL_Boolean:
		return SL_Object.Value
	case *SkyEnv.SL_String:
		return SL_Object.Value != ""
	case *SkyEnv.SL_NULL:
		return false
	case *SkyEnv.SL_Integer:
		if SL_Object.Value == 0 {
			return false
		}
		return true
	case *SkyEnv.SL_Float:
		if SL_Object.Value == 0.0 {
			return false
		}
		return true
	case *SkyEnv.SL_Array:
		if len(SL_Object.Elements) == 0 {
			return false
		}
		return true
	case *SkyEnv.SL_HashMap:
		if len(SL_Object.Pairs) == 0 {
			return false
		}
		return true
	default:
		return true
	}
}

func SkyLine_Eval_Apply_Function_To_Evaluation(Function SkyEnv.SL_Object, Env *SkyEnv.SkyLineEnvironment, Arguments []SkyEnv.SL_Object) SkyEnv.SL_Object {
	switch FuncT := Function.(type) {
	case *SkyEnv.SL_Functions:
		Extend := SkyLine_Eval_ExtendFunctionEnvironment(FuncT, Arguments)
		Evaled := SkyLine_Call_Eval(FuncT.Unit, Extend)
		return SkyLine_Eval_Unwrap_Returns(Evaled)
	case *SkyEnv.SL_Builtin:
		return FuncT.Function(Env, Arguments...)
	default:
		return SkyLine_Evaluator_CreateError("Body is not a function but is %s ", Function.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Eval_Unwrap_Returns(Object SkyEnv.SL_Object) SkyEnv.SL_Object {
	if ReturnValue, opk := Object.(*SkyEnv.SL_Return); opk {
		return ReturnValue.Value
	}
	return Object
}

func SkyLine_Eval_ExtendFunctionEnvironment(Function *SkyEnv.SL_Functions, Arguments []SkyEnv.SL_Object) *SkyEnv.SkyLineEnvironment {
	Environment := SkyEnv.SL_NewEnclosedEnvironment(Function.Env)

	for Key, Value := range Function.Defaults {
		Environment.Load(Key, SkyLine_Call_Eval(Value, Environment))
	}
	for ParamIDX, Param := range Function.Sky_Function_Arguments {
		if ParamIDX < len(Arguments) {
			Environment.Load(Param.Value, Arguments[ParamIDX])
		}
	}
	return Environment
}

func SkyLine_Eval_FromNativeBoolean_To_BooleanObject(Data bool) *SkyEnv.SL_Boolean {
	if Data {
		return SkyLine_True_ALLIAS
	}
	return SkyLine_False_ALLIAS
}

func SkyLine_Eval_Set_Context(SkyCTX context.Context) {
	SkyLine_Node_Context = SkyCTX
}

func RegisterVariable(name string, object SkyEnv.SL_Object) {
	switch object.(type) {
	case *SkyEnv.SL_String:
		BuiltInVariables_String[name] = &SkyEnv.SL_String{Value: object.SkyLine_ObjectFunction_GetTrueValue()}
	case *SkyEnv.SL_Float:
		conv, x := strconv.ParseFloat(object.SkyLine_ObjectFunction_GetTrueValue(), 64)
		if x != nil {
			log.Fatalf("SKYLINE DEVELOPER ERROR (Register variable): Could not register variable %s due to %s", name, fmt.Sprint(x))
		}
		BuiltInVariables_Float[name] = &SkyEnv.SL_Float{Value: conv}
	case *SkyEnv.SL_Integer:
		conv, x := strconv.ParseInt(object.SkyLine_ObjectFunction_GetTrueValue(), 10, 64)
		if x != nil {
			log.Fatalf("SKYLINE DEVELOPER ERROR (Register variable): Could not register variable %s due to %s", name, fmt.Sprint(x))
		}
		BuiltInVariables_Integer[name] = &SkyEnv.SL_Integer{Value: int(conv)}
	case *SkyEnv.SL_Boolean:
		conv, x := strconv.ParseBool(object.SkyLine_ObjectFunction_GetTrueValue())
		if x != nil {
			log.Fatalf("SKYLINE DEVELOPER ERROR (Register variable): Could not register variable %s due to %s", name, fmt.Sprint(x))
		}
		BuiltInVariables_Boolean[name] = &SkyEnv.SL_Boolean{Value: conv}
	}
}

func SkyLine_Register_Builtin(name string, Func SkyEnv.SkyLine_BuiltinFunction) {
	BuiltIn[name] = &SkyEnv.SL_Builtin{
		Function: Func,
	}
}

func SkyLine_Evaluator_CheckError(Object SkyEnv.SL_Object) bool {
	if Object != nil {
		return Object.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_ERROR_OBJECT
	}
	return false
}

func SkyLine_Evaluator_CreateError(Form string, Args ...interface{}) *SkyEnv.SL_Error {
	return &SkyEnv.SL_Error{
		Message: fmt.Sprintf(Form, Args...),
	}
}

func SkyLine_Evaluator_IsTruthy(Object SkyEnv.SL_Object) bool {
	switch Object {
	case SkyLine_Null_ALLIAS:
		return false
	case SkyLine_True_ALLIAS:
		return true
	case SkyLine_False_ALLIAS:
		return false
	default:
		return true
	}
}
