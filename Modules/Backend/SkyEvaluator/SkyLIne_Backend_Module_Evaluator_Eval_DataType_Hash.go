///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_DataType_Hash
// Extension         | .go ( golang source code file )
// Purpose           | Defines all functions for evaluating the hash data type
// Directory         | Modules/Backend/SkyEvaluator
// Modular Directory | github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEvaluator
// Package Name      | SkyLine_Backend_Module_Evaluation
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// The final part in standard interpretation inside of a programming language is to evaluate and execute the data or keys. In the case of SkyLine, it relies on different
//
// forms of engines which can use the byte code compiler or use the evaluator. The byte code compiler is a whole different story but the evaluator will take advantage of the
//
// AST and then check and execute conditions, statements, values, or modifications accordingly. The evaluator can also sometimes be complex to use but it still manages to stay
//
// one of the fastest ones to write.
//
package SkyLine_Backend_Evaluation

import (
	"fmt"

	SkyAST "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyAST"
	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

func SkyLine_Evaluator_Eval_Hash_Map_Literal(Node *SkyAST.SL_HashMap, SkyEnvironment *SkyEnv.SkyLineEnvironment) SkyEnv.SL_Object {
	HashPairs := make(map[SkyEnv.HashKey]SkyEnv.HashPair)

	for KeyNode, ValueNode := range Node.Pairs {
		KeyN := SkyLine_Call_Eval(KeyNode, SkyEnvironment)
		if SkyLine_Evaluator_CheckError(KeyN) {
			return KeyN
		}

		HashKey, ok := KeyN.(SkyEnv.SL_Hashable)
		if !ok {
			return SkyLine_Evaluator_CreateError("Unusable as a hash key: %s", KeyN.SkyLine_ObjectFunction_GetDataType())
		}

		ValueN := SkyLine_Call_Eval(ValueNode, SkyEnvironment)

		if SkyLine_Evaluator_CheckError(ValueN) {
			return ValueN
		}

		Hashed := HashKey.SL_HashKeyType()

		HashPairs[Hashed] = SkyEnv.HashPair{
			Key:   KeyN,
			Value: ValueN,
		}
	}
	return &SkyEnv.SL_HashMap{
		Pairs: HashPairs,
	}
}

func SkyLine_Evaluator_Eval_Hash_Map_IndexExpression(Map, Index SkyEnv.SL_Object) SkyEnv.SL_Object {
	key, ok := Index.(SkyEnv.SL_Hashable)
	if !ok {
		return SkyLine_Evaluator_CreateError("Unusable hash key - %s", Index.SkyLine_ObjectFunction_GetTrueValue())
	}
	hashObj := Map.(*SkyEnv.SL_HashMap)
	if pair, exists := hashObj.Pairs[key.SL_HashKeyType()]; exists {
		return pair.Value
	} else {
		fmt.Println("Exists? -> ", exists)
		fmt.Println("Tried indexing -> ", hashObj.Pairs[key.SL_HashKeyType()])
		fmt.Println("Index hash key -> ", key.SL_HashKeyType().Value)
	}
	return &SkyEnv.SL_NULL{}
}
