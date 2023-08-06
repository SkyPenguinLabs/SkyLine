///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_EvalStatements
// Extension         | .go ( golang source code file )
// Purpose           | Defines all statement functions that need to be evaluated
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

	SkyAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_Eval_UnitAndBlock_Statement(Unit *SkyAST.SL_UnitBlockStatement, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	var ResultingUnit SkyEnv.SL_Object

	for _, statement := range Unit.Statements {
		ResultingUnit = SkyLine_Call_Eval(statement, Env)

		if ResultingUnit != nil {
			RT := ResultingUnit.SkyLine_ObjectFunction_GetDataType()
			if RT == SkyEnv.SKYLINE_DATATYPE_RETURN_OBJECT || RT == SkyEnv.SKYLINE_DATATYPE_ERROR_OBJECT {
				return ResultingUnit
			}
		}
	}
	return ResultingUnit
}

func CheckIfLibBasedVar(Name string) bool {
	if _, ok := BuiltInVariables_String[Name]; ok {
		return true
	} else if _, ok := BuiltInVariables_Integer[Name]; ok {
		return true
	} else if _, ok := BuiltInVariables_Boolean[Name]; ok {
		return true
	} else if _, ok := BuiltInVariables_Float[Name]; ok {
		return true
	}
	return false
}

func SkyLine_Evaluator_Eval_VariableAssignment_Statement(Assignee *SkyAST.SL_EN_VariableAssignmentStatement, Env *SkyEnv.SkyLineEnvironment) (V SkyEnv.SL_Object) {
	Evaled := SkyLine_Call_Eval(Assignee.Value, Env)
	if SkyLine_Evaluator_CheckError(Evaled) {
		return Evaled
	}
	if CheckIfLibBasedVar(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value()) {
		var msg string
		msg = "ERROR: You are trying to modify a fluid variable registered into the environment\n"
		msg += "\t | [#] Registered variables are not allowed to be modified by users \n"
		msg += "\t | [#] they are controlled by the compiler (SkyVM) and interpreter (SkyEval)\n"
		msg += "\t | [#] - Please ensure that you meant to modify ` " + Assignee.Name.Value + " ` "
		fmt.Println(msg)
		return SkyLine_Evaluator_CreateError(msg)
	}
	switch Assignee.Operator {
	case "+=":
		Current_Value, ok := Env.Get(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())

		if !ok {
			return SkyLine_Evaluator_CreateError("%s does not exist within the environment", Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		}

		result := SkyLine_Evaluator_Eval_InfixExpression("+=", Current_Value, Evaled, Env)

		if SkyLine_Evaluator_CheckError(result) {
			fmt.Printf("Error when handling +=: %s", result.SkyLine_ObjectFunction_GetTrueValue())
			return result
		}

		Env.Load(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value(), result)

		return result

	case "-=":
		Current_Value, ok := Env.Get(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		if !ok {
			return SkyLine_Evaluator_CreateError("%s does not exist within the environment", Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		}
		result := SkyLine_Evaluator_Eval_InfixExpression("-=", Current_Value, Evaled, Env)
		if SkyLine_Evaluator_CheckError(result) {
			fmt.Printf("Error when handling -=: %s", result.SkyLine_ObjectFunction_GetTrueValue())
			return result
		}
		Env.Load(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value(), result)
		return result
	case "*=":
		Current_Value, ok := Env.Get(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		if !ok {
			return SkyLine_Evaluator_CreateError("%s does not exist within the environment", Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		}
		result := SkyLine_Evaluator_Eval_InfixExpression("*=", Current_Value, Evaled, Env)
		if SkyLine_Evaluator_CheckError(result) {
			fmt.Printf("Error when handling *=: %s", result.SkyLine_ObjectFunction_GetTrueValue())
			return result
		}
		Env.Load(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value(), result)
		return result
	case "/=":
		Current_Value, ok := Env.Get(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		if !ok {
			return SkyLine_Evaluator_CreateError("%s does not exist within the environment", Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		}
		result := SkyLine_Evaluator_Eval_InfixExpression("/=", Current_Value, Evaled, Env)
		if SkyLine_Evaluator_CheckError(result) {
			fmt.Printf("Error when handling /=: %s", result.SkyLine_ObjectFunction_GetTrueValue())
			return result
		}
		Env.Load(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value(), result)
		return result
	case "=":
		_, ok := Env.Get(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		if !ok {
			return SkyLine_Evaluator_CreateError("Cannot assign variable a new value, make sure you declare the variable first using `set/cause/let/allow data := value;`")
		}
		Env.Load(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value(), Evaled)
	case ":=":
		_, ok := Env.Get(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value())
		if ok {
			return SkyLine_Evaluator_CreateError("Sorry, but `:=` tokens are only to declare variables, not to change them, please use '=' as an assignment expression")
		} else {
			Env.Load(Assignee.Name.SkyLine_NodeInterface_Get_Node_Value(), Evaled)
		}
	}
	return Evaled
}

func EvalRegisterCall(Val []SkyEnv.SL_Object) SkyEnv.SL_Object {
	for _, mod := range Val {
		switch v := mod.(type) {
		case *SkyEnv.SL_String:
			if ok := StandardLibNames[v.SkyLine_ObjectFunction_GetTrueValue()]; ok {
				if res := RegisterStandard[v.SkyLine_ObjectFunction_GetTrueValue()]; res != nil {
					RegisterStandard[v.SkyLine_ObjectFunction_GetTrueValue()]()
				} else {
					var errorm string
					errorm = "SkyLine STD: Developer Error -> On line 335 of SkyLine_Evaluator_Evaluate_ObjectRet.go (Error when working on reigstry)"
					errorm += "|RegisterStandard[]->SkyLine_Scrips_Language_Backend_Models.go(L:878)??? "
					errorm += "| The StandardLibNames map returned to recongnize the library, but it seems that the RegisterStandard[] did not return a valid function"
					return SkyLine_Evaluator_CreateError(errorm)
				}
			} else {
				return SkyLine_Evaluator_CreateError("SkyLine STD: Could not register or find standard libraries under the name of %s ", v.SkyLine_ObjectFunction_GetTrueValue())
			}
		default:
			return &SkyEnv.SL_Error{Message: "Sorry this data type to import libraries is not supported, register requires STRING"}
		}
	}
	return &SkyEnv.SL_Boolean{Value: false}
}
