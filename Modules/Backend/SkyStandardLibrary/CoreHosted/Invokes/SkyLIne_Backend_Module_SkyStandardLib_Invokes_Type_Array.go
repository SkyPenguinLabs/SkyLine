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
// Defines       | This file defines all invoke and object call functions for the Array data type within the SkyLine programming language
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
	"encoding/hex"
	"fmt"
	"log"
	"strconv"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func DEPEND_INVOKE_ARRAY_COPY(elems []SkyEnv.SL_Object, Array *SkyEnv.SL_Array) SkyEnv.SL_Object {
	Array.Elements = append(Array.Elements, elems...)
	return &SkyEnv.SL_NULL{}
}

func DEPEND_INVOKE_ARRAY_REVERSE(Elements []SkyEnv.SL_Object) *SkyEnv.SL_Array {
	for idx, jex := 0, len(Elements)-1; idx < jex; idx, jex = idx+1, jex-1 {
		Elements[idx], Elements[jex] = Elements[jex], Elements[idx]
	}
	return &SkyEnv.SL_Array{Elements: Elements}
}

func DEPEND_INVOKE_ARRAY_PREPEND(Element SkyEnv.SL_Object, SL_Array *SkyEnv.SL_Array) {
	SL_Array.Elements = append([]SkyEnv.SL_Object{Element}, SL_Array.Elements...)
}

func DEPEND_INVOKE_ARRAY_APPEND(Element SkyEnv.SL_Object, SL_Array *SkyEnv.SL_Array) {
	SL_Array.Elements = append(SL_Array.Elements, Element)
}

func DEPEND_INVOKE_ARRAY_POPLEFT(SL_Array *SkyEnv.SL_Array) SkyEnv.SL_Object {
	if len(SL_Array.Elements) > 0 {
		e := SL_Array.Elements[0]
		SL_Array.Elements = SL_Array.Elements[1:]
		return e
	}
	return &SkyEnv.SL_NULL{}
}

func DEPEND_INVOKE_ARRAY_POPRIGHT(SL_Array *SkyEnv.SL_Array) SkyEnv.SL_Object {
	if len(SL_Array.Elements) > 0 {
		e := SL_Array.Elements[(len(SL_Array.Elements) - 1)]
		SL_Array.Elements = SL_Array.Elements[:(len(SL_Array.Elements) - 1)]
		return e
	}
	return &SkyEnv.SL_NULL{}
}

func DEPEND_INVOKE_ARRAY_SWAP(F_elem, S_elem int, SL_Array *SkyEnv.SL_Array) {
	SL_Array.Elements[F_elem], SL_Array.Elements[S_elem] = SL_Array.Elements[S_elem], SL_Array.Elements[F_elem]
}

func Array_Typeof(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_String{Value: SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT}
}

func Array_Length(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_Integer{Value: len(ArrayObject.Elements)}
}

func Array_View(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_String{Value: ArrayObject.SkyLine_ObjectFunction_GetTrueValue()}
}

func Array_Prepend(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyLine_Standard_Library_Helper_CheckArguments("ARRAY.Append", args, SkyLine_Standard_Library_Helper_ExactArguments(1)); x == nil {
		DEPEND_INVOKE_ARRAY_APPEND(args[0], ArrayObject)
		return &SkyEnv.SL_NULL{}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func Array_Append(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyLine_Standard_Library_Helper_CheckArguments("ARRAY.Append", args, SkyLine_Standard_Library_Helper_ExactArguments(1)); x == nil {
		DEPEND_INVOKE_ARRAY_APPEND(args[0], ArrayObject)
		return &SkyEnv.SL_NULL{}
	} else {
		return &SkyEnv.SL_Error{Message: fmt.Sprint(x)}
	}
}

func Array_PopRight(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return DEPEND_INVOKE_ARRAY_POPRIGHT(ArrayObject)
}

func Array_PopLeft(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return DEPEND_INVOKE_ARRAY_POPLEFT(ArrayObject)
}

func Array_Reverse(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return DEPEND_INVOKE_ARRAY_REVERSE(ArrayObject.Elements)
}

func Array_Copy(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return DEPEND_INVOKE_ARRAY_COPY(ArrayObject.Elements, ArrayObject)
}

func Array_Swap(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyLine_Standard_Library_Helper_CheckArguments("ARRAY.Swap", args, SkyLine_Standard_Library_Helper_ExactArguments(2), SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT)); x == nil {
		Swap1, x := strconv.Atoi(args[0].SkyLine_ObjectFunction_GetTrueValue())
		if x != nil {
			log.Fatal(x)
		}
		Swap2, x := strconv.Atoi(args[1].SkyLine_ObjectFunction_GetTrueValue())
		if x != nil {
			log.Fatal(x)
		}
		DEPEND_INVOKE_ARRAY_SWAP(Swap1, Swap2, ArrayObject)
	}
	return &SkyEnv.SL_NULL{}
}

func Array_Hex(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	arr := make([]byte, 0)
	for i := 0; i < len(Values); i++ {
		switch t := Values[i].(type) {
		case *SkyEnv.SL_Integer:
			arr = append(arr, byte(t.Value))
		case *SkyEnv.SL_Integer8:
			arr = append(arr, byte(t.Value))
		case *SkyEnv.SL_Integer16:
			arr = append(arr, byte(t.Value))
		case *SkyEnv.SL_Integer32:
			arr = append(arr, byte(t.Value))
		case *SkyEnv.SL_Integer64:
			arr = append(arr, byte(t.Value))
		case *SkyEnv.SL_String:
			if num, x := strconv.Atoi(t.Value); x == nil {
				arr = append(arr, byte(num))
			} else if num, x := strconv.ParseUint(t.Value, 0, 8); x == nil {
				arr = append(arr, byte(num))
			}
		}
	}
	return &SkyEnv.SL_String{Value: hex.EncodeToString(arr)}
}

func Array_Ascii(Values []SkyEnv.SL_Object, ArrayObject *SkyEnv.SL_Array, Env SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := SkyLine_Standard_Library_Helper_CheckArguments("ARRAY.ToAscii", args, SkyLine_Standard_Library_Helper_ExactArguments(0), SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT)); x == nil {
		var arr []int
		for i := 0; i < len(Values); i++ {
			switch t := Values[i].(type) {
			case *SkyEnv.SL_Integer:
				arr = append(arr, int(t.Value))
			case *SkyEnv.SL_Integer8:
				arr = append(arr, int(t.Value))
			case *SkyEnv.SL_Integer16:
				arr = append(arr, int(t.Value))
			case *SkyEnv.SL_Integer32:
				arr = append(arr, int(t.Value))
			case *SkyEnv.SL_Integer64:
				arr = append(arr, int(t.Value))
			case *SkyEnv.SL_String:
				if num, x := strconv.Atoi(t.Value); x == nil {
					arr = append(arr, num)
				}
			}
		}
		return &SkyEnv.SL_String{Value: ConvertToAscii(arr)}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func ConvertToAscii(input []int) string {
	converted := make([]byte, len(input))
	for i, num := range input {
		converted[i] = byte(num)
	}
	return string(converted)
}
