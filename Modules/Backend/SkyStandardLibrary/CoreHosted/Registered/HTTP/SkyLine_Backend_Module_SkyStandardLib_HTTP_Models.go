//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ _____ _____     _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |_   _|_   _|  _  |___| __  |___ ___ _ _ ___ ___| |_ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|     | | |   | | |   __|___|    -| -_| . | | | -_|_ -|  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__| |_|   |_| |__|      |__|__|___|_  |___|___|___|_| |___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|               |_|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> Like all model files, this file holds all models and data types and constants
//
//
package SkyLine_Standard_Library_HTTP

import (
	"io"
	"net/http"
	"net/url"
)

const (
	HTTPMETHOD_GET     = http.MethodGet
	HTTPMETHOD_HEAD    = http.MethodHead
	HTTPMETHOD_POST    = http.MethodPost
	HTTPMETHOD_DELETE  = http.MethodDelete
	HTTPMETHOD_PUTS    = http.MethodPut
	HTTPMETHOD_PATCH   = http.MethodPatch
	HTTPMETHOD_CONNECT = http.MethodConnect
	HTTPMETHOD_TRACE   = http.MethodTrace
	HTTPMETHOD_OPTIONS = http.MethodOptions
)

type (
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: CustomTransport is a method for making			  ::
	//:: your own custom transport and HTTP options       ::
	//:: which is typically used for new request builders ::
	//::::::::::::::::::::::::::::::::::::::::::::::::::::::
	CustomTransport struct {
		Transport         *http.Transport
		SkipCertificate   bool
		ProxyURL          *url.URL
		ClientURL         string
		Method            string
		ResponseReaderBod io.Reader
	}

	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	//:: HTTPStructure is a structure for easy manipulation  ::
	//:: of HTTP requests, this allows users to change       ::
	//:: the way http requests are done                      ::
	//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	HTTPStructure struct {
		Req_URL    string // Request URL
		Req_Method string // Request method
		Filename   string // If user wants output then this is the filename
		OutToFile  bool   // Output response body to file
		TorProx    string // SOCKS proxy to use
		Headers    []string
		Tor        bool // Use tor

	}
)

var (
	Httpstruct HTTPStructure
)
