///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_Readers
// Extension         | .go ( golang source code file )
// Purpose           | Define readers for the scanner during lexical analysis
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines a set of functions known as readers. The reader functions main jobs are to read specific values such as characters, read integers, read floats and data types
//
// like string values as well and data within strings. This can improve the performance of the language by not having to use constant conditional expressions under each time we
//
// want to parse a value or read a new character into the scanners stream.
package SkyLine_Backend_Scanner

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	SLTK "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyTokens"
)

func (l *SL_ScannerStructure) RCHAR() {
	if l.Scanner_RPOS >= len(l.Scanner_Characters) {
		l.Scanner_Character = rune(0)
	} else {
		l.Scanner_Character = l.Scanner_Characters[l.Scanner_RPOS]
	}
	l.Scanner_POS = l.Scanner_RPOS
	l.Scanner_RPOS++
}

func (SL_Scanner *SL_ScannerStructure) R_NUMBER() (string, string) {
	STRING := ""
	ACCEPTING := "0123456789"
	if SL_Scanner.Scanner_Character == '0' && SL_Scanner.CharacterPeek() == 'b' {
		ACCEPTING = "b01"
	}
	if SL_Scanner.Scanner_Character == '0' && SL_Scanner.CharacterPeek() == 'x' {
		ACCEPTING = "0x123456789abcdefABCDEF"
	}
	for strings.Contains(ACCEPTING, string(SL_Scanner.Scanner_Character)) {
		STRING += string(SL_Scanner.Scanner_Character)
		SL_Scanner.RCHAR()
	}
	var parseas string
	input := STRING
	if CheckintRange(input, math.MinInt32, math.MaxInt32) {
		parseas = "Int32"
	}
	if CheckintRange(input, math.MinInt8, math.MaxInt8) {
		parseas = "Int8"
	}

	if CheckintRange(input, math.MinInt16, math.MaxInt16) {
		parseas = "Int16"
	}
	if CheckintRange(input, math.MinInt32, math.MaxInt32) {
		parseas = "Int32"
	}
	if CheckintRange(input, math.MinInt64, math.MaxInt64) {
		parseas = "Int64"
	}
	return STRING, parseas
}

func CheckintRange(input string, min, max int64) bool {
	value, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return false
	}
	return value >= min && value <= max
}

func (SL_Scanner *SL_ScannerStructure) R_DECIMAL() SLTK.SL_TokenConstruct {
	INTEGER, ParseAs := SL_Scanner.R_NUMBER()
	if SL_Scanner.Scanner_Character == '.' && CharacterInputIsDigit(SL_Scanner.CharacterPeek()) {
		SL_Scanner.RCHAR()
		FRAC, _ := SL_Scanner.R_NUMBER()
		return SLTK.SL_TokenConstruct{
			Token_Type: SLTK.TOKEN_FLOAT,
			Literal:    INTEGER + "." + FRAC,
		}
	}
	switch ParseAs {
	case "Int8":
		return SLTK.SL_TokenConstruct{
			Token_Type: SLTK.TOKEN_INTEGER8,
			Literal:    INTEGER,
		}
	case "Int16":
		return SLTK.SL_TokenConstruct{
			Token_Type: SLTK.TOKEN_INTEGER16,
			Literal:    INTEGER,
		}
	case "Int32":
		return SLTK.SL_TokenConstruct{
			Token_Type: SLTK.TOKEN_INTEGER32,
			Literal:    INTEGER,
		}
	case "Int64":
		return SLTK.SL_TokenConstruct{
			Token_Type: SLTK.TOKEN_INTEGER64,
			Literal:    INTEGER,
		}
	default:
		return SLTK.SL_TokenConstruct{
			Token_Type: SLTK.TOKEN_INT,
			Literal:    INTEGER,
		}
	}
}

func (SL_Scanner *SL_ScannerStructure) R_STRING(Character rune) string {
	bytereader := &strings.Builder{}
	for {
		SL_Scanner.RCHAR()
		if SL_Scanner.Scanner_Character == '"' || SL_Scanner.Scanner_Character == 0 {
			break
		}
		if SL_Scanner.Scanner_Character == '\\' {
			SL_Scanner.RCHAR()
			switch SL_Scanner.Scanner_Character {
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
				SL_Scanner.RCHAR()
				SL_Scanner.RCHAR()
				SL_Scanner.RCHAR()
				s := string([]byte{SL_Scanner.Scanner_PreviousCharacter, byte(SL_Scanner.Scanner_Character)})
				bytereader.WriteString(fmt.Sprintf("%x", s))
				continue
			}
			SL_Scanner.RCHAR()
			continue
		}
		bytereader.WriteByte(byte(SL_Scanner.Scanner_Character))
	}
	return bytereader.String()
}

func (SL_Scanner *SL_ScannerStructure) R_IDENTIFIER() string {
	id := ""
	POS := SL_Scanner.Scanner_POS
	RPOS := SL_Scanner.Scanner_RPOS
	for CharacterIsIdentifier(SL_Scanner.Scanner_Character) {
		id += string(SL_Scanner.Scanner_Character)
		SL_Scanner.RCHAR()
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
			SL_Scanner.Scanner_POS = POS
			SL_Scanner.Scanner_RPOS = RPOS
			for offset > 0 {
				SL_Scanner.RCHAR()
				offset--
			}
		}
	}
	return id
}

func ReadVerifyIdentifier(ident string) SLTK.SL_TokenDataType {
	if tok, ok := SLTK.SkyLine_Keywords[ident]; ok {
		return tok
	}
	return SLTK.TOKEN_IDENT
}

func (SL_Scanner *SL_ScannerStructure) ReadBacktick() string {
	bytereader := &strings.Builder{}
	for {
		SL_Scanner.RCHAR()
		if SL_Scanner.Scanner_Character == '`' {
			break
		}
		bytereader.WriteString(string(SL_Scanner.Scanner_Character))
	}
	return bytereader.String()
}

func (SL_Scanner *SL_ScannerStructure) ReadLineNum() int {
	line := 0
	chars := len(SL_Scanner.Scanner_Characters)
	i := 0
	for i < SL_Scanner.Scanner_RPOS && i < chars {
		if SL_Scanner.Scanner_Characters[i] == rune('\n') {
			line++
		}
		i++
	}
	return line
}
