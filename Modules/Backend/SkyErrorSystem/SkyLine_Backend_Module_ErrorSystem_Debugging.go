///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
//
// This module plays a bit differently. Instead of just using standard error functions and basic algorithms to catch errors, we take a much more different approach. This module
//
// lays out specific situations where code can be debugged, suggested with fixes or even improve current and existing code. This also allows the debugging system to check
//
// for recomendations on what the user may have actually meant when working with the errors.
//
// ::::: Example
//
// Say we have a problem with identifiers not being properly found or called in the environment like Object(String).split()  | If this becomes an error, the interpreter will
//
// try to match it with current symbols in the environment and see if those symbols may match any existing ones or within the database for libraries.
//
//
package SkyLine_Error_System

import (
	"fmt"
	"regexp"
	"strings"
)

func Debug_Identifier(input string) {
	split := strings.Split(input, ".")
	fmt.Println(split)
	var f bool
	for i := 0; i < len(ErrorConstEnv.LibraryModules); i++ {
		if split[0] == ErrorConstEnv.LibraryModules[i] {
			fmt.Println("Module [ " + split[0] + " ] exists")
			f = true
			break
		} else {
			f = false
		}
	}
	if !f {
		fmt.Println("Module not found in STD -> ", split[0])
		matched := false
		for _, name := range ErrorConstEnv.LibraryIdentifiers {
			matched, _ = regexp.MatchString(fmt.Sprintf(`\b%s\b`, strings.Split(name, ".")[1]), split[1])
			if matched {
				break
			}
		}
		if matched {
			fmt.Printf("Function name '%s' is recognized.\n", input)
		} else {
			fmt.Printf("Function name '%s' is not recognized.\n", input)
			suggestions := SuggestFunctionNames(input)
			if len(suggestions) > 0 {
				fmt.Printf("Did you mean: %s?\n", strings.Join(suggestions, ", "))
			}
		}
	}
}

func SuggestFunctionNames(funcName string) []string {
	var suggestions []string
	for _, name := range ErrorConstEnv.LibraryIdentifiers {
		if strings.HasPrefix(name, funcName) {
			suggestions = append(suggestions, name)
		}
	}
	return suggestions
}
