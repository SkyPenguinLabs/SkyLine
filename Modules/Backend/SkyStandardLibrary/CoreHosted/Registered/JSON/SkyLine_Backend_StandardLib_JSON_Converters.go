//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _                __ _____ _____ _____
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___   |  |   __|     |   | |
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|  |  |__   |  |  | | | |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_____|_____|_|___|
//	primary units that define the functions to register the sub func's  //            |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This part of the standard library contains any and all functions for JSON data sets and functions. This includes dumping relative information, dumping file information,
//
// gnerating golang structures for plugins, generating infromation, generating strings, dumping to a hash map and so on from there.
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This file defines all of the functions to help with conversions
//
package SkyLine_Standard_Library_JSON

import SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"

func MapToHash(m interface{}) *SkyEnv.SL_HashMap {
	pairs := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
	if m, ok := m.(map[string]interface{}); ok {
		for k, v := range m {
			key := &SkyEnv.SL_String{Value: k}
			switch val := v.(type) {
			case string:
				pairs[key.SL_HashKeyType()] = SkyEnv.HashPair{
					Key:   key,
					Value: &SkyEnv.SL_String{Value: val},
				}
			case int:
				pairs[key.SL_HashKeyType()] = SkyEnv.HashPair{
					Key:   key,
					Value: &SkyEnv.SL_Integer{Value: val},
				}
			case bool:
				pairs[key.SL_HashKeyType()] = SkyEnv.HashPair{
					Key:   key,
					Value: &SkyEnv.SL_Boolean{Value: val},
				}
			case map[string]interface{}:
				pairs[key.SL_HashKeyType()] = SkyEnv.HashPair{
					Key:   key,
					Value: MapToHash(val),
				}
			case []interface{}:
				arr := make([]SkyEnv.SL_Object, 0)
				for _, elem := range val {
					switch elemVal := elem.(type) {
					case string:
						arr = append(arr, &SkyEnv.SL_String{Value: elemVal})
					case int:
						arr = append(arr, &SkyEnv.SL_Integer{Value: elemVal})
					case bool:
						arr = append(arr, &SkyEnv.SL_Boolean{Value: elemVal})
					case map[string]interface{}:
						arr = append(arr, MapToHash(elemVal))
					}
				}
				pairs[key.SL_HashKeyType()] = SkyEnv.HashPair{
					Key:   key,
					Value: &SkyEnv.SL_Array{Elements: arr},
				}
			default:
				pairs[key.SL_HashKeyType()] = SkyEnv.HashPair{
					Key:   key,
					Value: &SkyEnv.SL_NULL{},
				}
			}
		}
	}
	return &SkyEnv.SL_HashMap{Pairs: pairs}
}
