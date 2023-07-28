dnl | SLC Programmatic execution and interfaces for templating code
dnl | ----------------------------------------------------------------
dnl | SLC allows you to generate programs and code, one way we can do that
dnl | rather choose to do that is by using autoconf/M4-Sugar. In this case 
dnl | these files are used to generate C++/C wrappers and Golang file templates
dnl | for projects as well as json files for linux based operating systems that 
dnl | may have or use M4 installed on the system. 
dnl | ----------------------------------------------------------------------
dnl | This file will help SLC generate the main script for development and 
dnl | importing module files from the projects directory to auto set everything up
dnl |  

divert(0)

define(`SLM_Prog', `ENGINE("Backend/Conf/EngineFiles/Engine.slmod");


set Module1 := import("Backend/Core/Scripts/Module.sl");


define Main() {
    // Call main imported code
    Module1.HelloWorld();
};
')

define(`auto_configuration_SL_main', `SLM_Prog')

auto_configuration_SL_main