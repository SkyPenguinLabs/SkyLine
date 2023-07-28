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
// This file defines all of the functions for number theory based functions such as prime's and factorial based functions
//
package SkyLine_Standard_Library_Math

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	"math/big"
)

func Math_NumberTheory_IsPrime(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Theory_isprime", 1, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		n := args[0].(*SkyEnv.SL_Integer).Value
		return &SkyEnv.SL_Boolean{Value: big.NewInt(int64(n)).ProbablyPrime(0)}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_NumberTheory_Prime(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Theory_prime", 1, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		n := args[0].(*SkyEnv.SL_Integer).Value
		if n <= 0 {
			return nil
		}

		primes := make([]*big.Int, n)
		primes[0] = big.NewInt(2)

		for idx := 1; idx < n; idx++ {
			for j := primes[idx-1].Int64() + 1; ; j++ {
				if big.NewInt(j).ProbablyPrime(0) {
					primes[idx] = big.NewInt(j)
					break
				}
			}
		}
		return &SkyEnv.SL_Integer{Value: int(primes[n-1].Int64())}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_NumberTheory_GCD(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Theory_gcd", 2, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		X := args[0].(*SkyEnv.SL_Integer).Value
		Y := args[1].(*SkyEnv.SL_Integer).Value
		for Y != 0 {
			X, Y = Y, X%Y
		}
		return &SkyEnv.SL_Integer{Value: X}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_NumberTheory_Factor(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Theory_factor", 1, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		n := args[0].(*SkyEnv.SL_Integer).Value
		factors := []int{}
		for idx := 2; idx <= n/idx; idx++ {
			for n%idx == 0 {
				factors = append(factors, idx)
				n /= idx
			}
		}
		if n > 1 {
			factors = append(factors, n)
		}
		arr := make([]SkyEnv.SL_Object, 0)
		for i := 0; i < len(factors); i++ {
			arr = append(arr, &SkyEnv.SL_Integer{Value: factors[i]})
		}
		return &SkyEnv.SL_Array{Elements: arr}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_NumberTheory_Abs(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Theory_abs", 1, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		n := args[0].(*SkyEnv.SL_Integer).Value
		if n < 0 {
			n = n * -1
		}
		return &SkyEnv.SL_Integer{Value: n}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}
