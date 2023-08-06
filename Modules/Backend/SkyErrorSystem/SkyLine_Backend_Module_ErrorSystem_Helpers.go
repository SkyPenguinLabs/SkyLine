///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
//
// This part of the error system will create error constructs based on the type of error or message or the configuration
//
package SkyLine_Error_System

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	SkyFS "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyFS"
)

// We want the tree to output the following
// Error : This defines the error message
// Type  : if it is a warning or a fatal error
// Tech  : This defines the technology it came from (Parser, Engine-Parser, Engine-Lexer, Regex-Lexer, Regex-Engine, Engine, Configuration, AST, Configuration, Initation etc)
// Line  : Line number or range
// BoxOCode : Box of range of code

func ParseRGB(code string) string {
	code = strings.Trim(code, "#")
	if len(code) != 6 {
		log.Fatal("Code is not RGB")
	}
	decR, _ := strconv.ParseUint(code[0:2], 16, 8)
	decG, _ := strconv.ParseUint(code[2:4], 16, 8)
	decB, _ := strconv.ParseUint(code[4:6], 16, 8)
	r := uint8(decR)
	g := uint8(decG)
	b := uint8(decB)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func CreateStandardErrorTree(Typeof, Technology, Code, Message, Line, Fix, Box string) string {
	var cerr string
	fmt.Println(Box)
	code := ParseRGB("fce803")
	cerr += "\033[38;5;56m%s- \033[38;5;196m[%s]:\033[38;5;242m(\033[38;5;209m%s\033[38;5;242m)\033[39m \033[38;5;242m| %s"
	cerr += "\n          |"
	cerr += "\n \033[38;5;196m[%s] ->  \033[38;5;242m | \033[38;5;196m%s \033[38;5;242m"
	cerr += "\n          |"
	cerr += "\n \033[38;5;86m[L] -> \033[38;5;242m  |\033[38;5;86m %s \033[38;5;242m"
	cerr += "\n          |     "
	if Fix != "" {
		cerr += "\n" + code + " [Fix] \033[38;5;242m   | " + code + " " + Fix + "\033[38;5;242m"
	}
	cerr = fmt.Sprintf(cerr, SkyFS.Current.SL_FileName, Typeof, Technology, Code, Typeof, Message, Line)
	return cerr
}

func CheckArguments(args []string, exact int) bool {
	if len(args) == exact {
		return true
	} else {
		return false
	}
}
