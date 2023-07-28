///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_AST_StatementNode_Construction
// Extension         | .go ( golang source code file )
// Purpose           | Defines all statement nodes (AST)
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
	"strings"
)

func (SN_Constant *Assignment_Constant_Const) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString(SN_Constant.SkyLine_NodeInterface_Token_Literal() + " ")
	Out.WriteString(SN_Constant.Name.SkyLine_NodeInterface_Token_Literal())
	Out.WriteString(" := ")
	if SN_Constant.Value != nil {
		Out.WriteString(SN_Constant.Name.SkyLine_NodeInterface_Get_Node_Value())
	} else {
		Out.WriteString("null")
	}
	Out.WriteString(";")
	return Out.String()
} // AST Statement Node | Constant statement
func (SN_VariableDecl *Assignment_Cause_Set_Allow) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString(SN_VariableDecl.SkyLine_NodeInterface_Token_Literal() + " ")
	Out.WriteString(SN_VariableDecl.Name.SkyLine_NodeInterface_Token_Literal())
	Out.WriteString(" := ")
	if SN_VariableDecl.Value != nil {
		Out.WriteString(SN_VariableDecl.Value.SkyLine_NodeInterface_Get_Node_Value())
	} else {
		Out.WriteString("null")
	}
	Out.WriteString(";")
	return Out.String()
} // AST Statement Node | Variable declaration
func (SN_Return *Return_Ret_Return_Information) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString(SN_Return.SkyLine_NodeInterface_Token_Literal() + " ")
	if SN_Return.Expression != nil {
		Out.WriteString(SN_Return.Expression.SkyLine_NodeInterface_Get_Node_Value())
	} else {
		Out.WriteString("null")
	}
	Out.WriteString(";")
	return Out.String()
} // AST Statement Node | Return statement
func (SN_Expression *Expression_Statement) SkyLine_NodeInterface_Get_Node_Value() string {
	if SN_Expression.Expression != nil {
		return SN_Expression.SkyLine_NodeInterface_Get_Node_Value()
	}
	return ""
} // AST Statement Node | Expression statement
func (SN_UnitBlock *SL_UnitBlockStatement) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	for _, s := range SN_UnitBlock.Statements {
		Out.WriteString(s.SkyLine_NodeInterface_Get_Node_Value())
	}
	return Out.String()
} // AST Statement Node | Unit block

func (SN_Register *SL_Register) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	var values []string
	for _, s := range SN_Register.RegistryValue {
		values = append(values, s.SkyLine_NodeInterface_Get_Node_Value())
	}
	Out.WriteString("register(")
	Out.WriteString(strings.Join(values, ","))
	Out.WriteString(")")
	return Out.String()
}

func (SN_SLC *SL_ENGINE) SkyLine_NodeInterface_Get_Node_Value() string {
	var Out bytes.Buffer
	Out.WriteString("ENGINE(")
	Out.WriteString(SN_SLC.EngineValue.SkyLine_NodeInterface_Get_Node_Value())
	Out.WriteString(");")
	return Out.String()
}
