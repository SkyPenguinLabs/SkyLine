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
package SkyLine_Module_Std_Lib_Invokes

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkyEnvironment "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func String_BArr(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	Vals := []byte(Value)
	arr := make([]SkyEnv.SL_Object, 0)
	for i := 0; i < len(Vals); i++ {
		arr = append(arr, &SkyEnvironment.SL_String{Value: fmt.Sprint(Vals[i])})
	}
	return &SkyEnvironment.SL_Array{Elements: arr}
}

func String_Outln(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	fmt.Println(Value)
	return &SkyEnvironment.SL_NULL{}
}

func String_Length(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_Integer{Value: len(Value)}
}

func String_Ord(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_Integer{Value: int(Value[0])}
}

func String_Upper(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_String{Value: strings.ToUpper(Value)}
}

func String_Lower(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_String{Value: strings.ToLower(Value)}
}

func String_ToInt(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	Integer, x := strconv.ParseInt(Value, 0, 64)
	if x != nil {
		log.Fatal(x)
	}
	return &SkyEnv.SL_Integer{Value: int(Integer)}
}

func String_ToBoolean(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	Boolean, x := strconv.ParseBool(Value)
	if x != nil {
		log.Fatal(x)
	}
	return &SkyEnv.SL_Boolean{Value: Boolean}
}

type Formatter struct {
	format string
}

func (f *Formatter) Format(args string) {
	placeholder := "{}"
	f.format = strings.Replace(f.format, placeholder, fmt.Sprintf("%v", args), 1)
}

func String_Format(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	formatstr := Value
	x := Formatter{
		format: formatstr,
	}
	for _, arg := range args {
		x.Format(arg.SkyLine_ObjectFunction_GetTrueValue())

	}
	return &SkyEnv.SL_String{
		Value: x.format,
	}
}

func String_ToFloat(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	Float64, x := strconv.ParseFloat(Value, 64)
	if x != nil {
		log.Fatal(x)
	}
	return &SkyEnv.SL_Float{Value: Float64}
}

func String_Split(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyLine_Standard_Library_Helper_CheckArguments("STRING.Split", args, SkyLine_Standard_Library_Helper_ExactArguments(1), SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	Arr := strings.Split(Value, args[0].SkyLine_ObjectFunction_GetTrueValue())
	NewerArr := make([]SkyEnv.SL_Object, 0)
	for _, idx := range Arr {
		NewerArr = append(NewerArr, &SkyEnv.SL_String{Value: idx})
	}
	return &SkyEnv.SL_Array{Elements: NewerArr}
}

func String_Methods(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	Dynamic_Environment := Env.Names("string.")
	var Callable []string
	Callable = append(Callable, Static_String_Methods...)
	for _, envi := range Dynamic_Environment {
		bits := strings.Split(envi, ".")
		Callable = append(Callable, bits[1])
	}
	sort.Strings(Callable)
	result := make([]SkyEnvironment.SL_Object, len(Callable))
	for idx, txt := range Callable {
		result[idx] = &SkyEnvironment.SL_String{Value: txt}
	}
	return &SkyEnvironment.SL_Array{Elements: result}
}

func String_Hex(Value string, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	fmt.Println("Value : ", Value)
	return &SkyEnvironment.SL_String{Value: fmt.Sprintf("%x", Value)}
}

func String_ToHex(Value string, Env SkyEnvironment.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	ObjArr := make([]string, 0)
	if Value != "" {
		codes := strings.Split(Value, " ")
		for _, code := range codes {
			hexCode := fmt.Sprintf("0x%s", code)
			ObjArr = append(ObjArr, hexCode)
		}
	} else {
		return nil
	}
	return &SkyEnvironment.SL_String{Value: strings.Join(ObjArr, ", ")}
}

func String_SubString(Value string, Env SkyEnvironment.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyLine_Standard_Library_Helper_CheckArguments("STRING.Subtr", args, SkyLine_Standard_Library_Helper_ExactArguments(2)); x != nil {
		return &SkyEnvironment.SL_Error{Message: x.Error()}
	} else {
		if CheckIntegerType(args[0]) && CheckIntegerType(args[1]) {
			var start int
			var length int
			switch s := args[0].(type) {
			case *SkyEnv.SL_Integer:
				start = s.Value
			case *SkyEnv.SL_Integer8:
				start = int(s.Value)
			case *SkyEnv.SL_Integer16:
				start = int(s.Value)
			case *SkyEnv.SL_Integer32:
				start = int(s.Value)
			case *SkyEnv.SL_Integer64:
				start = int(s.Value)
			}
			switch s := args[1].(type) {
			case *SkyEnv.SL_Integer:
				length = s.Value
			case *SkyEnv.SL_Integer8:
				length = int(s.Value)
			case *SkyEnv.SL_Integer16:
				length = int(s.Value)
			case *SkyEnv.SL_Integer32:
				length = int(s.Value)
			case *SkyEnv.SL_Integer64:
				length = int(s.Value)
			}
			fmt.Println("Value -> ", start, length)
			end := start + length
			if start < 0 || start >= len(Value) || end > len(Value) {
				return &SkyEnvironment.SL_NULL{}
			}
			return &SkyEnvironment.SL_String{Value: Value[start:end]}
		} else {
			return &SkyEnvironment.SL_Error{Message: "Function `SubStr()` takes two arguments of type Integer | 8/16/32/64"}
		}
	}
}
