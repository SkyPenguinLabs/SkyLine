///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Parser_SLDataTypes
// Extension         | .go ( golang source code file )
// Purpose           | Defines all parsing functions for data types
// Directory         | Modules/Backend/SkyEnvironment
// Modular Directory | SkyLine/Modules/Backend/SkyEnvironment
// Package Name      | SkyLine_Backend_Module_Parser
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// The second major part of a programming language is parsing. Parsing does not necessarily execute the tokens but rather `parses` the tokens themselves and can pass them onto
//
// the evaluation step. In this step, we parse statements and expressions such as let, set, cause, engine, allow, call, etc.
//
package SkyLine_Backend_Module_Parser

import (
	SLAST "SkyLine/Modules/Backend/SkyAST"
	SL_TK "SkyLine/Modules/Backend/SkyTokens"
	"fmt"
	"log"
	"math"
	"strconv"
)

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Boolean parsing function, translates to "bool" in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Boolean() SLAST.SL_Expression {
	return &SLAST.SL_Boolean{
		TokenConstruct: SLP.SL_CurrentToken,
		Value:          SLP.SkyLine_Parser_Helper_CurrentTokenIs(SL_TK.TOKEN_TRUE),
	}
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: String parsing function, translates to "string" in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_String() SLAST.SL_Expression {
	return &SLAST.SL_String{
		TokenConstruct: SLP.SL_CurrentToken,
		Value:          SLP.SL_CurrentToken.Literal,
	}
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Byte parsing function, translates to "byte" in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Byte() SLAST.SL_Expression {
	if SLP.SL_CurrentToken.Token_Type != SL_TK.TOKEN_BYTESTART {
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, "Missing ( ' ) in statement for byte parsing around line "+fmt.Sprint(SLP.SL_Scanner.ReadLineNum()))
		return nil
	}
	BV := &SLAST.SL_Byte{
		TokenConstruct: SLP.SL_CurrentToken,
		Value:          SLP.SL_CurrentToken.Literal[0],
	}
	SLP.SkyLine_Parser_Helper_LoadNextToken()
	return BV
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Integer parsing function, translates to "int" in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func FitIntoDataType(VAL int64) interface{} {
	if VAL >= math.MinInt8 && VAL <= math.MaxInt8 {
		return &SLAST.SL_Integer8{Value: int8(VAL)}
	} else if VAL >= math.MinInt16 && VAL <= math.MaxInt16 {
		return &SLAST.SL_Integer16{Value: int16(VAL)}
	} else if VAL >= math.MinInt32 && VAL <= math.MaxInt32 {
		return &SLAST.SL_Integer32{Value: int32(VAL)}
	} else if VAL >= math.MinInt && VAL <= math.MaxInt {
		return &SLAST.SL_Integer{Value: int(VAL)}
	} else {
		return &SLAST.SL_Integer64{Value: VAL}
	}
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Integer() SLAST.SL_Expression {
	current := SLP.SL_CurrentToken

	VAL, err := strconv.ParseInt(SLP.SL_CurrentToken.Literal, 0, 64)
	if err != nil {
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, fmt.Sprintf("Could not parse (%s) as an integer value, may be too large -> %s", SLP.SL_CurrentToken, err.Error()))
		return nil
	}
	inttoken, x := strconv.ParseInt(current.Literal, 0, 64)
	if x != nil {
		log.Fatal(x)
	}
	dataType := FitIntoDataType(inttoken)
	switch dataType.(type) {
	case *SLAST.SL_Integer8:
		return &SLAST.SL_Integer8{Value: int8(VAL), TokenConstruct: SLP.SL_CurrentToken}
	case *SLAST.SL_Integer16:
		return &SLAST.SL_Integer16{Value: int16(VAL), TokenConstruct: SLP.SL_CurrentToken}
	case *SLAST.SL_Integer32:
		return &SLAST.SL_Integer32{Value: int32(VAL), TokenConstruct: SLP.SL_CurrentToken}
	case *SLAST.SL_Integer:
		return &SLAST.SL_Integer{Value: int(VAL), TokenConstruct: SLP.SL_CurrentToken}
	default:
		return &SLAST.SL_Integer64{Value: int64(VAL), TokenConstruct: SLP.SL_CurrentToken}
	}
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Float parsing function, translates to "float32" in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Float64() SLAST.SL_Expression {
	VALUE := &SLAST.SL_Float{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	VAL, err := strconv.ParseFloat(SLP.SL_CurrentToken.Literal, 64)
	if err != nil {
		SLP.SL_Parser_Errors = append(SLP.SL_Parser_Errors, fmt.Sprintf("could not parse %q as float around line %d", SLP.SL_CurrentToken.Literal, SLP.SL_Scanner.ReadLineNum()))
		return nil
	}
	VALUE.Value = VAL
	return VALUE
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Null parsing function, translates to "nil" in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Null() SLAST.SL_Expression {
	return &SLAST.SL_NULL{
		TokenConstruct: SLP.SL_CurrentToken,
	}
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Array parsing function, translates to '[]interface' in golang's type system
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Array() SLAST.SL_Expression {
	Arr := &SLAST.SL_Array{
		TokenConstruct: SLP.SL_CurrentToken,
	}
	Arr.Elements = SLP.SkyLine_Parser_Functions_Parse_ExpressionList(SL_TK.TOKEN_RBRACKET)
	return Arr
}

func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_HashMap() SLAST.SL_Expression {
	Hash := &SLAST.SL_HashMap{TokenConstruct: SLP.SL_CurrentToken}
	Hash.Pairs = make(map[SLAST.SL_Expression]SLAST.SL_Expression)
	for !SLP.SkyLine_Parser_Helper_PeekTokenCmp(SL_TK.TOKEN_RBRACE) {
		SLP.SkyLine_Parser_Helper_LoadNextToken()

		Key := SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
		if !SLP.SkyLine_Parser_Helper_ExpectPeek(SL_TK.TOKEN_COLON) {
			return nil
		}
		SLP.SkyLine_Parser_Helper_LoadNextToken()
		Value := SLP.SkyLine_Parser_Expressions_Parse_Expression(LOWEST)
		Hash.Pairs[Key] = Value
		if !SLP.SkyLine_Parser_Helper_PeekTokenCmp(SL_TK.TOKEN_RBRACE) && !SLP.SkyLine_Parser_Helper_ExpectPeek(SL_TK.TOKEN_COMMA) {
			return nil
		}
	}
	if !SLP.SkyLine_Parser_Helper_ExpectPeek(SL_TK.TOKEN_RBRACE) {
		return nil
	}
	return Hash

}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:: Identifier parsing function
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//

func (SLP *SkyLine_Parser) SkyLine_Parser_Function_DataType_Identifier() SLAST.SL_Expression {
	return &SLAST.SL_Identifier{
		TokenConstruct: SLP.SL_CurrentToken,
		Value:          SLP.SL_CurrentToken.Literal,
	}
}
