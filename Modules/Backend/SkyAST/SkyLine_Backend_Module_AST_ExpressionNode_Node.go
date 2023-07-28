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

func (EN_I8 *SL_Integer8) SL_EN()                                       {} // AST Expression Note | 8 Bit Integer Alias
func (EN_I16 *SL_Integer16) SL_EN()                                     {} // AST Expression Note | 16 Bit Integer Alias
func (EN_I32 *SL_Integer32) SL_EN()                                     {} // AST Expression Note | 32 Bit Integer Alias
func (EN_I64 *SL_Integer64) SL_EN()                                     {} // AST Expression Note | 64 Bit Integer Alias
func (EN_Byte *SL_Byte) SL_EN()                                         {} // AST Expression Note | Byte alias
func (EN_Identifier *SL_Identifier) SL_EN()                             {} // AST Expression Node | Identifier
func (EN_String *SL_String) SL_EN()                                     {} // AST Expression Node | String Type
func (EN_Boolean *SL_Boolean) SL_EN()                                   {} // AST Expression Node | Boolean Type
func (EN_Integer *SL_Integer) SL_EN()                                   {} // AST Expression Node | Integer Type
func (EN_Float *SL_Float) SL_EN()                                       {} // AST Expression Node | Float Type
func (EN_Null *SL_NULL) SL_EN()                                         {} // AST Expression Node | Null Type
func (EN_Array *SL_Array) SL_EN()                                       {} // AST Expression Node | Array Type
func (EN_HashMap *SL_HashMap) SL_EN()                                   {} // AST Expression Node | Hash Map Type
func (EN_VariableAssignment *SL_EN_VariableAssignmentStatement) SL_EN() {} // AST Expression Node | Variable Assignment
func (EN_CaseExpr *SL_EN_Case_ExpressionStatement) SL_EN()              {} // AST Expression Node | Case in switch
func (EN_SwitchExpr *SL_EN_Switch_ExpressionStatement) SL_EN()          {} // AST Expression Node | switch decl
func (EN_Index *SL_EN_Index_Expression) SL_EN()                         {} // AST Expression Node | Index Expression
func (EN_ObjectCallExpr *SL_EN_Object_Call_Expression) SL_EN()          {} // AST Expression Node | Object Call
func (EN_FunctionCallExpr *SL_EN_Call_Expression) SL_EN()               {} // AST Expression Node | Function call
func (EN_FunctionDecl *SL_EN_Function_Definition) SL_EN()               {} // AST Expression Node | Function declaration
func (EN_FunctionLiteral *SL_EN_Function_Literal) SL_EN()               {} // AST Expression Node | Function literal
func (EN_ForLoop *SL_EN_Conditional_Loop) SL_EN()                       {} // AST Expression Node | For loop
func (EN_ForEachLoop *SL_EN_For_Each_Loop) SL_EN()                      {} // AST Expression Node | For each loop
func (EN_Conditional *SL_EN_Conditional_IfElse) SL_EN()                 {} // AST Expression Node | Conditional statement ( if else and else if )
func (EN_PostfixOP *SL_EN_Postfix) SL_EN()                              {} // AST Expression Node | Postfix Expression
func (EN_PrefixOP *SL_EN_Prefix) SL_EN()                                {} // AST Expression Node | Prefix Expression
func (EN_InfixOP *SL_EN_Infix) SL_EN()                                  {} // AST Expression Node | Infix Expression
func (EN_Import *SL_ImportExpression) SL_EN()                           {} // AST Expression Node | Import Expression
func (EN_Lable *SL_EN_Lable_JumpDef) SL_EN()                            {} // AST Expression Node | Lable definition (@lable) / goto

