//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ _____ _____     _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |_   _|_   _|  _  |___| __  |___ ___ _ _ ___ ___| |_ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|     | | |   | | |   __|___|    -| -_| . | | | -_|_ -|  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__| |_|   |_| |__|      |__|__|___|_  |___|___|___|_| |___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|               |_|
//////////////////////////////////////////////////////////////////////////
//
//
// This file defines all transport related functions or methods to create and auto use specific HTTP requests
//
//
package SkyLine_Standard_Library_HTTP

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"time"

	SkyEnv "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyEnvironment"
	Helpers "github.com/SkyPenguinLabs/SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

var TransportNewSL CustomTransport

func NewTransport(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"http.TransportNew",
		args,
		Helpers.SkyLine_Standard_Library_Helper_WithinRangeOFArguments(1, 2),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT),
	); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}

	URL := args[0].(*SkyEnv.SL_String).Value
	var Proxy string
	if len(args) > 1 {
		Proxy = args[1].(*SkyEnv.SL_String).Value
	}

	CustomTransport, x := CreateTransport(URL)
	if x != nil {
		return &SkyEnv.SL_Error{Message: "Could not create transport : " + x.Error()}
	}

	if Proxy != "" {
		x := CustomTransport.SetProxy(Proxy)
		if x != nil {
			return &SkyEnv.SL_Error{Message: "Could not set proxy URL : " + x.Error()}
		}
	}
	TransportNewSL = *CustomTransport
	return &SkyEnv.SL_NULL{}
}

func CreateTransport(requestURL string) (*CustomTransport, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	if strings.HasPrefix(requestURL, "https") {
		transport.TLSClientConfig.InsecureSkipVerify = true
	}
	customTransport := &CustomTransport{
		Transport: transport,
		ClientURL: requestURL,
	}
	return customTransport, nil
}

func (c *CustomTransport) SetProxy(proxyURL string) error {
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	c.Transport.Proxy = http.ProxyURL(parsedURL)
	c.ProxyURL = parsedURL
	return nil
}

func (c *CustomTransport) MakeRequest(method string) (*http.Response, error) {
	client := &http.Client{
		Transport: c.Transport,
		Timeout:   10 * time.Second,
	}
	request, err := http.NewRequest(method, c.ClientURL, c.ResponseReaderBod)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
