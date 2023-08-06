////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
//
// The SkyLine configuration language is a language and engine designed to act as a modification language to the SkyLine programming language. This language is
// very minimal and contains a regex base lexer, a very basic parser, a few keywords, a base interpreter and that is all as well as some backend engine code. This
// language is purely modified to be an extension to the SkyLine programming language, something that can be a pre processor language post processing for the main
// SkyLine script. Below is more technical information for the language
//
// Lexer       : Regex based lexer with minimal tokens and base syntax
// Parser      : Base parser with minimal tokens and base syntax with simple error systems
// REPL        : Does not exist
// Environment : Extremely minimal
// Types       : String, Boolean, Integer
// Statements  : set, import, use, errors, output, system, constant/const
//
//
//
// File    ->  SkyLine_Configuration_Language_Frontend_ArtWork.go
// Apart   ->  SLC/Modules/Backend
// Source  -> .go, .mod, .SL-Modify
//
// File contains -> This file is dedicated to the entire configuration language's core library. The configuration language for SkyLine mainly runs off of standard functions
//                  or rather standard library based functions which include functions to verify systems, load json files and much more amongst that. So to do that and to
//                  allow for specific functions we will be registering all of them into the environment at once using a init() function. This file will hold all proper types
//				    as well as all proper functions, init calls, descriptions, results, parsers and more. This may be one of the longest files in this entire development directory
//
// WARN: Each section is split up into its own little part and function.
//
package SkyLine_Configuration_Engine_Backend_Source

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//			         										   ┏━━━━━
//                  										  ┃┃┏━┓
//                										      ┃┃┃ ┃
//           										      ━━━━┛┃
//	              											   ┗━━━━
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// Title -> OX[M:1]
//
// This section contains all of the type structures, constants, variables, maps and other forms of data that the library may actually require. For context, the engine will allow
//
// users to load requirements.json files and Program.json files which all contain various data. In order to do this though we need to pre define json structures. These structures
//
// are going to be held within this section of the engine and will be exported if held capital.
//

type Load_Requirements struct {
	Requirements struct {
		Libraries       []string `json:"Libraries"`
		OperatingSystem string   `json:"Operating-System"`
		SLCVersion      string   `json:"SLC-Version"`
		SLVersion       string   `json:"SL-Version"`
	} `json:"Requirements"`
}

/*

Load project data will all be data and variables and constants that will be pre loaded into the environment of the application when parsed.

When working with the engine in skyline, there will be a keyword which works similar to import but is rather ENGINE. When this keyword is found

and looks like ENGINE("main.slc") the parser will load and execute the engine, if the engine comes across a Project.json file within the respected

and expected directories, it will load the below structure which can then be accessed by SL external to SLC. Which means instead of setting variables

within the engine;s environment, you can register it within a projects environment.

*/
type Load_ProjectData struct {
	ProjectInformation struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
		SupportedOS string `json:"Supported-OS"`
		Languages   string `json:"Languages"`
		Version     string `json:"Version"`
	} `json:"Project-Information"`
}

// Exportable is data that can be read externally to SLC
var Exportable_data Exportable

type Exportable struct {
	ProjectData struct {
		Name        string   // Name of the project
		Description string   // Description
		SuportedOS  string   // Supported operating system
		Languages   string   // Language of the program
		Version     string   // Version of the program
		Require     []string // Libraries to register
	}
}

/*

The systems map allows us to define how much arguments the right side of the modifier '->' can have.

In the SkyLine Configuration Language, a modifier otherwise known as a passer is defined by '->'.

The idea of a modifier is to take the left value of type STRING and then the right side of type ARRAY

and modify the system respectively or in this case store the results so SL can access it before running

the program and all of the data properly. The right side of the array will be how much you can modify.

For example, say we want to modify the error system, well, there are three things you can modify in the error system

output  of type string
color of type boolean
warnings of type boolean

warnings | defines whether or not the user wants to see warnings or output warnings ( true=on or false=off)
color    | defines whether or not the user wants to turn on color ( true ) or turn it off ( false )
output   | defines the error systems form of output, should it be basic, verbose, or tree or masters.




when we want to parse these values, given the array on the right side of the arrow is a array of arguments,

we need to be able to monitor how much arguments are being passed. if the array in the case of

system("warnings") -> [];

does not have any arguments, the engine will output a warning, if the array has arguments but not enough the engine will assume values
and if the array has too much arguments according to the system the engine will stop parsing and call to kill the process of both skyline
and the engine sending syscall's too its process.

The purpose of this map allows us to define the limit of how much arguments are allows per system.

*/
var SystemsArguments = map[string]int{
	"error_system":       3, // color=bool (1), Output_Format=integer (2), verbose=integer (3)
	"import_system":      1, // import_with=string
	"strict_system":      1, // strict_warnings=bool
	"parser_code_system": 2, // code=integer, message=string
	"output_system":      2, // color=bool, verbose=bool,
}

/*

This MAP lets us specify the systems allowed to be modified currently

*/

var Systems_Modifable = map[string]string{
	"errors": "error_system",
	"parser": "parser_code_system",
	"output": "output_system",
	"import": "import_system",
}

/*

This map shows all active lists and systems that work with the libraries or all active skyline standard libraries anyway

*/

var LibrariesMap = map[string]bool{
	"math":                true,
	"io":                  true,
	"File":                true,
	"http":                true,
	"Google/Smart_API_DB": true,
	"Amazon/Smart_API_DB": true,
	"Roku/Smart_API_DB":   true,
	"Apple/Smart_API_DB":  true,
	"Other/Smart_API_DB":  true,
	"xml":                 true,
	"json":                true,
	"env":                 true,
	"console":             true,
	"forensics/image":     true,
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//			         										   ┏━━━━━
//                  										  ┃┃┏━┓
//                										      ┃┃┃ ┃
//           										      ━━━━┛┃
//	              											   ┗━━━━
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// Title -> OX[M:2]
//
// This section defines helper functions which are dedicated to helping out the calls such as grabbing the length of the arguments, grabbing the data type,
//
// Checking and verifying the data type and much more among that list
//

func EngineGetArgLength(args ...SLC_Object) int {
	return len(args)
}

func EngineGetDataType(args SLC_Object) string {
	var typeexp string
	switch args.(type) {
	case *ObjectString:
		typeexp = "string"
	case *ObjectArray:
		typeexp = "array"
	case *ObjectInteger:
		typeexp = "integer"
	case *ObjectBoolean:
		typeexp = "bool"
	default:
		// Engine error
		typeexp = "E_ERR"
	}
	return typeexp
}

// this tekks the engine if the argument count is 0 or NULL
func EngineIsZero(arguments ...SLC_Object) bool {
	return len(arguments) == 0
}

// This tells the engine if the argument count is too much
// in the case that a function is called with too many arguments
// this function can be called to check the length and if it goes
// over or is greater than the max number of arguments allowed
func EngineMeetsArgLimit(max int, arguments ...SLC_Object) bool {
	return len(arguments) > max
}

// Check sif the argument list is below the minimum requirement
// if a function is called with () and the minimum requirement is 1
// argument, then this function can be used to check that.
func EngineBelowArgLimit(minimum int, arguments ...SLC_Object) bool {
	return len(arguments) < minimum
}

// Grabs the actual system to modify
func EngineGrabSystemToModify(system string) string {
	return Systems_Modifable[strings.ToLower(system)]
}

func EngineCheckNil(data interface{}) bool {
	return data == nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//			         										   ┏━━━━━
//                  										  ┃┃┏━┓
//                										      ┃┃┃ ┃
//           										      ━━━━┛┃
//	              											   ┗━━━━
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// Title -> OX[M:3]
//
// This section defines all of the libraries main functions. The purpose of these functions is to aid in the result of the arrays. When working with SLC, in order to modify data you
//
// must use the -> operator which is defined as a modifier. A modifier requires two types on the left and right side, left is of type string returned by a function call known as system
//
// and the left is of type Array which holds the arguments and settings to modify the system. When working with modifiers, you can use base operations like this
//
// "error_system" -> [10, 1, false]
//
// However, it is suggested by the developers to use the standard library functions because these functions verify and rather prepare the output of the system calls. For example
//
// instead of having to type error_system each time you can do the following to automate this process
//
//`````CODE UNIT`
/*


ENGINE {
   INIT {
	load("Requirements.json")
    constant DEFINE_CODE_MISSING_SEMICOLON = 12;
    constant DEFINE_CODE_MISSING_LEFT_BRCE = 109;

    system("errors") -> modify_sys[basic(true), verbosity(true), depth(0)];
    system("output") -> modify_sys[debug(debuglev)];
    system("import") -> modify_sys[expect("directories")];
    system("parser") -> modify_parser[DEFINE_CODE_MISSING_SEMICOLON, "Missing semicolon in statement"]
  }
};

*/
// This is because, the standard function calls such as system() allow the engine to verify that the system is a real system that you are trying to modify and will also
// check the argument detail. If the system is not existent then the arguments and modification call will be immediately disabled and a warning will appear. Using the system
// call as well as expect(), debug(), basic(), verbosity(), depth() etc function calls allow the engine to properly decode and transfer and convert the data types of the info.
// mis use of this will cause the engine to have an internal fault and fix or rather self heal that fault by ignoring any and all values in that array then continuing onto the next
// major conflict. It is highly suggested to USE the function calls rather than leaving them out

func S_Engine_Call_SYSTEM(arguments ...SLC_Object) SLC_Object {
	if !EngineIsZero(arguments...) {
		switch arguments[0].(type) {
		case *ObjectString:
			return &ObjectString{Value: Systems_Modifable[arguments[0].ObjectInspectFunc()]}
		}
		// Check system
	}
	return &ObjectNULL{}
}

func S_Engine_Call_Load(arguments ...SLC_Object) SLC_Object {
	if !EngineIsZero(arguments...) {
		switch arguments[0].(type) {
		case *ObjectString:
			// run and parse files
			file := arguments[0].ObjectInspectFunc()
			if IsRenderForCmdDbg {
				fmt.Printf("\033[38;5;57m[\033[38;5;50m+\033[38;5;57m] \033[38;5;249m Verified File        | \033[38;5;210m(<%s>) \033[38;5;249m\n", file)
			}
			return S_Engine_Call_Loader_CheckerAndVerifier(file)
		}
	}
	return &ObjectNULL{}
}

func S_Engine_Call_Loader_CheckerAndVerifier(filename string) SLC_Object {
	f, x := os.Open(filename)
	if x != nil {
		return &ObjectERROR{Message: "Engine Error: Could not open the file due to -> " + fmt.Sprint(x)}
	}
	defer f.Close()
	var Settings Load_Requirements
	decoded := json.NewDecoder(f)
	if x = decoded.Decode(&Settings); x != nil {
		return &ObjectERROR{Message: "Engine Error: Could not process or decode the json data due to -> " + fmt.Sprint(x)}
	}
	// run through all the libraries and check if they exist
	var ActivateList []string
	for i := 0; i < len(Settings.Requirements.Libraries); i++ {
		if LibrariesMap[Settings.Requirements.Libraries[i]] {
			ActivateList = append(ActivateList, Settings.Requirements.Libraries[i])
			if IsRenderForCmdDbg {
				fmt.Printf("\033[38;5;57m[\033[38;5;50m+\033[38;5;57m] \033[38;5;249mVerified Library      | \033[38;5;210m(<%s>) \033[38;5;249m\n", Settings.Requirements.Libraries[i])
			}
		} else {
			fmt.Println("Engine: FAIL -> ", Settings.Requirements.Libraries[i])
			os.Exit(0)
		}
	}
	if ActivateList != nil {
		Exportable_data.ProjectData.Require = ActivateList
	}
	if !strings.EqualFold(Settings.Requirements.OperatingSystem, runtime.GOOS) {
		return &ObjectERROR{Message: "Engine Verification: Sorry, operating system is not supported. You are on " + runtime.GOOS + " but the project requires " + Settings.Requirements.OperatingSystem}
	}
	// Now parse or check for exportable data.
	Path := "Backend/Conf/ProjectData/Project.json"
	f, x = os.Open(Path)
	if x != nil {
		return &ObjectERROR{
			Message: "Engine Error: Could not open the project directory of ( " + Path + " ) -> " + fmt.Sprint(x),
		}
	} else {
		if IsRenderForCmdDbg {
			fmt.Printf("\033[38;5;57m[\033[38;5;50m+\033[38;5;57m] \033[38;5;249mVerified Project Info | \033[38;5;210m(<%s>) \033[38;5;249m\n", Path)
		}
	}
	defer f.Close()
	var ProjectLoad Load_ProjectData
	decoder := json.NewDecoder(f)
	if x = decoder.Decode(&ProjectLoad); x != nil {
		return &ObjectERROR{
			Message: "Engine Error: Sorry, could not verify or decode the json data in this file due to -> " + fmt.Sprint(x),
		}
	}
	if !EngineCheckNil(ProjectLoad.ProjectInformation.Description) {
		Exportable_data.ProjectData.Description = ProjectLoad.ProjectInformation.Description
	}
	if !EngineCheckNil(ProjectLoad.ProjectInformation.Name) {
		Exportable_data.ProjectData.Name = ProjectLoad.ProjectInformation.Name
	}
	if !EngineCheckNil(ProjectLoad.ProjectInformation.SupportedOS) {
		Exportable_data.ProjectData.SuportedOS = ProjectLoad.ProjectInformation.SupportedOS
	}
	if !EngineCheckNil(ProjectLoad.ProjectInformation.Version) {
		Exportable_data.ProjectData.Version = ProjectLoad.ProjectInformation.Version
	}
	if !EngineCheckNil(ProjectLoad.ProjectInformation.Languages) {
		Exportable_data.ProjectData.Languages = ProjectLoad.ProjectInformation.Languages
	}
	// In later versions check the versions of SLC and Sl required to run the project

	return &ObjectNULL{}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//			         										   ┏━━━━━
//                  										  ┃┃┏━┓
//                										      ┃┃┃ ┃
//           										      ━━━━┛┃
//	              											   ┗━━━━
//  													 _____ _       __    _
// 														|   __| |_ _ _|  |  |_|___ ___
// 													    |__   | '_| | |  |__| |   | -_|
// 														|_____|_,_|_  |_____|_|_|_|___|
//		   														  |___|
//
// Title -> OX[M:4]
//
// This section is super light as it defines the init function which will be called to register all of the standard libraries functions
//

func init() {
	RegisterBuiltin(
		"system",
		S_Engine_Call_SYSTEM,
	)
	RegisterBuiltin(
		"load",
		S_Engine_Call_Load,
	)
}
