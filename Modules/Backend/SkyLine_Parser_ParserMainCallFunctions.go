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
// Filename      |  SkyLine_Parser_ParserMainCallFunctions.go
// Project       |  SkyLine programming language
// Line Count    |  1,200+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This entire file is dedicated to every function for the parser or all main functions anyway which return their respective data types. It is important
//                 to mention that these functions are very important and need to be organized.
//
// State         | Working and secure but needs major changes
//
// Resolution    | These functions can be named better, worked better, modified better and even moved into another specific file while also having changes such as
//                 not flooding the file with trees
//
//
package SkyLine_Backend

import (
	"fmt"
	"strconv"
	"strings"
)

var root *TreeNode

// Parser helper functions
func ParserErrorSystem_GetFileName() (fname string) {
	if FileCurrent.Filename == "" && FileCurrent.Get_Name() == "" {
		fname = "REPL_main.skyline"
	} else {
		fname = FileCurrent.Filename
	}
	return fname
}

func (parser *Parser) ParseEngineStatement() *ENGINE {
	statement := &ENGINE{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing ( in import statement")
	}
	parser.NT()
	statement.EngineValue = parser.SkyLine_Expression(LOWEST)
	for parser.PeekTokenIs(TOKEN_SEMICOLON) {
		parser.NT()
	}
	return statement
}

func (parser *Parser) SkyLine_Statement() Statement {
	switch parser.CurrentToken.Token_Type {
	case TOKEN_KEYWORD_ENGINE:
		return parser.ParseEngineStatement()
	case TOKEN_REGISTER:
		return parser.ParseRegisterStatement()
	case TOKEN_LET:
		return parser.SkyLine_VariableCreation()
	case TOKEN_RETURN:
		return parser.SkyLine_Return()
	case TOKEN_CONSTANT:
		return parser.SkyLine_Constants()
	default:
		return parser.SkyLine_ExpressionStatement()
	}
}

func (parser *Parser) SkyLine_SwitchCase_Expression() Expression {
	EXP := &Switch{Token: parser.CurrentToken}
	EXP.Value = parser.SkyLine_Arguments()
	if EXP.Value == nil {
		return nil
	}
	if !parser.ExpectPeek(TOKEN_LBRACE) {
		return nil
	}
	parser.NT()
	for !parser.CurrentTokenIs(TOKEN_RBRACE) {
		if parser.CurrentTokenIs(TOKEN_EOF) {
			parser.Errors = append(parser.Errors, "unterminated switch statement")
		}
		exp := &Case{Token: parser.CurrentToken}
		if parser.CurrentTokenIs(TOKEN_DEFAULT) {
			exp.Def = true

		} else if parser.CurrentTokenIs(TOKEN_CASE) {

			parser.NT()
			if parser.CurrentTokenIs(TOKEN_DEFAULT) {
				exp.Def = true
			} else {
				exp.Expr = append(exp.Expr, parser.SkyLine_Expression(LOWEST))
				for parser.PeekTokenIs(TOKEN_COMMA) {
					parser.NT()
					parser.NT()
					exp.Expr = append(exp.Expr, parser.SkyLine_Expression(LOWEST))
				}
			}
		} else {
			parser.Errors = append(parser.Errors, fmt.Sprintf("expected case|default, got %s >>> %s ", parser.CurrentToken.Token_Type, parser.CurrentToken))
			return nil
		}

		if !parser.ExpectPeek(TOKEN_LBRACE) {

			msg := fmt.Sprintf("expected token to be '{', got %s instead", parser.CurrentToken.Token_Type)
			parser.Errors = append(parser.Errors, msg)
			fmt.Printf("error\n")
			return nil
		}

		// parse the block
		exp.Block = parser.SkyLine_BlockStatement()

		if !parser.CurrentTokenIs(TOKEN_RBRACE) {
			msg := fmt.Sprintf("Syntax Error: expected token to be '}', got %s instead", parser.CurrentToken.Token_Type)
			parser.Errors = append(parser.Errors, msg)
			fmt.Printf("error\n")
			return nil

		}
		parser.NT()
		EXP.Choices = append(EXP.Choices, exp)
	}
	count := 0
	for _, c := range EXP.Choices {
		if c.Def {
			count++
		}
	}
	if count > 1 {
		msg := "A switch-statement should only have one default block"
		parser.Errors = append(parser.Errors, msg)
		return nil

	}
	return EXP

}

func (parser *Parser) SkyLine_Arguments() Expression {
	if !parser.ExpectPeek(TOKEN_LPAREN) {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Unexpected token | %s | I need %s for this argument list ", parser.CurrentToken.Literal, TOKEN_LPAREN))
		return nil
	}
	parser.NT()
	exp := parser.SkyLine_Expression(LOWEST)
	if exp == nil {
		return nil
	}
	if !parser.ExpectPeek(TOKEN_RPAREN) {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Unexpected token | %s|  I need %s for this statement ", parser.CurrentToken.Literal, TOKEN_RPAREN))
		return nil
	}
	return exp
}

func (parser *Parser) SkyLine_VariableCreation() *LetStatement {
	stmt := &LetStatement{Token: parser.CurrentToken}
	if !parser.ExpectPeek(TOKEN_IDENT) {
		return nil
	}
	stmt.Name = &Ident{Token: parser.CurrentToken, Value: parser.CurrentToken.Literal}
	if !parser.ExpectPeek(TOKEN_ASSIGN) {
		return nil
	}
	parser.NT()
	stmt.Value = parser.SkyLine_Expression(LOWEST)
	for !parser.CurrentTokenIs(TOKEN_SEMICOLON) {
		if parser.CurrentTokenIs(TOKEN_EOF) {
			var fname string
			if FileCurrent.Filename == "" {
				fname = "REPL.skyline"
			} else {
				fname = FileCurrent.Filename
			}

			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + fname + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprintf("%v", ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											"Parser Error (parse create assignment)" +
											SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											Map_Parser[ERROR_MISSING_SEMICOLON_IN_STATEMENT_AT](stmt.SL_ExtractStringValue(), SKYLINE_HIGH_DEFPURPLE, SKYLINE_RESTORE).Message +
											SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(parser.GetLineCound()) + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated statement   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + strings.Trim(stmt.SL_ExtractStringValue(), ";"),
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] " + "Consider adding a semicolon to the end of the statement (';')" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Fixed Satatement" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_FIXBLUE + "[F] " + stmt.SL_ExtractStringValue(),
									},
								},
							},
						},
					},
				},
			}
			if ErrorSys.TreeValid() {
				RetTreeSys(root, "", false)
			}
			if ErrorSys.LineValid() {
				var msg string
				msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
				msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Error when trying to parse statement, missing semicolon" + "\n"
				msg += SKYLINE_HIGH_FIXBLUE + "[F] | " + SKYLINE_RESTORE + stmt.SL_ExtractStringValue() + "\n"
				msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
				msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + "Suggested to put a semicolon at the end of the statement"
				fmt.Println(msg)
			}
			parser.Errors = append(parser.Errors, "")
			return nil
		}
		parser.NT()
	}
	return stmt

}

// wwork on this error here
func (parser *Parser) SkyLine_Assignment(name Expression) Expression {
	stmt := &AssignmentStatement{Token: parser.CurrentToken}
	if StatementName, ok := name.(*Ident); ok {
		stmt.Name = StatementName
	} else {
		parser.Errors = append(parser.Errors, fmt.Sprintf("Expected assignment token before operator to be an IDENTIFIER not %s", name.SL_ExtractNodeValue()))
	}
	opperand := parser.CurrentToken
	parser.NT()
	switch opperand.Token_Type {
	case TOKEN_PLUS_EQUALS:
		stmt.Operator = "+="
	case TOKEN_MINUS_EQUALS:
		stmt.Operator = "-="
	case TOKEN_DIVEQ:
		stmt.Operator = "/="
	case TOKEN_ASTERISK_EQUALS:
		stmt.Operator = "*="
	default:
		stmt.Operator = "="
	}

	stmt.Value = parser.SkyLine_Expression(LOWEST)
	return stmt
}

func (parser *Parser) SkyLine_Constants() *Constant {
	statement := &Constant{Token: parser.CurrentToken}
	if !parser.ExpectPeek(TOKEN_IDENT) {
		return nil
	}
	statement.Name = &Ident{Token: parser.CurrentToken, Value: parser.CurrentToken.Literal}
	if !parser.ExpectPeek(TOKEN_ASSIGN) {
		return nil
	}
	parser.NT()
	statement.Value = parser.SkyLine_Expression(LOWEST)
	for !parser.CurrentTokenIs(TOKEN_SEMICOLON) {
		if parser.CurrentTokenIs(TOKEN_EOF) {
			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											fmt.Sprintf(
												"%v",
												ERROR_UNTERMINATED_CONSTANT_VALUE,
											) +
											SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse Constant)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											Map_Parser[ERROR_UNTERMINATED_CONSTANT_VALUE](
												statement.SL_ExtractStringValue(),
												SKYLINE_HIGH_DEFPURPLE,
												SKYLINE_RESTORE,
											).Message,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(parser.GetLineCound()) + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated statement   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + strings.Trim(statement.SL_ExtractStringValue(), ";"),
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA +
											Map_Parser[ERROR_UNTERMINATED_CONSTANT_VALUE](
												statement.SL_ExtractStringValue(),
												SKYLINE_HIGH_DEFPURPLE,
												SKYLINE_RESTORE,
											).Suggestion,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Fixed Statement" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_FIXBLUE + "[F] " + statement.SL_ExtractStringValue(),
									},
								},
							},
						},
					},
				},
			}
			if ErrorSys.TreeValid() {
				RetTreeSys(root, "", false)
			}
			if ErrorSys.LineValid() {
				fmt.Println(parser.Prepare_Base_Error_Message(
					"Missing semicolon at the end of constant assignment",
					strings.Trim(statement.SL_ExtractStringValue(), ";"),
					true,
					statement.SL_ExtractStringValue(),
				))
			}
			parser.Errors = append(parser.Errors, "")
			return nil
		}
		parser.NT()
	}
	return statement
}

func (parser *Parser) SkyLine_Return() *ReturnStatement {
	stmt := &ReturnStatement{
		Token: parser.CurrentToken,
	}
	parser.NT()

	stmt.ReturnValue = parser.SkyLine_Expression(LOWEST)

	for parser.PeekTokenIs(TOKEN_SEMICOLON) {
		parser.NT()
	}

	return stmt
}

func (parser *Parser) SkyLine_ExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{
		Token:      parser.CurrentToken,
		Expression: parser.SkyLine_Expression(LOWEST),
	}

	if parser.PeekTokenIs(TOKEN_SEMICOLON) {
		parser.NT()
	}

	return stmt
}

func (SL_Parser *Parser) ParsePostfixExpression() Expression {
	expression := &PostfixExpression{
		Token:    SL_Parser.PreviousToken,
		Operator: SL_Parser.CurrentToken.Literal,
	}
	return expression
}

func (parser *Parser) SkyLine_Expression(precedence int) Expression {
	postfix := parser.PostfixParseFns[parser.CurrentToken.Token_Type]
	if postfix != nil {
		return (postfix())
	}
	prefix := parser.PrefixParseFns[parser.CurrentToken.Token_Type]

	if prefix == nil {
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA +
										fmt.Sprintf(
											"%v",
											ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV,
										) +
										SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse Expression) " + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA +
										Map_Parser[ERROR_PREFIX_PARSE_FUNCTION_NOT_LOADED_INTO_ENV]().Message +
										SKYLINE_RESTORE,
								},
								{
									Type: SKYLINE_HIGH_DEFAQUA + "[Sub Branch] Token " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_HIGH_DEFAQUA + parser.CurrentToken.Literal + SKYLINE_RESTORE,
										},
									},
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(parser.GetLineCound()),
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute(),
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Token       : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.CurrentToken.Literal,
								},
							},
						},
					},
				},
			},
		}
		if ErrorSys.TreeValid() {
			RetTreeSys(root, "", false)
		}
		if ErrorSys.LineValid() {
			var msg string
			msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
			msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Could not locate a prefix parse function for the parsed token" + "\n"
			msg += SKYLINE_HIGH_FIXBLUE + "[F] | " + SKYLINE_RESTORE + parser.CurrentToken.Literal + "\n"
			msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
			msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + "Check your token?"
			fmt.Println(msg)
		}
		parser.Errors = append(parser.Errors, "")
		return nil
	}
	leftExp := prefix()

	for !parser.CurrentTokenIs(TOKEN_SEMICOLON) && precedence < parser.peekPrecedence() {
		infix := parser.InfixParseFns[parser.PeekToken.Token_Type]
		if infix == nil {
			return leftExp
		}

		parser.NT()

		leftExp = infix(leftExp)
	}
	return leftExp
}

func (parser *Parser) SkyLine_Identifier() Expression {
	return &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

//TODO: Remake this error system for 90%
func (parser *Parser) SkyLine_IntegerLiteral() Expression {
	lit := &IntegerLiteral{Token: parser.CurrentToken}
	var value int64
	var x error
	var errtype string
	if strings.HasPrefix(parser.CurrentToken.Literal, "0b") {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal[2:], 2, 64)
		if x != nil {
			errtype = "binary"
		}
	} else if strings.HasPrefix(parser.CurrentToken.Literal, "0x") {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal[2:], 16, 64)
		if x != nil {
			errtype = "hex"
		}
	} else {
		value, x = strconv.ParseInt(parser.CurrentToken.Literal, 0, 64)
		if x != nil {
			errtype = "int"
		}
	}
	// Yes we check again, simple, dont ask
	if x != nil {
		_, x := CheckParseValue(parser.CurrentToken.Literal, errtype)
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_TYPE_INTEGRITY_PARSE_INTEGER_ERROR) + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse Integer Literal)" + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser could not parse the given integer" + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[Sub Branch] Debug " + SKYLINE_RESTORE,
											Children: []*TreeNode{
												{

													Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Error In Type  ? " + SUNRISE_HIGH_DEFINITION + x.Type + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Error In Value ? " + SUNRISE_HIGH_DEFINITION + x.Value + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Error In Parse ? " + SUNRISE_HIGH_DEFINITION + x.Err.Error() + SKYLINE_RESTORE,
												},
											},
										},
									},
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + lit.SL_ExtractStringValue(),
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Max possible value for int64    ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Max + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Lowest possible value for int64 ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Lowest + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_SICK_BLUE + "[S] " + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Suggest + SKYLINE_RESTORE,
								},
							},
						},
					},
				},
			},
		}
		if ErrorSys.TreeValid() {
			RetTreeSys(root, "", false)
		}
		if ErrorSys.LineValid() {
			var msg string
			msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
			msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Could not parse the integer/hex/binary number given" + "\n"
			msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
			msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Suggest + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Max    (int64)  -> " + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Max + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Lowest (int64)  -> " + CheckParsedError(x.Err.Error(), lit.SL_ExtractStringValue()).Lowest + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Type of (Error) -> " + x.Type + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Error Message   -> " + x.Err.Error() + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Value to parse  -> " + x.Value + SKYLINE_RESTORE
			fmt.Println(msg)
		}
		parser.Errors = append(parser.Errors, "")
	}
	lit.Value = value
	return lit
}

func (parser *Parser) SkyLine_FloatLiteral() Expression {
	val, err := strconv.ParseFloat(parser.CurrentToken.Literal, 64)
	if err != nil {
		duringparse := CheckAndVerify(parser.CurrentToken.Literal)
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_COULD_NOT_PARSE_FLOAT_VALUE) + SKYLINE_RESTORE,
						},
					},
				},
				{
					Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFAQUA + "Parser Error ( Parse Float Literal ) " + SKYLINE_RESTORE,
						},
					},
				},
				{
					Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFAQUA + "Parser was not able to process the provided input as a float64",
							Children: []*TreeNode{
								{
									Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[Sub Branch] Debug " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Value too large   ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(duringparse.TooLong) + SKYLINE_RESTORE,
										},
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Value too small   ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(duringparse.TooShort) + SKYLINE_RESTORE,
										},
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[SL-Info] Value parsed      ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(duringparse.Parsed) + SKYLINE_RESTORE,
										},
									},
								},
							},
						},
					},
				},
				{
					Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.CurrentToken.Literal,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] MAX (Float64)          ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(duringparse.Max) + SKYLINE_RESTORE,
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "[UW] MIN (Float64)          ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + fmt.Sprint(duringparse.Low) + SKYLINE_RESTORE,
						},
					},
				},
				{
					Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_SICK_BLUE + "[S] " + duringparse.Recomend + SKYLINE_RESTORE,
						},
					},
				},
			},
		}
		if ErrorSys.TreeValid() {
			RetTreeSys(root, "", false)
		}
		if ErrorSys.LineValid() {
			var msg string
			msg += SKYLINE_HIGH_DEFRED + "[E] | " + SKYLINE_RESTORE + FileCurrent.GetAbsolute() + "\n"
			msg += SKYLINE_HIGH_DEFRED + "[1] | " + SKYLINE_RESTORE + "Could not parse the integer/hex/binary number given" + "\n"
			msg += SUNRISE_LIGHT_DEFINITION + "[L] | " + SKYLINE_RESTORE + parser.GetLineCound() + "\n"
			msg += SKYLINE_SICK_BLUE + "[S] | " + SKYLINE_RESTORE + duringparse.Recomend + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Max    (int64)  -> " + fmt.Sprint(duringparse.Max) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Lowest (int64)  -> " + fmt.Sprint(duringparse.Low) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Too long        -> " + fmt.Sprint(duringparse.TooLong) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Too short       -> " + fmt.Sprint(duringparse.TooShort) + SKYLINE_RESTORE
			msg += SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "\n[W] | " + SKYLINE_RESTORE + " Value to parse  -> " + parser.CurrentToken.Literal + SKYLINE_RESTORE
			fmt.Println(msg)
		}
		parser.Errors = append(parser.Errors, "")
		return nil
	}

	return &FloatLiteral{
		Token: parser.CurrentToken,
		Value: val,
	}
}

func (parser *Parser) SkyLine_PrefixExpression() Expression {
	expr := &PrefixExpression{
		Token:    parser.CurrentToken,
		Operator: parser.CurrentToken.Literal,
	}

	parser.NT()

	expr.Right = parser.SkyLine_Expression(PREFIX)
	return expr
}

func (parser *Parser) peekPrecedence() int {
	if p, ok := Precedences[parser.PeekToken.Token_Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) curPrecedence() int {
	if p, ok := Precedences[parser.CurrentToken.Token_Type]; ok {
		return p
	}
	return LOWEST
}

func (parser *Parser) SkyLine_InfixExpression(left Expression) Expression {
	expr := &InfixExpression{
		Token:    parser.CurrentToken,
		Operator: parser.CurrentToken.Literal,
		Left:     left,
	}

	prec := parser.curPrecedence()

	parser.NT()

	expr.Right = parser.SkyLine_Expression(prec)
	return expr
}

func (parser *Parser) SkyLine_Boolean() Expression {
	return &Boolean_AST{
		Token: parser.CurrentToken,
		Value: parser.CurrentTokenIs(TOKEN_TRUE),
	}
}

func (parser *Parser) SkyLine_GroupImportExpression() Expression {
	parser.NT()

	expr := parser.SkyLine_Expression(LOWEST)

	if !parser.ExpectPeek(TOKEN_LINE) {
		return nil
	}

	return expr
}

func (parser *Parser) SkyLine_GroupedExpression() Expression {
	parser.NT()

	expr := parser.SkyLine_Expression(LOWEST)

	if !parser.ExpectPeek(TOKEN_RPAREN) {
		return nil
	}

	return expr
}

func (parser *Parser) SkyLine_ConditionalExpression() Expression {
	expr := &ConditionalExpression{Token: parser.CurrentToken}

	parser.NT()
	expr.Condition = parser.SkyLine_Expression(LOWEST)
	if !parser.ExpectPeek(TOKEN_LBRACE) {
		return nil
	}

	expr.Consequence = parser.SkyLine_ConditionalBlock()
	if expr.Consequence == nil {
		root = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_PARSER_FOUND_NIL_EXPRESSION_UNEXPECTED) + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse If Expression)" + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Parser found an empty or nil consequence (UNEXPECT:NIL->EXPRESSION)" + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + expr.Condition.SL_ExtractStringValue(),
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_SICK_BLUE + "[S] Make sure the consequence is not empty for the conditional" + SKYLINE_RESTORE,
								},
							},
						},
					},
				},
			},
		}
		RetTreeSys(root, "", true)
		parser.Errors = append(parser.Errors, "")
		return nil
	}
	if parser.PeekTokenIs(TOKEN_ELSE) {
		parser.NT()

		if parser.PeekTokenIs(TOKEN_IF) {

			parser.NT()
			expr.Alternative = &BlockStatement{
				Statements: []Statement{
					&ExpressionStatement{
						Expression: parser.SkyLine_ConditionalBlock(),
					},
				},
			}
			return expr
		}

		// parse else

		if !parser.ExpectPeek(TOKEN_LBRACE) {
			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_PARSER_EXPECTED_LBRACE_BUT_GOT_SOMETHING) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse If Expression)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser found an unexpected token but needs '{' (UNEXPECT:Token)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number   : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path     : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + expr.Condition.SL_ExtractStringValue() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Parser Found Unexpected : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.CurrentToken.Literal + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] Make sure the conditional includes the proper statement" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Fixed Satatement " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_FIXBLUE + "[F] " + expr.SL_ExtractStringValue() + SKYLINE_RESTORE,
									},
								},
							},
						},
					},
				},
			}
			RetTreeSys(root, "", false)
			parser.Errors = append(parser.Errors, "")
			return nil
		}

		expr.Alternative = parser.SkyLine_ConditionalBlock()
		if expr.Alternative == nil {
			root = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree" + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_PARSER_FOUND_NIL_EXPRESSION_UNEXPECTED) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser Error (Parse If Expression)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Parser found an empty or nil consequence (UNEXPECT:NIL->EXPRESSION)" + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated line number  : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + parser.GetLineCound() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path    : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated Value Parsed : " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + expr.Condition.SL_ExtractStringValue(),
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Suggestion " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] Make sure the consequence is not empty for the conditional" + SKYLINE_RESTORE,
									},
								},
							},
						},
					},
				},
			}
			RetTreeSys(root, "", true)
			return nil
		}
	}

	return expr
}

func (parser *Parser) SkyLine_ConditionalBlock() *BlockStatement {
	block := &BlockStatement{
		Token:      parser.CurrentToken,
		Statements: []Statement{},
	}
	parser.NT()
	for !parser.CurrentTokenIs(TOKEN_RBRACE) && !parser.CurrentTokenIs(TOKEN_EOF) {
		stmt := parser.SkyLine_Statement()
		block.Statements = append(block.Statements, stmt)
		parser.NT()
	}
	return block
}

func (parser *Parser) SkyLine_BlockStatement() *BlockStatement {
	block := &BlockStatement{
		Token:      parser.CurrentToken,
		Statements: []Statement{},
	}

	parser.NT()
	for !parser.CurrentTokenIs(TOKEN_RBRACE) && !parser.CurrentTokenIs(TOKEN_EOF) {
		stmt := parser.SkyLine_Statement()
		block.Statements = append(block.Statements, stmt)
		parser.NT()
	}
	return block
}

func (parser *Parser) SkyLine_FunctionLiteral() Expression {
	lit := &FunctionLiteral{Token: parser.CurrentToken}

	if !parser.ExpectPeek(TOKEN_LPAREN) {
		return nil
	}

	lit.Parameters = parser.SkyLine_FunctionParameters()

	if !parser.ExpectPeek(TOKEN_LBRACE) {
		return nil
	}

	lit.Body = parser.SkyLine_BlockStatement()

	return lit
}

func (parser *Parser) SkyLine_FunctionParameters() []*Ident {
	idents := []*Ident{}

	if parser.PeekTokenIs(TOKEN_RPAREN) {
		parser.NT()
		return idents
	}

	parser.NT()

	ident := &Ident{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
	idents = append(idents, ident)

	for parser.PeekTokenIs(TOKEN_COMMA) || parser.PeekTokenIs(TOKEN_COLON) {
		parser.NT()
		parser.NT()
		ident := &Ident{
			Token: parser.CurrentToken,
			Value: parser.CurrentToken.Literal,
		}
		idents = append(idents, ident)
	}

	if !parser.ExpectPeek(TOKEN_RPAREN) {
		return nil
	}

	return idents
}

func (parser *Parser) SkyLine_ExpressionList(end Token_Type) []Expression {
	list := make([]Expression, 0)

	if parser.PeekTokenIs(end) {
		parser.NT()
		return list
	}

	parser.NT()
	list = append(list, parser.SkyLine_Expression(LOWEST))

	for parser.PeekTokenIs(TOKEN_COMMA) {
		parser.NT()
		parser.NT()
		list = append(list, parser.SkyLine_Expression(LOWEST))
	}

	if !parser.ExpectPeek(end) {
		return nil
	}

	return list
}

func (parser *Parser) parseCallExpression(function Expression) Expression {
	return &CallExpression{
		Token:     parser.CurrentToken,
		Function:  function,
		Arguments: parser.SkyLine_ExpressionList(TOKEN_RPAREN),
	}
}

func (parser *Parser) parseStringLiteral() Expression {
	return &StringLiteral{
		Token: parser.CurrentToken,
		Value: parser.CurrentToken.Literal,
	}
}

func (parser *Parser) parseArrayLiteral() Expression {
	return &ArrayLiteral{
		Token:    parser.CurrentToken,
		Elements: parser.SkyLine_ExpressionList(TOKEN_RBRACKET),
	}
}

func (parser *Parser) parseIndexExpression(left Expression) Expression {
	expr := &IndexExpression{
		Token: parser.CurrentToken,
		Left:  left,
	}

	parser.NT()
	expr.Index = parser.SkyLine_Expression(LOWEST)

	if !parser.ExpectPeek(TOKEN_RBRACKET) {
		return nil
	}

	return expr
}

func (parser *Parser) parseHashLiteral() Expression {
	hash := &HashLiteral{
		Token: parser.CurrentToken,
		Pairs: make(map[Expression]Expression),
	}

	for !parser.PeekTokenIs(TOKEN_RBRACE) {
		parser.NT()
		key := parser.SkyLine_Expression(LOWEST)

		if !parser.ExpectPeek(TOKEN_COLON) {
			return nil
		}

		parser.NT()
		value := parser.SkyLine_Expression(LOWEST)
		hash.Pairs[key] = value

		if !parser.PeekTokenIs(TOKEN_RBRACE) && !parser.ExpectPeek(TOKEN_COMMA) {
			return nil
		}
	}

	if !parser.ExpectPeek(TOKEN_RBRACE) {
		return nil
	}

	return hash
}

func (parser *Parser) parseMethodCallExpression(obj Expression) Expression {
	methodcall := &ObjectCallExpression{
		Token:      parser.CurrentToken,
		SLC_Object: obj,
	}
	parser.NT()
	name := parser.SkyLine_Identifier()
	parser.NT()
	methodcall.Call = parser.parseCallExpression(name)
	return methodcall
}

func (parser *Parser) ParseRegisterStatement() *Register {
	statement := &Register{Token: parser.CurrentToken}
	if !parser.PeekTokenIs("(") {
		fmt.Println("Missing ( in import statement")
	}
	parser.NT()
	statement.RegistryValue = parser.SkyLine_Expression(LOWEST)
	for parser.PeekTokenIs(TOKEN_SEMICOLON) {
		parser.NT()
	}
	return statement
}

func (p *Parser) SkyLine_ImportExpression() Expression {
	expression := &ImportExpression{Token: p.CurrentToken}

	if !p.ExpectPeek(TOKEN_LPAREN) {
		return nil
	}

	p.NT()
	expression.Name = p.SkyLine_Expression(LOWEST)
	if !p.ExpectPeek(TOKEN_RPAREN) {
		return nil
	}

	return expression
}

func (p *Parser) SkyLine_SelectorExpression(exp Expression) Expression {
	p.ExpectPeek(TOKEN_IDENT)
	index := &StringLiteral{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
	return &IndexExpression{Left: exp, Index: index}
}

func (parser *Parser) SkyLine_ForLoop() Expression {
	expression := &ForLoopExpression{
		Token: parser.CurrentToken,
	}
	if !parser.ExpectPeek(TOKEN_LPAREN) {
		return nil
	}
	parser.NextLoadFaultToken()
	expression.Condition = parser.SkyLine_Expression(LOWEST)
	if !parser.ExpectPeek(TOKEN_RPAREN) {
		return nil
	}
	if !parser.ExpectPeek(TOKEN_LBRACE) {
		return nil
	}
	expression.Consequence = parser.SkyLine_BlockStatement()
	return expression
}

func (p *Parser) SkyLine_ForEach() Expression {
	expression := &ForeachStatement{Token: p.CurrentToken}
	p.NT()
	expression.Ident = p.CurrentToken.Literal
	if p.PeekTokenIs(TOKEN_COMMA) {
		p.NT()

		if !p.PeekTokenIs(TOKEN_IDENT) {
			p.Errors = append(p.Errors, fmt.Sprintf("second argument to foreach must be ident, got %v", p.PeekToken))
			return nil
		}
		p.NT()

		expression.Index = expression.Ident
		expression.Ident = p.CurrentToken.Literal
	}
	if !p.ExpectPeek(TOKEN_INSIDE) {
		return nil
	}
	p.NT()
	expression.Value = p.SkyLine_Expression(LOWEST)
	if expression.Value == nil {
		return nil
	}
	p.NT()
	expression.Body = p.SkyLine_BlockStatement()
	return expression
}
