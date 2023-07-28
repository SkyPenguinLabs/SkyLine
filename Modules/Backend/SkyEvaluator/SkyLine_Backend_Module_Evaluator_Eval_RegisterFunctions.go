///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//
//																 ┏━┓
//																┃┃ ┃
//															━━━━┛
//
// Filename          | SkyLine_Backend_Module_Evaluator_Eval_RegisterFunctions
// Extension         | .go ( golang source code file )
// Purpose           | Defines all of the registry functions that can be used to register library based functions that are all standard
// Directory         | Modules/Backend/SkyEvaluator
// Modular Directory | SkyLine/Modules/Backend/SkyEvaluator
// Package Name      | SkyLine_Backend_Module_Evaluation
//
//
//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: Module Description / Learners Activity :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//
//
// The final part in standard interpretation inside of a programming language is to evaluate and execute the data or keys. In the case of SkyLine, it relies on different
//
// forms of engines which can use the byte code compiler or use the evaluator. The byte code compiler is a whole different story but the evaluator will take advantage of the
//
// AST and then check and execute conditions, statements, values, or modifications accordingly. The evaluator can also sometimes be complex to use but it still manages to stay
//
// one of the fastest ones to write.
//
package SkyLine_Backend_Evaluation

import (
	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	SkyBrands_DB_Amazon "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/BrandNames/Amazon/SmartDevices"
	SkyBrands_DB_Apple "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/BrandNames/Apple/SmartDevices"
	SkyBrands_DB_Google "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/BrandNames/Google/SmartDevices"
	SkyBrands_DB_Other "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/BrandNames/Other/SmartDevices"
	SkyBrands_DB_Roku "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/BrandNames/Roku/SmartDevices"
	SkyConsole "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/Console"
	SkyEnviron "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/Environment"
	SkyFile "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/File"
	SkyForensicsImage "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/Forensics/Image"
	SkyHTTP "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/HTTP"
	SkyIO "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/IO"
	SkyJSON "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/JSON"
	SkyMath "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/Mathematics"
	SkyXML "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/Registered/XML"
)

func RegisterMath() {
	// math.rand based functions
	SkyLine_Register_Builtin("math.Rand_Chars", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Randomization_Chars(InvokeArgs...))
	})
	// math.theory based functions
	SkyLine_Register_Builtin("math.Theory_isprime", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_NumberTheory_IsPrime(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Theory_prime", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_NumberTheory_Prime(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Theory_gcd", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_NumberTheory_GCD(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Theory_factor", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_NumberTheory_Factor(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Theory_abs", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_NumberTheory_Abs(InvokeArgs...))
	})
	// math.trig based functions
	SkyLine_Register_Builtin("math.Tan", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Trig_Tan(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Cos", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Trig_Cos(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Sin", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Trig_Sin(InvokeArgs...))
	})
	// math.geo based functions
	SkyLine_Register_Builtin("math.Cbrt", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Geoemtry_CBRT(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Sqrt", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Geometry_SQRT(InvokeArgs...))
	})
	// math.modular based functions
	SkyLine_Register_Builtin("math.Modular_mod", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Modular_Arith_Mod(InvokeArgs...))
	})
	SkyLine_Register_Builtin("math.Modular_modexp", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Modular_Arith_Mod_Exp(InvokeArgs...))
	})
	// math.std based functions
	SkyLine_Register_Builtin("math.Pow", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyMath.Math_Generic_PowerOf(InvokeArgs...))
	})
}

func IOPUT() SkyEnv.SL_Object {
	return &SkyEnv.SL_String{Value: "Hello world"}
}

func RegisterIO() {
	SkyLine_Register_Builtin("io.clear", func(env *SkyEnv.SkyLineEnvironment, args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyIO.IO_Clear())
	})
	SkyLine_Register_Builtin("io.put", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (IOPUT())
	})
	SkyLine_Register_Builtin("io.restore", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyIO.IO_ReturnColor())
	})
	SkyLine_Register_Builtin("io.input", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyIO.IO_Input(InvokeArgs...))
	})
}

func RegisterFile2() {
	SkyLine_Register_Builtin("File.Carve", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_Carve(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.New", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_IniateNewFunction(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.Open", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_OpenAndOutFile(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.Write", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_WriteFile(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.Overwrite", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_OverWrite_WriteFile(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.Head", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_GetFileHeader(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.Mime", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_GetFileHeader(InvokeArgs...))
	})
	SkyLine_Register_Builtin("File.Extractlables", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyFile.FileLib_ExtractLables_FileOut(InvokeArgs...))
	})
}

func RegisterHTTP() {
	SkyLine_Register_Builtin("http.Get", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.MakeGet(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.NullPost", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.MakeNullPOST(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.Post", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.MakeBasePost(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.TransportNew", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.NewTransport(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.TransportNewReader", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.Creator_CreateNewBody(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.TransportReq", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.MakeTransPortRequest(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.New", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.Creator_HttpNew(InvokeArgs...))
	})
	SkyLine_Register_Builtin("http.Req", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyHTTP.MakeRequest())
	})
	RegisterVariable("http.MethodDelete", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_DELETE})
	RegisterVariable("http.MethodPatch", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_PATCH})
	RegisterVariable("http.MethodPuts", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_PUTS})
	RegisterVariable("http.MethodPost", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_POST})
	RegisterVariable("http.MethodGet", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_GET})
	RegisterVariable("http.MethodHead", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_HEAD})
	RegisterVariable("http.MethodOptions", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_OPTIONS})
	RegisterVariable("http.MethodTrace", &SkyEnv.SL_String{Value: SkyHTTP.HTTPMETHOD_TRACE})
}

func RegisterAppleIoTDatabase() {
	RegisterVariable("Apple.AirPlayServerInfo", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_AIRPLAY_SERVERINFO})
	RegisterVariable("Apple.AirPlayPlayBackInfo", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_AIRPLAY_PLAYBACKINFO})
	RegisterVariable("Apple.AirPlayScrubInfo", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_AIRPLAY_SCRUB})
	RegisterVariable("Apple.AirPlayStreamInfo", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_AIRPLAY_STREAMINFO})
	RegisterVariable("Apple.AirPlayInfo", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_AIRPLAY_INFO})
	RegisterVariable("Apple.DAAPMain", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_DAAP_PATH})
	RegisterVariable("Apple.DAAPLogin", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_DAAP_LOGIN})
	RegisterVariable("Apple.DAAPDatabase", &SkyEnv.SL_String{Value: SkyBrands_DB_Apple.APPLE_DAAP_DATABASE})
	//SkyBrands_DB_Roku .
}
func RegisterAmazonIoTDatabase() {
	RegisterVariable("Amazon.TvDeviceInformation", &SkyEnv.SL_String{Value: SkyBrands_DB_Amazon.AMAZON_FIRE_TV_DEVICE_INFORMATION})
	RegisterVariable("Amazon.TvDeviceDescription", &SkyEnv.SL_String{Value: SkyBrands_DB_Amazon.AMAZON_FIRE_TV_DEVICE_DESCRIPTION})
}

func RegisterRokuIoTDatabase() {
	RegisterVariable("Roku.KeyPressHome", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_HOME})
	RegisterVariable("Roku.KeyPressPlay", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_PLAY})
	RegisterVariable("Roku.KeyPressUp", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_UP})
	RegisterVariable("Roku.KeyPressDown", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_DOWN})
	RegisterVariable("Roku.KeyPressLeft", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_LEFT})
	RegisterVariable("Roku.KeyPressRight", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_RIGHT})
	RegisterVariable("Roku.KeyPressSelect", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_OK})
	RegisterVariable("Roku.KeyPressRewind", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_REWIND})
	RegisterVariable("Roku.KeyPressFFW", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_FFW})
	RegisterVariable("Roku.KeyPressOptions", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_OPTIONS})
	RegisterVariable("Roku.KeyPressPause", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_PAUSE})
	RegisterVariable("Roku.KeyPressBack", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_BACK})
	RegisterVariable("Roku.KeyPressPoweroff", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_POWEROFF})
	RegisterVariable("Roku.KeyPressVolumeUp", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_VUP})
	RegisterVariable("Roku.KeyPressVolumeDown", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_VDOWN})
	RegisterVariable("Roku.KeyPressVolumeMute", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_MUTE})
	RegisterVariable("Roku.DeviceDown", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_DEVDOWN})
	RegisterVariable("Roku.DeviceUp", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_KEYPRESS_DEVUP})
	RegisterVariable("Roku.DevAppLaunch", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_LAUNCH})
	RegisterVariable("Roku.DevAppInstall", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_INSTALL})
	RegisterVariable("Roku.DevDisableSGR", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_DISABLE_SGR})
	RegisterVariable("Roku.DevEnableSGR", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_ENABLE_SGR})
	RegisterVariable("Roku.DevTV", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_TV})
	RegisterVariable("Roku.DevSGNODES", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_SGNODE})
	RegisterVariable("Roku.DevActiveTVS", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_ACTIVETV})
	RegisterVariable("Roku.DevDial", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_DIAL})
	RegisterVariable("Roku.DeviceBrowse", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_BROWSE})
	RegisterVariable("Roku.DeviceInformation", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_INFO})
	RegisterVariable("Roku.DeviceApplications", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_APPS})
	RegisterVariable("Roku.DeviceActiveApplications", &SkyEnv.SL_String{Value: SkyBrands_DB_Roku.ROKU_DEVICE_ACTIVE})
}

func RegisterGoogleIoTDatabase() {
	RegisterVariable("Google.CastDeviceInfo", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_INFORMATION})
	RegisterVariable("Google.CastDeviceReboot", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_REBOOT})
	RegisterVariable("Google.CastDeviceDescription", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_DESCRIPTION})
	RegisterVariable("Google.CastDeviceWiFiForget", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_WIFI_FORGET})
	RegisterVariable("Google.CastDeviceWiFiScan", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_WIFI_SCAN})
	RegisterVariable("Google.CastDeviceWiFiScanResults", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_WIFI_SCAN_RESULTS})
	RegisterVariable("Google.CastDeviceWiFiConfigured", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_CONFIGURED_NETWORK})
	RegisterVariable("Google.CastDeviceAlarms", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_ALARMS})
	RegisterVariable("Google.CastDeviceTimezones", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_TIMEZONES})
	RegisterVariable("Google.CastDeviceLegacyConf", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_LEGACYCONFIG})
	RegisterVariable("Google.CastDeviceBleStat", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_STAT})
	RegisterVariable("Google.CastDeviceBlePaired", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_PAIRED})
	RegisterVariable("Google.CastDeviceSetName", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_DEVICE_NAME})
	RegisterVariable("Google.CastDeviceApplications", &SkyEnv.SL_String{Value: SkyBrands_DB_Google.GOOGLE_CAST_DEVICE_APPLICATION_URL})
}

func RegisterOtherIoTDatabases_ARRIS() {
	RegisterVariable("Arris.WANIP_SCPD", &SkyEnv.SL_String{Value: SkyBrands_DB_Other.Arris_WANIP_SCPD_XML_FILE})
	RegisterVariable("Arris.WANCI_SCPD", &SkyEnv.SL_String{Value: SkyBrands_DB_Other.Arris_WANCI_SCPD_XML_FILE})
	RegisterVariable("Arris.IDG_Description", &SkyEnv.SL_String{Value: SkyBrands_DB_Other.Arris_URL_IDG_XML_DESCRIPTION_FILE})
	RegisterVariable("Arris.L3_SCPD", &SkyEnv.SL_String{Value: SkyBrands_DB_Other.Arris_L3_Forwarding_SCPD_XML_Configuration})
}

func RegisterXML() {
	SkyLine_Register_Builtin("xml.Parse", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyXML.XML_CallToParse(InvokeArgs...))
	})
	SkyLine_Register_Builtin("xml.FromJson", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyXML.XML_CallToConveryJsonToXML(InvokeArgs...))
	})
	SkyLine_Register_Builtin("xml.ToJson", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return SkyXML.XML_CallToConvertXMLtoJson(InvokeArgs...)
	})
}

func RegisterJson() {
	SkyLine_Register_Builtin("json.ToGo", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyJSON.IniateCallToStructure(false, InvokeArgs...))
	})
	SkyLine_Register_Builtin("json.Parse", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyJSON.IniateMapToHash(InvokeArgs...))
	})
}

func RegisterEnv() {
	SkyLine_Register_Builtin("env.Setenv", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyEnviron.Environment_SetEnvironment(InvokeArgs...))
	})
	SkyLine_Register_Builtin("env.Getenv", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyEnviron.Environment_GetEnvironmentPath(InvokeArgs...))
	})
	SkyLine_Register_Builtin("env.Environment", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyEnviron.Environment_GetEnvironment(InvokeArgs...))
	})
}

func RegisterConsole() {
	SkyLine_Register_Builtin("console.HtmlToAnsi", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyConsole.Console_Frontend_GenerateAnsiEscapeSequence(InvokeArgs...))
	})
	SkyLine_Register_Builtin("console.TableNew", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyConsole.Console_Frontend_TableNewCreation(InvokeArgs...))
	})
	SkyLine_Register_Builtin("console.Table", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyConsole.Console_Frontend_TableCreation(InvokeArgs...))
	})
}

func RegisterForensicsImagePath() {
	SkyLine_Register_Builtin("image.ChunkCount", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallGetNumChunks(InvokeArgs...))
	})
	SkyLine_Register_Builtin("image.Create", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallCreateImage(InvokeArgs...))
	})
	SkyLine_Register_Builtin("image.LoadOffsets", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallGetOffsets(InvokeArgs...))
	})
	SkyLine_Register_Builtin("image.ChunkType", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallGGetChunkName(InvokeArgs...))
	})
	SkyLine_Register_Builtin("image.DumpData", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallGetMetaChunkData(InvokeArgs...))
	})
	SkyLine_Register_Builtin("image.FindArchive", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallLocateArchive(InvokeArgs...))
	})
	SkyLine_Register_Builtin("image.InjectFiles", func(Environ *SkyEnv.SkyLineEnvironment, InvokeArgs ...SkyEnv.SL_Object) SkyEnv.SL_Object {
		return (SkyForensicsImage.SkyLine_Forensics_Image_CallInjectImage(InvokeArgs...))
	})
}
