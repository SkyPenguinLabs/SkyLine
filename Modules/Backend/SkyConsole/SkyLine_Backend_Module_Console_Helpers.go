///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Console_Helpers
// Extension         | .go ( golang source code file )
// Purpose           | Define the helper functions for the REPL
// Directory         | Modules/Backend/SkyConsole
// Modular Directory | SkyLine/Modules/Backend/SkyConsole
// Package Name      | SkyLine_Backend_Module_Console
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
// This is known as the REPL otherwise and commonly known as a Read Eval Print Loop which goes straight to executing the next input and spins up a mock environment.
//
// This is not a necessary step within language development and is more of a feature. REPL helps users easily and quickly test environments and allows them to save workspaces
//
// and environments depending on the language. Languages like R-Script allow you to save the workspace while languages like ruby might not allow you to do so as easily or at all.
//
// SkyLine's REPL will be fully customized and much more modern than people may expect!
//
package SkyLine_Backend_Module_Console

import "io"

func ExecuteParserErrors(line string, Out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(Out, msg)
		io.WriteString(Out, "\n")
	}
}
