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
// Filename      |  Skyline_Evaluator_Evaluate_Registry_Function.go
// Project       |  SkyLine programming language
// Line Count    |  150+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all of the evaluation functions that are being called or in other words the main evaluation step and call for the language.
//                 This function takes a node input value and checks its types then calls the appropriate functions tied to the resulting data type.
//
// STATE         | Needs to be organized and worked on
// Resolution    | Functions need to be renamed to be SL specific, functions need to be organized, functions need to be automated.
//
package SkyLine_Backend

import (
	"fmt"
	"os"
)

func Check() bool {
	return EngineCallVal.Prepped
}

func Eval(node Node, Env *Environment_of_environment) SLC_Object {
	if Check() {
		Env.Set(EngineCallVal.Name+"Version", &String{Value: EngineCallVal.Version})  // Set version
		Env.Set(EngineCallVal.Name+"Name", &String{Value: EngineCallVal.Name})        // Set Name
		Env.Set(EngineCallVal.Name+"Desc", &String{Value: EngineCallVal.Description}) // Set Description
		Env.Set(EngineCallVal.Name+"Supported", &String{Value: EngineCallVal.SOS})    // Set Supported operating systems
		for _, lib := range EngineCallVal.Require {
			if regf, ok := RegisterStandard[lib]; ok {
				regf()
			}
		}
		// Syntax: ProjectNameVersion
		// Reset to false, so loop does not repeat
		EngineCallVal.Prepped = false
	}
	switch node := node.(type) {
	case *PostfixExpression:
		return EvalPostfixExpression(Env, node.Operator, node)
	case *ForeachStatement:
		return SkyLine_Eval_Unit_Foreach(node, Env)
	case *ForLoopExpression:
		return SkyLine_Eval_Unit_ForLoop(node, Env)
	case *Program:
		return evalProgram(node, Env)
	case *ExpressionStatement:
		return Eval(node.Expression, Env)
	case *ReturnStatement:
		value := Eval(node.ReturnValue, Env)
		if isError(value) {
			return value
		}
		return &ReturnValue{Value: value}
	case *Constant:
		value := Eval(node.Value, Env)
		if isError(value) {
			return value
		}
		Env.Set(node.Name.Value, value)
		return value
	case *ObjectCallExpression:
		res := Evaluate_ObjectCall(node, Env)
		if isError(res) {
			fmt.Fprintf(os.Stderr, "Error calling method: %s", res.SL_InspectObject())
		}
		return res
	case *ImportExpression:
		return EvalImportingExpression(node, Env)
	case *BlockStatement:
		return Eval_Block_Statement(node, Env)
	case *Switch:
		return EvalSwitch(node, Env)
	case *LetStatement:
		value := Eval(node.Value, Env)
		if isError(value) {
			return value
		}
		Env.Set(node.Name.Value, value)
	case *AssignmentStatement:
		return EvalAssignModify(node, Env)
	case *IntegerLiteral:
		return &Integer{Value: node.Value}

	case *FloatLiteral:
		return &Float{Value: node.Value}

	case *Boolean_AST:
		return nativeBoolToBooleanObject(node.Value)

	case *PrefixExpression:
		right := Eval(node.Right, Env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *InfixExpression:
		left := Eval(node.Left, Env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, Env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ConditionalExpression:
		return Eval_Conditional(node, Env)
	case *Register:
		value := Eval(node.RegistryValue, Env)
		if isError(value) {
			return value
		}
		return EvalRegisterCall(value)
	case *ENGINE:
		value := Eval(node.EngineValue, Env)
		if isError(value) {
			return value
		}
		//C1::CHANGE
		return EvalEngineCall(value)
	case *Ident:
		return evalIdent(node, Env)

	case *FunctionLiteral:
		return &Function{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        Env,
		}
	case *CallExpression:
		function := Eval(node.Function, Env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, Env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(Env, function, args)

	case *StringLiteral:
		return &String{Value: node.Value}

	case *ArrayLiteral:
		elems := evalExpressions(node.Elements, Env)
		if len(elems) == 1 && isError(elems[0]) {
			return elems[0]
		}
		return &Array{Elements: elems}

	case *IndexExpression:
		left := Eval(node.Left, Env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, Env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	case *HashLiteral:
		return evalHashLiteral(node, Env)
	}
	return nil
}
