package SkyLine_Utilities

import (
	"strconv"
	"strings"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

type PwnBuffer struct {
	Initated bool
	data     []byte
}

var (
	Pwnbuf PwnBuffer
)

func CheckBufCtx() bool {
	return Pwnbuf.Initated
}

func (buf *PwnBuffer) Add(data []byte) {
	buf.data = append(buf.data, data...)
}

func (buf *PwnBuffer) Unget(data []byte) {
	buf.data = append(data, buf.data...)
}

func (buf *PwnBuffer) Retrieve(length int) []byte {
	if length > len(buf.data) || length == 0 {
		length = len(buf.data)
	}
	result := buf.data[:length]
	buf.data = buf.data[length:]
	return result
}

func (buf *PwnBuffer) Index(target string) int {
	str := string(buf.data)
	index := strings.Index(str, target)
	if index == -1 {
		return -1
	}
	return index
}

func (buf *PwnBuffer) Destroy() {
	for i := 0; i < len(buf.data); i++ {
		buf.data[i] = 0
	}
	Pwnbuf.Initated = false
}

func BufferImpl_New(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	Pwnbuf = PwnBuffer{}
	Pwnbuf.Initated = true
	return &SkyEnv.SL_NULL{}
}

func BufferImpl_Get(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if !CheckBufCtx() {
		return &SkyEnv.SL_Error{Message: "Error when implementing get: Buffer was not created, use NewBuf() to create one"}
	}
	var v int
	switch t := args[0].(type) {
	case *SkyEnv.SL_Integer:
		v = int(t.Value)
	case *SkyEnv.SL_Integer16:
		v = int(t.Value)
	case *SkyEnv.SL_Integer8:
		v = int(t.Value)
	case *SkyEnv.SL_Integer32:
		v = int(t.Value)
	case *SkyEnv.SL_Integer64:
		v = int(t.Value)
	case *SkyEnv.SL_String:
		if vt, ok := strconv.Atoi(t.Value); ok == nil {
			v = int(vt)
		} else {
			return &SkyEnv.SL_Error{Message: "Data type was string in argument to BufGet() which requires (Integer|String) however if it is a string it must represent a number that can be converted to an integer data type"}
		}
	}
	return &SkyEnv.SL_String{Value: string(Pwnbuf.Retrieve(v))}
}

func BufferImpl_Set(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if !CheckBufCtx() {
		return &SkyEnv.SL_Error{Message: "Error when implementing set: Buffer was not created, use NewBuf() to create one"}
	}
	if args != nil {
		var str string
		if ct, ok := args[0].(*SkyEnv.SL_String); ok {
			str = ct.Value
		} else {
			return &SkyEnv.SL_Error{Message: "Data type used as an argument could not be converted to a string"}
		}
		Pwnbuf.Add([]byte(str))
	}
	return &SkyEnv.SL_NULL{}
}

func BufferImpl_Unget(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if !CheckBufCtx() {
		return &SkyEnv.SL_Error{Message: "Error when implementing unget: Buffer was not created, use NewBuf() to create to create one"}
	}
	if args != nil {
		var str string
		if ct, ok := args[0].(*SkyEnv.SL_String); ok {
			str = ct.Value
		} else {
			return &SkyEnv.SL_Error{Message: "Data type used as argument could not be converted to type string"}
		}
		Pwnbuf.Unget([]byte(str))
	}
	return &SkyEnv.SL_String{Value: string(Pwnbuf.data)}
}

func BufferImpl_Destroy(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if !CheckBufCtx() {
		return &SkyEnv.SL_Error{Message: "Error when implementing destroy: Buffer was not created, use NewBuf() to create one"}
	}
	Pwnbuf.Destroy()
	return &SkyEnv.SL_NULL{}
}
