///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_EvalInfix
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions to evaluate and execute infix expressions
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
	"math"
	"strings"
)

var InfixSameTypeFunctions = map[SkyEnv.ObjectDataType]func(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object{
	SkyEnv.SKYLINE_DATATYPE_INTEGER8_OBJECT:  SkyLine_Evaluator_Eval_InfixExpression_Integer8,  // If L = int8 && R = int7
	SkyEnv.SKYLINE_DATATYPE_INTEGER16_OBJECT: SkyLine_Evaluator_Eval_InfixExpression_Integer16, // If L = int16 && R = int16
	SkyEnv.SKYLINE_DATATYPE_INTEGER32_OBJECT: SkyLine_Evaluator_Eval_InfixExpression_Integer32, // If L = int32 && R = int32
	SkyEnv.SKYLINE_DATATYPE_INTEGER64_OBJECT: SkyLine_Evaluator_Eval_InfixExpression_Integer64, // If L = int64 && R = int64
	SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT:   SkyLine_Evaluator_Eval_InfixExpression_Integer,   // If L = int  && R = int
	SkyEnv.SKYLINE_DATATYPE_BOOLEAN_OBJECT:   SkyLine_Evaluator_Eval_InfixExpression_Boolean,   // If L = bool && R = bool
	SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT:     SkyLine_Evaluator_Eval_InfixExpression_Float,     // If L = float && R == float
	SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT:    SkyLine_Evaluator_Eval_InfixExpression_String,    // If L = string && R = string
}

func CheckIntType(Value SkyEnv.SL_Object) bool {
	switch Value.(type) {
	case *SkyEnv.SL_Integer8:
		return true
	case *SkyEnv.SL_Integer:
		return true
	case *SkyEnv.SL_Integer16:
		return true
	case *SkyEnv.SL_Integer32:
		return true
	case *SkyEnv.SL_Integer64:
		return true
	default:
		return false
	}
}

func SkyLine_Evaluator_Eval_InfixExpression(Op string, Left, Right SkyEnv.SL_Object, Env *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	leftType := Left.SkyLine_ObjectFunction_GetDataType()
	rightType := Right.SkyLine_ObjectFunction_GetDataType()

	// Check if the data types are supported
	if leftType == rightType {
		if _, ok := InfixSameTypeFunctions[leftType]; ok {
			if evf, ok := InfixSameTypeFunctions[rightType]; ok {
				return evf(Op, Left, Right)
			}
		}
	}
	switch {
	case Left.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT && Right.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT:
		return SkyLine_Evaluator_Eval_InfixExpression_FloatInteger(Op, Left, Right)
	case Left.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT && Right.SkyLine_ObjectFunction_GetDataType() == SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT:
		return SkyLine_Evaluator_Eval_InfixExpression_IntegerFloat(Op, Left, Right)
	case Op == "+" && leftType == SkyEnv.SKYLINE_DATATYPE_HASH_OBJECT && rightType == SkyEnv.SKYLINE_DATATYPE_HASH_OBJECT:
		LV := Left.(*SkyEnv.SL_HashMap).Pairs
		RV := Right.(*SkyEnv.SL_HashMap).Pairs
		pairs := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
		for Key, Val := range LV {
			pairs[Key] = Val
		}
		for Key, Val := range RV {
			pairs[Key] = Val
		}
		return &SkyEnv.SL_HashMap{Pairs: pairs}
	case Op == "+" && leftType == SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT && rightType == SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT:
		LV := Left.(*SkyEnv.SL_Array).Elements
		RV := Right.(*SkyEnv.SL_Array).Elements
		elements := make([]SkyEnv.SL_Object, len(LV)+len(RV))
		elements = append(LV, RV...)
		return &SkyEnv.SL_Array{Elements: elements}
	case Op == "*" && leftType == SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT && CheckIntType(Right):
		LV := Left.(*SkyEnv.SL_Array).Elements
		elements := LV
		var rv int
		switch RVal := Right.(type) {
		case *SkyEnv.SL_Integer8:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer16:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer32:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer64:
			rv = int(RVal.Value)
		}
		for i := rv; i > 1; i-- {
			elements = append(elements, LV...)
		}
		return &SkyEnv.SL_Array{Elements: elements}

	case Op == "*" && CheckIntType(Left) && rightType == SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT:
		var lv int
		switch RVal := Left.(type) {
		case *SkyEnv.SL_Integer8:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer16:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer32:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer64:
			lv = int(RVal.Value)
		}
		rightVal := Right.(*SkyEnv.SL_Array).Elements
		elements := rightVal
		for i := lv; i > 1; i-- {
			elements = append(elements, rightVal...)
		}
		return &SkyEnv.SL_Array{Elements: elements}
	case Op == "*" && leftType == SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT && CheckIntType(Right):
		LV := Left.(*SkyEnv.SL_String).Value
		var rv int
		switch RVal := Right.(type) {
		case *SkyEnv.SL_Integer8:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer16:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer32:
			rv = int(RVal.Value)
		case *SkyEnv.SL_Integer64:
			rv = int(RVal.Value)
		}
		return &SkyEnv.SL_String{Value: strings.Repeat(LV, rv)}
	case Op == "*" && CheckIntType(Left) && rightType == SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT:
		var lv int
		switch RVal := Left.(type) {
		case *SkyEnv.SL_Integer8:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer16:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer32:
			lv = int(RVal.Value)
		case *SkyEnv.SL_Integer64:
			lv = int(RVal.Value)
		}
		rv := Right.(*SkyEnv.SL_String).Value
		return &SkyEnv.SL_String{Value: strings.Repeat(rv, lv)}
	case Op == "&&":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(SkyLine_Evaluator_ObjectToNativeBoolean(Left) && SkyLine_Evaluator_ObjectToNativeBoolean(Right))
	case Op == "||":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(SkyLine_Evaluator_ObjectToNativeBoolean(Left) || SkyLine_Evaluator_ObjectToNativeBoolean(Right))
	case Op == "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(Left == Right)
	case Op == "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(Left != Right)
	case Left.SkyLine_ObjectFunction_GetDataType() != Right.SkyLine_ObjectFunction_GetDataType():
		return SkyLine_Evaluator_CreateError("data type mismatch in Infix Expression : %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for infix expression: %s %s %s", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_String(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_String)
	RV := Right.(*SkyEnv.SL_String)

	switch Op {
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV.Value == RV.Value)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV.Value != RV.Value)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV.Value >= RV.Value)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV.Value > RV.Value)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV.Value <= RV.Value)
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV.Value < RV.Value)
	case "+":
		return &SkyEnv.SL_String{
			Value: LV.Value + RV.Value,
		}
	case "+=":
		return &SkyEnv.SL_String{
			Value: LV.Value + RV.Value,
		}
	case "*=":
		return &SkyEnv.SL_String{
			Value: strings.Repeat(LV.Value, len(RV.Value)),
		}
	case "-=":
		return &SkyEnv.SL_String{
			Value: LV.Value[len(RV.Value):],
		}
	}
	return SkyLine_Evaluator_CreateError("Unknown operator for infix string operation : %s %s %s", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
}

func SkyLine_Evaluator_Eval_InfixExpression_Float(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Float).Value
	RV := Right.(*SkyEnv.SL_Float).Value

	switch Op {
	case "+":
		return &SkyEnv.SL_Float{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Float{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Float{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Float{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Float{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Float{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Float{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Float{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Float{Value: math.Pow(LV, RV)}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for float infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_Integer32(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Integer32).Value
	RV := Right.(*SkyEnv.SL_Integer32).Value
	switch Op {
	case "%":
		return &SkyEnv.SL_Integer32{Value: LV % RV}
	case "+":
		return &SkyEnv.SL_Integer32{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Integer32{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Integer32{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Integer32{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Integer32{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Integer32{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Integer32{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Integer32{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Integer32{Value: int32(math.Pow(float64(LV), float64(RV)))}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	case "..":
		len := int(RV-LV) + 1
		array := make([]SkyEnv.SL_Object, len)
		i := 0
		for i < len {
			array[i] = &SkyEnv.SL_Integer32{Value: LV}
			LV++
			i++
		}
		return &SkyEnv.SL_Array{Elements: array}
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for Integer infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_Integer16(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Integer16).Value
	RV := Right.(*SkyEnv.SL_Integer16).Value
	switch Op {
	case "%":
		return &SkyEnv.SL_Integer16{Value: LV % RV}
	case "+":
		return &SkyEnv.SL_Integer16{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Integer16{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Integer16{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Integer16{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Integer16{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Integer16{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Integer16{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Integer16{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Integer16{Value: int16(math.Pow(float64(LV), float64(RV)))}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	case "..":
		len := int(RV-LV) + 1
		array := make([]SkyEnv.SL_Object, len)
		i := 0
		for i < len {
			array[i] = &SkyEnv.SL_Integer16{Value: LV}
			LV++
			i++
		}
		return &SkyEnv.SL_Array{Elements: array}
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for Integer infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_Integer8(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Integer8).Value
	RV := Right.(*SkyEnv.SL_Integer8).Value
	switch Op {
	case "%":
		return &SkyEnv.SL_Integer8{Value: LV % RV}
	case "+":
		return &SkyEnv.SL_Integer8{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Integer8{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Integer8{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Integer8{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Integer8{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Integer8{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Integer8{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Integer8{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Integer8{Value: int8(math.Pow(float64(LV), float64(RV)))}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	case "..":
		len := int(RV-LV) + 1
		array := make([]SkyEnv.SL_Object, len)
		i := 0
		for i < len {
			array[i] = &SkyEnv.SL_Integer8{Value: LV}
			LV++
			i++
		}
		return &SkyEnv.SL_Array{Elements: array}
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for Integer infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_Integer64(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Integer64).Value
	RV := Right.(*SkyEnv.SL_Integer64).Value
	switch Op {
	case "%":
		return &SkyEnv.SL_Integer64{Value: LV % RV}
	case "+":
		return &SkyEnv.SL_Integer64{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Integer64{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Integer64{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Integer64{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Integer64{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Integer64{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Integer64{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Integer64{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Integer64{Value: int64(math.Pow(float64(LV), float64(RV)))}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	case "..":
		len := int(RV-LV) + 1
		array := make([]SkyEnv.SL_Object, len)
		i := 0
		for i < len {
			array[i] = &SkyEnv.SL_Integer64{Value: LV}
			LV++
			i++
		}
		return &SkyEnv.SL_Array{Elements: array}
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for Integer infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_Integer(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Integer).Value
	RV := Right.(*SkyEnv.SL_Integer).Value
	switch Op {
	case "%":
		return &SkyEnv.SL_Integer{Value: LV % RV}
	case "+":
		return &SkyEnv.SL_Integer{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Integer{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Integer{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Integer{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Integer{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Integer{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Integer{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Integer{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Integer{Value: int(math.Pow(float64(LV), float64(RV)))}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	case "..":
		len := int(RV-LV) + 1
		array := make([]SkyEnv.SL_Object, len)
		i := 0
		for i < len {
			array[i] = &SkyEnv.SL_Integer{Value: LV}
			LV++
			i++
		}
		return &SkyEnv.SL_Array{Elements: array}
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for Integer infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_IntegerFloat(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := float64(Left.(*SkyEnv.SL_Integer).Value)
	RV := Right.(*SkyEnv.SL_Float).Value

	switch Op {
	case "+":
		return &SkyEnv.SL_Float{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Float{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Float{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Float{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Float{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Float{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Float{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Float{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Float{Value: math.Pow(float64(LV), float64(RV))}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for Integer->Float infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())

	}
}

func SkyLine_Evaluator_Eval_InfixExpression_FloatInteger(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LV := Left.(*SkyEnv.SL_Float).Value
	RV := float64(Right.(*SkyEnv.SL_Integer).Value)

	switch Op {
	case "+":
		return &SkyEnv.SL_Float{Value: LV + RV}
	case "+=":
		return &SkyEnv.SL_Float{Value: LV + RV}
	case "-":
		return &SkyEnv.SL_Float{Value: LV - RV}
	case "-=":
		return &SkyEnv.SL_Float{Value: LV - RV}
	case "/":
		return &SkyEnv.SL_Float{Value: LV / RV}
	case "/=":
		return &SkyEnv.SL_Float{Value: LV / RV}
	case "*":
		return &SkyEnv.SL_Float{Value: LV * RV}
	case "*=":
		return &SkyEnv.SL_Float{Value: LV * RV}
	case "**":
		return &SkyEnv.SL_Float{Value: math.Pow(LV, RV)}
	case "<":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV < RV)
	case "<=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV <= RV)
	case ">":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV > RV)
	case ">=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV >= RV)
	case "==":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV == RV)
	case "!=":
		return SkyLine_Eval_FromNativeBoolean_To_BooleanObject(LV != RV)
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for String->Integer infix expression: %s %s %s ", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}

func SkyLine_Evaluator_Eval_InfixExpression_Boolean(Op string, Left, Right SkyEnv.SL_Object) SkyEnv.SL_Object {
	LeftValue := &SkyEnv.SL_String{Value: string(Left.SkyLine_ObjectFunction_GetTrueValue())}
	RightValue := &SkyEnv.SL_String{Value: string(Right.SkyLine_ObjectFunction_GetTrueValue())}

	switch Op {
	case "<":
		return SkyLine_Evaluator_Eval_InfixExpression_String(Op, LeftValue, RightValue)
	case "<=":
		return SkyLine_Evaluator_Eval_InfixExpression_String(Op, LeftValue, RightValue)
	case ">=":
		return SkyLine_Evaluator_Eval_InfixExpression_String(Op, LeftValue, RightValue)
	case ">":
		return SkyLine_Evaluator_Eval_InfixExpression_String(Op, LeftValue, RightValue)
	default:
		return SkyLine_Evaluator_CreateError("Unknown operator for boolean infix expression: %s %s %s", Left.SkyLine_ObjectFunction_GetDataType(), Op, Right.SkyLine_ObjectFunction_GetDataType())
	}
}
