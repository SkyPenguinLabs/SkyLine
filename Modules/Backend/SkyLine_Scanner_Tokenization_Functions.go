package SkyLine_Backend

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// This file contains all scanner and categorization functions for tokens that were caught or found during lexical analysis.
//
// Belongs to   : TokenizerMap of type map[string]->TokenizerFunc func(*ScannerStructure) Token
// Within File  : SkyLine_Scanner_Verification_TokenizerAndCategorizer.go - SkyLine_Scanner_Tokenization_Functions
// Funcs        : 31 total functions
//
//
// This file should not be touched much other than when a new token, function or token type is executed or found. This code should be relatively
// faster and much more simple and efficient to run than chaining if statements. Working with maps just make it genuinely better and more efficient
// to work with. Also in the case that we needed to modify something functions are held in a seperate file
//

// Categorize AND AND (&&)
func Categorize_AndAnd(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_ANDAND, lex.Char)
}

// Categorize equal =
func Categorize_Eq(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_EQ, lex.Char)
}

// Categorize not equal !=
func Categorize_ze_Neq(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_NEQ, lex.Char)
}

// Categorize Colon :
func Categorize_Colon(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_COLON, lex.Char)
}

// Categorize semicolons ;
func Categorize_Semicolon(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_SEMICOLON, lex.Char)
}

// Categorize left parentheses (
func Categorize_LParen(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_LPAREN, lex.Char)
}

// Categorize OR OR boolean algebra symbol
func Categorize_OrOr(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_OROR, lex.Char)
}

// Categorize commas ,
func Categorize_Comma(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_COMMA, lex.Char)
}

// Categorize right parentheses )
func Categorize_RParen(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_RPAREN, lex.Char)
}

// Categorize Addition
func Categorize_Plus(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_PLUS, lex.Char)
}

// Categorize addition assignment
func Categorize_PlusEquals(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_PLUS_EQUALS, lex.Char)
}

// Categorize subtraction
func Categorize_Minus(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_MINUS, lex.Char)
}

// Categorize minus equals (-=)
func Categorize_MinusEquals(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_MINUS_EQUALS, lex.Char)
}

// Categorize multiply equals (*=)
func Categorize_AsteriskEquals(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_ASTERISK_EQUALS, lex.Char)
}

// Categorize power of **
func Categorize_PowerOf(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_POWEROF, lex.Char)
}

// Categorize asterisk characters (*)
func Categorize_Asterisk(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_ASTARISK, lex.Char)
}

// Categorize modulo (%)
func Categorize_Modulo(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_MODULO, lex.Char)
}

// Categorize division equal (/=)
func Categorize_DivEq(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_DIVEQ, lex.Char)
}

// Categorize slash (/)
func Categorize_Slash(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_SLASH, lex.Char)
}

// Categorize less than or equal to (<=)
func Categorize_Lteq(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_LTEQ, lex.Char)
}

// Categorize less than (<)
func Categorize_Lt(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_LT, lex.Char)
}

// Categorize Greater Than or Equal to (>=)
func Categorize_Gteq(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_GTEQ, lex.Char)
}

// Categorize greater than (>)
func Categorize_Gt(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_GT, lex.Char)
}

// Categorize {
func Categorize_LBrace(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_LBRACE, lex.Char)
}

// Categorize }
func Categorize_RBrace(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_RBRACE, lex.Char)
}

// Categorize right bracket ]
func Categorize_LBracket(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_LBRACKET, lex.Char)
}

// Categorize left bracket [
func Categorize_RBracket(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_RBRACKET, lex.Char)
}

// Categorize bangs (!)
func Categorize_Bang(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_BANG, lex.Char)
}

// Categorize assignment tokens
func Categorize_Assign(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_ASSIGN, lex.Char)
}

// Categorize periods (.)
func Categorize_Period(lex *ScannerStructure) Token {
	return ScanNewToken(TOKEN_PERIOD, lex.Char)
}
