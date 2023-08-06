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
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the functions that will help retrieve the literal value of tokens within the statement nodes. This will help during parsing and evaluation of statements
//
// and also helps in some cases for extra functionality.
//
package SkyLine_Backend_Abstract_Synatax_Tree

func (EN_Identifier *SL_Identifier) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Identifier.TokenConstruct.Literal
} // AST Expression Node | Identifier
func (EN_Byte *SL_Byte) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Byte.TokenConstruct.Literal
} // AST Expression Node | Byte | Character
func (EN_String *SL_String) SkyLine_NodeInterface_Token_Literal() string {
	return EN_String.TokenConstruct.Literal
} // AST Expression Node | String Type
func (EN_Boolean *SL_Boolean) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Boolean.TokenConstruct.Literal
} // AST Expression Node | Boolean Type
func (EN_Integer *SL_Integer) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Integer.TokenConstruct.Literal
} // AST Expression Node | Integer Type
func (EN_I8 *SL_Integer8) SkyLine_NodeInterface_Token_Literal() string {
	return EN_I8.TokenConstruct.Literal
} // AST Expression Node | Integer 8
func (EN_I16 *SL_Integer16) SkyLine_NodeInterface_Token_Literal() string {
	return EN_I16.TokenConstruct.Literal
} // AST Expression Node | Integer 16
func (EN_I32 *SL_Integer32) SkyLine_NodeInterface_Token_Literal() string {
	return EN_I32.TokenConstruct.Literal
} // AST Expression Node | Integer 32
func (EN_I64 *SL_Integer64) SkyLine_NodeInterface_Token_Literal() string {
	return EN_I64.TokenConstruct.Literal
} // AST Expression Node | Integer 64
func (EN_Float *SL_Float) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Float.TokenConstruct.Literal
} // AST Expression Node | Float Type
func (EN_Null *SL_NULL) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Null.TokenConstruct.Literal
} // AST Expression Node | Null Type
func (EN_Array *SL_Array) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Array.TokenConstruct.Literal
} // AST Expression Node | Array Type
func (EN_HashMap *SL_HashMap) SkyLine_NodeInterface_Token_Literal() string {
	return EN_HashMap.TokenConstruct.Literal
} // AST Expression Node | Hash Map Type
func (EN_VariableAssignment *SL_EN_VariableAssignmentStatement) SkyLine_NodeInterface_Token_Literal() string {
	return EN_VariableAssignment.TokenConstruct.Literal
} // AST Expression Node | Variable Assignment
func (EN_CaseExpr *SL_EN_Case_ExpressionStatement) SkyLine_NodeInterface_Token_Literal() string {
	return EN_CaseExpr.TokenConstruct.Literal
} // AST Expression Node | Case in switch
func (EN_SwitchExpr *SL_EN_Switch_ExpressionStatement) SkyLine_NodeInterface_Token_Literal() string {
	return EN_SwitchExpr.TokenConstruct.Literal
} // AST Expression Node | switch decl
func (EN_Index *SL_EN_Index_Expression) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Index.TokenConstruct.Literal
} // AST Expression Node | Index expression
func (EN_ObjectCallExpr *SL_EN_Object_Call_Expression) SkyLine_NodeInterface_Token_Literal() string {
	return EN_ObjectCallExpr.TokenConstruct.Literal
} // AST Expression Node | Object Call
func (EN_FunctionCallExpr *SL_EN_Call_Expression) SkyLine_NodeInterface_Token_Literal() string {
	return EN_FunctionCallExpr.TokenConstruct.Literal
} // AST Expression Node | Function call
func (EN_FunctionDecl *SL_EN_Function_Definition) SkyLine_NodeInterface_Token_Literal() string {
	return EN_FunctionDecl.TokenConstruct.Literal
} // AST Expression Node | Function declaration
func (EN_FunctionLiteral *SL_EN_Function_Literal) SkyLine_NodeInterface_Token_Literal() string {
	return EN_FunctionLiteral.TokenConstruct.Literal
} // AST Expression Node | Function literal
func (EN_ForLoop *SL_EN_Conditional_Loop) SkyLine_NodeInterface_Token_Literal() string {
	return EN_ForLoop.TokenConstruct.Literal
} // AST Expression Node | For loop
func (EN_ForEachLoop *SL_EN_For_Each_Loop) SkyLine_NodeInterface_Token_Literal() string {
	return EN_ForEachLoop.TokenConstruct.Literal
} // AST Expression Node | For each loop
func (EN_Conditional *SL_EN_Conditional_IfElse) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Conditional.TokenConstruct.Literal
} // AST Expression Node | Conditional statement ( if else and else if )
func (EN_PostfixOP *SL_EN_Postfix) SkyLine_NodeInterface_Token_Literal() string {
	return EN_PostfixOP.TokenConstruct.Literal
} // AST Expression Node | Postfix expression
func (EN_PrefixOP *SL_EN_Prefix) SkyLine_NodeInterface_Token_Literal() string {
	return EN_PrefixOP.TokenConstruct.Literal
} // AST Expression Node | Prefix expression
func (EN_InfixOP *SL_EN_Infix) SkyLine_NodeInterface_Token_Literal() string {
	return EN_InfixOP.TokenConstruct.Literal
} // AST Expression Node | Infix Expression
func (EN_Program *SL_Prog) SkyLine_NodeInterface_Token_Literal() string {
	if len(EN_Program.ProgramStatements) > 0 {
		return EN_Program.ProgramStatements[0].SkyLine_NodeInterface_Token_Literal()
	}
	return ""
} // AST Expression Node | Program
func (EN_Import *SL_ImportExpression) SkyLine_NodeInterface_Token_Literal() string {
	return EN_Import.TokenConstruct.Literal
}

func (ENLable *SL_EN_Lable_JumpDef) SkyLine_NodeInterface_Token_Literal() string {
	return ENLable.TokenConstruct.Literal
}
