/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//                              _____ _       __    _
//                             |   __| |_ _ _|  |  |_|___ ___
//                             |__   | '_| | |  |__| |   | -_|
//                             |_____|_,_|_  |_____|_|_|_|___|
//                                       |___|
//
// These sections are to help yopu better understand what each section is or what each file represents within the SkyLine programming language. These sections can also
//
// help seperate specific values so you can better understand what a specific section or specific set of values of functions is doing.
//
// These sections also give information on the file, project and status of the section
//
//
// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// Filename      |  SkyLine_Evaluator_Evaluate_ObjectRet.go
// Project       |  SkyLine programming language
// Line Count    |  700+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the Evaluator functions that return a type structure SLC_Object which will return SL_Object based values. These functions all
//                 belong to the evaluator
//
// STATE         | Needs to be organized and worked on
// Resolution    | Functions need to be renamed to be SL specific, functions need to be organized, functions need to be automated.
//
package SkyLine_Backend

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	SLC "github.com/SkyPenguin-Solutions/SkyLineConfigurationEngine/Engine/Backend"
)

func EvalSwitch(Switch *Switch, Env *Environment_of_environment) SLC_Object {
	object := Eval(Switch.Value, Env)
	for _, options := range Switch.Choices {
		if options.Def {
			continue
		}
		for _, values := range options.Expr {
			output := Eval(values, Env)
			if object.SL_RetrieveDataType() == output.SL_RetrieveDataType() && (object.SL_InspectObject() == output.SL_InspectObject()) {
				Block := Eval_Block_Statement(options.Block, Env)
				return Block
			}
		}
	}
	for _, opt := range Switch.Choices {
		if opt.Def {
			output := Eval_Block_Statement(opt.Block, Env)
			return output
		}
	}
	return nil
}

func Evaluate_ObjectCall(call *ObjectCallExpression, env *Environment_of_environment) SLC_Object {
	var meth *CallExpression
	obj := Eval(call.SLC_Object, env)
	if method, ok := call.Call.(*CallExpression); ok {
		args := evalExpressions(call.Call.(*CallExpression).Arguments, env)
		ret := obj.InvokeMethod(method.Function.SL_ExtractStringValue(), *env, args...)
		if ret != nil {
			return ret
		}
		attempts := []string{}
		attempts = append(attempts, strings.ToLower(string(obj.SL_RetrieveDataType())))
		attempts = append(attempts, "object")
		for _, prefix := range attempts {
			name := prefix + "." + method.Function.SL_ExtractStringValue()
			if fn, ok := env.Get(name); ok {
				extendEnv := extendFunctionEnv(fn.(*Function), args)
				extendEnv.Set("self", obj)
				evaluated := Eval(fn.(*Function).Body, extendEnv)
				obj = unwrapReturnValue(evaluated)
				return obj
			}
		}
	} else {
		meth = method
	}
	return NewError("Failed to invoke method: %s with method %s and object %s",
		call.Call.(*CallExpression).Function.SL_ExtractStringValue(),
		meth,
		obj,
	)
}

func applyFunction(Env *Environment_of_environment, fn SLC_Object, args []SLC_Object) SLC_Object {
	switch fn := fn.(type) {
	case *Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *Builtin:
		return fn.Fn(Env, args...)
	default:
		// expecting crash if it is not a function type object | Macros are bugged
		defer func() {
			if x := recover(); x != nil {
				fmt.Print("\n\n Assuming this is a macro, there was something wrong. Macro does not exist")
				os.Exit(1)
			}
		}()
		return NewError(Map_Eval[ERROR_INVALID_FUNCTION_NOT_FOUND_OR_UNSUPPORTED_TYPE](string(fn.SL_InspectObject())).Message)
	}
}

func unwrapReturnValue(obj SLC_Object) SLC_Object {
	if returnValue, ok := obj.(*ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalIndexExpression(left, index SLC_Object) SLC_Object {
	switch {
	case left.SL_RetrieveDataType() == ArrayType && index.SL_RetrieveDataType() == IntegerType:
		return evalArrayIndexExpression(left, index)
	case left.SL_RetrieveDataType() == HashType:
		return evalHashIndexExpression(left, index)
	case left.SL_RetrieveDataType() == TOKEN_MODULE:
		return evalModuleIndexExpression(left, index)
	default:
		return NewError(Map_Eval[ERROR_INVALID_INDEX_EXPRESSION_OR_UNSUPPORTED_INDEX](string(left.SL_RetrieveDataType())).Message)
	}
}
func evalModuleIndexExpression(module, index SLC_Object) SLC_Object {
	moduleObject := module.(*Module)
	return evalHashIndexExpression(moduleObject.Attrs, index)
}

func evalArrayIndexExpression(array, index SLC_Object) SLC_Object {
	arrObj := array.(*Array)
	idx := index.(*Integer).Value
	max := int64(len(arrObj.Elements) - 1)

	if idx < 0 || idx > max {
		return NilValue
	}

	return arrObj.Elements[idx]
}

func evalHashLiteral(node *HashLiteral, Env *Environment_of_environment) SLC_Object {
	pairs := make(map[HashKey]HashPair, len(node.Pairs))

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, Env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(Hashable)
		if !ok {
			return NewError(Map_Eval[ERROR_INVALID_HASH_KEY_COULD_NOT_PARSE_HASHKEY](string(key.SL_RetrieveDataType())).Message)
		}

		value := Eval(valueNode, Env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = HashPair{
			Key:   key,
			Value: value,
		}
	}

	return &Hash{Pairs: pairs}
}

func evalHashIndexExpression(left, index SLC_Object) SLC_Object {
	key, ok := index.(Hashable)
	if !ok {
		return NewError(Map_Eval[ERROR_INVALID_HASH_KEY_COULD_NOT_BE_USED_AS_KEY](string(index.SL_RetrieveDataType())).Message)
	}

	hashObj := left.(*Hash)
	if pair, exists := hashObj.Pairs[key.HashKey()]; exists {
		return pair.Value
	}
	return NilValue
}

func evalIdent(node *Ident, Env *Environment_of_environment) SLC_Object {
	if val, ok := Env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := Builtins[node.Value]; ok {
		return builtin
	} else {
		if builtin_var, ok := BuiltInVariables_String[node.Value]; ok {
			return builtin_var
		} else if builtin_var_i, ok := BuiltInVariables_Integer[node.Value]; ok {
			return builtin_var_i
		} else if builtin_var_b, ok := BuiltInVariables_Boolean[node.Value]; ok {
			return builtin_var_b
		} else if builtin_var_f, ok := BuiltInVariables_Float[node.Value]; ok {
			return builtin_var_f
		}
	}
	return NewError(Map_Eval[ERROR_INVALID_IDENTIFIER_IDENTIFIER_WAS_NOT_FOUND_OR_KNOWN](node.Value).Message)
}

func evalExpressions(exprs []Expression, Env *Environment_of_environment) []SLC_Object {
	result := make([]SLC_Object, 0, len(exprs))

	for _, expr := range exprs {
		evaluated := Eval(expr, Env)
		if isError(evaluated) {
			return []SLC_Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func EvalAssignModify(Assign *AssignmentStatement, Env *Environment_of_environment) SLC_Object {
	evaled := Eval(Assign.Value, Env)
	if isError(evaled) {
		return evaled
	}
	if Assign.Operator == "+=" || Assign.Operator == "-=" || Assign.Operator == "*=" || Assign.Operator == "/=" {
		currentvalue, ok := Env.Get(Assign.Name.Value)
		if !ok {
			return NewError("%s might not exist or is unknown to the parser -> \n SUGGESTION |+ allow %s = ...", Assign.Name.SL_ExtractStringValue(), Assign.Name.SL_ExtractStringValue())
		}
		res := evalInfixExpression(Assign.Operator, currentvalue, evaled)
		if isError(res) {
			fmt.Printf("Error during handle and evaluation to symbol %s %s ", Assign.Operator, res.SL_InspectObject())
			return res
		}
		Env.Set(Assign.Name.SL_ExtractStringValue(), res)
	} else if Assign.Operator == "=" {
		Env.Set(Assign.Name.SL_ExtractStringValue(), evaled)
	}
	return evaled
}

func evalProgram(program *Program, Env *Environment_of_environment) SLC_Object {
	var result SLC_Object

	for _, stmt := range program.Statements {
		result = Eval(stmt, Env)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *Boolean_Object {
	if input {
		return TrueValue
	}
	return FalseValue
}

func evalPrefixExpression(operator string, right SLC_Object) SLC_Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		// Eventually add support for prefix and infix seperate errors
		return NewError(fmt.Sprintf("Operator '%s' was not supported, (- and !) are the only ones currently supported.", operator))
	}
}

func evalBangOperatorExpression(right SLC_Object) SLC_Object {
	if right == NilValue || right == FalseValue {
		return TrueValue
	}
	return FalseValue
}

func EvalModifyCall(val SLC_Object) SLC_Object {
	var modwhat, modsettings string
	switch v := val.(type) {
	case *String:
		if !strings.Contains(v.SL_InspectObject(), ":") {
			var outmsg string
			outmsg += "Sorry but the function to modify needs a : seperating the environment and the type parameter"
			outmsg += "\n\n"
			outmsg += "should be called like so -> modify(Environmental_system:setting)"
			return NewError(outmsg)
		} else {
			strs := strings.Split(v.SL_InspectObject(), ":")
			if len(strs) != 2 {
				var newmsg string
				newmsg += "Whoops! that was not good, the length from the splitter returned not equal to two which means there was more arguments than needed\n"
				newmsg += "| Greater than 2? " + fmt.Sprint((len(strs) > 2))
				newmsg += "| ----------------"
				newmsg += "| Should be 2 arguments only and 2 statements which are shown in the example below\n"
				newmsg += "|\n"
				newmsg += "|modify('Environmental_system:setting')\n"
				return NewError(newmsg)
			} else {
				modwhat = strs[0]
				modsettings = strs[1]
			}
		}
	default:
		return NewError(
			"Sorry but the type supplied to modify needs to be STRING not %s", fmt.Sprint(val.SL_RetrieveDataType()))
	}
	VerifyModification(modwhat, modsettings)
	return &Nil{}
}

func EvalRegisterCall(val SLC_Object) SLC_Object {
	switch v := val.(type) {
	case *String:
		// Continue to parse the next statement or library
		if ok := StandardLibNames[val.SL_InspectObject()]; ok {
			if res := RegisterStandard[val.SL_InspectObject()]; res != nil {
				RegisterStandard[val.SL_InspectObject()]()
			} else {
				var errorm string
				errorm = "SkyLine STD: Developer Error -> On line 335 of SkyLine_Evaluator_Evaluate_ObjectRet.go (Error when working on reigstry)"
				errorm += "|RegisterStandard[]->SkyLine_Scrips_Language_Backend_Models.go(L:878)??? "
				errorm += "| The StandardLibNames map returned to recongnize the library, but it seems that the RegisterStandard[] did not return a valid function"
				return NewError(errorm)
			}
		} else {
			return NewError("SkyLine STD: Could not register or find standard libraries under the name of %s ", val.SL_InspectObject())
		}
	default:
		return NewError("Sorry but this type is unsupported, MUST be a string in call to register() -> got %s", v.SL_RetrieveDataType())
	}
	return &Nil{}
}

func evalMinusPrefixOperatorExpression(right SLC_Object) SLC_Object {
	switch right := right.(type) {
	case *Integer:
		return &Integer{Value: -right.Value}
	case *Float:
		return &Float{Value: -right.Value}
	default:
		return NewError("unknown operator: ( %s ) ", right.SL_RetrieveDataType())
	}
}

func EnvalBooleanInfixEXPRESSION_NODE(operator string, left, right SLC_Object) SLC_Object {
	lft := &String{Value: string(left.SL_InspectObject())}
	rft := &String{Value: string(right.SL_InspectObject())}
	switch operator {
	case "<":
		return Eval_String_Infix_Expression(operator, lft, rft)
	case "<=":
		return Eval_String_Infix_Expression(operator, left, rft)
	case ">":
		return Eval_String_Infix_Expression(operator, lft, rft)
	case ">=":
		return Eval_String_Infix_Expression(operator, left, rft)
	default:
		return NewError("Unknown operator: %s %s %s", lft.SL_RetrieveDataType(), operator, rft.SL_RetrieveDataType())

	}
}

func evalInfixExpression(operator string, left, right SLC_Object) SLC_Object {
	switch operator {
	}
	switch {
	case operator == "+" && left.SL_RetrieveDataType() == HashType && right.SL_RetrieveDataType() == HashType:
		lv := left.(*Hash).Pairs
		rv := right.(*Hash).Pairs
		pair := make(map[HashKey]HashPair)
		for idx, x := range lv {
			pair[idx] = x
		}
		for idx, c := range rv {
			pair[idx] = c
		}
		return &Hash{Pairs: pair}
	case operator == "+" && left.SL_RetrieveDataType() == ArrayType && right.SL_RetrieveDataType() == ArrayType:
		lv, rv := left.(*Array).Elements, right.(*Array).Elements
		elements := make([]SLC_Object, len(lv)+len(rv))
		elements = append(elements, rv...)
		return &Array{Elements: elements}
	case operator == "*" && left.SL_RetrieveDataType() == ArrayType && right.SL_RetrieveDataType() == IntegerType:
		lv, rv := left.(*Array).Elements, int(right.(*Integer).Value)
		elements := lv
		for idx := rv; idx > 1; idx-- {
			elements = append(elements, lv...)
		}
		return &Array{Elements: elements}
	case operator == "*" && left.SL_RetrieveDataType() == IntegerType && right.SL_RetrieveDataType() == ArrayType:
		lv, rv := int(left.(*Integer).Value), right.(*Array).Elements
		elements := rv
		for idx := lv; idx > 1; idx-- {
			elements = append(elements, rv...)
		}
		return &Array{Elements: elements}
	case operator == "*" && left.SL_RetrieveDataType() == TOKEN_STRING && right.SL_RetrieveDataType() == IntegerType:
		lv := left.(*String).Value
		rv := right.(*Integer).Value
		return &String{Value: strings.Repeat(lv, int(rv))}
	case operator == "*" && left.SL_RetrieveDataType() == IntegerType && right.SL_RetrieveDataType() == StringType:
		lv, rv := left.(*Integer).Value, right.(*String).Value
		return &String{Value: strings.Repeat(rv, int(lv))}
	case left.SL_RetrieveDataType() == IntegerType && right.SL_RetrieveDataType() == IntegerType:
		return evalIntegerInfixExpression(operator, left, right)
	case left.SL_RetrieveDataType() == FloatType && right.SL_RetrieveDataType() == FloatType:
		return Eval_Float_Infix_Expression(operator, left, right)
	case left.SL_RetrieveDataType() == FloatType && right.SL_RetrieveDataType() == IntegerType:
		return evalFloatIntegerInfixExpression(operator, left, right)
	case left.SL_RetrieveDataType() == IntegerType && right.SL_RetrieveDataType() == FloatType:
		return EvalSpecialFloatInfix(operator, left, right)
	case left.SL_RetrieveDataType() == StringType && right.SL_RetrieveDataType() == StringType:
		return Eval_String_Infix_Expression(operator, left, right)
	case operator == "&&":
		return nativeBoolToBooleanObject(OBJECT_CONV_TO_NATIVE_BOOL(left) && OBJECT_CONV_TO_NATIVE_BOOL(right))
	case operator == "||":
		return nativeBoolToBooleanObject(OBJECT_CONV_TO_NATIVE_BOOL(left) || OBJECT_CONV_TO_NATIVE_BOOL(right))
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.SL_RetrieveDataType() == BooleanType && right.SL_RetrieveDataType() == BooleanType:
		return EnvalBooleanInfixEXPRESSION_NODE(operator, left, right)
	case left.SL_RetrieveDataType() != right.SL_RetrieveDataType():
		return NewError(Map_Eval[ERROR_INVALID_DATATYPE_INIFX_OPERATION_LR](string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType()), operator).Message)
	default:
		return NewError(Map_Eval[ERROR_INVALID_OPERATOR_DURING_EVALUATION](operator, string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType())).Message)
	}
}

// FLOATS ARE NOT INTEGERS THEY ARE THEIR OWN DATA TYPE | Modify later
func EvalSpecialFloatInfix(operator string, left, right SLC_Object) SLC_Object {
	leftVal := left.(*Float).Value
	rightVal := right.(*Float).Value
	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "+=":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "-=":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "*=":
		return &Float{Value: leftVal * rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "/=":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NewError("unknown operator: %s %s %s",
			left.SL_RetrieveDataType(), operator, right.SL_RetrieveDataType())
	}
}

func evalFloatIntegerInfixExpression(operator string, left, right SLC_Object) SLC_Object {
	leftVal := left.(*Float).Value
	rightVal := float64(right.(*Integer).Value)
	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "+=":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "-=":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "*=":
		return &Float{Value: leftVal * rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "/=":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NewError("unknown operator: %s %s %s",
			left.SL_RetrieveDataType(), operator, right.SL_RetrieveDataType())
	}
}

func evalIntegerInfixExpression(operator string, left, right SLC_Object) SLC_Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "%":
		return &Integer{Value: leftVal % rightVal}
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "-=":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "**":
		return &Integer{Value: int64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "*=":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		return &Integer{Value: leftVal / rightVal}
	case "/=":
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "..":
		len := int(rightVal-leftVal) + 1
		array := make([]SLC_Object, len)
		i := 0
		for i < len {
			array[i] = &Integer{Value: leftVal}
			leftVal++
			i++
		}
		return &Array{Elements: array}
	default:
		println("DEBUG -> OPERATOR FOR INFIX EXPRESSION -> ", operator)

		return NewError(fmt.Sprintf("Integer Infix Expression does not exist %s and %s", string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType())))
	}
}

func Eval_Float_Infix_Expression(operator string, left, right SLC_Object) SLC_Object {
	var leftVal, rightVal float64

	switch left := left.(type) {
	case *Integer:
		leftVal = float64(left.Value)
	case *Float:
		leftVal = left.Value
	default:
		return NewError(fmt.Sprintf("Please make sure the type of the left value is equal to the right type value in -> %s %s %s", operator, string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType())))
	}

	switch right := right.(type) {
	case *Integer:
		rightVal = float64(right.Value)
	case *Float:
		rightVal = right.Value
	default:
		return NewError(fmt.Sprintf("Please make sure the type of the left value is equal to the right type value in -> %s %s %s", operator, string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType())))
	}

	switch operator {
	case "+":
		return &Float{Value: leftVal + rightVal}
	case "+=":
		return &Float{Value: leftVal + rightVal}
	case "-":
		return &Float{Value: leftVal - rightVal}
	case "-=":
		return &Float{Value: leftVal - rightVal}
	case "*":
		return &Float{Value: leftVal * rightVal}
	case "**":
		return &Float{Value: math.Pow(leftVal, rightVal)}
	case "*=":
		return &Float{Value: leftVal * rightVal}
	case "/":
		return &Float{Value: leftVal / rightVal}
	case "/=":
		return &Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NewError(fmt.Sprintf("Make sure the operator you are using is supported for float infix expressions in -> (%s %s %s)", operator, string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType())))
	}
}

func Eval_String_Infix_Expression(operator string, left, right SLC_Object) SLC_Object {
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value

	switch operator {
	case "+":
		return &String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case "+=":
		return &String{Value: leftVal + rightVal}
	case "*=":
		return &String{Value: leftVal + strings.Repeat(leftVal, len(rightVal))}
	default:
		return NewError(fmt.Sprintf("Make sure the operator you used is support for string infix expressiong, you gave an unknown operator in -> (%s %s %s)", operator, string(left.SL_RetrieveDataType()), string(right.SL_RetrieveDataType())))
	}
}

func Eval_Block_Statement(block *BlockStatement, Env *Environment_of_environment) SLC_Object {
	var result SLC_Object

	for _, stmt := range block.Statements {
		result = Eval(stmt, Env)
		if result != nil {
			rt := result.SL_RetrieveDataType()
			if rt == ReturnValueType || rt == ErrorType {
				return result
			}
		}
	}

	return result
}

func Eval_Conditional(ie *ConditionalExpression, Env *Environment_of_environment) SLC_Object {
	condition := Eval(ie.Condition, Env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, Env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, Env)
	}
	return NilValue
}

func Eval_Final_Module(pathname string) SLC_Object {
	byter, x := ioutil.ReadFile(pathname)
	//fmt.Println(string(byter))
	// modular debug
	if x != nil {
		return NewError("IO Error: Could not read module '%s' -> %s", pathname, x)
	}
	lex := LexNew(string(byter))
	parser := New_Parser(lex)
	mdod := parser.ParseProgram()
	if len(parser.Errors) != 0 {
		return NewError("Parser error: %s", parser.Errors)
	}
	env := NewEnvironment()
	Eval(mdod, env)
	return env.ExportedHash()
}

func EvalModule(pathname string) SLC_Object {
	_, x := os.Stat(pathname)
	if x != nil {
		return NewError("SkyLine Statistics: Could not get the file name to be parsed -> %s", fmt.Sprint(x))
	}
	return Eval_Final_Module(pathname)
}

func EvalImportingExpression(iexp *ImportExpression, Env *Environment_of_environment) SLC_Object {
	name := Eval(iexp.Name, Env)
	if isError(name) {
		return name
	}
	if s, ok := name.(*String); ok {
		attrs := EvalModule(s.Value)
		if isError(attrs) {
			return attrs
		}
		return &Module{Name: s.Value, Attrs: attrs}
	}
	return NewError("ImportError: Invalid import path -> %s", iexp.Name)

}

func SkyLine_Eval_Unit_Foreach(fle *ForeachStatement, env *Environment_of_environment) SLC_Object {
	val := Eval(fle.Value, env)
	helper, ok := val.(Iterable)
	if !ok {
		return NewError("%s object doesn't implement the Iterable interface", val.SL_RetrieveDataType())
	}
	var permit []string
	permit = append(permit, fle.Ident)
	if fle.Index != "" {
		permit = append(permit, fle.Index)
	}
	child := NewTempScop(env, permit)
	helper.Reset()
	ret, idx, ok := helper.Next()
	for ok {
		child.Set(fle.Ident, ret)
		idxName := fle.Index
		if idxName != "" {
			child.Set(fle.Index, idx)
		}
		rt := Eval(fle.Body, child)
		if !isError(rt) && (rt.SL_RetrieveDataType() == ReturnValueType || rt.SL_RetrieveDataType() == ErrorType) {
			return rt
		}
		ret, idx, ok = helper.Next()
	}
	return &Nil{}
}

// AST

// TO INTERFACE : ENGINE

var EngineCallVal EngineCallValues

func Assign() {
	EngineCallVal.Name = SLC.Exportable_data.ProjectData.Name
	EngineCallVal.SOS = SLC.Exportable_data.ProjectData.SuportedOS
	EngineCallVal.Description = SLC.Exportable_data.ProjectData.Description
	EngineCallVal.Languages = SLC.Exportable_data.ProjectData.Languages
	EngineCallVal.Require = SLC.Exportable_data.ProjectData.Require
	EngineCallVal.Version = SLC.Exportable_data.ProjectData.Version
	EngineCallVal.Prepped = true // Indicating this function was run
}

func EvalEngineCall(val SLC_Object) SLC_Object {
	switch Value := val.(type) {
	case *String:
		SLC.Start(Value.SL_InspectObject())
		Assign() // Call to assign
	default:
		return NewError("Sorry but the data type provided %s is not `String` as the ENGINE keyword requires a string value which is the filename SLC can parse", Value)
	}
	return &Nil{}
}

func SkyLine_Eval_Unit_ForLoop(loop *ForLoopExpression, Env *Environment_of_environment) SLC_Object {
	Val := &Boolean_Object{Value: true}
	for {
		condition := Eval(loop.Condition, Env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			newval := Eval(loop.Consequence, Env)
			if !isError(newval) && (newval.SL_RetrieveDataType() == ReturnValueType || newval.SL_RetrieveDataType() == ErrorType) {
				return newval
			}
		} else {
			break
		}
	}
	return Val
}
