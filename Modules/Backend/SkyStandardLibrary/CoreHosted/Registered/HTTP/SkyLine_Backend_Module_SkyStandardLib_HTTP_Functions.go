//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ _____ _____     _____                     _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|  |  |_   _|_   _|  _  |___| __  |___ ___ _ _ ___ ___| |_ ___
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|     | | |   | | |   __|___|    -| -_| . | | | -_|_ -|  _|_ -|
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |__|__| |_|   |_| |__|      |__|__|___|_  |___|___|___|_| |___|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|               |_|
//////////////////////////////////////////////////////////////////////////
//
//
// Def -> This file defines all functions to make a request to specific targets or to craft specific HTTP requests
//
//
package SkyLine_Standard_Library_HTTP

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	SkyEnv "SkyLine/Modules/Backend/SkyEnvironment"
	Helpers "SkyLine/Modules/Backend/SkyStandardLibrary/CoreHosted/GenuineCoreHelpers"
)

func ReadResponseBody(bod *http.Response) map[SkyEnv.HashKey]SkyEnv.HashPair {
	defer bod.Body.Close()
	Mapper := make(map[string]interface{}, 0)
	RetHash := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
	Mapper["Status"] = bod.Status
	Mapper["StatusCode"] = bod.StatusCode
	Mapper["Proto"] = bod.Proto
	Mapper["ProtoMajor"] = bod.ProtoMajor
	Mapper["ProtoMinor"] = bod.ProtoMinor
	var buffer bytes.Buffer
	_, x := io.Copy(&buffer, bod.Body)
	if x != nil {
		log.Fatal(x)
	}
	Mapper["ResponseBody"] = buffer.String()
	Mapper["TranserEncoding"] = bod.TransferEncoding
	Mapper["ContentLength"] = bod.ContentLength
	for k, v := range Mapper {
		key := &SkyEnv.SL_String{Value: k}
		val := &SkyEnv.SL_String{Value: fmt.Sprint(v)}
		NewHashCreate := SkyEnv.HashPair{Key: key, Value: val}
		RetHash[key.SL_HashKeyType()] = NewHashCreate
	}
	return RetHash
}

//var proxy string = "socks5://127.0.0.1:9050"
func TorReq(proxy, target string) interface{} {
	torProxyUrl, x := url.Parse(proxy)
	if x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	tort := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: tort, Timeout: time.Second * 30}
	resp, err1 := client.Get(target)
	if err1 != nil {
		return &SkyEnv.SL_Error{Message: err1.Error()}
	}
	return resp
}

func MakeRequest() SkyEnv.SL_Object {
	hashed := make(map[SkyEnv.HashKey]SkyEnv.HashPair)
	var currentbod *http.Response
	if Httpstruct.Req_URL != "" && Httpstruct.Req_Method != "" {

		//:::::::::::::::::::::::::
		//:: Check for transport ::
		//:::::::::::::::::::::::::
		if TransportNewSL.Transport != nil {
			resp, x := TransportNewSL.MakeRequest(Httpstruct.Req_Method)
			if x != nil {
				return &SkyEnv.SL_Error{Message: x.Error()}
			}
			hashed = ReadResponseBody(resp)
			currentbod = resp
		} else {
			//::::::::::::::::::::::::::::::::
			//:: Check for Tor requirements ::
			//::::::::::::::::::::::::::::::::
			if Httpstruct.Tor && Httpstruct.TorProx != "" {
				Response := TorReq(Httpstruct.TorProx, Httpstruct.Req_URL)
				if f, ok := Response.(*SkyEnv.SL_Error); ok {
					return f
				} else {
					if Body, ok := Response.(*http.Response); ok {
						// call body filler
						hashed = ReadResponseBody(Body)
						currentbod = Body
					}
				}
			} else {
				client := &http.Client{}
				req, err := http.NewRequest(Httpstruct.Req_Method, Httpstruct.Req_URL, nil)
				if err != nil {
					return &SkyEnv.SL_Error{Message: err.Error()}
				}
				for i := 0; i < len(Httpstruct.Headers); i++ {
					parts := strings.SplitN(Httpstruct.Headers[i], ":", 2)
					if len(parts) != 2 {
						return &SkyEnv.SL_Error{Message: "Header Parse error: Header format must be (key:value) or (header:value), the lenggth was not two for splitting parts"}
					}
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					req.Header.Set(key, value)
				}
				response, x := client.Do(req)
				if x != nil {
					return &SkyEnv.SL_Error{Message: "Could not make client execute request -> " + x.Error()}
				}
				currentbod = response
				defer response.Body.Close()
				hashed = ReadResponseBody(currentbod)
			}
		}
		//::::::::::::::::::::::::::::
		//:: Check for file output  ::
		//::::::::::::::::::::::::::::
		if Httpstruct.Filename != "" && Httpstruct.OutToFile {
			file, err := os.Create(Httpstruct.Filename)
			if err != nil {
				fmt.Println("Error creating file:", err)
				os.Exit(0)
			}
			defer file.Close()
			_, err = io.Copy(file, currentbod.Body)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(0)
			}
			fmt.Println("[True] Data written to a file")
		}
	} else {
		return &SkyEnv.SL_Error{Message: "Please use http.New() to make sure that configuration is sent | MISSING(URL, METHOD)..."}
	}
	return &SkyEnv.SL_HashMap{Pairs: hashed}
}

func MakeTransPortRequest(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("http.TransportReq", args, Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	if TransportNewSL.Transport == nil {
		return &SkyEnv.SL_Error{Message: "Please make sure you run http.TransportNew() before trying to make requests with the client"}
	} else {
		method := args[0].(*SkyEnv.SL_String).Value
		resp, x := TransportNewSL.MakeRequest(method)
		if x != nil {
			return &SkyEnv.SL_Error{Message: x.Error()}
		}
		return &SkyEnv.SL_HashMap{Pairs: ReadResponseBody(resp)}
	}
}

func MakeGet(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("http.Get", args, Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	bod, x := http.Get(args[0].SkyLine_ObjectFunction_GetTrueValue())
	if x != nil {
		return &SkyEnv.SL_Error{Message: "HTTP request failed"}
	}
	defer bod.Body.Close()
	return &SkyEnv.SL_HashMap{Pairs: ReadResponseBody(bod)}
}

func MakeNullPOST(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments("http.NullPost", args, Helpers.SkyLine_Standard_Library_Helper_ExactArguments(1), Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	uri := args[0].(*SkyEnv.SL_String).Value
	data := []byte(``)
	request, x := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	if x != nil {
		fmt.Println("Could not make a request -> ", x)
		return &SkyEnv.SL_Boolean{Value: false}
	}
	client := &http.Client{}
	response, x := client.Do(request)
	if x != nil {
		fmt.Println("Could not fufil the request -> ", x)
		return &SkyEnv.SL_Boolean{Value: false}
	}
	return &SkyEnv.SL_HashMap{Pairs: ReadResponseBody(response)}
}

func MakeBasePost(args ...SkyEnv.SL_Object) SkyEnv.SL_Object {
	if x := Helpers.SkyLine_Standard_Library_Helper_CheckArguments(
		"http.Post",
		args,
		Helpers.SkyLine_Standard_Library_Helper_ExactArguments(2),
		Helpers.SkyLine_Standard_Library_Helper_ArgumentCheckDataType(SkyEnv.SKYLINE_DATATYPE_STRING_OBJECT, SkyEnv.SKYLINE_DATATYPE_HASH_OBJECT)); x != nil {
		return &SkyEnv.SL_Error{Message: x.Error()}
	}
	URL := args[0].(*SkyEnv.SL_String).Value
	transport := &http.Transport{}
	TransPortData := args[1].(*SkyEnv.SL_HashMap).Pairs
	JsonData := make(map[string]string)
	for _, v := range TransPortData {
		JsonData[v.Key.SkyLine_ObjectFunction_GetTrueValue()] = v.Value.SkyLine_ObjectFunction_GetTrueValue()
	}
	jsonValue, _ := json.Marshal(JsonData)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Could not create a new http.POST request -> ", err)
		return &SkyEnv.SL_Boolean{Value: false}
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not execute http.Client->POST request: ", err)
		return &SkyEnv.SL_Error{Message: "http.Client->POST [FAIL]"}
	}
	defer resp.Body.Close()
	return &SkyEnv.SL_HashMap{Pairs: ReadResponseBody(resp)}
}
