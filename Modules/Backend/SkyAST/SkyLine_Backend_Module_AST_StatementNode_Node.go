///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_AST_StatementNode_Node
// Extension         | .go ( golang source code file )
// Purpose           | Defines all statement nodes (AST)
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file helps define all of the statement node functions or rather nodes for all statements within SkyLine.
//
//
//
package SkyLine_Backend_Abstract_Synatax_Tree

func (SN_Constant *Assignment_Constant_Const) SL_SN()      {} // AST Statement Node | Constant statement
func (SN_VariableDecl *Assignment_Cause_Set_Allow) SL_SN() {} // AST Statement Node | Variable declaration
func (SN_Return *Return_Ret_Return_Information) SL_SN()    {} // AST Statement Node | Return statement
func (SN_Expression *Expression_Statement) SL_SN()         {} // AST Statement Node | Expression statement
func (SN_UnitBlock *SL_UnitBlockStatement) SL_SN()         {} // AST Statement Node | Unit block
func (SN_Register *SL_Register) SL_SN()                    {} // AST Statement Node | Register statement
func (SN_SLC *SL_ENGINE) SL_SN()                           {} // AST Statement Node | Engine statement
