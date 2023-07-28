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
// Defines       | This file defines all invoke and object call functions for the String data type within the SkyLine programming language
//
//
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//
// This is the registry module, this module apart of SkyLine_Backend allows us to register standard library based functions which are called with
//
// class.functionname
//
// this is pretty simple to understand however class and module keywords have not been implemented which means that you can not register your
// own custom standard module / library. Asides from that, we use the init() function because init functions will always run or be called before
// the main() function in go. Using registers under the init function we can ensure the environment has standard functions registered and placed
// into the environment before it is fully started and the input program is parsed. This eliminates the need to import("math") however in the
// further future import keywords will need to be added for standard library functions. This is becausethe bigger our standard library gets the
// more imports will need to be added and the harder the program will be to parse. Currently, due to the factor of how small the standard library
// is, it is not that bad to register the built in functions before a new environment for the input program is started which means it will not slow
// down runtime. However, as this again gets bigger we will need to eliminate registering before runtime unless they are standard functions such as
// .str, .int, integer, boolean, empty?, nil?, carries?, exported? etc which allow for a much heavier use case and do not require imports
// Using the import keyword will give the user the option to allow the program to import and register the standard library functions before
// runtime and parsing. This may cause collisions within the environment so we can actually cause another keyword to exist known as "register"
// followed by the library name. This keyword may be called like so register("math") pr register<<"math">> for a much more complex and parsed
// syntax. Allowing both register and import keywords allow the user to register the library functions before runtime and import files before
// runtime.
//
//
// - Mon 27 Feb 2023 10:23:16 PM EST
//
// as of the given date and time of writing this, SkyLine now will ask you to register the library before you use them if it is standard
// this includes crypt, math, net, http and much other built in libraries used within the SkyLine programming language
//
// ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This unit is the prime unit for any invoke or object call based functions that get called or pushed into the environment when the SkyEnvironment is loaded
//
// you do not need to register data types unless you are using specific functions that are invokes but not always going to be invoke calls but again are used as one.
//
//
package SkyLine_Module_SkyStandardLib_Direct

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	SkyHelpers "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SwitchMode(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyHelpers.SkyLine_Standard_Library_Helper_CheckArguments("Mode", args, SkyHelpers.SkyLine_Standard_Library_Helper_ExactArguments(1), SkyHelpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x != nil {
		log.Fatal(x)
	}
	arg1 := args[0].(*SkyEnv.SL_String).Value
	if arg1 != "" {
		switch arg1 {
		case "pwn":
			InitateCallRegisterPwnFuncs()
		}
	} else {
		return &SkyEnv.SL_Error{Message: "Invalid mode and argument, mode must be supported and not an empty string"}
	}
	return &SkyEnv.SL_NULL{}
}

func Version(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_String{Value: "0.10.0"}
}

func Arguments(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	l := len(os.Args[1:])
	result := make([]SkyEnv.SL_Object, l)
	for i, txt := range os.Args[1:] {
		result[i] = &SkyEnv.SL_String{Value: txt}
	}
	return &SkyEnv.SL_Array{Elements: result}
}

func Println(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	for _, k := range args {
		fmt.Println(k.SkyLine_ObjectFunction_GetTrueValue())
	}
	return &SkyEnv.SL_NULL{}
}

func Print(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	for _, k := range args {
		fmt.Print(k.SkyLine_ObjectFunction_GetTrueValue())
	}
	return &SkyEnv.SL_NULL{}
}

func Sprint(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if len(args) == 0 {
		return &SkyEnv.SL_Error{Message: "Sprint takes at least one positional argument"}
	}
	var returnedout string
	for _, k := range args {
		returnedout += fmt.Sprint(k.SkyLine_ObjectFunction_GetTrueValue())
	}
	return &SkyEnv.SL_String{Value: returnedout}
}

func GetDataType(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if len(args) == 0 {
		return &SkyEnv.SL_Error{Message: "Typeof takes at least one positional argument, but only one"}
	}
	return &SkyEnv.SL_String{
		Value: string(args[0].SkyLine_ObjectFunction_GetDataType()),
	}
}

func convertToInteger(arg *SkyEnv.SL_Object) (SkyEnv.SL_Object, error) {
	switch c := (*arg).(type) {
	case *SkyEnv.SL_Integer16:
		return &SkyEnv.SL_Integer{Value: int(c.Value)}, nil
	case *SkyEnv.SL_Integer32:
		return &SkyEnv.SL_Integer{Value: int(c.Value)}, nil
	case *SkyEnv.SL_Integer64:
		return &SkyEnv.SL_Integer{Value: int(c.Value)}, nil
	case *SkyEnv.SL_Integer8:
		return &SkyEnv.SL_Integer{Value: int(c.Value)}, nil
	case *SkyEnv.SL_String:
		conv, err := strconv.ParseInt(c.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		return &SkyEnv.SL_Integer{Value: int(conv)}, nil
	default:
		return nil, errors.New("Sorry, int() takes only integer types and string types")
	}
}

func ConvertInt8_16_32_64_to_int(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyHelpers.SkyLine_Standard_Library_Helper_CheckArguments("int", args, SkyHelpers.SkyLine_Standard_Library_Helper_ExactArguments(1)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}

	result, err := convertToInteger(&args[0])
	if err != nil {
		return &SkyEnv.SL_Error{Message: err.Error()}
	}

	return result
}

func ConvertInt8_16_32_64_to_int8(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyHelpers.SkyLine_Standard_Library_Helper_CheckArguments("int", args, SkyHelpers.SkyLine_Standard_Library_Helper_ExactArguments(1)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}

	result, err := convertToInteger(&args[0])
	if err != nil {
		return &SkyEnv.SL_Error{Message: err.Error()}
	}

	if intValue, ok := result.(*SkyEnv.SL_Integer); ok {
		return &SkyEnv.SL_Integer8{Value: int8(intValue.Value)}
	}

	return &SkyEnv.SL_Error{Message: "Sorry, int() takes only integer types and string types"}
}

//

func Sprintf(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if len(args) == 0 {
		return &SkyEnv.SL_String{Value: ""}
	}

	format := args[0].SkyLine_ObjectFunction_GetTrueValue()
	formattedArgs := make([]interface{}, len(args)-1)
	for i := 1; i < len(args); i++ {
		formattedArgs[i-1] = convertToNativeType(args[i])
	}

	result := fastSprintf(format, formattedArgs...)
	return &SkyEnv.SL_String{Value: result}
}

func convertToNativeType(arg SkyEnv.SL_Object) interface{} {
	switch v := arg.(type) {
	case *SkyEnv.SL_Integer:
		return v.Value
	case *SkyEnv.SL_Float:
		return v.Value
	case *SkyEnv.SL_Boolean:
		return v.Value
	case *SkyEnv.SL_String:
		return v.Value
	default:
		return fmt.Sprintf("%v", arg)
	}
}

func fastSprintf(format string, args ...interface{}) string {
	var builder strings.Builder
	builder.Grow(len(format) + 32) // Preallocate initial capacity to avoid resizing

	argIndex := 0
	runes := []rune(format)
	for i := 0; i < len(runes); i++ {
		if runes[i] != '%' {
			builder.WriteRune(runes[i])
			continue
		}

		if i == len(runes)-1 {
			builder.WriteRune(runes[i])
			break
		}

		switch runes[i+1] {
		case 'd', 'i':
			builder.WriteString(strconv.FormatInt(toInt64(args[argIndex]), 10))
			argIndex++
		case 'o':
			builder.WriteString(strconv.FormatInt(toInt64(args[argIndex]), 8))
			argIndex++
		case 'x':
			builder.WriteString(strconv.FormatInt(toInt64(args[argIndex]), 16))
			argIndex++
		case 'X':
			builder.WriteString(strings.ToUpper(strconv.FormatInt(toInt64(args[argIndex]), 16)))
			argIndex++
		case 'f':
			builder.WriteString(strconv.FormatFloat(toFloat64(args[argIndex]), 'f', -1, 64))
			argIndex++
		case 'e', 'E':
			builder.WriteString(strconv.FormatFloat(toFloat64(args[argIndex]), 'e', -1, 64))
			argIndex++
		case 'g', 'G':
			builder.WriteString(strconv.FormatFloat(toFloat64(args[argIndex]), 'g', -1, 64))
			argIndex++
		case 's':
			builder.WriteString(fmt.Sprintf("%v", args[argIndex]))
			argIndex++
		case 't':
			builder.WriteString(strings.Title(fmt.Sprintf("%v", args[argIndex])))
			argIndex++
		case '%':
			builder.WriteRune('%')
		default:
			builder.WriteRune(runes[i])
			continue
		}

		i++
	}

	return builder.String()
}

func toInt64(arg interface{}) int64 {
	switch v := arg.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	default:
		return 0
	}
}

func toFloat64(arg interface{}) float64 {
	switch v := arg.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}
