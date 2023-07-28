///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Evaluation
// Extension         | .go ( golang source code file )
// Purpose           | Defines all of the evaluation main and prime functions
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
	"os"

	SkyAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func Check() bool {
	return EngineCallVal.Prepped
}

func SkyLineEvaluateContext(SkyLineCTX context.Context, SkyNode SkyAST.SL_Node, SkyEnvironment *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	select {
	case <-SkyLineCTX.Done():
		return &SkyEnv.SL_Error{
			Message: SkyLineCTX.Err().Error(),
		}
	default:
		// No Operation
	}
	if Check() {
		SkyEnvironment.Load(EngineCallVal.Name+"Version_", &SkyEnv.SL_String{Value: EngineCallVal.Version})  // Set version
		SkyEnvironment.Load(EngineCallVal.Name+"Name_", &SkyEnv.SL_String{Value: EngineCallVal.Name})        // Set Name
		SkyEnvironment.Load(EngineCallVal.Name+"Desc_", &SkyEnv.SL_String{Value: EngineCallVal.Description}) // Set Description
		SkyEnvironment.Load(EngineCallVal.Name+"Supported_", &SkyEnv.SL_String{Value: EngineCallVal.SOS})    // Set Supported operating systems
		for _, lib := range EngineCallVal.Require {
			if regf, ok := RegisterStandard[lib]; ok {
				regf()
			}
		}
		// Syntax: ProjectNameVersion
		// Reset to false, so loop does not repeat
		EngineCallVal.Prepped = false
	}
	switch N := SkyNode.(type) {
	case *SkyAST.SL_ENGINE:
		Value := SkyLine_Call_Eval(N.EngineValue, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Value) {
			return Value
		}
		return EvalEngineCall(Value)
	case *SkyAST.SL_Register:
		Element_s := SkyLine_Evaluator_Eval_Expression(N.RegistryValue, SkyEnvironment)
		if len(Element_s) == 1 && SkyLine_Evaluator_CheckError(Element_s[0]) {
			return Element_s[0]
		}
		return EvalRegisterCall(Element_s)
	case *SkyAST.SL_Identifier:
		return SkyLine_Evaluator_Eval_Identifier(N, SkyEnvironment)
	case *SkyAST.SL_Prog:
		return SkyLine_Evaluator_Eval_Program(N, SkyEnvironment)
	case *SkyAST.SL_Array:
		Element_s := SkyLine_Evaluator_Eval_Expression(N.Elements, SkyEnvironment)
		if len(Element_s) == 1 && SkyLine_Evaluator_CheckError(Element_s[0]) {
			return Element_s[0]
		}
		return &SkyEnv.SL_Array{
			Elements: Element_s,
		}
	case *SkyAST.SL_HashMap:
		return SkyLine_Evaluator_Eval_Hash_Map_Literal(N, SkyEnvironment)
	case *SkyAST.SL_Integer:
		return &SkyEnv.SL_Integer{
			Value: N.Value,
		}
	case *SkyAST.SL_String: // String
		return &SkyEnv.SL_String{
			Value: N.Value,
		}
	case *SkyAST.SL_Byte: // Byte c
		return &SkyEnv.SL_Byte{
			Value: N.Value,
		}
	case *SkyAST.SL_Integer16: // Integer 16
		return &SkyEnv.SL_Integer16{
			Value: N.Value,
		}
	case *SkyAST.SL_Integer32: // integer 32
		return &SkyEnv.SL_Integer32{
			Value: N.Value,
		}
	case *SkyAST.SL_Integer64: // Integer 64
		return &SkyEnv.SL_Integer64{
			Value: N.Value,
		}
	case *SkyAST.SL_Integer8: // Integer 8
		return &SkyEnv.SL_Integer8{
			Value: N.Value,
		}
	case *SkyAST.SL_Float:
		return &SkyEnv.SL_Float{
			Value: N.Value,
		}
	case *SkyAST.SL_Boolean:
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(N.Value)
	case *SkyAST.SL_NULL:
		return SkyLine_Null_ALLIAS
	case *SkyAST.SL_EN_Index_Expression:
		LeftOf := SkyLine_Call_Eval(N.Left, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(LeftOf) {
			return LeftOf
		}
		IDX := SkyLine_Call_Eval(N.Index, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(IDX) {
			return IDX
		}
		return SkyLine_Evaluator_Eval_Index_Expression(LeftOf, IDX)
	case *SkyAST.SL_EN_Postfix:
		return SkyLine_Evaluator_Eval_PostFix_PFExpression(N.Operator, SkyEnvironment, N)
	case *SkyAST.SL_EN_Infix:
		LeftOf := SkyLine_Call_Eval(N.Left, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(LeftOf) {
			return LeftOf
		}

		RightOf := SkyLine_Call_Eval(N.Right, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(RightOf) {
			return RightOf
		}

		Result := SkyLine_Evaluator_Eval_InfixExpression(N.Operator, LeftOf, RightOf, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Result) {
			fmt.Printf("Error[EVAL_INFIX]: %s\n", Result.SkyLine_ObjectFunction_GetTrueValue())
			os.Exit(0)
		}
		return (Result)
	case *SkyAST.SL_EN_Prefix:
		Rightof := SkyLine_Call_Eval(N.Right, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Rightof) {
			return Rightof
		}
		return SkyLine_Evaluator_Eval_Prefix_Expression(N.Operator, Rightof)
	case *SkyAST.SL_EN_Conditional_IfElse:
		return SkyLine_Evaluator_Eval_Conditional_Statement(N, SkyEnvironment)
	case *SkyAST.SL_EN_Conditional_Loop:
		return SkyLine_Evaluator_Eval_Conditional_ForLoop(N, SkyEnvironment)
	case *SkyAST.SL_EN_For_Each_Loop:
		return SkyLine_Evaluator_Eval_ForEach_Range_Loop(N, SkyEnvironment)
	case *SkyAST.SL_EN_Function_Literal:
		return &SkyEnv.SL_Functions{
			Defaults:               N.Defaults,
			Unit:                   N.Unit,
			Sky_Function_Arguments: N.Parameters,
			Env:                    SkyEnvironment,
		}
	case *SkyAST.SL_EN_Lable_JumpDef: //Lables
	case *SkyAST.SL_EN_Function_Definition:
		params := N.FunctionArguments
		body := N.Unit
		defaults := N.Defaults
		F := &SkyEnv.SL_Functions{
			Sky_Function_Arguments: params,
			Env:                    SkyEnvironment,
			Unit:                   body,
			Defaults:               defaults,
		}

		SkyEnvironment.Load(N.SkyLine_NodeInterface_Token_Literal(), F)
		return SkyLine_Null_ALLIAS
	case *SkyAST.SL_EN_Object_Call_Expression:
		Call := SKyLine_Evaluator_Eval_ObjectCallingExpression(N, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Call) {
			fmt.Printf("Error[Evaluation]: %s", Call.SkyLine_ObjectFunction_GetTrueValue())
		}
		return Call
	case *SkyAST.SL_EN_Call_Expression:
		Function := SkyLine_Call_Eval(N.Function, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Function) {
			return Function
		}
		Arguments := SkyLine_Evaluator_Eval_Expression(N.Arguments, SkyEnvironment)
		if len(Arguments) == 1 && SkyLine_Evaluator_CheckError(Arguments[0]) {
			return Arguments[0]
		}
		Application := SkyLine_Eval_Apply_Function_To_Evaluation(Function, SkyEnvironment, Arguments)
		if SkyLine_Evaluator_CheckError(Application) {
			fmt.Printf("Error calling function `%s` : %s \n", N.Function, Application.SkyLine_ObjectFunction_GetTrueValue())
			return Application
		}
		return Application
	case *SkyAST.SL_EN_Switch_ExpressionStatement:
		return SkyLine_Evaluator_Eval_Conditional_Switch(N, SkyEnvironment)
	case *SkyAST.SL_EN_VariableAssignmentStatement:
		return SkyLine_Evaluator_Eval_VariableAssignment_Statement(N, SkyEnvironment)
	case *SkyAST.Assignment_Constant_Const:
		Value := SkyLine_Call_Eval(N.Value, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Value) {
			return Value
		}
		SkyEnvironment.LoadConstant(N.Name.Value, Value)
		return Value
	case *SkyAST.Assignment_Cause_Set_Allow:
		Value := SkyLine_Call_Eval(N.Value, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Value) {
			return Value
		}
		SkyEnvironment.Load(N.Name.Value, Value)
		return Value
	case *SkyAST.Return_Ret_Return_Information:
		Value := SkyLine_Call_Eval(N.Expression, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(Value) {
			return Value
		}
		return &SkyEnv.SL_Return{
			Value: Value,
		}
	case *SkyAST.SL_UnitBlockStatement:
		return SkyLine_Evaluator_Eval_UnitAndBlock_Statement(N, SkyEnvironment)
	case *SkyAST.Expression_Statement:
		return SkyLine_Call_Eval(N.Expression, SkyEnvironment)
	case *SkyAST.SL_ImportExpression:
		Element_s := EvalImportingExpression(N, SkyEnvironment)
		switch c := Element_s.(type) {
		case *SkyEnv.SL_Array:
			Mapper := make(map[string]*SkyEnv.SL_Module, 0)
			RetHash := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
			if len(c.Elements) == 1 && SkyLine_Evaluator_CheckError(c.Elements[0]) {
				return c.Elements[0]
			}
			for _, c := range c.Elements {
				switch t := c.(type) {
				case *SkyEnv.SL_Module:
					Mapper[t.GetModuleFile()] = c.(*SkyEnv.SL_Module)
				}
			}
			for k, v := range Mapper {
				key := &SkyEnv.SL_String{Value: k}
				val := v
				NewHashCreate := SkyEnv.HashPair{Key: key, Value: val}
				RetHash[key.SL_HashKeyType()] = NewHashCreate
			}
			return &SkyEnv.SL_HashMap{Pairs: RetHash}
		case *SkyEnv.SL_Module:
			return Element_s
		}
	}

	return nil
}

func EvalImportingExpression(iexp *SkyAST.SL_ImportExpression, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {

	if len(iexp.Name) > 1 {
		evaled := make([]SkyEnv.SL_Object, 0)
		modules := make([]SkyEnv.SL_Object, 0)
		for _, name := range iexp.Name {
			var e bool
			name := SkyLine_Call_Eval(name, Env)
			if SkyLine_Evaluator_CheckError(name) {
				evaled = append(evaled, name)
				return &SkyEnv.SL_Array{Elements: evaled}
			}
			if s, ok := name.(*SkyEnv.SL_String); ok {
				attrs := EvalModule(s.Value)
				if SkyLine_Evaluator_CheckError(attrs) {
					evaled = append(evaled, attrs)
					return &SkyEnv.SL_Array{Elements: evaled}
				}
				modules = append(modules, &SkyEnv.SL_Module{SL_ModuleNames: s.Value, SL_Attributes: attrs})
			} else {
				fmt.Println("Set error to true during conversion")
				e = true
			}
			if e {
				log.Fatalf("ImportError: Invalid import path -> %s", iexp.Name)
			}
		}
		return &SkyEnv.SL_Array{Elements: modules}
	} else {
		name := SkyLine_Call_Eval(iexp.Name[0], Env)
		if SkyLine_Evaluator_CheckError(name) {
			return name
		}
		if s, ok := name.(*SkyEnv.SL_String); ok {
			attrs := EvalModule(s.Value)
			if SkyLine_Evaluator_CheckError(attrs) {
				return attrs
			}
			return &SkyEnv.SL_Module{SL_ModuleNames: s.Value, SL_Attributes: attrs}
		}
		return SkyLine_Evaluator_CreateError("ImportError: Invalid import path -> %s", iexp.Name[0])
	}
}
