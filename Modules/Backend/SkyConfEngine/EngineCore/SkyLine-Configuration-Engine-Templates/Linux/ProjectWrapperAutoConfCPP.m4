dnl | SLC Programmatic execution and interfaces for templating code
dnl | ----------------------------------------------------------------
dnl | SLC allows you to generate programs and code, one way we can do that
dnl | rather choose to do that is by using autoconf/M4-Sugar. In this case 
dnl | these files are used to generate C++/C wrappers and Golang file templates
dnl | for projects as well as json files for linux based operating systems that 
dnl | may have or use M4 installed on the system. 
dnl | ----------------------------------------------------------------------
dnl | This file is the file that is used to template the cpp file for C/C++ wrappers 
dnl | which are generated by SLC to give users the option of hybrid development
dnl | 

divert(0)

define(`CPP_Prog', `
// cpp_wrapper.cpp
#include <iostream>

extern "C" {
    // C interface functions
    int callCPPFunction();
}

int cppFunction() {
    // C++ logic for functionality
    return 20 - 100;
}

int callCPPFunction() {
    return cppFunction();
}
divert(0)')

define(`auto_configuration_cpp', `CPP_Prog')

auto_configuration_cpp