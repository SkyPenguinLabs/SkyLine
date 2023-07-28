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
package main 
/*
#include "cpp_wrapper.hpp"
*/
import "C"
func main() {
    res := C.callCPPFunction()
    println(msg + res)
}
')

define(`auto_configuration_go', `GoProg')

auto_configuration_go