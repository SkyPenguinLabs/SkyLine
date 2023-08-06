//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  _  | |_ _ ___|_|___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|   __| | | | . | |   |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|  |_|___|_  |_|_|_|
//	primary units that define the functions to register the sub func's  //            |___|                               |___|
//////////////////////////////////////////////////////////////////////////
//
// This section of files contain mathemaical functions that can be registered into the SkyLine programming language. Most of these automate the backend of golang's basic
//
// interfaces for math and automate most of the backend. However, some other functions may be algorithmic implementations, specific sets or tweaks that can also be added.
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//				^...^
//			   <_* *_>
//				 \_/
//
// This file defines all of the functions dedicated to foreign function interfaces such as lua or go plugins
//
package SkyLine_Backend_Native_Plugins

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"plugin"
	"reflect"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkyEval "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEvaluator"
	Helpers "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"

	"log"
)

var ReflectMap = map[reflect.Kind]string{
	reflect.Array:         "Array",
	reflect.Bool:          "Boolean",
	reflect.Chan:          "Channel",
	reflect.Complex64:     "Complex64",
	reflect.Complex128:    "Complex128",
	reflect.Float32:       "Float32",
	reflect.Float64:       "Float64",
	reflect.Func:          "Function",
	reflect.Int:           "Int",
	reflect.Int8:          "Int8",
	reflect.Int16:         "Int16",
	reflect.Int32:         "Int32",
	reflect.Int64:         "Int64",
	reflect.Interface:     "Interface",
	reflect.Map:           "Map",
	reflect.Ptr:           "Pointer",
	reflect.Slice:         "Slice",
	reflect.String:        "String",
	reflect.Struct:        "Struct",
	reflect.Uint:          "Uint",
	reflect.Uint8:         "Uint8",
	reflect.Uint16:        "Uint16",
	reflect.Uint32:        "Uint32",
	reflect.Uint64:        "Uint64",
	reflect.Uintptr:       "Uintptr",
	reflect.UnsafePointer: "UnsafePointer",
}

type PluginSession struct {
	LoadedPluginFile string
	CurrentPlugin    *plugin.Plugin
}

var (
	VariableList = make(map[string]interface{})
	PlugSes      PluginSession
)

func GenerateSomeRandomName() string {
	randomBytes := make([]byte, 6)
	_, x := rand.Read(randomBytes)
	if x != nil {
		return "null:error"
	}
	randomID := base64.RawURLEncoding.EncodeToString(randomBytes)
	randomID = randomID[:6]
	return randomID
}

func InitatePlugin(Environ *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"InitiatePlugin",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	// Now initate the plugin
	var x error
	PlugSes.LoadedPluginFile = args[0].(*SkyEnv.SL_String).Value
	PlugSes.CurrentPlugin, x = plugin.Open(PlugSes.LoadedPluginFile)
	if x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	return &SkyEnv.SL_Boolean{Value: true}
}

func ExecutePluginFunction(Environ *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		// execute foreign function
		"ExecuteFF",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	symbol := args[0].(*SkyEnv.SL_String).Value
	// now parse the type
	if PlugSes.CurrentPlugin != nil {
		// Now lets go ahead and see if it exists
		Value, x := PlugSes.CurrentPlugin.Lookup(symbol)
		if x != nil {
			return &SkyEnv.SL_Error{Message: x.Error()}
		}
		FunctionValues := reflect.ValueOf(Value)
		FunctionSignature := FunctionValues.Type()
		if ReflectMap[FunctionSignature.Kind()] != "Function" {
			return &SkyEnv.SL_Error{Message: "[!] Error: plugin symbol " + symbol + " is not a function"}
		}

		expectedFnType := reflect.TypeOf((*func(*SkyEnv.SkyLineEnvironment, ...SkyEnv.SL_Object) SkyEnv.SL_Object)(nil)).Elem()
		if FunctionSignature != expectedFnType {
			fmt.Println("Function type -> ", FunctionSignature)
			log.Fatalf("symbol has unexpected function signature")
		}
		argValues := make([]reflect.Value, 0, len(args)+1)
		argValues = append(argValues, reflect.ValueOf(Environ))
		for _, arg := range args {
			argValues = append(argValues, reflect.ValueOf(arg))
		}
		resultValues := FunctionValues.Call(argValues)
		return resultValues[0].Interface().(SkyEnv.SL_Object)
	} else {
		return &SkyEnv.SL_Error{Message: "Plugin buffer was null, please use `InitiatePlugin` to create one"}
	}
}

func ExecutePluginLoadBuffer(Environ *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"PluginPopLoadBuff",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	// loading file
	Plugin := args[0].(*SkyEnv.SL_String).Value
	if Plugin != "" {
		// replace plugin
		var x error
		PlugSes.LoadedPluginFile = Plugin
		PlugSes.CurrentPlugin, x = plugin.Open(PlugSes.LoadedPluginFile)
		if x != nil {
			return &SkyEnv.SL_Error{Message: x.Error()}
		}
		return &SkyEnv.SL_Boolean{Value: true}
	} else {
		return &SkyEnv.SL_Error{Message: "Plugin value was empty, failed to reload plugin buffer"}
	}
}

// Functions point is to view every exported variable and its name
func ExecutePluginEnvironmentBufferView(Environ *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	// view and return nothing
	for k, v := range VariableList {
		println(k, v)
	}
	return &SkyEnv.SL_NULL{}
}

func ExecutePluginExtractAllVarList(Environ *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"PluginExtractVarList",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	SymbolsStrings := make([]string, 0)
	argsarr := args[0].(*SkyEnv.SL_Array).Elements
	for _, v := range argsarr {
		if valuestr, ok := v.(*SkyEnv.SL_String); ok {
			SymbolsStrings = append(SymbolsStrings, valuestr.Value)
		}
	}
	if PlugSes.CurrentPlugin != nil {
		for i := 0; i < len(SymbolsStrings); i++ {
			Value, x := PlugSes.CurrentPlugin.Lookup(SymbolsStrings[i])
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			ExtractValue := reflect.ValueOf(Value)
			ExtractSignature := ExtractValue.Type()
			if ReflectMap[ExtractSignature.Kind()] != "Pointer" {
				return &SkyEnv.SL_Error{Message: "[!] Error: plugin symbol " + SymbolsStrings[i] + " is not a pointer"}
			}
			name := GenerateSomeRandomName()
			fmt.Println("registering name - ", name)
			interfaceable := ExtractValue.Elem().Interface()
			if v, ok := interfaceable.(string); ok {
				VariableList[name] = v
				SkyEval.RegisterVariable(name, &SkyEnv.SL_String{Value: v})
			} else if v, ok := interfaceable.([]string); ok {
				VariableList[name] = v
				newarr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					newarr = append(newarr, &SkyEnv.SL_String{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: newarr})
			} else if v, ok := interfaceable.(int); ok {
				VariableList[name] = v
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Integer{Value: v})
			} else if v, ok := interfaceable.([]int); ok {
				VariableList[name] = v
				IntegerArr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					IntegerArr = append(IntegerArr, &SkyEnv.SL_Integer{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: IntegerArr})
			} else if v, ok := interfaceable.([]int8); ok {
				VariableList[name] = v
				Integer8Arr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					Integer8Arr = append(Integer8Arr, &SkyEnv.SL_Integer8{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: Integer8Arr})
			} else if v, ok := interfaceable.([]int16); ok {
				VariableList[name] = v
				Integer16Arr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					Integer16Arr = append(Integer16Arr, &SkyEnv.SL_Integer16{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: Integer16Arr})
			} else if v, ok := interfaceable.([]int32); ok {
				VariableList[name] = v
				Integer32Arr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					Integer32Arr = append(Integer32Arr, &SkyEnv.SL_Integer32{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: Integer32Arr})
			} else if v, ok := interfaceable.([]int64); ok {
				VariableList[name] = v
				Integer64Arr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					Integer64Arr = append(Integer64Arr, &SkyEnv.SL_Integer64{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: Integer64Arr})
			} else if v, ok := interfaceable.([]float64); ok {
				VariableList[name] = v
				FloatArr := make([]SkyEnv.SL_Object, 0)
				for i := 0; i < len(v); i++ {
					FloatArr = append(FloatArr, &SkyEnv.SL_Float{Value: v[i]})
				}
				SkyEval.RegisterVariable(name, &SkyEnv.SL_Array{Elements: FloatArr})
			} else {
				return &SkyEnv.SL_Error{Message: "Sorry, that data type is not supported during export"}
			}
		}
	} else {
		return &SkyEnv.SL_Error{Message: "Plugin buffer was null, please use `InitiatePlugin` to create one"}
	}
	return &SkyEnv.SL_Boolean{Value: true}
}
