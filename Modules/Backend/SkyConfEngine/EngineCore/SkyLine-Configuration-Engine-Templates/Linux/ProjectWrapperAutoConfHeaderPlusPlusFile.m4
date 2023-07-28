dnl | SLC Programmatic execution and interfaces for templating code
dnl | ----------------------------------------------------------------
dnl | SLC allows you to generate programs and code, one way we can do that
dnl | rather choose to do that is by using autoconf/M4-Sugar. In this case 
dnl | these files are used to generate C++/C wrappers and Golang file templates
dnl | for projects as well as json files for linux based operating systems that 
dnl | may have or use M4 installed on the system. 
dnl | ----------------------------------------------------------------------
dnl | This file is for the HPP file when working with C++
divert(0)

define(`CPP_Prog', `
// cpp_wrapper.hpp
#ifndef CPP_WRAPPER_H
#define CPP_WRAPPER_H
#ifdef __cplusplus
extern "C" {
#endif
int callCPPFunction();
#ifdef __cplusplus
}
#endif
#endif  // CPP_WRAPPER_HPP
divert(0)')

define(`auto_configuration_cpp', `CPP_Prog')

auto_configuration_cpp