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
// This file contains function definitions dedicated to modular aithmetic operations
//
package SkyLine_Standard_Library_Math

import (
	"math/big"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func Math_Modular_Arith_Mod(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Modular_mod", 3, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		x, y, m := int64(args[0].(*SkyEnv.SL_Integer).Value), int64(args[1].(*SkyEnv.SL_Integer).Value), int64(args[2].(*SkyEnv.SL_Integer).Value)
		base := big.NewInt(x)
		exponent := big.NewInt(y)
		modulus := big.NewInt(m)
		result := new(big.Int)
		result.Exp(base, exponent, modulus)
		return &SkyEnv.SL_Integer{Value: int(result.Int64())}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}

func Math_Modular_Arith_Mod_Exp(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := AutomateError_ExactArgCheckAndType("math.Modular_modexp", 3, SkyEnv.SKYLINE_DATATYPE_INTEGER_OBJECT, args...); x == nil {
		base := big.NewInt(0).SetInt64(int64(args[0].(*SkyEnv.SL_Integer).Value))
		exp := big.NewInt(0).SetInt64(int64(args[1].(*SkyEnv.SL_Integer).Value))
		mod := big.NewInt(0).SetInt64(int64(args[2].(*SkyEnv.SL_Integer).Value))
		result := big.NewInt(1)
		base = base.Mod(base, mod)
		for exp.BitLen() > 0 {
			if exp.Bit(0) == 1 {
				result = result.Mul(result, base).Mod(result, mod)
			}
			exp = exp.Rsh(exp, 1)
			base = base.Mul(base, base).Mod(base, mod)
		}
		return &SkyEnv.SL_Integer{Value: int(result.Int64())}
	} else {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
}
