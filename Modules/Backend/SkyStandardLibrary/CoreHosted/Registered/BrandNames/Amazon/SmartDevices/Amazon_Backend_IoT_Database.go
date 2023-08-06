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

package Amazon_Smart_Services

//Amazon FireTV API endpoint paths
const (
	AMAZON_FIRE_TV_DEVICE_INFORMATION = "http://%s:53917/zc?action=getInfo&version=2.7.1"
	AMAZON_FIRE_TV_DEVICE_DESCRIPTION = "http://%s:60000/upnp/dev/%s/desc"
)
