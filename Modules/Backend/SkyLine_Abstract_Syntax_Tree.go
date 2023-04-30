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
// Filename      |  SkyLine_Abstract_Syntax_Tree_Expression_Nodes.go
// Project       |  SkyLine programming language
// Line Count    |  35 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
//
// Defines       | This file defines all of the AST functions within Skyline
//
//
package SkyLine_Backend

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x1]
//
// This section defines all of the AST expression node functions within Skyline
//
//
func (SL_PostFix *PostfixExpression) EN()                   {} // Expression Node    | Postfix Expressions (-- || ++)
func (SL_Foreach *ForeachStatement) EN()                    {} // Expression Node    | For each expression
func (SL_ForLoop *ForLoopExpression) EN()                   {} // Expression Node    | For Loop expression
func (SL_ModularImport *ImportExpression) EN()              {} // Expression Node    | Import expression
func (SL_Identifier *Ident) EN()                            {} // Expression Node    | Identifier
func (SL_IntegerLiteral *IntegerLiteral) EN()               {} // Expression Node    | Integer Literal
func (SL_FloatLiteral *FloatLiteral) EN()                   {} // Expression Node    | Float Literal
func (SL_PrefixExpression *PrefixExpression) EN()           {} // Expression Node    | Prefix Expression
func (SL_InfixExpression *InfixExpression) EN()             {} // Expression Node    | Infix Expression
func (SL_BooleanValue *Boolean_AST) EN()                    {} // Expression Node    | Boolean
func (SL_ConditionalExpression *ConditionalExpression) EN() {} // Expression Node    | Conditional Expression
func (SL_UnitBlockStatement *BlockStatement) EN()           {} // Expression Node    | Block
func (SL_FunctionLiteral *FunctionLiteral) EN()             {} // Expression Node    | Function Literal
func (SL_FunctionCallExpression *CallExpression) EN()       {} // Expression Node    | Call
func (SL_StringLiteral *StringLiteral) EN()                 {} // Expression Node    | String
func (SL_ArrayLiteral *ArrayLiteral) EN()                   {} // Expression Node    | Array Literal
func (SL_IndexExpression *IndexExpression) EN()             {} // Expression Node    | Index expression
func (SL_HashLiteral *HashLiteral) EN()                     {} // Expression Node    | Hash Literal
func (SL_CaseExpression *Case) EN()                         {} // Expression Node    | Case Expression
func (SL_SwitchExpression *Switch) EN()                     {} // Expression Node    | Switch Expression
func (SL_VariableAssignment *AssignmentStatement) EN()      {} // Expression Node    | Assign Expression
func (SL_ObjectCallExpression *ObjectCallExpression) EN()   {} // Expression Node    | SLC_Object Call

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x2]
//
// This section contains information around definitions for the statement nodes within the AST.
//
//

func (SL_VariableDecleration *LetStatement) SN()        {} // Statement Node     | Allow condition
func (SL_ReturnStatement *ReturnStatement) SN()         {} // Statement Node     | Allow Return
func (SL_ExpressionStatement *ExpressionStatement) SN() {} // Statement Node     | Allow Expression
func (SL_ConstantVariableDecleration *Constant) SN()    {} // Statement Node     | Allow Constants
func (SL_RegistryFunction *Register) SN()               {} // Statement Node     | Allow register
func (SL_ENGINE *ENGINE) SN()                           {} // Statement Node     | Allow engine call

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x3]
//
// The code below is there to define the retrievedatatype function apart of the Type_Object structure. It is there to help define and return the data type of a value
//
//

func (SL_ENGINE_VAL *ENGINE_Value) SL_RetrieveDataType() Type_Object        { return TOKEN_ENGINE_TYPE } // SLC_Object interface   | Engine value data type
func (SL_ModuleDataType *Module) SL_RetrieveDataType() Type_Object          { return TOKEN_MODULE }      // SLC_Object interface   | Module data type
func (SL_DataTypeIntger *Integer) SL_RetrieveDataType() Type_Object         { return IntegerType }       // SLC_Object Interface   | Integer type object
func (SL_DataTypeFloat *Float) SL_RetrieveDataType() Type_Object            { return FloatType }         // SLC_Object Interface   | Float type object
func (SL_DataTypeBoolean *Boolean_Object) SL_RetrieveDataType() Type_Object { return BooleanType }       // SLC_Object Interface   | Boolean type object
func (SL_DataTypeNULL *Nil) SL_RetrieveDataType() Type_Object               { return NilType }           // SLC_Object Interface   | Null type object
func (SL_DataTypeReturn *ReturnValue) SL_RetrieveDataType() Type_Object     { return ReturnValueType }   // SLC_Object Interface   | return value type object
func (SL_DataTypeError *Error) SL_RetrieveDataType() Type_Object            { return ErrorType }         // SLC_Object Interface   | Error type object
func (SL_DataTypeFunction *Function) SL_RetrieveDataType() Type_Object      { return FunctionType }      // SLC_Object Interface   | Function type object
func (SL_DataTypeString *String) SL_RetrieveDataType() Type_Object          { return StringType }        // SLC_Object Interface   | String type object
func (SL_DataTypeBuiltin *Builtin) SL_RetrieveDataType() Type_Object        { return BuiltinType }       // SLC_Object Interface   | Built in function type object
func (SL_DataTypeArray *Array) SL_RetrieveDataType() Type_Object            { return ArrayType }         // SLC_Object Interface   | Array type object
func (SL_DataTypeHash *Hash) SL_RetrieveDataType() Type_Object              { return HashType }          // SLC_Object Interface   | Hash type object
func (SL_RegisterValue *RegisterValue) SL_RetrieveDataType() Type_Object    { return ImportingType }     // SLC_Object Interface   | Register type object

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x4]
//
// This next code section defines the SkyLine_ExtractNodeValue which is a function to get the literal value of a given node or type,
//
//

func (fes *ForeachStatement) SL_ExtractNodeValue() string {
	return fes.Token.Literal
} // Token Literal        | Returns the token literal for the values within foreach
func (SL_ForLoop *ForLoopExpression) SL_ExtractNodeValue() string {
	return SL_ForLoop.Token.Literal
} // Token Literal        | Returns the token literal of a for loop
func (SL_Engine *ENGINE) SL_ExtractNodeValue() string {
	return SL_Engine.Token.Literal
} // Token Literal        | Returns the token literal value of engine
func (SL_VariableDecleration *LetStatement) SL_ExtractNodeValue() string {
	return SL_VariableDecleration.Token.Literal
} // Token Literal 	      | Returns literal of an allow statement
func (SL_Identifier *Ident) SL_ExtractNodeValue() string {
	return SL_Identifier.Token.Literal
} // Token Literal 	      | Returns literal of an identifier
func (SL_ReturnStatement *ReturnStatement) SL_ExtractNodeValue() string {
	return SL_ReturnStatement.Token.Literal
} // Token Literal 	      | Returns literal of a return statement
func (SL_ExpressionStatement *ExpressionStatement) SL_ExtractNodeValue() string {
	return SL_ExpressionStatement.Token.Literal
} // Token Literal 	      | Returns literal of expression statement
func (SL_IntegerLiteral *IntegerLiteral) SL_ExtractNodeValue() string {
	return SL_IntegerLiteral.Token.Literal
} // Token Literal 	      | Returns literal of integer
func (SL_FloatLiteral *FloatLiteral) SL_ExtractNodeValue() string {
	return SL_FloatLiteral.Token.Literal
} // Token Literal 	      | Returns literal of a float
func (SL_PrefixExpression *PrefixExpression) SL_ExtractNodeValue() string {
	return SL_PrefixExpression.Token.Literal
} // Token Literal 	      | Returns literal of PrefixExpression
func (SL_InfixExpression *InfixExpression) SL_ExtractNodeValue() string {
	return SL_InfixExpression.Token.Literal
} // Token Literal 	      | Returns literal of InfixExpression
func (SL_PostFixExpression *PostfixExpression) SL_ExtractNodeValue() string {
	return SL_PostFixExpression.Token.Literal
} // Token Literal        | Returns literal of a postfix expression
func (SL_BooleanValue *Boolean_AST) SL_ExtractNodeValue() string {
	return SL_BooleanValue.Token.Literal
} // Token Literal 	      | Returns literal of Boolean statement
func (SL_ConditionalExpression *ConditionalExpression) SL_ExtractNodeValue() string {
	return SL_ConditionalExpression.Token.Literal
} // Token Literal 	   	  | Returns literal of Conditionals
func (SL_UnitBlockStatement *BlockStatement) SL_ExtractNodeValue() string {
	return SL_UnitBlockStatement.Token.Literal
} // Token Literal 	   	  | Returns literal of code block statements
func (SL_FunctionLiteral *FunctionLiteral) SL_ExtractNodeValue() string {
	return SL_FunctionLiteral.Token.Literal
} // Token Literal 	      | Returns literal of function
func (SL_FunctionCallExpression *CallExpression) SL_ExtractNodeValue() string {
	return SL_FunctionCallExpression.Token.Literal
} // Token Literal 	      | Returns literal of CallExpression
func (SL_SwitchExpression *Switch) SL_ExtractNodeValue() string {
	return SL_SwitchExpression.Token.Literal
} // Token Literal 	      | Returns literal of switch
func (SL_CaseExpression *Case) SL_ExtractNodeValue() string {
	return SL_CaseExpression.Token.Literal
} // Token Literal 	      | Returns literal of Case
func (SL_VariableAssignment *AssignmentStatement) SL_ExtractNodeValue() string {
	return SL_VariableAssignment.Token.Literal
} // Token Literal        | Returns literal of assignment statement
func (SL_ConstantVariableDecleration *Constant) SL_ExtractNodeValue() string {
	return SL_ConstantVariableDecleration.Token.Literal
} // Token Literal  	  | Returns literal of a constant
func (SL_ObjectCallExpression *ObjectCallExpression) SL_ExtractNodeValue() string {
	return SL_ObjectCallExpression.Token.Literal
} // Token Literal        | Returns literal of an object call
func (SL_RegistryFunction *Register) SL_ExtractNodeValue() string {
	return SL_RegistryFunction.Token.Literal
} // Token Literal        | Returns literal of registry list

func (SL_ImportExpression *ImportExpression) SL_ExtractNodeValue() string {
	return SL_ImportExpression.Token.Literal
} // Token Literal        | Returns literal of import token

func (SL_StringLiteral *StringLiteral) SL_ExtractNodeValue() string {
	if SL_StringLiteral == nil {
		return ""
	}
	return SL_StringLiteral.Token.Literal
}

func (SL_ArrayLiteral *ArrayLiteral) SL_ExtractNodeValue() string {
	if SL_ArrayLiteral == nil {
		return ""
	}
	return SL_ArrayLiteral.Token.Literal
}

func (SL_IndexExpression *IndexExpression) SL_ExtractNodeValue() string {
	if SL_IndexExpression == nil {
		return ""
	}
	return SL_IndexExpression.Token.Literal
}

func (SL_HashLiteral *HashLiteral) SL_ExtractNodeValue() string {
	if SL_HashLiteral == nil {
		return ""
	}
	return SL_HashLiteral.Token.Literal
}

func (SL_Program *Program) SL_ExtractNodeValue() string {
	if len(SL_Program.Statements) == 0 {
		return ""
	}
	return SL_Program.Statements[0].SL_ExtractNodeValue()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x5]
//
// This code section defines the inspect functions to grab the values of a specific object or return the inspect
//
//

func (ENGINE_VAL *ENGINE_Value) SL_InspectObject() string {
	return ENGINE_VAL.Value.SL_InspectObject()
} // Methodize Inspect  | Inspect the value of the engine call ENGINE() <-
func (SL_BuiltInFunction *Builtin) SL_InspectObject() string {
	return "builtin function"
} // Methodize Inspect  | Inspect built in function call
func (SL_DataTypeString *String) SL_InspectObject() string {
	return SL_DataTypeString.Value
} //  Methodize Inspect | Inspect string data type or value
func (SL_DataTypeError *Error) SL_InspectObject() string {
	return SL_DataTypeError.Message
} //  Methodize Inspect | Inspect error type
func (SL_DataTypeNULL *Nil) SL_InspectObject() string {
	return ""
} //  Methodize Inspect | Inspect empty/NULL/0x00 value
func (SL_DataTypeFloat *Float) SL_InspectObject() string {
	return strconv.FormatFloat(SL_DataTypeFloat.Value, 'f', -1, 64)
} //  Methodize Inspect | Inspect Float types
func (SL_DataTypeInteger *Integer) SL_InspectObject() string {
	return strconv.FormatInt(SL_DataTypeInteger.Value, 10)
} //  Methodize Inspect | Inspect integer types
func (SL_DataTypeBoolean *Boolean_Object) SL_InspectObject() string {
	return strconv.FormatBool(SL_DataTypeBoolean.Value)
} //  Methodize Inspect | Inspect boolean values
func (SL_RegistryValue *RegisterValue) SL_InspectObject() string {
	return SL_RegistryValue.Value.SL_InspectObject()
} //  Methodize Inspect | Inspect register values
func (SL_RetrunValue *ReturnValue) SL_InspectObject() string {
	return SL_RetrunValue.Value.SL_InspectObject()
} //  Methodize Inspect | Inspect return values

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x6]
//
// This next code section defines the ExtractStringValue function which helps with better defining the Skyline syntax upon inspecting the object.
//
//

func (m *Module) SL_ExtractStringValue() string {
	return m.SL_InspectObject()
}

func (SL_StringLiteral *StringLiteral) SL_ExtractStringValue() string {
	return SL_StringLiteral.SL_ExtractNodeValue()
}

func (SL_BooleanValue *Boolean_AST) SL_ExtractStringValue() string {
	return SL_BooleanValue.Token.Literal
}

func (SL_FloatLiteral *FloatLiteral) SL_ExtractStringValue() string {
	return SL_FloatLiteral.Token.Literal
}

func (SL_IntegerLiteral *IntegerLiteral) SL_ExtractStringValue() string {
	return SL_IntegerLiteral.Token.Literal
}

func (SL_Identifier *Ident) SL_ExtractStringValue() string {
	return SL_Identifier.Value
}

func (SL_HashLiteral *HashLiteral) SL_ExtractStringValue() string {
	if SL_HashLiteral == nil {
		return ""
	}
	pairs := make([]string, len(SL_HashLiteral.Pairs))
	for key, value := range SL_HashLiteral.Pairs {
		pairs = append(pairs, key.SL_ExtractStringValue()+": "+value.SL_ExtractStringValue()+"\n")
	}
	var Out bytes.Buffer
	Out.WriteString("{")
	Out.WriteString(strings.Join(pairs, ", "))
	Out.WriteString("}")
	return Out.String()
}

func (SL_IndexExpression *IndexExpression) SL_ExtractStringValue() string {
	if SL_IndexExpression == nil {
		return ""
	}
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(SL_IndexExpression.Left.SL_ExtractStringValue())
	Out.WriteString("[")
	Out.WriteString(SL_IndexExpression.Index.SL_ExtractStringValue())
	Out.WriteString("])")
	return Out.String()
}

func (SL_PostFix *PostfixExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(SL_PostFix.Token.Literal)
	Out.WriteString(SL_PostFix.Operator)
	Out.WriteString(")")
	return Out.String()
}

func (SL_Foreach *ForeachStatement) SL_ExtractStringValue() string {
	var out bytes.Buffer
	out.WriteString("foreach ")
	out.WriteString(SL_Foreach.Ident)
	out.WriteString(" ")
	out.WriteString(SL_Foreach.Value.SL_ExtractStringValue())
	out.WriteString(SL_Foreach.Body.SL_ExtractStringValue())
	return out.String()
}

func (SL_ArrayLiteral *ArrayLiteral) SL_ExtractStringValue() string {
	if SL_ArrayLiteral == nil {
		return ""
	}
	elements := make([]string, 0, len(SL_ArrayLiteral.Elements))
	for _, el := range SL_ArrayLiteral.Elements {
		elements = append(elements, el.SL_ExtractStringValue())
	}
	var Out bytes.Buffer
	Out.WriteString("[")
	Out.WriteString(strings.Join(elements, ", "))
	Out.WriteString("]")
	return Out.String()
}

func (SL_RegisterValue *Register) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("register (")
	Out.WriteString(SL_RegisterValue.RegistryValue.SL_ExtractStringValue())
	Out.WriteString(")")
	return Out.String()
}

func (SL_ObjectCallExpression *CallExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	args := make([]string, 0, len(SL_ObjectCallExpression.Arguments))
	for _, arg := range SL_ObjectCallExpression.Arguments {
		args = append(args, arg.SL_ExtractStringValue())
	}
	Out.WriteString(SL_ObjectCallExpression.Function.SL_ExtractStringValue())
	Out.WriteString("(")
	Out.WriteString(strings.Join(args, ", "))
	Out.WriteString(")")
	return Out.String()
}

func (ENGINEK *ENGINE) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("ENGINE|")
	Out.WriteString(`"` + ENGINEK.EngineValue.SL_ExtractStringValue() + `"`)
	Out.WriteString("|")
	return Out.String()
}

func (SL_FunctionLiteral *FunctionLiteral) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	params := make([]string, 0, len(SL_FunctionLiteral.Parameters))
	for _, param := range SL_FunctionLiteral.Parameters {
		params = append(params, param.SL_ExtractStringValue())
	}
	Out.WriteString(SL_FunctionLiteral.SL_ExtractNodeValue())
	Out.WriteString("(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") ")
	Out.WriteString(SL_FunctionLiteral.Body.SL_ExtractStringValue())
	return Out.String()
}

func (SL_UnitBlockStatement *BlockStatement) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	for _, s := range SL_UnitBlockStatement.Statements {
		Out.WriteString(s.SL_ExtractStringValue())
	}
	return Out.String()
}

func (SL_ConditionalExpression *ConditionalExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("if")
	Out.WriteString(SL_ConditionalExpression.Condition.SL_ExtractStringValue())
	Out.WriteString(" ")
	Out.WriteString(SL_ConditionalExpression.Consequence.SL_ExtractStringValue())
	if SL_ConditionalExpression.Alternative != nil {
		Out.WriteString("else ")
		Out.WriteString(SL_ConditionalExpression.Alternative.SL_ExtractStringValue())
	}
	return Out.String()
}

func (SL_InfixExpression *InfixExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(SL_InfixExpression.Left.SL_ExtractStringValue())
	Out.WriteString(" " + SL_InfixExpression.Operator + " ")
	Out.WriteString(SL_InfixExpression.Right.SL_ExtractStringValue())
	Out.WriteString(")")
	return Out.String()
}

func (SL_PrefixExpression *PrefixExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(SL_PrefixExpression.Operator)
	Out.WriteString(SL_PrefixExpression.Right.SL_ExtractStringValue())
	Out.WriteString(")")
	return Out.String()
}

func (SL_ExpressionStatement *ExpressionStatement) SL_ExtractStringValue() string {
	if SL_ExpressionStatement.Expression != nil {
		return SL_ExpressionStatement.Expression.SL_ExtractStringValue()
	}
	return ""
}

func (SL_ConstantVariableDecleration *Constant) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(SL_ConstantVariableDecleration.SL_ExtractNodeValue() + " ")
	Out.WriteString(SL_ConstantVariableDecleration.Name.SL_ExtractNodeValue())
	Out.WriteString(" = ")
	if SL_ConstantVariableDecleration.Value != nil {
		Out.WriteString(SL_ConstantVariableDecleration.Value.SL_ExtractStringValue())
	}
	Out.WriteString(";")
	return Out.String()
}

func (SL_ReturnStatement *ReturnStatement) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(SL_ReturnStatement.SL_ExtractNodeValue() + " ")
	if SL_ReturnStatement.ReturnValue != nil {
		Out.WriteString(SL_ReturnStatement.ReturnValue.SL_ExtractStringValue())
	}
	Out.WriteString(";")
	return Out.String()
}

func (SL_VariableAssignment *AssignmentStatement) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(SL_VariableAssignment.Name.SL_ExtractStringValue())
	Out.WriteString(SL_VariableAssignment.Operator)
	Out.WriteString(SL_VariableAssignment.Value.SL_ExtractStringValue())
	return Out.String()
}

func (SL_VariableDecleration *LetStatement) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(SL_VariableDecleration.SL_ExtractNodeValue() + " ")
	Out.WriteString(SL_VariableDecleration.Name.SL_ExtractStringValue())
	Out.WriteString(" = ")
	if SL_VariableDecleration.Value != nil {
		Out.WriteString(SL_VariableDecleration.Value.SL_ExtractStringValue())
	}
	Out.WriteString(";")
	return Out.String()
}

func (SL_Program *Program) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	for _, s := range SL_Program.Statements {
		Out.WriteString(s.SL_ExtractStringValue())
	}
	return Out.String()
}

func (SL_SwitchExpression *Switch) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("\nswitch ( ")
	Out.WriteString(SL_SwitchExpression.Value.SL_ExtractStringValue())
	Out.WriteString(")\n{\n")
	for _, bod := range SL_SwitchExpression.Choices {
		if bod != nil {
			Out.WriteString(bod.SL_ExtractStringValue())
		}
	}
	Out.WriteString("}\n")
	return Out.String()
}

func (SL_CaseExpression *Case) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	if SL_CaseExpression.Def {
		Out.WriteString("defualt")
	} else {
		Out.WriteString("case ")
		bod := []string{}
		for _, body := range SL_CaseExpression.Expr {
			bod = append(bod, body.SL_ExtractStringValue())
		}
		Out.WriteString(strings.Join(bod, ","))
	}
	Out.WriteString(SL_CaseExpression.Block.SL_ExtractStringValue())
	return Out.String()
}

func (SL_ObjectCallExpression *ObjectCallExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(SL_ObjectCallExpression.SLC_Object.SL_ExtractStringValue())
	Out.WriteString(".")
	Out.WriteString(SL_ObjectCallExpression.Call.SL_ExtractStringValue())
	return Out.String()
}

func (SL_Function *Function) SL_InspectObject() string {
	var Out bytes.Buffer
	params := make([]string, 0, len(SL_Function.Parameters))
	for _, parser := range SL_Function.Parameters {
		params = append(params, parser.SL_ExtractStringValue())
	}
	Out.WriteString(fmt.Sprint(SL_Function.SL_RetrieveDataType()) + "(")
	Out.WriteString(strings.Join(params, ", "))
	Out.WriteString(") {")
	Out.WriteString(SL_Function.Body.SL_ExtractStringValue())
	Out.WriteString("}")
	return Out.String()
}

func (SL_Array *Array) SL_InspectObject() string {
	if SL_Array == nil {
		return ""
	}
	elements := make([]string, 0, len(SL_Array.Elements))
	for _, e := range SL_Array.Elements {
		elements = append(elements, e.SL_InspectObject())
	}
	var Out bytes.Buffer
	Out.WriteString("")
	Out.WriteString("[")
	Out.WriteString(strings.Join(elements, ", "))
	Out.WriteString("]")
	return Out.String()
}

func (SL_Hash *Hash) SL_InspectObject() string {
	if SL_Hash == nil {
		return ""
	}
	pairs := make([]string, 0, len(SL_Hash.Pairs))
	for _, pair := range SL_Hash.Pairs {
		pairs = append(pairs, pair.Key.SL_InspectObject()+": "+pair.Value.SL_InspectObject())
	}
	var Out bytes.Buffer
	Out.WriteString("{")
	Out.WriteString(strings.Join(pairs, ", "))
	Out.WriteString("}")
	return Out.String()
}

func (SL_ImportExpression *ImportExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString(SL_ImportExpression.SL_ExtractNodeValue())
	Out.WriteString("(")
	Out.WriteString(fmt.Sprintf("\"%s\"", SL_ImportExpression.Name))
	return Out.String()
}

func (SL_ForLoop *ForLoopExpression) SL_ExtractStringValue() string {
	var Out bytes.Buffer
	Out.WriteString("for (")
	Out.WriteString(SL_ForLoop.Condition.SL_ExtractStringValue())
	Out.WriteString(") {")
	Out.WriteString(SL_ForLoop.Consequence.SL_ExtractStringValue())
	Out.WriteString("}")
	return Out.String()
}

func (m *Module) SL_InspectObject() string { return fmt.Sprintf("<module '%s'>", m.Name) }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x7]
//
// This section defines the hash key of specific data types such as numerical data types or alpha data types.
//
//

func (SL_DataTypeInteger *Integer) HashKey() HashKey {
	return HashKey{
		Type_Object: SL_DataTypeInteger.SL_RetrieveDataType(),
		Value:       uint64(SL_DataTypeInteger.Value),
	}
}

func (SL_DataTypeFloat *Float) HashKey() HashKey {
	s := strconv.FormatFloat(SL_DataTypeFloat.Value, 'f', -1, 64)
	h := fnv.New64a()
	h.Write([]byte(s))
	return HashKey{
		Type_Object: SL_DataTypeFloat.SL_RetrieveDataType(),
		Value:       h.Sum64(),
	}
}

func (SL_DataTypeBoolean *Boolean_Object) HashKey() HashKey {
	key := HashKey{Type_Object: SL_DataTypeBoolean.SL_RetrieveDataType()}
	if SL_DataTypeBoolean.Value {
		key.Value = 1
	}
	return key
}

func (SL_DataTypeString *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(SL_DataTypeString.Value))
	return HashKey{
		Type_Object: SL_DataTypeString.SL_RetrieveDataType(),
		Value:       h.Sum64(),
	}
}
