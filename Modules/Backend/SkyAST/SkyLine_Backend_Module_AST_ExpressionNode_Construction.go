///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_AST_ExpressionNode_Construction
// Extension         | .go ( golang source code file )
// Purpose           | Defines all Expression nodes (AST)
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// One important part of the AST is to finally construct the full syntax of a given statement or expression. In order to do so, each node within the AST ( statement or expression )
//
// must have a function that can help it better define that output and structure the code to read it into parsing. This will help also when outputting and formatting error messages
//
// in cases like missing semicolons. You can use the AST for also fixing statements and telling users what it should look like versus what they used for a more robust error system.
//
package SkyLine_Backend_Abstract_Synatax_Tree

import (
	"bytes"
	"fmt"
	"strings"
)

func (EN_Identifier *SL_Identifier) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_Identifier.Value
} // AST Expression Node | Identifier

func (EN_Byte *SL_Byte) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_Byte.TokenConstruct.Literal
}

func (EN_I8 *SL_Integer8) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_I8.TokenConstruct.Literal
} // AST Expression Node | Integer 8
func (EN_I16 *SL_Integer16) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_I16.TokenConstruct.Literal
} // AST Expression Node | Integer 16
func (EN_I32 *SL_Integer32) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_I32.TokenConstruct.Literal
} // AST Expression Node | Integer 32
func (EN_I64 *SL_Integer64) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_I64.TokenConstruct.Literal
} // AST Expression Node | Integer 64

func (EN_String *SL_String) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_String.TokenConstruct.Literal
} // AST Expression Node | String Type

func (EN_Boolean *SL_Boolean) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_Boolean.TokenConstruct.Literal
} // AST Expression Node | Boolean Type

func (EN_Integer *SL_Integer) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_Integer.TokenConstruct.Literal
} // AST Expression Node | Integer Type

func (EN_Float *SL_Float) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_Float.TokenConstruct.Literal
} // AST Expression Node | Float Type

func (EN_Null *SL_NULL) SkyLine_NodeInterface_Get_Node_Value() string {
	return EN_Null.TokenConstruct.Literal
} // AST Expression Node | Null Type

func (EN_Array *SL_Array) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	ArrayElements := make([]string, 0)
	for _, element := range EN_Array.Elements {
		ArrayElements = append(ArrayElements, element.SkyLine_NodeInterface_Get_Node_Value())
	}
	Out.WriteString("[")
	Out.WriteString(strings.Join(ArrayElements, ", "))
	Out.WriteString("];")
	return Out.String()
} // AST Expression Node | Array Type

func (EN_HashMap *SL_HashMap) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Pairs := make([]string, 0)
	for Key, Val := range EN_HashMap.Pairs {
		Pairs = append(Pairs, Key.SkyLine_NodeInterface_Get_Node_Value()+" , "+Val.SkyLine_NodeInterface_Get_Node_Value())
	}
	Out.WriteString("map\n")
	for _, n := range Pairs {
		Out.WriteString(n + ",")
	}
	Out.WriteString("\n};")
	return Out.String()
} // AST Expression Node | Hash Map Type

func (EN_VariableAssignment *SL_EN_VariableAssignmentStatement) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString(EN_VariableAssignment.Name.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(EN_VariableAssignment.Operator)
	Out.WriteString(EN_VariableAssignment.Value.SkyLine_NodeInterface_Get_Node_Value())
	return Out.String()
} // AST Expression Node | Variable Assignment

func (EN_CaseExpr *SL_EN_Case_ExpressionStatement) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	if EN_CaseExpr.Default {
		Out.WriteString("default")
	} else {
		Out.WriteString("case")
		UNIT_TMP := []string{}
		for _, expression := range EN_CaseExpr.Expression {
			UNIT_TMP = append(UNIT_TMP, expression.SkyLine_NodeInterface_Get_Node_Value())
		}
		Out.WriteString(strings.Join(UNIT_TMP, ","))
	}
	Out.WriteString(EN_CaseExpr.Unit.SkyLine_NodeInterface_Get_Node_Value())
	return Out.String()
} // AST Expression Node | Case in switch

func (EN_SwitchExpr *SL_EN_Switch_ExpressionStatement) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("switch (")
	Out.WriteString(EN_SwitchExpr.Value.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("){")
	for _, Unit := range EN_SwitchExpr.Conditions {
		if Unit != nil {
			Out.WriteString(Unit.SkyLine_NodeInterface_Get_Node_Value())
		} else {
			println("Switch case units should not be empty")
			//TODO: Implement rule for NULL units
		}
	}
	Out.WriteString("};")
	return Out.String()
} // AST Expression Node | switch decl

func (EN_Index *SL_EN_Index_Expression) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(EN_Index.Left.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("[")
	Out.WriteString(EN_Index.Index.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("]);")
	return Out.String()
} // AST Expression Node | Index expression

func (EN_ObjectCallExpr *SL_EN_Object_Call_Expression) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString(EN_ObjectCallExpr.Object.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(".")
	Out.WriteString(EN_ObjectCallExpr.Call.SkyLine_NodeInterface_Get_Node_Value())
	return Out.String()
} // AST Expression Node | Object Call

func (EN_FunctionCallExpr *SL_EN_Call_Expression) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Params := make([]string, 0)
	for _, param := range EN_FunctionCallExpr.Arguments {
		Params = append(Params, param.SkyLine_NodeInterface_Get_Node_Value())
	}
	Out.WriteString(EN_FunctionCallExpr.Function.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("(")
	Out.WriteString(strings.Join(Params, ","))
	Out.WriteString(")")
	return Out.String()
} // AST Expression Node | Function call

func (EN_LableDecl *SL_EN_Lable_JumpDef) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("@")
	Out.WriteString(EN_LableDecl.SkyLine_NodeInterface_Token_Literal())
	Out.WriteString("{\n")
	Out.WriteString(EN_LableDecl.Unit.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("\n};")
	return Out.String()
}

func (EN_FunctionDecl *SL_EN_Function_Definition) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Params := make([]string, 0)
	for _, param := range EN_FunctionDecl.FunctionArguments {
		Params = append(Params, param.SkyLine_NodeInterface_Get_Node_Value())
	}
	Out.WriteString(EN_FunctionDecl.SkyLine_NodeInterface_Token_Literal())
	Out.WriteString("(")
	Out.WriteString(strings.Join(Params, ", "))
	Out.WriteString(") ")
	Out.WriteString(EN_FunctionDecl.Unit.SkyLine_NodeInterface_Get_Node_Value())
	return Out.String()
} // AST Expression Node | Function declaration

func (EN_FunctionLiteral *SL_EN_Function_Literal) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Params := make([]string, 0)
	for _, param := range EN_FunctionLiteral.Parameters {
		Params = append(Params, param.SkyLine_NodeInterface_Get_Node_Value())
	}
	Out.WriteString(EN_FunctionLiteral.SkyLine_NodeInterface_Token_Literal())
	Out.WriteString("(")
	Out.WriteString(strings.Join(Params, ", "))
	Out.WriteString(") { ")
	Out.WriteString(EN_FunctionLiteral.Unit.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("};")
	return Out.String()

} // AST Expression Node | Function literal

func (EN_ForLoop *SL_EN_Conditional_Loop) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("for (")
	Out.WriteString(EN_ForLoop.Condition.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(") {")
	Out.WriteString(EN_ForLoop.Consequence.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString("};")
	return Out.String()
} // AST Expression Node | For loop

func (EN_ForEachLoop *SL_EN_For_Each_Loop) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("foreach ")
	Out.WriteString(EN_ForEachLoop.Identifier)
	Out.WriteString(" ")
	Out.WriteString(EN_ForEachLoop.Value.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(EN_ForEachLoop.Unit.SkyLine_NodeInterface_Get_Node_Value())
	return Out.String()
} // AST Expression Node | For each loop

func (EN_Conditional *SL_EN_Conditional_IfElse) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("if")
	Out.WriteString(EN_Conditional.Condition.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(" ")
	Out.WriteString(EN_Conditional.Consequence_Unit.SkyLine_NodeInterface_Get_Node_Value())
	if EN_Conditional.Alternative_Unit != nil {
		Out.WriteString("else")
		Out.WriteString(EN_Conditional.Alternative_Unit.SkyLine_NodeInterface_Get_Node_Value())
	}
	return Out.String()
} // AST Expression Node | Conditional statement ( if else and else if )

func (EN_PostfixOP *SL_EN_Postfix) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(EN_PostfixOP.TokenConstruct.Literal)
	Out.WriteString(EN_PostfixOP.Operator)
	Out.WriteString(")")
	return Out.String()
} // AST Expression Node | Postfix expression

func (EN_PrefixOP *SL_EN_Prefix) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(EN_PrefixOP.Operator)
	Out.WriteString(EN_PrefixOP.Right.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(")")
	return Out.String()
} // AST Expression Node | Prefix expression

func (EN_InfixOP *SL_EN_Infix) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("(")
	Out.WriteString(EN_InfixOP.Left.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(" " + EN_InfixOP.Operator + " ")
	Out.WriteString(EN_InfixOP.Right.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(")")
	return Out.String()
} // AST Expression Node | Infix Expression

func (ImportExpression *SL_ImportExpression) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("import")
	Out.WriteString("(\n")
	var count int
	for i, n := range ImportExpression.Name {
		count++
		if count > i {
			Out.WriteString(fmt.Sprintf("\"%s\"\n", n))

		} else {
			Out.WriteString(fmt.Sprintf("\"%s\",\n", n))
		}
	}
	Out.WriteString(");")
	return Out.String()
}

func (EN_Prog *SL_Prog) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	for _, statement := range EN_Prog.ProgramStatements {
		Out.WriteString(statement.SkyLine_NodeInterface_Get_Node_Value())
	}
	return Out.String()
}
