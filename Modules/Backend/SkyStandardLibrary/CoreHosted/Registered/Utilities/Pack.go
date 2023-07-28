package SkyLine_Utilities

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	"encoding/binary"
	"fmt"
)

func Pwn_Pack64(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	b := make([]byte, 8)
	switch argt := args[0].(type) {
	case *SkyEnv.SL_Integer:
		binary.LittleEndian.PutUint64(b, uint64(argt.Value))
	case *SkyEnv.SL_Integer16:
		binary.LittleEndian.PutUint64(b, uint64(argt.Value))
	case *SkyEnv.SL_Integer32:
		binary.LittleEndian.PutUint64(b, uint64(argt.Value))
	case *SkyEnv.SL_Integer64:
		binary.LittleEndian.PutUint64(b, uint64(argt.Value))
	case *SkyEnv.SL_Integer8:
		binary.LittleEndian.PutUint64(b, uint64(argt.Value))
	default:
		return &SkyEnv.SL_Error{Message: "Sorry, integer data type is the only one allowed or supported"}
	}
	barr := make([]SkyEnv.SL_Object, 0)
	for i := 0; i < len(b); i++ {
		value := fmt.Sprintf("%#v", b[i])
		barr = append(barr, &SkyEnv.SL_String{Value: value})
	}
	return &SkyEnv.SL_Array{Elements: barr}
}

func Pwn_Pack32(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	b := make([]byte, 4)
	switch argt := args[0].(type) {
	case *SkyEnv.SL_Integer:
		binary.LittleEndian.PutUint32(b, uint32(argt.Value))
	case *SkyEnv.SL_Integer16:
		binary.LittleEndian.PutUint32(b, uint32(argt.Value))
	case *SkyEnv.SL_Integer32:
		binary.LittleEndian.PutUint32(b, uint32(argt.Value))
	case *SkyEnv.SL_Integer64:
		binary.LittleEndian.PutUint32(b, uint32(argt.Value))
	case *SkyEnv.SL_Integer8:
		binary.LittleEndian.PutUint32(b, uint32(argt.Value))
	default:
		return &SkyEnv.SL_Error{Message: "Sorry, integer data type is the only one allowed or supported"}
	}
	var dt string
	for i := 0; i < len(b); i++ {
		value := fmt.Sprintf("%#v", b[i])
		fmt.Println(value)
	}
	dt += fmt.Sprintf("%#v", b)
	return &SkyEnv.SL_String{Value: dt}
}
