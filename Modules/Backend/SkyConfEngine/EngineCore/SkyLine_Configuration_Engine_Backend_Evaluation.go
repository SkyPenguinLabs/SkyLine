package SkyLine_Configuration_Engine_Backend_Source

import (
	"fmt"
	"strings"
)

func Eval(node AbstractSyntaxTree_Node, env *Engine_Environment_Of_Environment) SLC_Object {
	switch node := node.(type) {
	case *ENGINE_Expression_AbstractSyntaxTree:
		return SkyLine_Configuration_Engine_Evaluate_Iniation(node, env)
	case *Engine_Prog:
		return SkyLine_Configuration_Engine_Evaluate_Program(node, env)
	case *BlockStatement_Statement_AbstractSyntaxTree:
		return SkyLine_Configuration_Engine_Evaluate_BlockStatement(node, env)
	case *Expression_Statement_AbstractSyntaxTree:
		return Eval(node.Expression, env)
	case *Constant_Statement_AbstractSyntaxTree:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.StoreConstant(node.Name.Value, val)
		return val
	case *Identifier_Expression_AbstractSyntaxTree:
		return SkyLine_Configuration_Engine_Evaluate_Identifier(node, env)
	case *CallFunction_Expression_AbstractSyntaxTree:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := SkyLine_Configuration_Engine_Evaluate_Expressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return SkyLine_Configuration_Engine_Evaluate_ApplicationalFunction(function, args)
	case *ArrayLiteral_Expression_AbstractSyntaxTree:
		elements := SkyLine_Configuration_Engine_Evaluate_Expressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &ObjectArray{Elements: elements}
	case *IndexLit_Expression_AbstractSyntaxTree:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return SkyLine_Configuration_Engine_Evaluate_IndexExpression(left, index)
	case *Assignment_Statement_AbstractSyntaxTree:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Engine_Set_Varname(node.Name.Value, val)
	case *IntegerDataType_Expression_AbstractSyntaxTree:
		return &ObjectInteger{Value: node.Value}
	case *StringDataType_Expression_AbstractSyntaxTree:
		return &ObjectString{Value: node.Value}
	case *BooleanDataType_Expression_AbstractSyntaxTree:
		return nativeBoolToBooleanObject(node.Value)
	case *InfixExpression_Expression_AbstractSyntaxTree:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return SkyLine_Configuration_Engine_Evaluate_InfixExpression(node.Operator, left, right)
	}

	return nil
}

func SkyLine_Configuration_Engine_Evaluate_Program(program *Engine_Prog, env *Engine_Environment_Of_Environment) SLC_Object {
	var result SLC_Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *ObjectERROR:
			return result
		}
	}

	return result
}

func SkyLine_Configuration_Engine_Evaluate_BlockStatement(block *BlockStatement_Statement_AbstractSyntaxTree, env *Engine_Environment_Of_Environment) SLC_Object {
	var result SLC_Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *ObjectBoolean {
	if input {
		return TRUE
	}
	return FALSE
}

func SkyLine_Configuration_Engine_Evaluate_InfixExpression(
	operator string,
	left, right SLC_Object,
) SLC_Object {
	switch {
	case left.ObjectDataType() == STRING_OBJ && right.ObjectDataType() == STRING_OBJ:
		return SkyLine_Configuration_Engine_Evaluate_StringInfixExpression(operator, left, right)
	case left.ObjectDataType() == STRING_OBJ && right.ObjectDataType() == ARRAY_OBJ:
		return SkyLine_Configuration_Engine_Evaluate_StringInfixExpression(operator, left, right)
	case left.ObjectDataType() != right.ObjectDataType():
		statement := fmt.Sprintf("%s %s %s", left.ObjectInspectFunc(), operator, right.ObjectInspectFunc())
		Message := CallErrorStr(
			fmt.Sprint(SLC_Evaluator_DataType_Mismatch),
			"Data type mismatch",
			fmt.Sprintf("%s %s %s", left.ObjectDataType(), operator, right.ObjectDataType())+" in("+statement+")",
		)
		return newError(Message)
	default:
		statement := fmt.Sprintf("%s %s %s", left.ObjectInspectFunc(), operator, right.ObjectInspectFunc())
		Message := CallErrorStr(
			fmt.Sprint(SLC_Evaluator_Unknown_Token_Operator),
			fmt.Sprintf("Unknown operator (%s)", operator),
			fmt.Sprintf("%s %s %s", left.ObjectDataType(), operator, right.ObjectDataType())+" in("+statement+")",
		)
		return newError(Message)
	}
}

type EvaluationForModification struct {
	System         string
	SystemSettings string
}

func SkyLine_Configuration_Engine_Evaluate_SysModify(left string, right []SLC_Object) bool {
	if strings.ToLower(left) == "errors" {
		if len(right) > 0 {
			fmt.Println("Argument for system modify ", left, " = ", right[0].ObjectInspectFunc())
			return true
		} else {
			return true
		}
	} else {
		return false
	}
}

func SkyLine_Configuration_Engine_Evaluate_StringInfixExpression(operator string, left, right SLC_Object) SLC_Object {
	if operator != "+" && operator != "->" {
		statement := fmt.Sprintf("%s %s %s", left.ObjectInspectFunc(), operator, right.ObjectInspectFunc())
		Message := CallErrorStr(
			fmt.Sprint(SLC_Evaluator_Unknown_Token_Operator),
			fmt.Sprintf("Unknown operator (%s)", operator),
			fmt.Sprintf("%s %s %s", left.ObjectDataType(), operator, right.ObjectDataType())+" in("+statement+")",
		)
		return newError(Message)
	}
	if operator == "->" {
		leftval := left.(*ObjectString).Value
		rightVal := right.(*ObjectArray).Elements
		return &ObjectBoolean{Value: SkyLine_Configuration_Engine_Evaluate_SysModify(leftval, rightVal)}
	} else {
		leftVal := left.(*ObjectString).Value
		rightVal := right.(*ObjectString).Value
		return &ObjectString{Value: leftVal + rightVal}
	}
}

func SkyLine_Configuration_Engine_Evaluate_Identifier(node *Identifier_Expression_AbstractSyntaxTree, env *Engine_Environment_Of_Environment) SLC_Object {
	if val, ok := env.Engine_Grab_Varname(node.Value); ok {
		return val
	}

	if builtin, ok := Builtins[node.Value]; ok {
		return builtin
	}
	Message := CallErrorStr(
		fmt.Sprint(SLC_Evaluator_Identifier_Not_Found),
		fmt.Sprintf("Unknown IDENTIFIER (%s)", node.Value),
		fmt.Sprintf("%s", node),
	)
	return newError(Message)
}

func newError(format string, a ...interface{}) *ObjectERROR {
	return &ObjectERROR{Message: fmt.Sprintf(format, a...)}
}

func isError(obj SLC_Object) bool {
	if obj != nil {
		return obj.ObjectDataType() == ERROR_OBJ
	}
	return false
}

func SkyLine_Configuration_Engine_Evaluate_Expressions(exps []AbstractSyntaxTree_Expression, env *Engine_Environment_Of_Environment) []SLC_Object {
	var result []SLC_Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []SLC_Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func SkyLine_Configuration_Engine_Evaluate_ApplicationalFunction(fn SLC_Object, args []SLC_Object) SLC_Object {
	switch fn := fn.(type) {

	case *ObjectBUILTINFUNCTION:
		return fn.Function(args...)

	default:
		Message := CallErrorStr(
			fmt.Sprint(SLC_Evaluator_ValueCall_Not_A_Function),
			"Node value is not a function, stop calling it -> "+fmt.Sprintf("called as %s but is really of type %s ", fn, fn.ObjectDataType()),
			""+fn.ObjectInspectFunc(),
		)
		return newError(Message)
	}
}

func SkyLine_Configuration_Engine_Evaluate_IndexExpression(left, index SLC_Object) SLC_Object {
	switch {
	case left.ObjectDataType() == ARRAY_OBJ && index.ObjectDataType() == INTEGER_OBJ:
		return SkyLine_Configuration_Engine_Evaluate_ArrayIndexExpression(left, index)
	default:
		Message := CallErrorStr(
			fmt.Sprint(SLC_Evaluator_Array_Index_Operator_Unsupported),
			fmt.Sprintf("Index operator not supported (%s)", left.ObjectDataType()),
			fmt.Sprintf("with index (%s) and inspect at %s ", index.ObjectInspectFunc(), left.ObjectInspectFunc()),
		)
		return newError(Message)
	}
}

func SkyLine_Configuration_Engine_Evaluate_ArrayIndexExpression(array, index SLC_Object) SLC_Object {
	arr := array.(*ObjectArray)
	idx := index.(*ObjectInteger).Value
	max := int64(len(arr.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arr.Elements[idx]
}

func SkyLine_Configuration_Engine_Evaluate_Iniation(se *ENGINE_Expression_AbstractSyntaxTree, env *Engine_Environment_Of_Environment) SLC_Object {
	obj := Eval(se.Value, env)
	for _, opt := range se.SubUnits {
		if opt.Default {
			continue
		}
		for _, val := range opt.Expression {
			out := Eval(val, env)
			if obj.ObjectDataType() == out.ObjectDataType() &&
				(obj.ObjectInspectFunc() == out.ObjectInspectFunc()) {
				blockOut := SkyLine_Configuration_Engine_Evaluate_BlockStatement(opt.Sub_UNIT, env)
				return blockOut
			}
		}
	}
	for _, opt := range se.SubUnits {
		if opt.Default {
			out := SkyLine_Configuration_Engine_Evaluate_BlockStatement(opt.Sub_UNIT, env)
			return out
		}
	}
	return nil
}
