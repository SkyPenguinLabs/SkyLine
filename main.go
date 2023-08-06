package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	SkyConfEngine "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyConfEngine/EngineCore"
	SkyREPL "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyConsole"
	SkyEnvironment "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	SkyEval "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEvaluator"
	SkyFS "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyFS"
	SkyParser "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyParser"
	SkyScanner "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyScanner"
	SkySettings "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkySettings"
	SkyInvokers "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Invokes"
	SkyStdCall "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/StandardCalls"
)

type ProjectInitSLCGen struct {
	ProjectGenData []struct {
		Projectname      string `json:"Projectname"`
		HybridWrap       bool   `json:"Hybrid-wrap?"`
		UsingSLModifiers bool   `json:"UsingSLModifiers?"`
		UsingJSONData    bool   `json:"UsingJsonData?"`
		UsingEngine      bool   `json:"UsingEngine?"`
		UsingMakefile    bool   `json:"UsingMakefile?"`
		DesignatedPath   string `json:"DesignatedPath"`
	} `json:"ProjectGenData"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
	SkyStdCall.InitateCallRegisterStandardCall()
	SkyInvokers.RegisterInvokes()
}

func Execute(input string) int {
	Environment := SkyEnvironment.SL_NewEnvironment()
	Scan := SkyScanner.New(input)
	Parse := SkyParser.SkyLineNewParser(Scan)
	Prog := Parse.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
	if len(Parse.SkyLine_Parser_Helper_Ret_Errors()) != 0 {
		fmt.Printf("%s ", Parse.SkyLine_Parser_Helper_Ret_Errors()[0])
		fmt.Print("\n")
		os.Exit(0)
	}
	InitateScanner := SkyScanner.New("")
	InitateParser := SkyParser.SkyLineNewParser(InitateScanner)
	InitateProg := InitateParser.SkyLine_Parser_Expressions_And_Statements_ExtraUnit_ProgramaticParse()
	SkyEval.SkyLine_Call_Eval(InitateProg, Environment)
	SkyEval.SkyLine_Call_Eval(Prog, Environment)
	return 0
}

func main() {
	if *SkySettings.SkySettings.SkyLine_Run_SLC && *SkySettings.SkySettings.SkyLine_EngineFile != "" {
		fmt.Println(SkyREPL.LINUX_CLS)
		SkyConfEngine.OutputBanner()
		SkyConfEngine.OutputBoxOCode(*SkySettings.SkySettings.SkyLine_EngineFile)
		SkyConfEngine.StartEngine_RenderFile(*SkySettings.SkySettings.SkyLine_EngineFile, true)
		os.Exit(0)
	}
	if *SkySettings.SkySettings.SkyLine_Help {
		fmt.Printf(SkySettings.HelpMenu, SkySettings.SkySession.Go_Version, SkySettings.SkySession.OS_Name, SkySettings.SkySession.OS_Arch)
		os.Exit(0)
	} else if *SkySettings.SkySettings.SkyLine_Console {
		SkyREPL.Start(os.Stdin, os.Stdout, false)
	} else if *SkySettings.SkySettings.SkyLine_Tooling_SLC_AutoGen && *SkySettings.SkySettings.SkyLine_Tooling_SLC_AutoGenSrc != "" {
		jsonData, err := ioutil.ReadFile(*SkySettings.SkySettings.SkyLine_Tooling_SLC_AutoGenSrc)
		if err != nil {
			log.Fatal(err)
		}
		var data ProjectInitSLCGen
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Fatal(err)
		}
		SkyConfEngine.SetupCall(data.ProjectGenData[len(data.ProjectGenData)-1].DesignatedPath)
		for i := 0; i < len(SkyConfEngine.AllTemplates); i++ {
			SkyConfEngine.GenerateContentsOfFile(SkyConfEngine.AllTemplates[i], data.ProjectGenData[len(data.ProjectGenData)-1].DesignatedPath)
		}
	} else if *SkySettings.SkySettings.SkyLine_ReplExecEnd {
		fmt.Println(SkyREPL.LINUX_CLS)
		SkyREPL.SkyLine_Console_Banner()
		SkyREPL.SkyLine_CodeBoxWithNumFree(os.Args[2])
		input, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		Execute(string(input))
	} else {
		input, err := ioutil.ReadFile(os.Args[1])
		SkyFS.Current.FileSystem_Modify_Name(os.Args[1])
		SkyFS.Current.FileSystem_ModifyDir()
		SkyFS.Current.FileSystem_ModifyExtension()
		SkyFS.Current.FileSystem_ModifyInfo()
		if err != nil {
			log.Fatal(err)
		}
		Execute(string(input))
	}
}
