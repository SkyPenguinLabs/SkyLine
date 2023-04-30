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
// Filename      |  SkyLine_Scanner_Verification_Readers.go
// Project       |  SkyLine programming language
// Line Count    |  175 active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines the reader functions for the SkyLine scanner.
//
//
//
package SkyLine_Backend

import (
	"fmt"
	"strings"
	"unicode"
)

func (lex *ScannerStructure) ReadIdentifier() string {
	id := ""
	POS := lex.POS
	RPOS := lex.RPOS
	for Verify_Ident_Mark(lex.Char) {
		id += string(lex.Char)
		lex.ReadChar()
	}
	if strings.Contains(id, ".") {
		ok := ConstantIdents[id]
		if !ok {
			for _, idx := range Datatypes {
				if strings.HasPrefix(id, idx) {
					ok = true
				}
			}
		}
		if !ok {
			offset := strings.Index(id, ".")
			id = id[:offset]
			lex.POS = POS
			lex.RPOS = RPOS
			for offset > 0 {
				lex.ReadChar()
				offset--
			}
		}
	}
	return id
}

func (lex *ScannerStructure) ReadBacktick() string {
	bytereader := &strings.Builder{}
	for {
		lex.ReadChar()
		if lex.Char == '`' {
			break
		}
		bytereader.WriteByte(lex.Char)
	}
	return bytereader.String()
}

func Verify_Ident_Mark(ch byte) bool {
	if unicode.IsLetter(rune(ch)) || unicode.IsDigit(rune(ch)) || ch == '.' {
		return true
	}
	return false
}

func (lex *ScannerStructure) ReadNumber() string {
	NewCharLst := ""
	InterpreterAccepts := "0123456789"
	if lex.Char == '0' && lex.Peek() == 'x' {
		InterpreterAccepts = "0x123456789abcdefABCDEF"
	}
	if lex.Char == '0' && lex.Peek() == 'b' {
		InterpreterAccepts = "b01"
	}
	for strings.Contains(InterpreterAccepts, string(lex.Char)) {
		NewCharLst += string(lex.Char)
		lex.ReadChar()
	}
	return NewCharLst

}

func (lex *ScannerStructure) ReadChar() {
	lex.Prevch = lex.Char
	if lex.RPOS >= len(lex.CharInput) {
		lex.Char = 0
	} else {
		lex.Char = lex.CharInput[lex.RPOS]
	}
	lex.POS = lex.RPOS
	lex.RPOS++
}

func (lex *ScannerStructure) ReadString() string {
	bytereader := &strings.Builder{}
	for {
		lex.ReadChar()
		if lex.Char == '"' || lex.Char == 0 {
			break
		}
		if lex.Char == '\\' {
			lex.ReadChar()
			switch lex.Char {
			case 'n':
				bytereader.WriteByte('\n')
			case 'r':
				bytereader.WriteByte('\r')
			case 't':
				bytereader.WriteByte('\t')
			case 'f':
				bytereader.WriteByte('\f')
			case 'v':
				bytereader.WriteByte('\v')
			case '\\':
				bytereader.WriteByte('\\')
			case '"':
				bytereader.WriteByte('"')
			case 'x':
				lex.ReadChar()
				lex.ReadChar()
				lex.ReadChar()
				s := string([]byte{lex.Prevch, lex.Char})
				bytereader.WriteString(fmt.Sprintf("%x", s))
				continue
			}
			lex.ReadChar()
			continue
		}
		bytereader.WriteByte(lex.Char)
	}
	return bytereader.String()
}

func (lex *ScannerStructure) ReadIntToken() Token {
	intPart := lex.ReadNumber()
	if lex.Char != '.' {
		return Token{
			Token_Type: TOKEN_INT,
			Literal:    intPart,
		}
	}

	lex.ReadChar()
	fracPart := lex.ReadNumber()
	return Token{
		Token_Type: TOKEN_FLOAT,
		Literal:    intPart + "." + fracPart,
	}
}

func LexNew(input string) *ScannerStructure {
	lex := &ScannerStructure{CharInput: input}
	lex.ReadChar()
	return lex
}

func (lex *ScannerStructure) Peek() byte {
	if lex.RPOS >= len(lex.CharInput) {
		return 0
	}
	return lex.CharInput[lex.RPOS]
}
