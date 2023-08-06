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
package Roku_Backend_Smart_Devices

//Google Chrome Cast API endpoint paths
const (
	ROKU_KEYPRESS_HOME      = "http://%s:8060/keypress/home"
	ROKU_KEYPRESS_PLAY      = "http://%s:8060/keypress/play"
	ROKU_KEYPRESS_DOWN      = "http://%s:8060/keypress/down"
	ROKU_KEYPRESS_UP        = "http://%s:8060/keypress/up"
	ROKU_KEYPRESS_LEFT      = "http://%s:8060/keypress/left"
	ROKU_KEYPRESS_RIGHT     = "http://%s:8060/keypress/right"
	ROKU_KEYPRESS_OK        = "http://%s:8060/keypress/select"
	ROKU_KEYPRESS_REWIND    = "http://%s:8060/keypress/rewind"
	ROKU_KEYPRESS_FFW       = "http://%s:8060/keypress/fastforward"
	ROKU_KEYPRESS_OPTIONS   = "http://%s:8060/keypress/options"
	ROKU_KEYPRESS_PAUSE     = "http://%s:8060/keypress/pause"
	ROKU_KEYPRESS_BACK      = "http://%s:8060/keypress/back"
	ROKU_KEYPRESS_POWEROFF  = "http://%s:8060/keypress/poweroff"
	ROKU_KEYPRESS_VUP       = "http://%s:8060/keypress/volumeup"
	ROKU_KEYPRESS_VDOWN     = "http://%s:8060/keypress/volumedown"
	ROKU_KEYPRESS_MUTE      = "http://%s:8060/keypress/volumemute"
	ROKU_KEYPRESS_DEVDOWN   = "http://%s:8060/keypress/powerOff"
	ROKU_KEYPRESS_DEVUP     = "http://%s:8060/keypress/powerOn"
	ROKU_DEVICE_LAUNCH      = "http://%s:8060/launch/%s"
	ROKU_DEVICE_INSTALL     = "http://%s:8060/install/%s?contentid=%s&MediaType=%s"
	ROKU_DEVICE_DISABLE_SGR = "http://%s:8060/query/sqrendezvous/untrack"
	ROKU_DEVICE_ENABLE_SGR  = "http://%s:8060/query/sqrendezvous/track"
	ROKU_DEVICE_TV          = "http://%s:8060/query/tv-channels"
	ROKU_DEVICE_SGNODE      = "http://%s:8060/query/sgnodes"
	ROKU_DEVICE_ACTIVETV    = "http://%s:8060/query/tv-active-channel"
	ROKU_DEVICE_DIAL        = "http://%s:8060/dial/dd.xml"
	ROKU_DEVICE_BROWSE      = "http://%s:8060/search/browse?keyword=%s"
	ROKU_DEVICE_INFO        = "http://%s:8060/query/device-info"
	ROKU_DEVICE_APPS        = "http://%s:8060/query/apps"
	ROKU_DEVICE_ACTIVE      = "http://%s:8060/query/active-app"
)
