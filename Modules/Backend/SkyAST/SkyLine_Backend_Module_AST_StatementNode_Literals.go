///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_AST_StatementNode_Literals
// Extension         | .go ( golang source code file )
// Purpose           | Defines all statement nodes (AST)
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
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

func (SN_Constant *Assignment_Constant_Const) SkyLine_NodeInterface_Token_Literal() string {
	return SN_Constant.TokenConstruct.Literal
} // AST Statement Node | Constant statement
func (SN_VariableDecl *Assignment_Cause_Set_Allow) SkyLine_NodeInterface_Token_Literal() string {
	return SN_VariableDecl.TokenConstruct.Literal
} // AST Statement Node | Variable declaration
func (SN_Return *Return_Ret_Return_Information) SkyLine_NodeInterface_Token_Literal() string {
	return SN_Return.TokenConstruct.Literal
} // AST Statement Node | Return statement
func (SN_Expression *Expression_Statement) SkyLine_NodeInterface_Token_Literal() string {
	return SN_Expression.TokenConstruct.Literal
} // AST Statement Node | Expression statement
func (SN_UnitBlock *SL_UnitBlockStatement) SkyLine_NodeInterface_Token_Literal() string {
	return SN_UnitBlock.TokenConstruct.Literal
} // AST Statement Node | Unit block
func (SN_Register *SL_Register) SkyLine_NodeInterface_Token_Literal() string {
	return SN_Register.TokenConstruct.Literal
} // AST Statement Node | Register statement
func (SN_SLC *SL_ENGINE) SkyLine_NodeInterface_Token_Literal() string {
	return SN_SLC.TokenConstruct.Literal
} // AST Statement Node | Engine statement
