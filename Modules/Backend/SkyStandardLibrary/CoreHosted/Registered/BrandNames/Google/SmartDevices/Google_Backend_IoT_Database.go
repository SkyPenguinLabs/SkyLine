//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____     _                   _         _____ ___       _____ _   _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|     |___| |_ ___ ___ ___ ___| |_      |     |  _|     |_   _| |_|_|___ ___ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|-   -|   |  _| -_|  _|   | -_|  _|     |  |  |  _|       | | |   | |   | . |_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____|_|_|_| |___|_| |_|_|___|_|  _____|_____|_|  _____  |_| |_|_|_|_|_|_  |___|
//	primary units that define the functions to register the sub func's  //            |___|                                                     |_____|         |_____|               |___|
//////////////////////////////////////////////////////////////////////////
//
//
// Defines -> The IoT library within the SkyLine programming language is more or less defined for IoT devices and is an experimental library. This library contains informaiton
//
// such as endpoints, parsers, expected codes, headers and expected responses from servers and much more for IoT devices such as Samsung smart devices, google home's, FireSticks,
//
// home routers, AppleTV's, RokuTV's and other various IoT devices. This mainly works on the API's and focuses on information gathering from them or utilizing them in specific ways
//
// as these API's are either open or were reverse engineered. Make sure when using this library YOU KNOW WHAT YOU ARE DOING. SkyLine is still an experimental programming language
//
// and it has MANY bugs in it. Right now, this library may not be the best thing to use if you do not understand it.
//
package Google_Smart_Devcies

//Google Chrome Cast API endpoint paths
const (
	GOOGLE_CAST_DEVICE_INFORMATION             = "http://%s:%s/setup/eureka_info?options=detail"
	GOOGLE_CAST_DEVICE_WIFI_SCAN               = "https://%s:%s/setup/scan_wifi"
	GOOGLE_CAST_DEVICE_WIFI_SCAN_RESULTS       = "https://%s:%s/setup/scan_results"
	GOOGLE_CAST_DEVICE_WIFI_FORGET             = "https://%s:%s/setup/forget_wifi"
	GOOGLE_CAST_DEVICE_CONFIGURED_NETWORK      = "https://%s:%s/setup/configured_networks"
	GOOGLE_CAST_DEVICE_APPLICATION_URL         = "http://%s:%s/apps/%s"
	GOOGLE_CAST_DEVICE_REBOOT                  = "http://%s:%s/setup/reboot"
	GOOGLE_CAST_DEVICE_DEVICE_DESCRIPTION      = "http://%s:%s/ssdp/device-desc.xml"
	GOOGLE_CAST_DEVICE_DEVICE_NAME             = "https://%s:%s/setup/set_eureka_info"
	GOOGLE_CAST_DEVICE_DEVICE_TIMEZONES        = "http://%s:%s/setup/supported_timezones"
	GOOGLE_CAST_DEVICE_DEVICE_ALARMS           = "http://%s:%s/setup/assistant/alarms"
	GOOGLE_CAST_DEVICE_DEVICE_LEGACYCONFIG     = "https://www.gstatic.com/eureka/config/legacy/config.json"
	GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_STAT   = "http://%s:%s/setup/bluetooth/status"
	GOOGLE_CAST_DEVICE_DEVICE_BLUETOOTH_PAIRED = "http://%s:%s/setup/bluetooth/get_bonded"
)
