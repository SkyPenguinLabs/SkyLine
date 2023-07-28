///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_AST_Models
// Extension         | .go ( golang source code file )
// Purpose           | Defines all variables and types for the AST ( Abstract Syntax Tree )
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the models that the AST will use. The AST will also be heavily used with our language to form specific statements like let, set, cause, f etc
//
// and even be there to load programs, statements, and expressions or to rather define functions to construct the sytax of them for the parser and evaluator
//
package SkyLine_Backend_Abstract_Synatax_Tree

import SL_Tokenization_Definitions "SkyLine/Modules/Backend/SkyTokens"

type (
	//::::::::::::::::::::::::::::::::::::::::::::::::::
	// Abstract Syntax Tree Node Structures | 1st module
	//::::::::::::::::::::::::::::::::::::::::::::::::::
	SL_Node interface {
		SkyLine_NodeInterface_Token_Literal() string
		SkyLine_NodeInterface_Get_Node_Value() string
	}

	SL_Statement interface {
		SL_Node
		SL_SN() // Statement node
	}

	SL_Expression interface {
		SL_Node
		SL_EN() // Expression node
	}

	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	// Abstract Syntax Tree Programatic Structures | 2nd module
	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	SL_Prog struct {
		ProgramStatements []SL_Statement
	}

	//:::::::::::::::::::::::::::::::::::::::::::::
	// Abstract Syntax Tree Dependants | 3rd module
	//:::::::::::::::::::::::::::::::::::::::::::::
	// Sub Info: Dependants is a function, variable, type, etc that
	// relies on another type such as expression which depends on statement

	SL_Identifier struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          string
	}

	//:::::::::::::::::::::::::::::::::::::::::::::
	// Abstract Syntax Tree Data Types | 3rd module
	//:::::::::::::::::::::::::::::::::::::::::::::

	SL_Byte struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          byte
	}

	SL_NULL struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
	}

	SL_Boolean struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          bool
	}

	SL_String struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          string
	}

	SL_Integer8 struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          int8
	}
	SL_Integer16 struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          int16
	}
	SL_Integer32 struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          int32
	}
	SL_Integer64 struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          int64
	}

	SL_Integer struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          int
	}

	SL_Float struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          float64
	}

	SL_Array struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Elements       []SL_Expression
	}

	SL_HashMap struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Pairs          map[SL_Expression]SL_Expression
	}

	//:::::::::::::::::::::::::::::::::::::::::::::
	// Abstract Syntax Tree Statements | 4th module
	//:::::::::::::::::::::::::::::::::::::::::::::
	// Requires :
	//          | SL_SN()
	//          |------| Node_
	//                 | 	SkyLine_NodeInterface_Token_Literal()  string
	//                      SkyLine_NodeInterface_Get_Node_Value() string

	// cause x = 10 ;
	// let   x = 10 ;
	// set   x = 10 ;
	Assignment_Cause_Set_Allow struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Name           *SL_Identifier
		Value          SL_Expression
	}
	// constant x = 10;
	// const    x = 10;
	Assignment_Constant_Const struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Name           *SL_Identifier
		Value          SL_Expression
	}

	// return (10 / 20) * 20
	// ret    (10 / 20) * 20
	Return_Ret_Return_Information struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Expression     SL_Expression
	}

	// Expression
	Expression_Statement struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Expression     SL_Expression
	}

	// Block | {}
	SL_UnitBlockStatement struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Statements     []SL_Statement
	}

	// Register statement | register(LIBNAME)
	SL_Register struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		RegistryValue  []SL_Expression
	}

	// ENGINE statement | ENGINE()
	SL_ENGINE struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		EngineValue    SL_Expression
	}

	//::::::::::::::::::::::::::::::::::::::::::::::
	// Abstract Syntax Tree Expressions | 4th module
	//::::::::::::::::::::::::::::::::::::::::::::::
	// Requires :
	//          | SL_EN()
	//          |------| Node_
	//                 | 	SkyLine_NodeInterface_Token_Literal()  string
	//                      SkyLine_NodeInterface_Get_Node_Value() string

	// Import expression
	SL_ImportExpression struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Name           []SL_Expression
	}

	// set x = 10;
	// x += 20;
	SL_EN_VariableAssignmentStatement struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Name           *SL_Identifier
		Operator       string
		Value          SL_Expression
	}

	// case condition {}
	SL_EN_Case_ExpressionStatement struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Default        bool
		Expression     []SL_Expression
		Unit           *SL_UnitBlockStatement
	}

	// switch(condition)
	SL_EN_Switch_ExpressionStatement struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Value          SL_Expression
		Conditions     []*SL_EN_Case_ExpressionStatement
	}

	// [1]
	SL_EN_Index_Expression struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Left           SL_Expression
		Index          SL_Expression
	}

	// Object.Object()
	SL_EN_Object_Call_Expression struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Object         SL_Expression
		Call           SL_Expression
	}

	// Object()
	SL_EN_Call_Expression struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Function       SL_Expression
		Arguments      []SL_Expression
	}

	//@lable { ... }
	//called with : jmp lable
	SL_EN_Lable_JumpDef struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Unit           *SL_UnitBlockStatement
	}

	// define function()
	SL_EN_Function_Definition struct {
		TokenConstruct    SL_Tokenization_Definitions.SL_TokenConstruct
		FunctionArguments []*SL_Identifier
		Defaults          map[string]SL_Expression
		Unit              *SL_UnitBlockStatement
	}

	// Func, function
	SL_EN_Function_Literal struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Parameters     []*SL_Identifier
		Defaults       map[string]SL_Expression
		Unit           *SL_UnitBlockStatement
	}

	// for ... in ... {}
	SL_EN_Conditional_Loop struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Condition      SL_Expression
		Consequence    *SL_UnitBlockStatement
	}

	// foreach x, y in ... {}
	SL_EN_For_Each_Loop struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Index          string
		Identifier     string
		Value          SL_Expression
		Unit           *SL_UnitBlockStatement
	}

	// if, else, else if
	SL_EN_Conditional_IfElse struct {
		TokenConstruct   SL_Tokenization_Definitions.SL_TokenConstruct
		Condition        SL_Expression
		Consequence_Unit *SL_UnitBlockStatement
		Alternative_Unit *SL_UnitBlockStatement
	}

	SL_EN_Postfix struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Operator       string
	}

	SL_EN_Infix struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Left           SL_Expression
		Operator       string
		Right          SL_Expression
	}

	SL_EN_Prefix struct {
		TokenConstruct SL_Tokenization_Definitions.SL_TokenConstruct
		Operator       string
		Right          SL_Expression
	}
)
