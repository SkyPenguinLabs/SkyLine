///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Scanner_IsVerifiers
// Extension         | .go ( golang source code file )
// Purpose           | Define verification functions during lexical analysis
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines functions that can verify the type of a given byte by comparing its ASCII values. For example if the byte number is greater than or equal to  b(a) and b(z) then we
//
// can say that it is a character, if it is greater than or equal to b(1) || b(9) then we can say that it is a digit. This can help categorize the type of byte so the scanner can know
//
// exactly what action or step to take based on a given series of conditions that are met when comparing. A table shows ASCII characters these verifiers use
//
// hex  | character
// ---- | ---------
// 0x61 | a
// 0x62 | b
// 0x63 | c
// 0x64 | d
// 0x65 | e
// 0x66 | f
// 0x67 | g
// 0x68 | h
// 0x69 | i
// 0x6A | j
// 0x6B | k
// 0x6C | l
// 0x6D | m
// 0x6E | n
// 0x6F | o
// 0x70 | p
// 0x71 | q
// 0x72 | r
// 0x73 | s
// 0x74 | t
// 0x75 | u
// 0x76 | v
// 0x77 | w
// 0x78 | x
// 0x79 | y
// 0x7A | z
// 0x41 | A
// 0x42 | B
// 0x43 | C
// 0x44 | D
// 0x45 | E
// 0x46 | F
// 0x47 | G
// 0x48 | H
// 0x49 | I
// 0x4A | J
// 0x4B | K
// 0x4C | L
// 0x4D | M
// 0x4E | N
// 0x4F | O
// 0x50 | P
// 0x51 | Q
// 0x52 | R
// 0x53 | S
// 0x54 | T
// 0x55 | U
// 0x56 | V
// 0x57 | W
// 0x58 | X
// 0x59 | Y
// 0x5A | Z
// 0x30 | 0
// 0x31 | 1
// 0x32 | 2
// 0x33 | 3
// 0x34 | 4
// 0x35 | 5
// 0x36 | 6
// 0x37 | 7
// 0x38 | 8
// 0x39 | 9
package SkyLine_Backend_Scanner

import "unicode"

func CharacterInputIsLetter(Character rune) bool {
	return 'a' <= Character && Character <= 'z' || 'A' <= Character && Character <= 'Z' || Character == '_'
}

func CharacterInputIsDigit(Character rune) bool {
	return '0' <= Character && Character <= '9'
}

func CharacterIsIdentifier(Character rune) bool {
	if unicode.IsLetter(Character) || unicode.IsDigit(Character) || Character == '.' || Character == '?' || Character == '$' || Character == '_' {
		return true
	}
	return false
}

func CharacterIsWhiteSpace(Character rune) bool {
	return Character == rune(' ') || Character == rune('\t') || Character == rune('\n') || Character == rune('\r')
}

func CharacterIsEscapedCharacter(Character rune) bool {
	return Character == '\\'
}
