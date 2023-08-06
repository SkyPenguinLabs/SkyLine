////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
//
// The SkyLine configuration language is a language and engine designed to act as a modification language to the SkyLine programming language. This language is
// very minimal and contains a regex base lexer, a very basic parser, a few keywords, a base interpreter and that is all as well as some backend engine code. This
// language is purely modified to be an extension to the SkyLine programming language, something that can be a pre processor language post processing for the main
// SkyLine script. Below is more technical information for the language
//
// Lexer       : Regex based lexer with minimal tokens and base syntax
// Parser      : Base parser with minimal tokens and base syntax with simple error systems
// REPL        : Does not exist
// Environment : Extremely minimal
// Types       : String, Boolean, Integer
// Statements  : set, import, use, errors, output, system, constant/const
//
// File contains -> This file contains all of the abstract syntax tree MODELS rather than functions, these models belong to the AST and allow the
//                  AST to work as it needs to when it contains a primary tree with a prime node and a leaf node
//
//	TokenConstructLiteral() string
//	TokenConstructToString() string
//

package SkyLine_Configuration_Engine_Backend_Source

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x1]
//
// This section of code defines all AST based functions for the engine program data structure and data interface
//
//
//
//

func (Prog *Engine_Prog) TokenConstructLiteral() string {
	if len(Prog.Statements) > 0 {
		return Prog.Statements[0].TokenConstructLiteral()
	} else {
		return ""
	}
}

func (Prog *Engine_Prog) TokenConstructToString() string {
	var Out bytes.Buffer
	for _, statement := range Prog.Statements {
		Out.WriteString(statement.TokenConstructToString())
	}
	return Out.String()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x2]
//
// This next section and brick of code will define all of the data type based AST functions which includes String, Error, NULL, Builtin, Error, Array, Boolean
//
// and Integer data types. These functions allow us to convert the data to a string and convert a base syntax. This defines all expression leaf nodes
//

func (AstString *StringDataType_Expression_AbstractSyntaxTree) ExpressionLeafNode()   {}
func (AstBoolean *BooleanDataType_Expression_AbstractSyntaxTree) ExpressionLeafNode() {}
func (AstInteger *IntegerDataType_Expression_AbstractSyntaxTree) ExpressionLeafNode() {}
func (AstArray *ArrayLiteral_Expression_AbstractSyntaxTree) ExpressionLeafNode()      {}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x2]->[0x3]
//
//
// This next section of code defines all of the other EXPRESSION nodes for different data types and other various calls or methods within the language

func (AstPrefix *PrefixExpression_Expression_AbstractSyntaxTree) ExpressionLeafNode() {}
func (AstInfix *InfixExpression_Expression_AbstractSyntaxTree) ExpressionLeafNode()   {}
func (AstCall *CallFunction_Expression_AbstractSyntaxTree) ExpressionLeafNode()       {}
func (AstIndex *IndexLit_Expression_AbstractSyntaxTree) ExpressionLeafNode()          {}
func (AstEngineINIT *INIT_Expression_AbstractSyntaxTree) ExpressionLeafNode()         {}
func (AstEngineENGINE *ENGINE_Expression_AbstractSyntaxTree) ExpressionLeafNode()     {}
func (AstIdentifier *Identifier_Expression_AbstractSyntaxTree) ExpressionLeafNode()   {}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x3]::[0x2]->[0x3]
//
//
// This code section defines all of the functions that are STATEMENTS within the AST such as constants, assignments, block, expressions and other various types

func (AstExpression *Expression_Statement_AbstractSyntaxTree) SatatementLeafNode()  {}
func (AstAssignment *Assignment_Statement_AbstractSyntaxTree) SatatementLeafNode()  {}
func (AstConstant *Constant_Statement_AbstractSyntaxTree) SatatementLeafNode()      {}
func (AstUnitSec *BlockStatement_Statement_AbstractSyntaxTree) SatatementLeafNode() {}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x04]
//
//
// Instead of just declaring nodes and different STATMENTS OR EXPRESSION functions, this section of code defines the actual Token Literal functions that allows us
//
// to grab the literal of the token.
//
//

func (AstString *StringDataType_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstString.TokenRegister.Literal
}
func (AstBoolean *BooleanDataType_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstBoolean.TokenRegister.Literal
}
func (AstInteger *IntegerDataType_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstInteger.TokenRegister.Literal
}
func (AstAssignment *Assignment_Statement_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstAssignment.TokenRegister.Literal
}
func (AstConstant *Constant_Statement_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstConstant.TokenRegister.Literal
}
func (AstArray *ArrayLiteral_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstArray.TokenRegister.Literal
}
func (AstIndexExp *IndexLit_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstIndexExp.TokenRegister.Literal
}
func (AstExpression *Expression_Statement_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstExpression.TokenRegister.Literal
}
func (AstInfixExp *InfixExpression_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstInfixExp.TokenRegister.Literal
}
func (AstPrefixExp *PrefixExpression_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstPrefixExp.TokenRegister.Literal
}
func (AstIdentifier *Identifier_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstIdentifier.TokenRegister.Literal
}
func (AstUnitSec *BlockStatement_Statement_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstUnitSec.TokenRegister.Literal
}
func (AstCallExp *CallFunction_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstCallExp.TokenRegister.Literal
}
func (AstENGINE *ENGINE_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstENGINE.TokenRegister.Literal
}
func (AstINIT *INIT_Expression_AbstractSyntaxTree) TokenConstructLiteral() string {
	return AstINIT.TokenRegister.Literal
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//[0x05]
//
// This next brick and unit of code allows us to define all tras-construct to string functions which convert the structures values into a syntax of strings
//
//

func (AstString *StringDataType_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	return AstString.TokenRegister.Literal
}
func (AstBoolean *BooleanDataType_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	return AstBoolean.TokenRegister.Literal
}
func (AstInteger *IntegerDataType_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	return AstInteger.TokenRegister.Literal
}
func (AstAssignment *Assignment_Statement_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString(AstAssignment.TokenConstructLiteral() + " ")
	Out.WriteString(AstAssignment.Name.TokenConstructToString())
	Out.WriteString(" := ")
	if AstAssignment.Value != nil {
		Out.WriteString(AstAssignment.Value.TokenConstructToString())
	} else {
		Out.WriteString("NULL")
	}
	Out.WriteString(";")
	return Out.String()
}
func (AstConstant *Constant_Statement_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString(AstConstant.TokenConstructLiteral() + " ")
	Out.WriteString(AstConstant.Name.TokenConstructLiteral())
	Out.WriteString(" := ")
	if AstConstant.Value != nil {
		Out.WriteString(AstConstant.Value.TokenConstructToString())
	} else {
		Out.WriteString(" NULL")
	}
	Out.WriteString(";")
	return Out.String()
}
func (AstArray *ArrayLiteral_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	elements := []string{}
	for _, elem := range AstArray.Elements {
		elements = append(elements, elem.TokenConstructToString())
	}
	Out.WriteString("modify[")
	Out.WriteString(strings.Join(elements, ", "))
	Out.WriteString("]")
	return Out.String()
}

func (AstIndexExp *IndexLit_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(AstIndexExp.Left.TokenConstructToString())
	Out.WriteString("[")
	Out.WriteString(AstIndexExp.Index.TokenConstructToString())
	Out.WriteString("])")
	return Out.String()
}

func (AstInfixExp *InfixExpression_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(AstInfixExp.Left.TokenConstructToString())
	Out.WriteString(" " + AstInfixExp.Operator + " ")
	Out.WriteString(AstInfixExp.Right.TokenConstructToString())
	Out.WriteString(")")
	return Out.String()
}

func (AstPrefixExp *PrefixExpression_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(AstPrefixExp.Operator)
	Out.WriteString(AstPrefixExp.Right.TokenConstructToString())
	Out.WriteString(")")
	return Out.String()
}

func (AstUnitSec *BlockStatement_Statement_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	defer func() {
		if x := recover(); x != nil {
			const SLC_Evaluator_NULL_Statements = "370"
			Message := CallErrorStr(
				fmt.Sprint(SLC_Evaluator_NULL_Statements),
				"The engine does not allow code to be run without a statement like system(SYS) -> [mod]...",
				"Ensure statements exist, this is a fatal error (NULL)",
			)
			println(Message)
			os.Exit(0)
		}
	}()
	if AstUnitSec.Statements == nil {
		Out.WriteString("NULL")
	} else {
		for _, statement := range AstUnitSec.Statements {
			Out.WriteString(statement.TokenConstructToString())
		}
	}
	return Out.String()
}

func (AstCallExp *CallFunction_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	arguments := []string{}
	for _, arg := range AstCallExp.Arguments {
		arguments = append(arguments, arg.TokenConstructToString())
	}
	Out.WriteString(AstCallExp.Function.TokenConstructToString())
	Out.WriteString("(")
	Out.WriteString(strings.Join(arguments, ", "))
	Out.WriteString(")")
	return Out.String()
}

func (AstENGINE *ENGINE_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString("\nENGINE {\n")
	for _, cs := range AstENGINE.SubUnits {
		if cs != nil {
			Out.WriteString(cs.TokenConstructToString())
		}
	}
	Out.WriteString("}::END_UNIT;\n")
	return Out.String()
}

func (AstINIT *INIT_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	var Out bytes.Buffer
	Out.WriteString("INIT ")
	CS := []string{}
	for _, expression := range AstINIT.Expression {
		CS = append(CS, expression.TokenConstructToString())
	}
	Out.WriteString(strings.Join(CS, ", "))
	Out.WriteString(AstINIT.Sub_UNIT.TokenConstructToString())
	return Out.String()
}

func (AstIdentifier *Identifier_Expression_AbstractSyntaxTree) TokenConstructToString() string {
	return AstIdentifier.Value
}

func (AstExpression *Expression_Statement_AbstractSyntaxTree) TokenConstructToString() string {
	if AstExpression.Expression != nil {
		return AstExpression.Expression.TokenConstructToString()
	}
	return ""
}
