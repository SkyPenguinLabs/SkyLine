///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
//
//
// This part of the error system will define updaters for the error systems environment to hold and check for environment variables
//
package SkyLine_Error_System

func (Updater *ConstantEnvironment) Update_LibraryFunctions(x []string, y []string) {
	Updater.LibraryIdentifiers = x
	Updater.LibraryModules = y
}
