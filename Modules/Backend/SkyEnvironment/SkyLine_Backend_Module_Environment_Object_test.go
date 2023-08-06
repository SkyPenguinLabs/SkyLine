//test:package
package SkyLine_Backend_Modules_Objects

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &SL_String{Value: "Hello World"}
	hello2 := &SL_String{Value: "Hello World"}
	diff1 := &SL_String{Value: "My name is johnny"}
	diff2 := &SL_String{Value: "My name is johnny"}
	if hello1.SL_HashKeyType() != hello2.SL_HashKeyType() {
		t.Errorf("string with same content have different keys")
	}
	if diff1.SL_HashKeyType() != diff2.SL_HashKeyType() {
		t.Errorf("string with same content have different keys")
	}
	if hello1.SL_HashKeyType() == diff1.SL_HashKeyType() {
		t.Errorf("string with different have same hash key")
	}
}
