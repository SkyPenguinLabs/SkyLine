///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Console_Models
// Extension         | .go ( golang source code file )
// Purpose           | Define all functions for the frontend work of the REPL
// Directory         | Modules/Backend/SkyConsole
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyConsole
// Package Name      | SkyLine_Backend_Module_Console
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This is known as the REPL otherwise and commonly known as a Read Eval Print Loop which goes straight to executing the next input and spins up a mock environment.
//
// This is not a necessary step within language development and is more of a feature. REPL helps users easily and quickly test environments and allows them to save workspaces
//
// and environments depending on the language. Languages like R-Script allow you to save the workspace while languages like ruby might not allow you to do so as easily or at all.
//
// SkyLine's REPL will be fully customized and much more modern than people may expect!
//
package SkyLine_Backend_Module_Console

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func SkyLine_CodeBoxWithNumFree(filename string) {
	fmt.Print("\033[38;5;249m")
	input, x := ioutil.ReadFile(filename)
	if x != nil {
		println(x)
		os.Exit(0)
	}
	lines := strings.Split(string(input), "\n")
	longest := 0
	for _, line := range lines {
		if len(line) > longest {
			longest = len(line)
		}
	}
	lastLineNumber := len(lines)
	LENLEN := int(math.Log10(float64(lastLineNumber))) + 1
	width := longest + 6 + LENLEN
	fmt.Printf("[SOF] ┌%s┐\n", strings.Repeat("─", width-2))
	for i, line := range lines {
		fmt.Printf("[%*d]   │ %s%s \n", LENLEN, i+1, line, strings.Repeat(" ", longest-len(line)))
	}
	fmt.Printf("[EOF] └%s┘\n", strings.Repeat("─", width-2))
}

func SkyLine_Console_Banner() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("\t\t\t	 \033[38;5;51m┏━┓\x1b[0m")
		fmt.Println("\t\t\t	\033[38;5;56m┃\033[38;5;51m┃ ┃\x1b[0m")
		fmt.Println("\t\t\t    \033[38;5;56m━━━━┛\x1b[0m")
		fmt.Println("\t\t\033[38;5;249m Sky Line Interpreter| V 0.10.0 \033[38;5;56m(Nightly)\x1b[0m")
		fmt.Print("\n\n\n\n\033[39m")
	default:
		fmt.Println("\t\t\t\t	 \u001b[38;5;51m┏━┓\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;56m┃\u001b[38;5;51m┃ ┃\u001b[0m")
		fmt.Println("\t\t\t\t    \u001b[38;5;56m━━━━┛\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;249mSky Line Interpreter| V 0.10.0")
	}
}

// Parses a HTML RGB color code
// without #
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

var SkyLine_KeywordColors = map[string]string{
	"define":   "FFDF00",
	"func":     "00FFFF",
	"Func":     "00FFFF",
	"function": "00FFFF",
	"let":      "FFA500",
	"set":      "FF00FF",
	"cause":    "FF00FF",
	"allow":    "FF00FF",
	"true":     "FFFF00",
	"false":    "FFFF00",
	"if":       "FF9900",
	"else":     "FF9900",
	"return":   "FF6600",
	"ret":      "FF6600",
	"const":    "008080",
	"constant": "008080",
	"switch":   "800080",
	"sw":       "800080",
	"case":     "FF3300",
	"cs":       "FF3300",
	"default":  "FF3366",
	"df":       "FF3366",
	"register": "FF00CC",
	"ENGINE":   "9900FF",
	"import":   "CC3300",
	"for":      "00CCFF",
	"STRING":   "FFFF00",
	"BOOLEANT": "FFFF00",
	"BOOLEANF": "FFFF00",
	":=":       "eb345b",
	"/":        "FF0000", // Deep Red
	"{":        "FF1493", // Magenta Glow
	"(":        "00FF00", // Neon Green
	")":        "00FFFF", // Cyan
	",":        "0000FF", // Blue
	"=":        "8A2BE2", // Blue Violet
	"!-":       "7FFF00", // Chartreuse
	".:":       "FF00FF", // Fuchsia
	":":        "FF4500", // Orange Red
	">=":       "FFFF00", // Yellow
	">":        "00FF7F", // Spring Green
	"<=":       "FFA500", // Orange
	"<":        "FF6347", // Tomato
	"*=":       "FF8C00", // Dark Orange
	"*":        "BDB76B", // Dark Khaki
	"%:":       "8B4513", // Saddle Brown
	"/=":       "006400", // Dark Green
	"-=":       "8FBC8F", // Dark Sea Green
	"--":       "2E8B57", // Sea Green
	"+":        "008080", // Teal
	"++":       "48D1CC", // Medium Turquoise
	"&&":       "20B2AA", // Light Sea Green
	"||":       "FF69B4", // Hot Pink
	"!=":       "9400D3", // Dark Violet
	"`":        "808080", // Gray
}

func SkyLine_Code_Box(line string) string {
	highlightedLine := line
	for key, value := range SkyLine_KeywordColors {
		if strings.HasPrefix(line, key) {
			highlightedLine = strings.ReplaceAll(highlightedLine, key, ParseRGB(strings.ToLower(value))+key+LINUX_GREY)
		} else if strings.Contains(line, key) {
			highlightedLine = strings.ReplaceAll(highlightedLine, key, ParseRGB(strings.ToLower(value))+key+LINUX_GREY)
		}
	}
	if REPL.LineCount >= 10 {
		return fmt.Sprintf("[%d] │ %s ", REPL.LineCount, highlightedLine)
	} else {
		return fmt.Sprintf("[%d]  │ %s ", REPL.LineCount, highlightedLine)
	}
}

func DrawUtilsBox(variable string) {
	BL := BoxFormat{
		TL: "┏",
		TR: "┓",
		BL: "┗",
		BR: "┛",
		HZ: "━",
		VT: "┃",
	}
	l := strings.Split(variable, "\n")
	var mlen int
	for _, lin := range l {
		if len(lin) > mlen {
			mlen = len(lin)
		}
	}
	fmt.Print("\033[38;5;93m" + BL.TL + strings.Repeat(BL.HZ, mlen) + BL.TR + "\n")
	for _, line := range l {
		fmt.Print(BL.VT + "\033[38;5;50m" + line + strings.Repeat(" ", mlen-len(line)) + "\033[38;5;93m" + BL.VT + "\n")
	}
	fmt.Print("\033[38;5;93m" + BL.BL + strings.Repeat(BL.HZ, mlen) + BL.BR + "\n")

}

func SkyLine_Draw_Around_REPL_ENV() {
	var str string
	for _, code := range REPL.Environment {
		str += code + "\n"
	}
	DrawUtilsBox(str)
}
