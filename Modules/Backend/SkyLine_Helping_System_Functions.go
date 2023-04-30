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
// Filename      |  SkyLine_Helping_System_Functions.go
// Project       |  SkyLine programming language
// Line Count    |  150+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | Defines function to reverse arrays, reverse file traceback, check argument lengths and much more
//
// STATE         | Needs to be moved
// Resolution    | Functions need to be erased and re written as well as moved into the appropriate files
//
package SkyLine_Backend

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ReverseArrayForFileTraceback(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func ReverseObjectArrayWithinCSCF(a []SLC_Object) {
	for k := len(a)/2 - 1; k >= 0; k-- {
		opp := len(a) - 1 - k
		a[k], a[opp] = a[opp], a[k]
	}
}

func CheckArgumentLength(ProperLength int, name string, Args []SLC_Object) {
	if len(Args) != ProperLength {
		fmt.Printf("Function %s must have %d positional arguments, function call got %d", name, ProperLength, len(Args))
		os.Exit(1)
	}
}

type ParseError struct {
	Value string
	Type  string
	Err   error
}

// Checks a parsed integer based error

func CheckParseValue(input string, errType string) (int64, *ParseError) {
	var (
		value  int64
		x      error
		errStr string
	)

	if x != nil {
		switch e := x.(type) {
		case *strconv.NumError:
			errStr = e.Err.Error()
		default:
			errStr = x.Error()
		}

		return 0, &ParseError{Value: input, Type: errType, Err: errors.New(errStr)}
	}

	if (errType == "binary" && len(input) > 66) ||
		(errType == "hex" && len(input) > 18) ||
		(len(input) > 19) {
		return 0, &ParseError{Value: input, Type: "length", Err: errors.New("value too long")}
	} else if value < math.MinInt64 || value > math.MaxInt64 {
		return 0, &ParseError{Value: input, Type: "range", Err: errors.New("value out of range")}
	}

	return value, nil
}

// Checks the type of integer and error then returns a message
type ParsedErrorInputParserParam struct {
	Fix     string
	Max     string
	Lowest  string
	Suggest string
}

func CheckParsedError(input string, number string) *ParsedErrorInputParserParam {
	var Caller ParsedErrorInputParserParam
	Caller.Max = "2^63 - 1"
	Caller.Lowest = "-2^63"
	if input == "value too long" {
		Caller.Suggest = "Try shortening the value within the range of int64 "
	} else if input == "value out of range" {
		Caller.Suggest = "Try making the value longer, as the value is too short to be int64 (LOWEST)"
	}
	return &Caller
}

// Checks the type of float and error then returns a message

type FloatParseError struct {
	TooShort bool
	TooLong  bool
	Parsed   bool
	Recomend string
	Max      int
	Low      int
}

func CheckAndVerify(number string) FloatParseError {
	M := 308
	S := 1
	if len(number) > M {
		return FloatParseError{
			Parsed:   false,
			TooLong:  true,
			TooShort: false,
			Recomend: "Make sure the float value given is shorter than the max (MAX)",
			Low:      S,
			Max:      M,
		}
	} else if len(number) < S {
		return FloatParseError{
			Parsed:   false,
			TooLong:  false,
			TooShort: true,
			Recomend: "Make sure the float value given is longer than the minimum (MIN)",
			Low:      S,
			Max:      M,
		}
	} else {
		// Wtf happened during parsing????
		return FloatParseError{
			Parsed:   false,
			TooShort: false,
			TooLong:  false,
			Low:      S,
			Max:      M,
		}
	}
}
