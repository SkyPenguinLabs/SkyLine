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
// Filename      |  SkyLine_Standard_Library_Core.go
// Project       |  SkyLine programming language
// Line Count    |  200+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all core library functions, kind of everything mashed into one giant file. The reason this file exists right now is because as of
//                 0.0.5 of the language, SkyLine does not have a proper modular structure to the language and is still in testing. When further development comes about
//                 and other core features are used or implemented, we will then require that the library be accessed externally to the `Modules` directory. This will make
//                 it much easier to access, read, and create a full fledged standard library.
//
//
// ::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file is split up into two major parts with sub units under them. The two sections come to the standard library itself and all the functions the standard library uses
//
// then the init or calls that can `register` these functions into the environment at runtime. The sub sections will be split into their own library functions suchas `math` or
//
// `forensics` or `IoT` and even `io`. Each sub unit will have a block of note's that will describe the topic and why it exists in this placement. It is important to note that
//
// all functions that are other routine based functions and do not return an Object under `SL_Object` will be put at the very bottom of this file.
//
//::END_UNIT
package SkyLine_Backend

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	SkyLine_External_Forensics "main/Modules/StandardLibraryExternal/Forensics"
	"net/http"
	"os/signal"
	"plugin"
	"sort"
	"syscall"
	"unicode/utf8"

	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____             _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_ _ ___ ___| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | |   |  _|  _| | . |   |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |___|_|_|___|_| |_|___|_|_|___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|
//////////////////////////////////////////////////////////////////////////
//
// DEF  -> This section defines the library functions for the mathematics library. It may be important to note that the functions that are registered into the registry function
//
// or in other words are called within `func RegisterBuiltin(name string, fun BuiltinFunction)` with (*FUNCTION)->CALL are going to all be at the top of each sub section. The sub
//
// functions or sub routines these functions depend on will all be thrown at the bottom of the main file. We do this because it helps with organization and location of specific functions
//
// and can be better to work around. These sub routines SHOULD ALL START WITH DEPENDANT_ because other functions under this unit all depend on them or use them in some shape or form

func MathAbs(args ...SLC_Object) SLC_Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &Integer{Value: v}
	case *Float:
		v := arg.Value
		if v < 0 {
			v = v * -1
		}
		return &Float{Value: v}
	default:
		return NewError("argument to `math.abs` not supported, got=%s",
			args[0].SL_RetrieveDataType())
	}
}

func MathCos(args ...SLC_Object) SLC_Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments. got=% but need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Cos(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Cos(arg.Value)}
	default:
		return NewError("Argument to `math.cos` must be integer or float")
	}
}

func MathSin(args ...SLC_Object) SLC_Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments. got=%d but need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Sin(float64(arg.Value)))}
	case *Float:
		return &Float{Value: float64(math.Sin(arg.Value))}
	default:
		return NewError("Argument in call to `math.cos` must be integer or float")
	}
}

func MathTan(args ...SLC_Object) SLC_Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments, got=%d bnut need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Tan(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Tan(arg.Value)}
	default:
		return NewError("Argument in call to `math.tan` must be integer or float")
	}
}

func MathSqrt(args ...SLC_Object) SLC_Object {
	if len(args) != 1 {
		return NewError("Wrong number of arguments, got=%d but need=1", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		return &Integer{Value: int64(math.Sqrt(float64(arg.Value)))}
	case *Float:
		return &Float{Value: math.Sqrt(arg.Value)}
	default:
		return NewError("Argument in call to `math.sqrt` must be integer or float")
	}
}

func MathCbrt(args ...SLC_Object) SLC_Object {
	switch arg := args[0].(type) {
	case *Integer:
		f, _ := strconv.Atoi(arg.SL_InspectObject())
		z := f / 3.0
		for i := 0; i < 10; i++ {
			z = z - ((z*z*z - f) / (3 * z * z))
		}
		return &Integer{Value: int64(z)}
	case *Float:
		f, _ := strconv.Atoi(arg.SL_InspectObject())
		z := f / 3.0
		for i := 0; i < 10; i++ {
			z = z - ((z*z*z - f) / (3 * z * z))
		}
		return &Float{Value: float64(z)}
	default:
		return NewError("Argument in call to `math.cbrt` must be  integer or float")
	}
}

func MathPow(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("math.pow", args, ExactArguments(2), ArgumentHasTypes(IntegerType, IntegerType)); x != nil {
		return NewError(x.Error())
	}
	x := args[0].(*Integer)
	y := args[1].(*Integer)
	value := Dependant_PowerOf(x.Value, y.Value)
	return &Integer{Value: value}
}

func MathRand(args ...SLC_Object) SLC_Object {
	return &Float{Dependant_Random()}
}

func Dependant_PowerOf(x, y int64) int64 {
	p := int64(1)
	for y > 0 {
		if y&1 != 0 {
			p *= x
		}
		y >>= 1
		x *= x
	}
	return p
}

func Dependant_Random() float64 {
	return rand.Float64()
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____             _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_ _ ___ ___| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | |   |  _|  _| | . |   |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |___|_|_|___|_| |_|___|_|_|___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This sub unit defines all of the environment related functions that will tell or inform  the user about the environment. This typically is its own package and depends
//
// on whether or not the user used `register("env")`. This sub unit follows the same ideas as the sub unit above
//
//

func Environment_GetEnvironment(arguments ...SLC_Object) SLC_Object {
	environmemt := os.Environ()
	Hashp := make(map[HashKey]HashPair)
	for idx := 1; idx < len(environmemt); idx++ {
		key := &String{Value: environmemt[idx]}
		val := &String{Value: os.Getenv(environmemt[idx])}
		NewHashCreate := HashPair{Key: key, Value: val}
		Hashp[key.HashKey()] = NewHashCreate
	}
	return &Hash{Pairs: Hashp}
}

func Environment_GetEnvironmentPath(arguments ...SLC_Object) SLC_Object {
	if len(arguments) != 1 {
		return NewError("Argument error (Standard Library): wwrong number of arguments! this function requires %d and you gave `%d` argument(s)", 1, len(arguments))
	}
	if arguments[0].SL_RetrieveDataType() != StringType {
		return NewError("Argument error (Standard Library): Data type error! This function requires a data type of type STRING but you gave %s", arguments[0].SL_RetrieveDataType())
	}
	input := arguments[0].(*String).Value
	return &String{Value: os.Getenv(input)}
}

func Environment_SetEnvironment(arguments ...SLC_Object) SLC_Object {
	if len(arguments) != 2 {
		return NewError("Argument error (Standard Library): Wrong number of arguments! This function requires %d and you gave `%d` argument(s)", 2, len(arguments))
	}
	if arguments[0].SL_RetrieveDataType() != StringType {
		return NewError("argument must be a string, got=%s",
			arguments[0].SL_RetrieveDataType())
	}
	if arguments[1].SL_RetrieveDataType() != StringType {
		return NewError("argument must be a string, got=%s",
			arguments[1].SL_RetrieveDataType())
	}
	name := arguments[0].(*String).Value
	value := arguments[1].(*String).Value
	os.Setenv(name, value)
	return NilValue
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____             _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     |___ _ _ ___| |_ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|   --|  _| | | . |  _| . |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_| |_  |  _|_| |___|
//	primary units that define the functions to register the sub func's  //            |___|                             |___|_|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This sub unit defines all of the functions for the cryptography library which can include hashing algorithm implementations, crypto frameworks and more which may aid in
//
// protecting or locking files. This sub unit follows the same ideas as the sub unit above
//
//
// This unit and crypto package exists not just for standard hashing and encryption algorithms but also to provide an easy way to encrypt files, decrypt files with keys, sign keys
//
// create certs, create pens and a bunch of other functions including the ability to easily create a hybrid encryption set. Right now, the crypto library does not contain much but
//
// the idea is being worked on

var HashMapper = map[string]func(string) string{
	"MD5": func(dt string) string {
		return fmt.Sprintf("%x", md5.Sum([]byte(dt)))
	},
	"SHA1": func(dt string) string {
		return fmt.Sprintf("%x", sha1.Sum([]byte(dt)))
	},
	"SHA256": func(dt string) string {
		return fmt.Sprintf("%x", sha256.Sum256([]byte(dt)))
	},
	"SHA512": func(dt string) string {
		return fmt.Sprintf("%x", sha512.Sum512([]byte(dt)))
	},
}

func Crypto_Hasher(args ...SLC_Object) SLC_Object {
	if len(args) != 2 {
		return NewError("Sorry but crypt.hash requires 2 positional based arguments, you gave %d but the function wants 2 ", len(args))
	}
	var hashtype, data string
	hashtype = strings.ToUpper(args[0].SL_InspectObject())
	data = args[1].SL_InspectObject()
	if HashMapper[hashtype] != nil {
		return &String{Value: HashMapper[hashtype](data)}
	} else {
		return NewError("crypt.hash requires a supported hash type, below are the supported hash types \n | -> MD5 \n | -> SHA1 \n | -> SHA256 \n | -> SHA512 ")
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____       _ _ ___ _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     |___ _| |_|  _|_|___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___| | | | . | . | |  _| | -_|  _|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_|_|_|___|___|_|_| |_|___|_|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This code brick and section defines everything for the modification of the skyline environment. This function might be deprecated soon considering SLC exists and it is
//
// currently a bit overly bloated and long. This sub unit follows the same ideas as the sub unit above
//
//

func ModifyEnviornment(args ...SLC_Object) SLC_Object {
	var mod, setting string
	var roottree *TreeNode
	var fixstatement, callto, errormsg string
	callto = "modify()"
	if len(args) != 1 {
		// 1 too many arguments or 1 too little arguments
		// This will be a tree by defualt
		roottree = &TreeNode{
			Type: SKYLINE_HIGH_DEFRED + " E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
			Children: []*TreeNode{
				{
					Type: SKYLINE_HIGH_DEFRED + "Error Information Tree" + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_STANDARD_FUNCTION_GOT__ARGUMENTS) + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + "Standard Library ( Modify Environment ) " + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFAQUA + " When calling 'modify()' the arguments given were NEGATIVE",
									Children: []*TreeNode{
										{
											Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[I] Sub Branch (info) " + SKYLINE_RESTORE,
											Children: []*TreeNode{
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[I] Arguments too long     ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint((len(args) > 1)) + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[I] Arguments too short    ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint((len(args) < 1)) + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[I] Required argument num  ? " + SUNRISE_HIGH_DEFINITION + fmt.Sprint(1) + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[I] In Call to function    ? " + SUNRISE_HIGH_DEFINITION + `modify()` + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[I] Syntax of function     ? " + SUNRISE_HIGH_DEFINITION + `modify("system:setting")` + SKYLINE_RESTORE,
												},
												{
													Type: SUNRISE_LIGHT_DEFINITION + "[I] Function requires type ? " + SUNRISE_HIGH_DEFINITION + `String` + SKYLINE_RESTORE,
												},
											},
										},
									},
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file name  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_SICK_BLUE + "[S] Ensure the amount of arguments is correct " + SKYLINE_RESTORE,
								},
							},
						},
						{
							Type: SUNRISE_LIGHT_DEFINITION + "Auto Detection" + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_FIXBLUE + "[F] " + DEPENDANT_DefineStatementBetweenArgs(1, len(args)) + SKYLINE_RESTORE,
								},
							},
						},
					},
				},
			},
		}
		errormsg = "Invalid amount of arguments in call to standard function"
		fixstatement = DEPENDANT_DefineStatementBetweenArgs(1, len(args))
	} else {
		switch argt := args[0].(type) {
		case *String:
			if strings.Contains(argt.Value, ":") {
				spliter := strings.Split(argt.Value, ":")
				if spliter[0] != "" && spliter[1] != "" {
					mod = spliter[0]
					setting = spliter[1]
				} else {
					// Nil error
					roottree = &TreeNode{
						Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "Error Information Tree" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
										Children: []*TreeNode{
											{
												Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_STANDARD_FUNCTION_COULD_NOT_SET_ARGS_VAL_AFTER_SEP_NIL) + SKYLINE_RESTORE,
											},
										},
									},
									{
										Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
										Children: []*TreeNode{
											{
												Type: SKYLINE_HIGH_DEFAQUA + " Standard Function ( Modify Environment ) " + SKYLINE_RESTORE,
											},
										},
									},
									{
										Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
										Children: []*TreeNode{
											{
												Type: SKYLINE_HIGH_DEFAQUA + "Argument after seperation was empty SEPTOKEN(':').Peek()",
												Children: []*TreeNode{
													{
														Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[I] Sub Branch (info) " + SKYLINE_RESTORE,
														Children: []*TreeNode{
															{
																Type: SUNRISE_LIGHT_DEFINITION + "[I] When calling the function, the function will seperate the module " + SKYLINE_RESTORE,
															},
															{
																Type: SUNRISE_LIGHT_DEFINITION + "[I] When the function tried seperating at :, it found that the argument" + SKYLINE_RESTORE,
															},
															{
																Type: SUNRISE_LIGHT_DEFINITION + "[I] Was found to be empty, this argument defines the setting for the module" + SKYLINE_RESTORE,
															},
														},
													},
												},
											},
										},
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
										Children: []*TreeNode{
											{
												Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
											},
											{
												Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file name  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
											},
											{
												Type: SUNRISE_LIGHT_DEFINITION + "[UW] Function Called      ?  " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "modify()",
											},
										},
									},
									{
										Type: SUNRISE_HIGH_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
										Children: []*TreeNode{
											{
												Type: SKYLINE_SICK_BLUE + "[S] Consider adding a value to this function after the `:` token in the argument list" + SKYLINE_RESTORE,
											},
										},
									},
								},
							},
						},
					}
				}
				errormsg = "Argument after symbol `:` in call to modify() was empty, expected String value"
				fixstatement = "Consider adding a value of type String after the `:` to define the modules setting"
			} else {
				// Strings did not contain expected seperator
				roottree = &TreeNode{
					Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
					Children: []*TreeNode{
						{
							Type: SKYLINE_HIGH_DEFRED + "Error Information Tree " + SKYLINE_RESTORE,
							Children: []*TreeNode{
								{
									Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_STANDARD_FUNCTION_COULD_NOT_LOCATE_ARGUMENT_SEPERATOR) + SKYLINE_RESTORE,
										},
									},
								},
								{
									Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_HIGH_DEFAQUA + "Standard Library ( Modify Environment ) " + SKYLINE_RESTORE,
										},
									},
								},
								{
									Type: SKYLINE_HIGH_DEFAQUA + "[E] Message " + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_HIGH_DEFRED + "The function argument sent over, did not contain a seperator (`:`)" + SKYLINE_RESTORE,
											Children: []*TreeNode{
												{
													Type: SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "[I] Sub Branch (info) " + SKYLINE_RESTORE,
													Children: []*TreeNode{
														{
															Type: SUNRISE_LIGHT_DEFINITION + "[I] This function requires 1 positional argument which holds two values" + SKYLINE_RESTORE,
														},
														{
															Type: SUNRISE_LIGHT_DEFINITION + "[I] These types are known as the module of the system you want to modify and the setting" + SKYLINE_RESTORE,
														},
														{
															Type: SUNRISE_LIGHT_DEFINITION + "[I] These are defined with (`:`) so, when you call modify() make sure it works like this" + SKYLINE_RESTORE,
														},
														{
															Type: SUNRISE_LIGHT_DEFINITION + `[I] Example -> modify("errors:basic")`,
														},
													},
												},
											},
										},
									},
								},
								{
									Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
										},
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file name  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
										},
										{
											Type: SUNRISE_LIGHT_DEFINITION + "[UW] Function Called      ?  " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "modify()",
										},
									},
								},
								{
									Type: SUNRISE_HIGH_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
									Children: []*TreeNode{
										{
											Type: SKYLINE_SICK_BLUE + "[S] Consider adding a the : to seperate the module and the setting" + SKYLINE_RESTORE,
										},
									},
								},
							},
						},
					},
				}
				errormsg = "SkyLine could not locate the seperator needed for the function"
				fixstatement = "Try adding : in call to the function like so " + `modify("system:setting")`
			}
		default:
			roottree = &TreeNode{
				Type: SKYLINE_HIGH_DEFRED + "E | " + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
				Children: []*TreeNode{
					{
						Type: SUNRISE_HIGH_DEFINITION + "Error Information Tree " + SKYLINE_RESTORE,
						Children: []*TreeNode{
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Code " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + fmt.Sprint(ERROR_STANDARD_FUNCITON_REQUIRES_SO_AND_SO_TYPE_NOT_PROVIDED) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Type " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Standard Library ( Modify Environment ) " + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SKYLINE_HIGH_DEFRED + "[E] Message " + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_HIGH_DEFAQUA + "Wrong data type, function requires String not " + fmt.Sprint(argt.SL_RetrieveDataType()) + SKYLINE_RESTORE,
									},
								},
							},
							{
								Type: SUNRISE_LIGHT_DEFINITION + "Information Branch" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file path  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + FileCurrent.GetAbsolute() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Estimated file name  ? " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + ParserErrorSystem_GetFileName() + SKYLINE_RESTORE,
									},
									{
										Type: SUNRISE_LIGHT_DEFINITION + "[UW] Function Called      ?  " + SKYLINE_SUNRISE_HIGH_DEF_YELLOW + "modify()",
									},
								},
							},
							{
								Type: SUNRISE_HIGH_DEFINITION + "Suggestion" + SKYLINE_RESTORE,
								Children: []*TreeNode{
									{
										Type: SKYLINE_SICK_BLUE + "[S] Ensure that the data type of the argument is String not " + fmt.Sprint(argt.SL_RetrieveDataType()) + SKYLINE_RESTORE,
									},
								},
							},
						},
					},
				},
			}
			errormsg = "In call to modify(), positional argument (1) was found to be the wrong data type"
			fixstatement = "Make sure this argument is of type String not " + fmt.Sprint(argt.SL_RetrieveDataType())
			// argument needs to be string
		}
	}
	if roottree == nil {
		VerifyModification(mod, setting)
	} else {
		if ErrorSys.Tree {
			RetTreeSys(roottree, "", true)
		} else {
			DEPENDANT_Construct_ErrorSystemSimple(callto, errormsg, true, fixstatement)
		}
	}
	return &Nil{}
} // see? told you, overly bloated

func DEPENDANT_DefineStatementBetweenArgs(lenrequired, lengiven int) string {
	if lenrequired < lengiven {
		return "Too many arguments, try taking some away"
	} else if lenrequired > lengiven {
		return "Too little arguments, function at least takes " + fmt.Sprint(lenrequired) + " argument(s)"
	} else {
		return "unsupported ( contact developers for this issue)"
	}
}

func DEPENDANT_Construct_ErrorSystemSimple(callto, errormsg string, fix bool, fixstatement string) string {
	var Out string
	Out += SKYLINE_HIGH_RES_VIS_RED + "E | " + SKYLINE_HIGH_RES_VIS_BLUE + ParserErrorSystem_GetFileName() + "\n"
	Out += "\n"
	Out += SKYLINE_HIGH_RES_VIS_SUNSET + "[N] | " + SKYLINE_RESTORE + callto + SKYLINE_RESTORE + "\n"
	if fix && fixstatement != "" {
		Out += SKYLINE_HIGH_RES_VIS_SUNSET + "[F] | " + SKYLINE_RESTORE + fixstatement + "\n"
	}
	Out += "\n"
	Out += SKYLINE_HIGH_RES_VIS_RED + "Error: " + SKYLINE_RESTORE + SKYLINE_HIGH_RES_VIS_BLUE + errormsg + "\n"
	return Out
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|___ ___ ___ ___ ___|_|___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|   __| . |  _| -_|   |_ -| |  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|  |___|_| |___|_|_|___|_|___|___|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This code section defines all of the standard functions for the forensics library. The forensics library needs ALOT of work and needs to be improved a lot more than what
//
// it is right now, however it is a start. These functions can get the number of chunks, get the data from the image, dissect the entire image, dump offsets, check for zip encodings
//
// and even check for specific signatures. There may be a debate here, with apple in the IoT section, there will be the ability to dissect and parse BPLIST and PLIST files. However
//
// the issue with this is that dissecting a BPLIST file is a forensics kind of topic or analysis topic so we might move apple related forensics into forensics but just in a specific
//
// file path like forensics/Apple/iOS or forensics/Apple/OSX or forensics/Apple/WatchoS(TVOS)
//

var metadata SkyLine_External_Forensics.PNG_Meta

func Run_Forensics_Metadata_PNG_GetNumChunks(args ...SLC_Object) SLC_Object {
	fmt.Println("Running function to get num chunks")
	dat, err := os.Open(args[0].SL_InspectObject())
	if err != nil {
		log.Fatal(err)
	}
	defer dat.Close()
	bReader, err := SkyLine_External_Forensics.Process_Given_Image(dat)
	if err != nil {
		log.Fatal(err)
	}
	return &String{
		Value: metadata.PNG_Meta_GetChunkCount(bReader),
	}
}

// Start settings for injection
var SessionSets SkyLine_External_Forensics.Settings

func RunSettings(args ...SLC_Object) SLC_Object {
	key := args[0].SL_InspectObject()
	out := args[1].SL_InspectObject()
	in := args[2].SL_InspectObject()
	filemode := 0
	offset := args[3].SL_InspectObject()
	payload := args[4].SL_InspectObject()
	chunktoinject := args[5].SL_InspectObject()
	SessionSets.Settings_Inject_New(
		key,
		out,
		in,
		fmt.Sprint(filemode),
		"false",
		"false",
		offset,
		payload,
		chunktoinject,
	)
	return &Nil{}
}

// Run injection
func RunInject(args ...SLC_Object) SLC_Object {
	dat, err := os.Open(args[0].SL_InspectObject())
	if err != nil {
		log.Fatal(err)
	}
	defer dat.Close()
	bReader, err := SkyLine_External_Forensics.Process_Given_Image(dat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(SessionSets.ImageOffset)
	SessionSets.Injection_Standard_Payload(bReader)
	return &Boolean_Object{Value: true}
}

// Run ZIP checking
func RunZipCheck(args ...SLC_Object) SLC_Object {
	return &Boolean_Object{Value: SkyLine_External_Forensics.Image_Controller_Function_External_1_Verify_Archive(args[0].SL_InspectObject())}
}

// Run file INJ
func RunFileInjFromFileToFile(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("fpng.InjectImage", args, ExactArguments(3), ArgumentHasTypes(StringType)); x == nil {
		return &Boolean_Object{
			Value: SkyLine_External_Forensics.Image_Controller_Function_External_2_Inject_Into_File(
				args[0].SL_InspectObject(), // input file for the source image.
				args[1].SL_InspectObject(), // output file to inject the data in the input file
				args[2].SL_InspectObject(), // // input file to copy data from
			),
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

// Run File verification
func RunFileVerification(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("fpng.FindSigUnknownFile", args, ExactArguments(1), ArgumentHasTypes(StringType)); x == nil {
		return &String{
			Value: SkyLine_External_Forensics.Image_Controller_Function_External_3_IdentifyUnknownFile(args[0].SL_InspectObject()),
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

// Run File creation
func RunFileCreationFromFile(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("ImageUtils.CreateImage", args, ExactArguments(1), ArgumentHasTypes(StringType)); x == nil {
		return &Boolean_Object{
			Value: SkyLine_External_Forensics.SkyLine_Forensics_Image_Creation_Utility(args[0].SL_InspectObject()),
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

// Run new file creation
func RunFileNewSettings(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("ImageUtils.CreationNew", args, RangeArguments(0, 3), ArgumentHasTypes(StringType)); x == nil {
		out := args[0].SL_InspectObject()
		pw := args[1].SL_InspectObject()
		ph := args[2].SL_InspectObject()
		if out == "" {
			out = "SkyLine_Created_File_Output_Image"
		}
		if pw == "" {
			pw = "600"
		}
		if ph == "" {
			ph = "1000"
		}
		SkyLine_External_Forensics.CreationSettings.Output = args[0].SL_InspectObject()
		SkyLine_External_Forensics.CreationSettings.PixelWidth = args[1].SL_InspectObject()
		SkyLine_External_Forensics.CreationSettings.PixelHeight = args[2].SL_InspectObject()
		return &Boolean_Object{Value: true}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

// Run payload injection
func RunPayloadInjectionRegular(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("fpng.InjectRegular", args, ExactArguments(2), ArgumentHasTypes(StringType)); x == nil {
		if args[1].SL_InspectObject() != "" && args[2].SL_InspectObject() != "" {
			return &Boolean_Object{
				Value: SkyLine_External_Forensics.InjectImage(args[1].SL_InspectObject(), args[0].SL_InspectObject()),
			}
		} else {
			var debug string
			debug = "POS #1 ??? [STRING] NIL???? -> " + fmt.Sprint(args[0].SL_InspectObject() == "")
			debug += "\n POS #2 ??? [STRING] NIL???? -> " + fmt.Sprint(args[1].SL_InspectObject() == "")
			return &Error{Message: "SkyLine PNG Forensics Library: Failed to inject regular image, one positional argument was empty\n" + debug}
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____ _ _     _____         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_| |___|   __|_ _ ___| |_ ___ ______
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | | -_|__   | | |_ -|  _| -_|     |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |_|_|___|_____|_  |___|_| |___|_|_|_|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|                   |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This defines all of the functions for the standard library of `File`. This allows you to open, hex dump, read and write data to files. This kind of manipulation is extremely
//
// important. The SkyLine file library eventually will be worked on and held as a primary source.
//

type FileInit struct {
	Filename string
	Mode     int
}

var File FileInit

func FileLib_WriteFile(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("file.Write", args, ExactArguments(1), ArgumentHasTypes(StringType)); x == nil {
		payload := args[0].SL_InspectObject()
		if File.Filename == "" {
			return &Error{
				Message: "File sys error: Sorry, you provided 1 argument in call to File.Write which means SkyLine assumed you ran File.New() but it seems as if you did not. Please place the following instruction above this call -> INSTRUCT (`%s`)" + `File.New("NAME_OF_FILE")`,
			}
		} else {
			f, x := os.OpenFile(File.Filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if x != nil {
				return &Error{Message: "File Sys Error: Failed to open file due to -> " + fmt.Sprint(x)}
			} else {
				defer f.Close()
				_, x = f.WriteString(payload)
				if x != nil {
					return &Error{Message: fmt.Sprint(x)}
				} else {
					return &Boolean_Object{Value: true}
				}
			}
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_OverWrite_WriteFile(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("file.Overwrite", args, ExactArguments(1), ArgumentHasTypes(StringType)); x == nil {
		payload := []byte(args[0].SL_InspectObject())
		x = ioutil.WriteFile(File.Filename, payload, fs.FileMode(File.Mode))
		if x != nil {
			return &Error{Message: "File Sys error: Could not write to file -> " + fmt.Sprint(x)}
		} else {
			return &Boolean_Object{Value: true}
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_IniateNewFunction(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("file.New", args, RangeArguments(1, 2), ArgumentHasTypes(StringType)); x == nil {
		if len(args) == 2 {
			var x error
			File.Mode, x = strconv.Atoi(args[1].SL_InspectObject())
			if x != nil {
				return &Error{Message: "Conversion Error: Could not conver `" + fmt.Sprint(args[0].SL_InspectObject()) + "` to integer -> " + fmt.Sprint(x)}
			}
		}
		if args[0].SL_InspectObject() != "" {
			_, x := os.Stat(args[0].SL_InspectObject())
			if x != nil {
				return &Error{Message: "File Sys error: Could not stat file -> " + args[0].SL_InspectObject() + " : because ( " + fmt.Sprint(x) + ")"}
			}
			File.Filename = args[0].SL_InspectObject()
			return &Boolean_Object{Value: true}
		} else {
			return &Error{Message: "File sys error: Refused to try opening file, data was NULL"}
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

func FileLib_OpenAndOutFile(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("file.Open", args, RangeArguments(0, 1), ArgumentHasTypes(StringType)); x == nil {
		var file string
		if len(args) == 1 {
			file = args[0].SL_InspectObject()
			File.Filename = file
		} else {
			if File.Filename != "" {
				file = File.Filename
			} else {
				return &Error{
					Message: "File sys error: Sorry, you provided 0 arguments which means SkyLine assumed you ran File.New() but it seems as if you did not. Please place the following instruction above this call -> INSTRUCT (`%s`)" + `File.New("NAME_OF_FILE")`}
			}
		}
		dt, x := ioutil.ReadFile(file)
		if x != nil {
			return &Error{Message: "File sys error: Sorry, could not read file due to -> " + fmt.Sprint(x)}
		} else {
			return &String{Value: string(dt)}
		}
	} else {
		return &Error{Message: fmt.Sprint(x)}
	}
}

func RegisterFile() {
	RegisterBuiltin("file.New", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (FileLib_IniateNewFunction(args...))
	})
	RegisterBuiltin("file.Open", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (FileLib_OpenAndOutFile(args...))
	})
	RegisterBuiltin("file.Overwrite", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (FileLib_OverWrite_WriteFile(args...))
	})
	RegisterBuiltin("file.Write", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (FileLib_WriteFile(args...))
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____             _   _____             _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |     |___ ___ _ _| |_|     |_ _ ___ _ _| |_
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |-   -|   | . | | |  _|  |  | | | . | | |  _|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|_____|_|_|  _|___|_| |_____|___|  _|___|_|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|         |_|                   |_|
//////////////////////////////////////////////////////////////////////////
//
// DEF -> This section defines all of the IO library functions which are quite light and small at the moment but work perfectly as they need to right now. You may be asking if you
//
// are even reading these, why do we store so much functions and require to register them? This is because of bloat, who wants 900+ useless functions just registered into their
//
// environment before developing or even writing a simple hello world program? Last time I checked (Toally_Not_A_Haxxer) noone does. Standard keywords should be the only thing
//
// registered into the environment other than a few invokes and base output statements.
//
//

func User_Input(args ...SLC_Object) SLC_Object {
	// this function takes a few positional arguments
	// POS-1 -> The prompt for the user to enter data
	// POS-2 -> The expected return type float, string, int
	// POS-3 -> When to send the input or to cancle the input tag, this can be n for new line or t for tab etc
	if len(args) != 3 {
		return NewError("Function `input(...)` requires 3 positional arguments of type string. \n Arg1=prompt \n Arg2=Expected return type\nArg3=Cancle when... ")
	}
	var prompt, expected_type, drop_when SLC_Object
	prompt = args[0].(*String)
	expected_type = args[1].(*String)
	drop_when = args[2].(*String)
	retret := bufio.NewReader(os.Stdin)
	var out string
	fmt.Print(prompt.SL_InspectObject())
	et := expected_type.SL_InspectObject()
	for {
		switch drop_when.SL_InspectObject() {
		case "n":
			out, _ = retret.ReadString('\n')
			out = strings.Replace(out, "\n", "", -1)
		default:
			return NewError("Unsupported argument in placement `3` final argument in call to input(...) -> supported=(n)")
		}
		if out != "" {
			switch strings.ToLower(et) {
			case "integer":
				c, x := strconv.ParseInt(out, 0, 64)
				if x != nil {
					return &Error{Message: "Could not return this value, it was not able to be parsed as a integer which means it was either a float or character but this function does not support that as input"}
				}
				return &Integer{Value: c}
			case "float":
				c, x := strconv.ParseFloat(out, 64)
				if x != nil {
					return &Error{Message: "Could not return this value, it was not able to parse as a float value which means it was either a character or an integer, this function input does not accept anything but a float value"}
				}
				return &Float{Value: c}
			case "string":
				return &String{Value: out}
			}
		}
	}
}

func IO_Clear() SLC_Object {
	// takes no positional arguments
	WIN := "\x1b[2J"
	LIN := "\x1b[H\x1b[2J\x1b[3J"
	if U.OperatingSystem == "windows" {
		fmt.Println(WIN)
	} else {
		fmt.Println(LIN)
	}
	return &String{Value: ""}
}

type Box struct {
	TL string
	TR string
	BL string
	BR string
	HZ string
	VT string
}

func IO_Box(args ...SLC_Object) SLC_Object {
	var BL Box
	var text string
	// Optionally takes 7 positonal arguments
	if len(args) == 7 {
		text = args[0].SL_InspectObject()  // Text for the box
		BL.TL = args[1].SL_InspectObject() // Top left
		BL.TR = args[2].SL_InspectObject() // Top right
		BL.BL = args[3].SL_InspectObject() // Bottom left
		BL.BR = args[4].SL_InspectObject() // Bottom right
		BL.HZ = args[5].SL_InspectObject() // Horizontal
		BL.VT = args[6].SL_InspectObject() // Verticle
	} else {
		if len(args) >= 1 {
			text = args[0].SL_InspectObject()
		} else {
			return NewError("Sorry this function takes 1 required positional argument and 6 other optional arguments")
		}
		BL = Box{
			TL: "┏",
			TR: "┓",
			BL: "┗",
			BR: "┛",
			HZ: "━",
			VT: "┃",
		}
	}
	l := strings.Split(text, "\n")
	var mlen int
	for _, lin := range l {
		if len(lin) > mlen {
			mlen = len(lin)
		}
	}
	var b string
	b += BL.TL + strings.Repeat(BL.HZ, mlen) + BL.TR + "\n"
	for _, line := range l {
		b += BL.VT + line + strings.Repeat(" ", mlen-len(line)) + BL.VT + "\n"
	}
	b += BL.BL + strings.Repeat(BL.HZ, mlen) + BL.BR + "\n"
	return &String{Value: b}
}

func IO_Listen(args ...SLC_Object) SLC_Object {
	// This function is a bit more complicated than the other IO functions
	// This will start a thread and simply return nothing but rather listen
	// for key based input such as CTRL+C
	if len(args) != 2 {
		return NewError("Sorry this function of call io does not support any other functions other than 2. ")
	}
	var sigtype string
	switch arg := args[0].(type) {
	case *String:
		sigtype = arg.SL_InspectObject()
	default:
		return NewError("Sorry first argument in call to io listen is not a string, this argument MUST be a string")
	}
	var sig os.Signal
	switch strings.ToLower(sigtype) {
	case "terminate":
		sig = syscall.SIGTERM
	case "kill":
		sig = syscall.SIGKILL
	case "hangup":
		sig = syscall.SIGHUP
	case "ctrl-c":
		sig = os.Interrupt
	case "user1":
		sig = syscall.SIGUSR1
	case "user2":
		sig = syscall.SIGUSR2
	default:
		return NewError("Sorry the first agrument in the list does not exist")
	}
	msg := args[1].SL_InspectObject()
	c := make(chan os.Signal)
	go func() {
		HandleListener(c, sig, msg, ExitGracefully)
	}()
	return &Nil{}
}

func ExitGracefully(msg string) {
	println(msg)
	os.Exit(0)
}

func HandleListener(c chan os.Signal, signalCHAN os.Signal, message string, run func(string)) {
	signal.Notify(c, signalCHAN)
	for s := <-c; ; s = <-c {
		switch {
		case signalCHAN == syscall.SIGUSR1 && s == syscall.SIGUSR1:
			run(message)
		case signalCHAN == os.Interrupt && s == os.Interrupt:
			run(message)
		case signalCHAN == syscall.SIGUSR2 && s == syscall.SIGUSR2:
			run(message)
		case signalCHAN == syscall.SIGHUP && s == syscall.SIGHUP:
			run(message)
		case signalCHAN == syscall.SIGTERM && s == syscall.SIGTERM:
			run(message)
		case signalCHAN == syscall.SIGKILL && s == syscall.SIGKILL:
			run(message)
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ _____ _____     _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |_   _|_   _|  _  |___| __  |___ ___ _ _ ___ ___| |_ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|     | | |   | | |   __|___|    -| -_| . | | | -_|_ -|  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__| |_|   |_| |__|      |__|__|___|_  |___|___|___|_| |___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|               |_|
//////////////////////////////////////////////////////////////////////////
//
// DEF -> This section defines all of the HTTP library functions and types. As of right now there is not much of a library for HTTP requests but there are some basic stats and
//
// functions. Mostly this is a section that is also the very first section to define built in variables and pre registered variables
//
//

func LoadPlugin(args ...SLC_Object) SLC_Object {
	if x := CheckArguments(
		"ffi", args,
		ExactArguments(2),
		ArgumentHasTypes(StringType, StringType),
	); x != nil {
		return NewError(x.Error())
	}

	name := args[0].(*String).Value
	symbol := args[1].(*String).Value

	p, err := plugin.Open(fmt.Sprintf("%s.so", name))
	if err != nil {
		return NewError("Got error when loading plugin -> %s", fmt.Sprint(err))
	}

	v, err := p.Lookup(symbol)
	fmt.Printf("%T", v)
	if err != nil {
		return NewError("error finding symbol: %s", err)
	}

	return &Builtin{Fn: v.(BuiltinFunction)}
}

// Methods | HTTP
const (
	HTTPMETHOD_GET     = http.MethodGet
	HTTPMETHOD_HEAD    = http.MethodHead
	HTTPMETHOD_POST    = http.MethodPost
	HTTPMETHOD_DELETE  = http.MethodDelete
	HTTPMETHOD_PUTS    = http.MethodPut
	HTTPMETHOD_PATCH   = http.MethodPatch
	HTTPMETHOD_CONNECT = http.MethodConnect
	HTTPMETHOD_TRACE   = http.MethodTrace
	HTTPMETHOD_OPTIONS = http.MethodOptions
)

func MakeGet(args ...SLC_Object) SLC_Object {
	if x := CheckArguments("http.Get", args, ExactArguments(1), ArgumentHasTypes(StringType)); x != nil {
		return NewError(x.Error())
	}
	bod, x := http.Get(args[0].SL_InspectObject())
	if x != nil {
		return &Error{Message: "HTTP request faile"}
	}
	defer bod.Body.Close()
	Mapper := make(map[string]interface{}, 0)
	RetHash := make(map[HashKey]HashPair)
	Mapper["Status"] = bod.Status
	Mapper["StatusCode"] = bod.StatusCode
	Mapper["Proto"] = bod.Proto
	Mapper["ProtoMajor"] = bod.ProtoMajor
	Mapper["ProtoMinor"] = bod.ProtoMinor
	var buffer bytes.Buffer
	_, x = io.Copy(&buffer, bod.Body)
	if x != nil {
		log.Fatal(x)
	}
	Mapper["ResponseBody"] = buffer.String()
	Mapper["TranserEncoding"] = bod.TransferEncoding
	Mapper["ContentLength"] = bod.ContentLength
	for k, v := range Mapper {
		key := &String{Value: k}
		val := &String{Value: fmt.Sprint(v)}
		NewHashCreate := HashPair{Key: key, Value: val}
		RetHash[key.HashKey()] = NewHashCreate
	}
	return &Hash{Pairs: RetHash}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____     _                   _         _____ ___       _____ _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     |___| |_ ___ ___ ___ ___| |_      |     |  _|     |_   _| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|-   -|   |  _| -_|  _|   | -_|  _|     |  |  |  _|       | | |   | |   | . |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_|_|_| |___|_| |_|_|___|_|  _____|_____|_|  _____  |_| |_|_|_|_|_|_  |___|
//	primary units that define the functions to register the sub func's  //            |___|                                                     |_____|         |_____|               |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Defines -> The IoT library within the SkyLine programming language is more or less defined for IoT devices and is an experimental library. This library contains informaiton
//
// such as endpoints, parsers, expected codes, headers and expected responses from servers and much more for IoT devices such as Samsung smart devices, google home's, FireSticks,
//
// home routers, AppleTV's, RokuTV's and other various IoT devices. This mainly works on the API's and focuses on information gathering from them or utilizing them in specific ways
//
// as these API's are either open or were reverse engineered. Make sure when using this library YOU KNOW WHAT YOU ARE DOING. SkyLine is still an experimental programming language
//
// and it has MANY bugs in it. Right now, this library may not be the best thing to use if you do not understand it.
//

// Apple API endpoint paths
const (
	APPLE_DAAP_PATH            = "http://%s:3689/"
	APPLE_DAAP_LOGIN           = "http://%s:3689/login?attempts=1"
	APPLE_DAAP_DATABASE        = "http://%s:3689/databases/1/containers/1/items"
	APPLE_AIRPLAY_INFO         = "http://%s:7000/info"
	APPLE_AIRPLAY_SCRUB        = "http://%s:7000/scrub"
	APPLE_AIRPLAY_SERVERINFO   = "http://%s:7000/server-info"
	APPLE_AIRPLAY_PLAYBACKINFO = "http://%s:7000/playback-info"
	APPLE_AIRPLAY_STREAMINFO   = "http://%s:7100/stream.xml"
)

//Roku API endpoint paths
const (
	ROKU_KEYPRESS_HOME      = "http://%s:8060/keypress/home"
	ROKU_KEYPRESS_PLAY      = "http://%s:8060/keypress/play"
	ROKU_KEYPRESS_DOWN      = "http://%s:8060/keypress/down"
	ROKU_KEYPRESS_UP        = "http://%s:8060/keypress/up"
	ROKU_KEYPRESS_LEFT      = "http://%s:8060/keypress/left"
	ROKU_KEYPRESS_RIGHT     = "http://%s:8060/keypress/right"
	ROKU_KEYPRESS_OK        = "http://%s:8060/keypress/select"
	ROKU_KEYPRESS_REWIND    = "http://%s:8060/keypress/rewind"
	ROKU_KEYPRESS_FFW       = "http://%s:8060/keypress/fastforward"
	ROKU_KEYPRESS_OPTIONS   = "http://%s:8060/keypress/options"
	ROKU_KEYPRESS_PAUSE     = "http://%s:8060/keypress/pause"
	ROKU_KEYPRESS_BACK      = "http://%s:8060/keypress/back"
	ROKU_KEYPRESS_POWEROFF  = "http://%s:8060/keypress/poweroff"
	ROKU_KEYPRESS_VUP       = "http://%s:8060/keypress/volumeup"
	ROKU_KEYPRESS_VDOWN     = "http://%s:8060/keypress/volumedown"
	ROKU_KEYPRESS_MUTE      = "http://%s:8060/keypress/volumemute"
	ROKU_KEYPRESS_DEVDOWN   = "http://%s:8060/keypress/powerOff"
	ROKU_KEYPRESS_DEVUP     = "http://%s:8060/keypress/powerOn"
	ROKU_DEVICE_LAUNCH      = "http://%s:8060/launch/%s"
	ROKU_DEVICE_INSTALL     = "http://%s:8060/install/%s?contentid=%s&MediaType=%s"
	ROKU_DEVICE_DISABLE_SGR = "http://%s:8060/query/sqrendezvous/untrack"
	ROKU_DEVICE_ENABLE_SGR  = "http://%s:8060/query/sqrendezvous/track"
	ROKU_DEVICE_TV          = "http://%s:8060/query/tv-channels"
	ROKU_DEVICE_SGNODE      = "http://%s:8060/query/sgnodes"
	ROKU_DEVICE_ACTIVETV    = "http://%s:8060/query/tv-active-channel"
	ROKU_DEVICE_DIAL        = "http://%s:8060/dial/dd.xml"
	ROKU_DEVICE_BROWSE      = "http://%s:8060/search/browse?keyword=%s"
	ROKU_DEVICE_INFO        = "http://%s:8060/query/device-info"
	ROKU_DEVICE_APPS        = "http://%s:8060/query/apps"
	ROKU_DEVICE_ACTIVE      = "http://%s:8060/query/active-app"
)

//Google Chrome Cast API endpoint paths
const (
	GOOGLE_CAST_DEVICE_INFORMATION             = "http://%s:%s/setup/eureka_info?options=detail"
	GOOGLE_CAST_DEVICE_WIFI_SCAN               = "https://%s:%s/setup/scan_wifi"
	GOOGLE_CAST_DEVICE_WIFI_SCAN_RESULTS       = "https://%s:%s/setup/scan_results"
	GOOGLE_CAST_DEVICE_WIFI_FORGET             = "https://%s:%s/setup/forget_wifi"
	GOOGLE_CAST_DEVICE_CONFIGURED_NETWORK      = "https://%s:%s/setup/configured_networks"
	GOOGLE_CAST_DEVICE_APPLICATION_URL         = "http://%s:%s/apps/%s"
	GOOGLE_CAST_DEVICE_REBOOT                  = "http://%s:%s/setup/reboot"
	GOOGLE_CAST_DEVICE_DEVICE_DESCRIPTION      = "http://%s:%s/ssdp/device-desc.xml"
	GOOGLE_CAST_DEVICE_DEVICE_NAME             = "https://%s:%s/setup/set_eureka_info"
	GOOGLE_CAST_DEVICE_DEVICE_TIMEZONES        = "http://%s:%s/setup/supported_timezones"
	GOOGLE_CAST_DEVICE_DEVICE_ALARMS           = "http://%s:%s/setup/assistant/alarms"
	GOOGLE_CAST_DEVICE_DEVICE_LEGACYCONFIG     = "https://www.gstatic.com/eureka/config/legacy/config.json"
	GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_STAT   = "http://%s:%s/setup/bluetooth/status"
	GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_PAIRED = "http://%s:%s/setup/bluetooth/get_bonded"
)

//Amazon FireTV API endpoint paths
const (
	AMAZON_FIRE_TV_DEVICE_INFORMATION = "http://%s:53917/zc?action=getInfo&version=2.7.1"
	AMAZON_FIRE_TV_DEVICE_DESCRIPTION = "http://%s:60000/upnp/dev/%s/desc"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____             _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_ _ ___ ___| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | |   |  _|  _| | . |   |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |___|_|_|___|_| |_|___|_|_|___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|
//////////////////////////////////////////////////////////////////////////
//
// DEF  -> This section defines all standard call functions which are registered into the environment and all of the functions that will be called without needing
//
// to use the `register` keyword to register the functions into the environment. All functions defined here should be defined in the init() function call in this file
//
//

func Sprintf(args ...SLC_Object) SLC_Object {
	format := args[0].SL_InspectObject()
	var res string
	for _, x := range args[0:] {
		res = fmt.Sprintf(format, x.SL_InspectObject())
	}
	return &String{Value: res}
}

func Println(args ...SLC_Object) SLC_Object {
	for _, k := range args {
		fmt.Println(k.SL_InspectObject())
	}
	return NilValue
}

func Print(args ...SLC_Object) SLC_Object {
	for _, k := range args {
		fmt.Print(k.SL_InspectObject())
	}
	return NilValue
}

func Sprint(args ...SLC_Object) SLC_Object {
	if len(args) == 0 {
		return NewError("Sprint takes at least one positional argunent")
	}
	var returnedout string
	for _, k := range args {
		returnedout += fmt.Sprint(k.SL_InspectObject())
	}
	return &String{Value: returnedout}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// This section known as the INVOKE_DEPEND section is filled of functions that depend on invokes. So if you invoke like .erase these are the functions that are used
//
// or that the invokes depend on!
//
//

func (SL_Array *Array) DEPEND_INVOKE_ARRAY_COPY(elems []SLC_Object) SLC_Object {
	SL_Array.Elements = append(SL_Array.Elements, elems...)
	return &Nil{}
}

func DEPEND_INVOKE_ARRAY_REVERSE(Elements []SLC_Object) *Array {
	for idx, jex := 0, len(Elements)-1; idx < jex; idx, jex = idx+1, jex-1 {
		Elements[idx], Elements[jex] = Elements[jex], Elements[idx]
	}
	return &Array{Elements: Elements}
}

func (SL_Array *Array) DEPEND_INVOKE_ARRAY_PREPEND(Element SLC_Object) {
	SL_Array.Elements = append([]SLC_Object{Element}, SL_Array.Elements...)
}

func (SL_Array *Array) DEPEND_INVOKE_ARRAY_APPEND(Element SLC_Object) {
	SL_Array.Elements = append(SL_Array.Elements, Element)
}

func (SL_Array *Array) DEPEND_INVOKE_ARRAY_POPLEFT() SLC_Object {
	if len(SL_Array.Elements) > 0 {
		e := SL_Array.Elements[0]
		SL_Array.Elements = SL_Array.Elements[1:]
		return e
	}
	return &Nil{}
}

func (SL_Array *Array) DEPEND_INVOKE_ARRAY_POPRIGHT() SLC_Object {
	if len(SL_Array.Elements) > 0 {
		e := SL_Array.Elements[(len(SL_Array.Elements) - 1)]
		SL_Array.Elements = SL_Array.Elements[:(len(SL_Array.Elements) - 1)]
		return e
	}
	return &Nil{}
}

func (SL_Array *Array) DEPEND_INVOKE_ARRAY_SWAP(F_elem, S_elem int) {
	SL_Array.Elements[F_elem], SL_Array.Elements[S_elem] = SL_Array.Elements[S_elem], SL_Array.Elements[F_elem]
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// The invoke method section which is the title of this section defines all invoke methods for each various SkyLine object. These methods are built in calls and functions that work
//
// by typing the data type or a variable of a specific data tyoe followed by `.` followed by the method. Each function and object allows you to work with the objects methods like so.
//
// STRING.methods() will get all of the methods for the string data type where STRING would be replaced with something like "hello world" and in this case "value".methods()
//
//
// This unit also contains MAPS for each individual function for the objects

type InvokeMethFunctionStr func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object
type InvokeMethFunctionArray func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object

var ArrayFunctionMap = map[string]InvokeMethFunctionArray{
	"Swap": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		if x := CheckArguments("ARRAY.Swap", args, ExactArguments(2), ArgumentHasTypes(IntegerType)); x == nil {
			Swap1, x := strconv.Atoi(args[0].SL_InspectObject())
			if x != nil {
				log.Fatal(x)
			}
			Swap2, x := strconv.Atoi(args[1].SL_InspectObject())
			if x != nil {
				log.Fatal(x)
			}
			ArrayObject.DEPEND_INVOKE_ARRAY_SWAP(Swap1, Swap2)
		}
		return &Nil{}
	},
	"Methods": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		dynamic := Env.Names("array.")
		var names []string
		names = append(names, Static_Array_Methods...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]SLC_Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	},
	"Copy": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return ArrayObject.DEPEND_INVOKE_ARRAY_COPY(args[0].(*Array).Elements)
	},
	"Reverse": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return DEPEND_INVOKE_ARRAY_REVERSE(Values)
	},
	"PopL": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return ArrayObject.DEPEND_INVOKE_ARRAY_POPLEFT()
	},
	"PopR": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return ArrayObject.DEPEND_INVOKE_ARRAY_POPRIGHT()
	},
	"Append": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		if x := CheckArguments("ARRAY.Append", args, ExactArguments(1), ArgumentHasTypes(
			StringType,
			IntegerType,
			NilType,
			FloatType,
			ReturnValueType,
			RegisterType,
			ImportingType,
			FunctionType,
			BooleanType,
			BuiltinType,
			ArrayType,
			HashType,
		)); x == nil {
			ArrayObject.DEPEND_INVOKE_ARRAY_APPEND(args[0])
			return &Nil{}
		} else {
			return &Error{Message: fmt.Sprint(x)}
		}
	},
	"Prepend": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		if x := CheckArguments("ARRAY.Prepend", args, ExactArguments(1), ArgumentHasTypes(
			StringType,
			IntegerType,
			NilType,
			FloatType,
			ReturnValueType,
			RegisterType,
			ImportingType,
			FunctionType,
			BooleanType,
			BuiltinType,
			ArrayType,
			HashType,
		)); x == nil {
			ArrayObject.DEPEND_INVOKE_ARRAY_PREPEND(args[0])
			return &Nil{}
		} else {
			return &Error{Message: fmt.Sprint(x)}
		}
	},
	"Typeof": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &String{Value: ArrayType}
	},
	"Length": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &Integer{Value: int64(len(ArrayObject.Elements))}
	},
	"View": func(Values []SLC_Object, ArrayObject *Array, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &String{Value: ArrayObject.SL_InspectObject()}
	},
}

var StringFunctionMap = map[string]InvokeMethFunctionStr{
	"UnlinkRegistry": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		switch strings.ToLower(Value) {
		case "math":
			for _, k := range MathLibFunctionsRegistered {
				UnlikeRegistryFunctions(k)
			}
			return &Boolean_Object{Value: true}
		case "http":
			for _, k := range HTTPLibFunctionsRegistered {
				UnlikeRegistryFunctions(k)
			}
			return &Boolean_Object{Value: true}
		case "io":
			for _, k := range IOLibFunctionsRegistered {
				UnlikeRegistryFunctions(k)
			}
			return &Boolean_Object{Value: true}
		case "crypt":
			for _, k := range CryptLibFunctionsRegistered {
				UnlikeRegistryFunctions(k)
			}
			return &Boolean_Object{Value: true}
		case "forensics":
			for _, k := range ForensicsFunctionsRegistered {
				UnlikeRegistryFunctions(k)
			}
			return &Boolean_Object{Value: true}
		case "env":
			for _, k := range EnvironmentFunctionsRegistered {
				UnlikeRegistryFunctions(k)
			}
			return &Boolean_Object{Value: true}
		}
		return &Nil{}
	},
	"Length": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &Integer{Value: int64(utf8.RuneCountInString(Value))}
	}, // Get the length of the value
	"Float": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		f, x := strconv.ParseFloat(Value, 64)
		if y := ReturnNewError_Convert(x, &String{Value: Value}, &Float{}); y != nil {
			return y
		}
		return &Integer{Value: int64(f)}
	}, // Convert to float
	"Integer": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		i, x := strconv.ParseInt(Value, 0, 64)
		if y := ReturnNewError_Convert(x, &String{Value: Value}, &Integer{}); y != nil {
			return y
		}
		return &Integer{Value: int64(i)}
	}, // Convert to integer
	"Boolean": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		b, x := strconv.ParseBool(Value)
		if y := ReturnNewError_Convert(x, &String{Value: Value}, &Boolean_Object{}); y != nil {
			return y
		}
		return &Boolean_Object{Value: b}
	}, // Convert to boolean
	"Trim": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		if x := CheckArguments("STRING.Trim", args, ExactArguments(1), ArgumentHasTypes(StringType)); x == nil {
			cutset := args[0].SL_InspectObject()
			newarr := make([]SLC_Object, 0)
			for _, k := range strings.Trim(Value, cutset) {
				newarr = append(newarr, &String{Value: string(k)})
			}
			return &Array{Elements: newarr}
		} else {
			return &Error{Message: fmt.Sprint(x)}
		}
	},
	"Ord": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &Integer{Value: int64(Value[0])}
	}, // short for "ordinal" and is a function that returns the Unicode code point of a character. In the provided code snippet, the ord method takes the first character of the String object (Str.Value[0]) and returns its Unicode code point as an Integer object.
	"Split": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		if x := CheckArguments("STRING.split", args, ExactArguments(1), ArgumentHasTypes(StringType)); x != nil {
			return NewError(x.Error())
		}
		Arr := strings.Split(Value, args[0].SL_InspectObject())
		newarr := make([]SLC_Object, 0)
		for _, idx := range Arr {
			newarr = append(newarr, &String{Value: idx})
		}
		return &Array{Elements: newarr}
	}, // split string
	"Upper": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &String{Value: strings.ToUpper(Value)}
	}, // Convert to upper case
	"Lower": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &String{Value: strings.ToLower(Value)}
	}, // Convert to lower case
	"Title": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		return &String{Value: strings.ToTitle(Value)}
	},
	"Methods": func(Value string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
		dynamic := Env.Names("string.")
		var names []string
		names = append(names, Static_String_Methods...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]SLC_Object, len(names))
		for Idx, txt := range names {
			result[Idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}, // Get all methods
}

func (Arr *Array) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if f, ok := ArrayFunctionMap[method]; ok {
		return f(Arr.Elements, Arr, Env, args...)
	}
	return nil
}

func (Str *String) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if f, ok := StringFunctionMap[method]; ok {
		return f(Str.Value, Env, args...)
	} else if method == "View" {
		return &String{Value: Str.SL_InspectObject()}
	} else {
		return &Error{Message: fmt.Sprintf("Invoke Error: (Object Call) -> Hm, it seems as if `%s` is not an actual function call? ", method)}
	}
}

func (ENGINE_VAL *ENGINE_Value) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (iexp *ImportExpression) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (mod *Module) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (null *Nil) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (Err *Error) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (RegisterVal *RegisterValue) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (ret *ReturnValue) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	return nil
}

func (f *Function) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if method == "methods" {
		static := []string{"methods"}
		dynamic := Env.Names("function.")
		var names []string
		names = append(names, static...)
		for _, environ := range dynamic {
			bits := strings.Split(environ, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		results := make([]SLC_Object, len(names))
		for IDX, TXT := range names {
			results[IDX] = &String{Value: TXT}
		}
		return &Array{Elements: results}
	}
	return nil
}

func (hash *Hash) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if method == "keys" {
		ents := len(hash.Pairs)
		arrays := make([]SLC_Object, ents)
		idx := 0
		for _, entity := range hash.Pairs {
			arrays[idx] = entity.Key
			idx++
		}
		return &Array{Elements: arrays}
	}
	return nil
}

func (Bool *Boolean_Object) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if method == "methods" {
		dynamic := Env.Names("bool.")
		var names []string
		names = append(names, Static_Boolean_Methods...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]SLC_Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (BuiltIn *Builtin) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if method == "methods" {
		names := []string{"methods"}
		result := make([]SLC_Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (float *Float) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if method == "methods" {

		dynamic := Env.Names("float.")
		var names []string
		names = append(names, Static_Float_Methods...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]SLC_Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

func (integer *Integer) InvokeMethod(method string, Env Environment_of_environment, args ...SLC_Object) SLC_Object {
	if method == "chr" {
		return &String{Value: string(rune(integer.Value))}
	}
	if method == "methods" {
		dynamic := Env.Names("integer.")
		var names []string
		names = append(names, Static_Integer_Methods...)
		for _, envi := range dynamic {
			bits := strings.Split(envi, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]SLC_Object, len(names))
		for idx, txt := range names {
			result[idx] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	return nil
}

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
// Filename      |  SkyLine_Standard_PreRegistered_Library_Mathematics.go
// Project       |  SkyLine programming language
// Line Count    |  220+ active lines
// Status        |  Working and Active
// Package       |  SkyLine_Backend
//
//
// Defines       | This file defines all core library functions, kind of everything mashed into one giant file. The reason this file exists right now is because as of
//                 0.0.5 of the language, SkyLine does not have a proper modular structure to the language and is still in testing. When further development comes about
//                 and other core features are used or implemented, we will then require that the library be accessed externally to the `Modules` directory. This will make
//                 it much easier to access, read, and create a full fledged standard library.
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
// This unit is the prime unit for any registration functions that go on or are executed when parsing and evaluating the `register()`keyword and the arguments inside
//
// of register. For context, register is similar to the import keyword but imstead of importing values and setting them into the environment from a file, register rather
//
// takes the contents of a library section and registers them using registry functions into the environment.
//
//
//::END_UNIT

// Register mathematical functions
func RegisterMath() {
	RegisterBuiltin("math.abs",
		func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
			return (MathAbs(args...))
		})
	RegisterBuiltin("math.cos",
		func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
			return (MathCos(args...))
		})
	RegisterBuiltin("math.tan",
		func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
			return (MathTan(args...))
		})
	RegisterBuiltin("math.sin",
		func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
			return (MathSin(args...))
		})
	RegisterBuiltin("math.sqrt",
		func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
			return (MathSqrt(args...))
		})
	RegisterBuiltin("math.cbrt",
		func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
			return (MathCbrt(args...))
		})
	RegisterBuiltin("math.pow", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (MathPow(args...))
	})
	RegisterBuiltin("math.rand", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (MathRand())
	})
}

// Register IO functions
func RegisterIO() {
	RegisterBuiltin("io.input", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (User_Input(args...))
	})
	RegisterBuiltin("io.clear", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (IO_Clear())
	})
	RegisterBuiltin("io.box", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (IO_Box(args...))
	})
	RegisterBuiltin("io.listen", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (IO_Listen(args...))
	})
	RegisterBuiltin("io.restore", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		fmt.Println(SKYLINE_RESTORE)
		return (&Nil{})
	})
}

// Register standard function calls.
// we use init() because this will be the prime standard keywords / built in functions like print or println that do not require any base
// registry

func init() {
	ErrorSys.Box = false
	ErrorSys.Line = false
	ErrorSys.Tree = true
	RegisterBuiltin("sprintf", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Sprintf(args...))
	})
	RegisterBuiltin("print", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Print(args...))
	})
	RegisterBuiltin("println", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Println(args...))
	})
	RegisterBuiltin("sprint", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Sprint(args...))
	})
	RegisterBuiltin("modify", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (ModifyEnviornment(args...))
	})
	RegisterBuiltin("plugin", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (LoadPlugin(args...))
	})
}

// Registers crypt library functions
func RegisterCrypt() {
	RegisterBuiltin("crypt.hash", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Crypto_Hasher(args...))
	})
}

// Register forensics standard and base functions
func RegisterForensicsSubFunctions() {
	RegisterBuiltin("ForensicsUtils.FindSigUnknownFile", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunFileVerification(args...))
	})
	RegisterBuiltin("ForensicsUtils.CheckZIPSig", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunZipCheck(args...))
	})
}

// Register forensics related functions |Apple|
func RegisterForensicsApple() {

}

// Register general image injection and image utilities |Forensics/Stego/Image|
func RegisterImageUtilitiesForensics() {
	RegisterBuiltin("ImageUtils.InjectImage", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunFileInjFromFileToFile(args...))
	})
	RegisterBuiltin("ImageUtils.CreationNew", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunFileNewSettings(args...))
	})
	RegisterBuiltin("ImageUtils.CreateImage", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunFileCreationFromFile(args...))
	})
}

// Register forensics related functions
func RegisterForensicsPNG() {
	RegisterBuiltin("ForensicsPng.PngChunkCount", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Run_Forensics_Metadata_PNG_GetNumChunks(args...))
	})
	RegisterBuiltin("ForensicsPng.PngSettingsNew", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunSettings(args...))
	})
	RegisterBuiltin("ForensicsPng.PngInject", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunInject(args...))
	})
	RegisterBuiltin("ForensicsPng.InjectRegular", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (RunPayloadInjectionRegular(args...))
	})

}

// Register environment functions
func RegisterEnvironment() {
	RegisterBuiltin("env.Getenv", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Environment_GetEnvironmentPath(args...))
	})
	RegisterBuiltin("env.Setenv", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Environment_SetEnvironment(args...))
	})
	RegisterBuiltin("env.Environment", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (Environment_GetEnvironment(args...))
	})
}

//Register HTTP functions
func RegisterHTTP() {
	RegisterBuiltin("http.Get", func(env *Environment_of_environment, args ...SLC_Object) SLC_Object {
		return (MakeGet(args...))
	})
	// Variables
	RegisterVariable("http.MethodDelete", &String{Value: HTTPMETHOD_DELETE})
	RegisterVariable("http.MethodPatch", &String{Value: HTTPMETHOD_PATCH})
	RegisterVariable("http.MethodPuts", &String{Value: HTTPMETHOD_PUTS})
	RegisterVariable("http.MethodPost", &String{Value: HTTPMETHOD_POST})
	RegisterVariable("http.MethodGet", &String{Value: HTTPMETHOD_GET})
	RegisterVariable("http.MethodHead", &String{Value: HTTPMETHOD_HEAD})
	RegisterVariable("http.MethodOptions", &String{Value: HTTPMETHOD_OPTIONS})
	RegisterVariable("http.MethodTrace", &String{Value: HTTPMETHOD_TRACE})
}

func RegisterAppleIoTDatabase() {
	RegisterVariable("Apple.AirPlayServerInfo", &String{Value: APPLE_AIRPLAY_SERVERINFO})
	RegisterVariable("Apple.AirPlayPlayBackInfo", &String{Value: APPLE_AIRPLAY_PLAYBACKINFO})
	RegisterVariable("Apple.AirPlayScrubInfo", &String{Value: APPLE_AIRPLAY_SCRUB})
	RegisterVariable("Apple.AirPlayStreamInfo", &String{Value: APPLE_AIRPLAY_STREAMINFO})
	RegisterVariable("Apple.AirPlayInfo", &String{Value: APPLE_AIRPLAY_INFO})
	RegisterVariable("Apple.DAAPMain", &String{Value: APPLE_DAAP_PATH})
	RegisterVariable("Apple.DAAPLogin", &String{Value: APPLE_DAAP_LOGIN})
	RegisterVariable("Apple.DAAPDatabase", &String{Value: APPLE_DAAP_DATABASE})
}

func RegisterRokuIoTDatabase() {
	RegisterVariable("Roku.KeyPressHome", &String{Value: ROKU_KEYPRESS_HOME})
	RegisterVariable("Roku.KeyPressPlay", &String{Value: ROKU_KEYPRESS_PLAY})
	RegisterVariable("Roku.KeyPressUp", &String{Value: ROKU_KEYPRESS_UP})
	RegisterVariable("Roku.KeyPressDown", &String{Value: ROKU_KEYPRESS_DOWN})
	RegisterVariable("Roku.KeyPressLeft", &String{Value: ROKU_KEYPRESS_LEFT})
	RegisterVariable("Roku.KeyPressRight", &String{Value: ROKU_KEYPRESS_RIGHT})
	RegisterVariable("Roku.KeyPressSelect", &String{Value: ROKU_KEYPRESS_OK})
	RegisterVariable("Roku.KeyPressRewind", &String{Value: ROKU_KEYPRESS_REWIND})
	RegisterVariable("Roku.KeyPressFFW", &String{Value: ROKU_KEYPRESS_FFW})
	RegisterVariable("Roku.KeyPressOptions", &String{Value: ROKU_KEYPRESS_OPTIONS})
	RegisterVariable("Roku.KeyPressPause", &String{Value: ROKU_KEYPRESS_PAUSE})
	RegisterVariable("Roku.KeyPressBack", &String{Value: ROKU_KEYPRESS_BACK})
	RegisterVariable("Roku.KeyPressPoweroff", &String{Value: ROKU_KEYPRESS_POWEROFF})
	RegisterVariable("Roku.KeyPressVolumeUp", &String{Value: ROKU_KEYPRESS_VUP})
	RegisterVariable("Roku.KeyPressVolumeDown", &String{Value: ROKU_KEYPRESS_VDOWN})
	RegisterVariable("Roku.KeyPressVolumeMute", &String{Value: ROKU_KEYPRESS_MUTE})
	RegisterVariable("Roku.DeviceDown", &String{Value: ROKU_KEYPRESS_DEVDOWN})
	RegisterVariable("Roku.DeviceUp", &String{Value: ROKU_KEYPRESS_DEVUP})
	RegisterVariable("Roku.DevAppLaunch", &String{Value: ROKU_DEVICE_LAUNCH})
	RegisterVariable("Roku.DevAppInstall", &String{Value: ROKU_DEVICE_INSTALL})
	RegisterVariable("Roku.DevDisableSGR", &String{Value: ROKU_DEVICE_DISABLE_SGR})
	RegisterVariable("Roku.DevEnableSGR", &String{Value: ROKU_DEVICE_ENABLE_SGR})
	RegisterVariable("Roku.DevTV", &String{Value: ROKU_DEVICE_TV})
	RegisterVariable("Roku.DevSGNODES", &String{Value: ROKU_DEVICE_SGNODE})
	RegisterVariable("Roku.DevActiveTVS", &String{Value: ROKU_DEVICE_ACTIVETV})
	RegisterVariable("Roku.DevDial", &String{Value: ROKU_DEVICE_DIAL})
	RegisterVariable("Roku.DeviceBrowse", &String{Value: ROKU_DEVICE_BROWSE})
	RegisterVariable("Roku.DeviceInformation", &String{Value: ROKU_DEVICE_INFO})
	RegisterVariable("Roku.DeviceApplications", &String{Value: ROKU_DEVICE_APPS})
	RegisterVariable("Roku.DeviceActiveApplications", &String{Value: ROKU_DEVICE_ACTIVE})
}

func RegisterGoogleIoTDatabase() {
	RegisterVariable("Google.CastDeviceInfo", &String{Value: GOOGLE_CAST_DEVICE_INFORMATION})
	RegisterVariable("Google.CastDeviceReboot", &String{Value: GOOGLE_CAST_DEVICE_REBOOT})
	RegisterVariable("Google.CastDeviceDescription", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_DESCRIPTION})
	RegisterVariable("Google.CastDeviceWiFiForget", &String{Value: GOOGLE_CAST_DEVICE_WIFI_FORGET})
	RegisterVariable("Google.CastDeviceWiFiScan", &String{Value: GOOGLE_CAST_DEVICE_WIFI_SCAN})
	RegisterVariable("Google.CastDeviceWiFiScanResults", &String{Value: GOOGLE_CAST_DEVICE_WIFI_SCAN_RESULTS})
	RegisterVariable("Google.CastDeviceWiFiConfigured", &String{Value: GOOGLE_CAST_DEVICE_CONFIGURED_NETWORK})
	RegisterVariable("Google.CastDeviceAlarms", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_ALARMS})
	RegisterVariable("Google.CastDeviceTimezones", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_TIMEZONES})
	RegisterVariable("Google.CastDeviceLegacyConf", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_LEGACYCONFIG})
	RegisterVariable("Google.CastDeviceBleStat", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_STAT})
	RegisterVariable("Google.CastDeviceBlePaired", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_PAIRED})
	RegisterVariable("Google.CastDeviceSetName", &String{Value: GOOGLE_CAST_DEVICE_DEVICE_NAME})
	RegisterVariable("Google.CastDeviceApplications", &String{Value: GOOGLE_CAST_DEVICE_APPLICATION_URL})
}

func RegisterAmazonIoTDatabase() {
	RegisterVariable("Amazon.TvDeviceInformation", &String{Value: AMAZON_FIRE_TV_DEVICE_INFORMATION})
	RegisterVariable("Amazon.TvDeviceDescription", &String{Value: AMAZON_FIRE_TV_DEVICE_DESCRIPTION})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____             _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_ _ ___ ___| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | |   |  _|  _| | . |   |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |___|_|_|___|_| |_|___|_|_|___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|
//////////////////////////////////////////////////////////////////////////
//
// DEF  -> This section defines all of the helper functions used within the standard library or is called within this file
//
//

type FunctionCall func(name string, arguments []SLC_Object) error

func ExactArguments(exact int) FunctionCall {
	return func(name string, arguments []SLC_Object) error {
		if len(arguments) != exact {
			return fmt.Errorf("Argument error (Standard Library): The standard function `%s()` takes %d argument(s) but you gave %s argument(s)", name, exact, fmt.Sprint(len(arguments)))
		}
		return nil
	}
}

func RangeArguments(minimum_args, maximum_args int) FunctionCall {
	return func(name string, arguments []SLC_Object) error {
		if len(arguments) < minimum_args || len(arguments) > maximum_args {
			return fmt.Errorf("Argument error (Standard Library): The standard function `%s()` takes at least %d argument(s) and at most %d argument(s) | You gave %d argument(s)", name, minimum_args, maximum_args, len(arguments))
		}
		return nil
	}
}

func ArgumentHasTypes(types ...Type_Object) FunctionCall {
	return func(name string, arguments []SLC_Object) error {
		for arg, datatype := range types {
			if arg < len(arguments) && arguments[arg].SL_RetrieveDataType() != datatype {
				return fmt.Errorf(
					"Argument error (Standard Library): The standard function `%s()` expects positional argument #%d to be of type `%s` but you gave `%s`",
					name, (arg + 1), datatype, arguments[arg].SL_RetrieveDataType(),
				)
			}
		}
		return nil
	}
}

func NotEnoughArguments(min int) FunctionCall {
	return func(name string, arguments []SLC_Object) error {
		if len(arguments) < min {
			return fmt.Errorf("Argument Error (Standard Library): The standard function `%s()` takes at least %d arguments but you provided %d", name, min, len(arguments))
		}
		return nil
	}
}

func CheckArguments(name string, arguments []SLC_Object, checks ...FunctionCall) error {
	for _, run := range checks {
		if x := run(name, arguments); x != nil {
			return x
		}
	}
	return nil
}

func ReturnNewError_Convert(x error, args ...SLC_Object) SLC_Object {
	if x != nil {
		return NewError("Argument Error (Invoke | Object Call): SkyLine could not convert the value (`%s`) into type (`%s`) ", args[0].SL_InspectObject(), args[1].SL_RetrieveDataType())
	}
	return nil
}
