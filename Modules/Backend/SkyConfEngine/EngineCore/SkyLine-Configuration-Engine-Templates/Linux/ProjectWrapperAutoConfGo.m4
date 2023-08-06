dnl | SLC Programmatic execution and interfaces for templating code
dnl | ----------------------------------------------------------------
dnl | SLC allows you to generate programs and code, one way we can do that
dnl | rather choose to do that is by using autoconf/M4-Sugar. In this case 
dnl | these files are used to generate C++/C wrappers and Golang file templates
dnl | for projects as well as json files for linux based operating systems that 
dnl | may have or use M4 installed on the system. 
dnl | ----------------------------------------------------------------------
dnl | This file is for golang wrapping code

divert(0)

dnl ifdef(`input', , `define(`input', `null')')

define(`GoProg', `
// declare main package
package main

// Import proper libs
import (
	"fmt"

	SkyEnvironment "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
)

/*
#include "CPPImplementation.hpp"
*/
import "C"

// Interact with SkyLines backend
func Call(env *SkyEnvironment.SkyLineEnvironment, args ...SkyEnvironment.SL_Object) SkyEnvironment.SL_Object {
	println(C.callCPPFunction()) // Call CPP function
	return &SkyEnvironment.SL_Integer{Value: 1}
}

func main() {
	result := Call(&SkyEnvironment.SkyLineEnvironment{}, &SkyEnvironment.SL_String{Value: "Hello world"})
	fmt.Println(result)
}
')

define(`auto_configuration_go', `GoProg')

auto_configuration_go