//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____     _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     |___| |_| |_
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___| | | | .'|  _|   |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_|_|_|__,|_| |_|_|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
// This section of files contain mathemaical functions that can be registered into the SkyLine programming language. Most of these automate the backend of golang's basic
//
// interfaces for math and automate most of the backend. However, some other functions may be algorithmic implementations, specific sets or tweaks that can also be added.
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the functions dedicated to randomization within the language
//
package SkyLine_Standard_Library_Math

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strconv"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

var (
	Charlist string
	numgen   int
	length   int
)

func Math_Randomization_Chars(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Rand_chars", 3, SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, arguments...); x == nil {
		Charlist = arguments[0].SkyLine_ObjectFunction_GetTrueValue()
		conv, x := strconv.Atoi(arguments[1].SkyLine_ObjectFunction_GetTrueValue())
		if x != nil {
			log.Fatal(x)
		}
		numgen = conv

		if Charlist == "" {
			Charlist = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+{}|[]\\:\";'<>?,./"
		}
		if numgen == 0 {
			numgen = 1
		}
		if length == 0 {
			length = 10
		}
		var Keys []string
		for Route := 0; Route < numgen; Route++ {
			var pass string
			for idx := 0; idx < length; idx++ {
				pass += string(Charlist[rand.Intn(len(Charlist))])
			}
			Keys = append(Keys, pass)
		}
		strs := make([]SkyEnv.SL_Object, 0)
		for i := 0; i < len(Keys); i++ {
			strs = append(strs, &SkyEnv.SL_String{Value: Keys[i]})
		}
		return &SkyEnv.SL_Array{Elements: strs}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Random_Choice(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Rand_choice", 1, SkyEnv.SKYLINE_DATATYPE_ARRAY_OBJECT, arguments...); x == nil {
		Array := make([]interface{}, 0)
		for i := 0; i < len(arguments[0].(*SkyEnv.SL_Array).Elements); i++ {
			Array = append(Array, arguments[0].(*SkyEnv.SL_Array).Elements[i])
		}
		choice := Array[rand.Intn(len(Array))]
		switch c := choice.(type) {
		case string:
			return &SkyEnv.SL_String{Value: c}
		case int:
			return &SkyEnv.SL_Integer{Value: c}
		case float64:
			return &SkyEnv.SL_Float{Value: c}
		case bool:
			return &SkyEnv.SL_Boolean{Value: c}
		default:
			return &SkyEnv.SL_Error{Message: "Sorry, return type is currently unsupported \n Supported | (string, int, float, bool)"}
		}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Random_Guassian(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Rand_gua", 2, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, arguments...); x == nil {
		gaussian := rand.NormFloat64()*arguments[0].(*SkyEnv.SL_Float).Value + arguments[1].(*SkyEnv.SL_Float).Value
		return &SkyEnv.SL_Float{Value: gaussian}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Random_Uniform(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Rand_uniform", 2, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, arguments...); x == nil {
		uniform := rand.Float64()*(arguments[1].(*SkyEnv.SL_Float).Value-arguments[0].(*SkyEnv.SL_Float).Value) + arguments[0].(*SkyEnv.SL_Float).Value
		return &SkyEnv.SL_Float{Value: uniform}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Random_Triangular(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Rand_tri", 3, SkyEnv.SKYLINE_DATATYPE_FLOAT_OBJECT, arguments...); x == nil {
		// Low
		L := arguments[0].(*SkyEnv.SL_Float).Value
		// High
		H := arguments[1].(*SkyEnv.SL_Float).Value
		// Mode
		M := arguments[2].(*SkyEnv.SL_Float).Value
		u := rand.Float64()
		if u < (M-L)/(H-L) {
			triangular := L + ((H - L) * u * ((M - L) / (H - L)))
			return &SkyEnv.SL_Float{Value: triangular}
		} else {
			triangular := H - ((H - L) * (1 - u) * ((H - M) / (H - L)))
			return &SkyEnv.SL_Float{Value: triangular}
		}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Random_Float(arguments ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	return &SkyEnv.SL_Float{Value: rand.Float64()}
}

func Math_Random_Sample(population []int) []int {
	sample := make([]int, 3)
	for i := range sample {
		index := rand.Intn(len(population))
		sample[i] = population[index]
		population[index] = population[len(population)-1]
		population = population[:len(population)-1]
	}
	return sample
}

func Math_Random_GenerateBits(n int) (*big.Int, error) {
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive")
	}
	b := make([]byte, (n+7)/8)
	_, x := rand.Read(b)
	if x != nil {
		return nil, x
	}
	if n%8 != 0 {
		b[0] &= byte(1<<uint(n%8)) - 1
	}
	return new(big.Int).SetBytes(b), nil
}
