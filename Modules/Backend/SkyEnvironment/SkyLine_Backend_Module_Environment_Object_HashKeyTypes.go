///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Environment_Object_HashKeyTypes
// Extension         | .go ( golang source code file )
// Purpose           | Defines all hashable functions for data types (integer, string, float) or any type that needs to be hashed
// Directory         | Modules/Backend/SkyScanner
// Modular Directory | SkyLine/Modules/Backend/SkyScanner
// Package Name      | SkyLine_Backend_Scanner
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This sector of files also known as the 'ObjectType' sector will be defining all of the functions for the current data types within SkyLine.
//
// WARN: This note exists for every single Object file in this directory
//
// These functions include
//						A - SkyLine_ObjectFunction_GetDataType
//						B - SkyLine_ObjectFunction_GetTrueValue
//						C - SkyLine_ObjectFunction_GetInterface
//                      D - SkyLine_ObjectFunction_InvokeObject
//
//
//
// A: Data Type return
//			- This function will tell the SkyLine interpreter to grab or return the data type of the current object.
//			  This helps during parsing and evaluation to ensure that the data type is what it is and to make sure that the object
//			  being executed has the proper type.
//
// B: Get True Value
//			- This function may be confusing at first, but it takes a value such as an array and makes the most syntactic appropriate
//			  representation of the data type and then returns it as a string. This function starts by first declaring a byte input and
//			  also writing those values of the data types proper specification and returning it as a string. For example, in SkyLine
//			  variables are dynamic, this means that everything is `interpreted` and variables have no constant type based values. If you
//			  declare a string variable like `set x := "name"; ` then the result of that from the function with lable `B(Get True Value)` will
//			  return `TYPE STRING VARIABLE (x) := STRING[name];` which can be helpful for optimization and syntactic purposes.`
//
// C: Get Interface:
//			- This function just returns a type interface representing the object
//
// D: Invoke Object:
//			- This allows objects to have their own specific functions such as ORD, LEN, SPLIT etc. Below is a table of common object
//			  call functions for different data types
//
//				| SkyLine Data Type | Object Call Function List  |
//   			------------------- | -------------------------- |
//				Integer             | [methods, ord]             |
//				String              | [methods, split, index ...]|
//              Boolean             | [methods]                  |
//              Float               | [methods]                  |
//               ...                | ...
//
// Object call functions are a ton easier and better to work with than requiring the user to require a data type such as `register(Array)`
//
// to get array object calls. In the further future this may actually be required due to the amount of object call functions that will exist
//
package SkyLine_Backend_Modules_Objects

import "hash/fnv"

// Byte
func (SL_ObjectByte *SL_Byte) SL_HashKeyType() HashKey {
	return HashKey{
		ObjectDataType: SL_ObjectByte.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(SL_ObjectByte.Value),
	}
}

// Float
func (SL_ObjectFloat *SL_Float) SL_HashKeyType() HashKey {
	HASH := fnv.New64a()
	HASH.Write([]byte(SL_ObjectFloat.SkyLine_ObjectFunction_GetTrueValue()))
	return HashKey{
		ObjectDataType: SL_ObjectFloat.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(HASH.Sum64()),
	}
}

// String
func (SL_ObjectString *SL_String) SL_HashKeyType() HashKey {
	HASH := fnv.New64a()
	HASH.Write([]byte(SL_ObjectString.Value))
	return HashKey{
		ObjectDataType: SL_ObjectString.SkyLine_ObjectFunction_GetDataType(),
		Value:          HASH.Sum64(),
	}
}

// Integer
func (SL_ObjectInteger *SL_Integer) SL_HashKeyType() HashKey {
	return HashKey{
		ObjectDataType: SL_ObjectInteger.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(SL_ObjectInteger.Value),
	}
}

// Integer 8
func (SL_ObjectInteger8 *SL_Integer8) SL_HashKeyType() HashKey {
	return HashKey{
		ObjectDataType: SL_ObjectInteger8.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(SL_ObjectInteger8.Value),
	}
}

// Integer 16
func (SL_ObjectInteger16 *SL_Integer16) SL_HashKeyType() HashKey {
	return HashKey{
		ObjectDataType: SL_ObjectInteger16.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(SL_ObjectInteger16.Value),
	}
}

// Integer32

func (SL_ObjectInteger32 *SL_Integer32) SL_HashKeyType() HashKey {
	return HashKey{
		ObjectDataType: SL_ObjectInteger32.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(SL_ObjectInteger32.Value),
	}
}

// Integer 64

func (SL_ObjectInteger64 *SL_Integer64) SL_HashKeyType() HashKey {
	return HashKey{
		ObjectDataType: SL_ObjectInteger64.SkyLine_ObjectFunction_GetDataType(),
		Value:          uint64(SL_ObjectInteger64.Value),
	}
}

// Boolean
func (SL_ObjectBoolean *SL_Boolean) SL_HashKeyType() HashKey {
	var val uint64
	switch SL_ObjectBoolean.Value {
	case true:
		val = 1
	default:
		val = 0
	}
	return HashKey{
		ObjectDataType: SL_ObjectBoolean.SkyLine_ObjectFunction_GetDataType(),
		Value:          val,
	}
}
