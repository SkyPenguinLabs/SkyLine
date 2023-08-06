//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//	This file contains all information for the SkyLine standard library //   _____ _       __    _             _____ _____ ____        _____ _ _     _____         _
//	such as the math, io, IoT, ECP, Socket, Web, HTTP and other various //  |   __| |_ _ _|  |  |_|___ ___ ___|   __|_   _|    \      |   __|_| |___|   __|_ _ ___| |_ ___ ______
//	library content types. This code section defines a sub unit under a //  |__   | '_| | |  |__| |   | -_|___|__   | | | |  |  |     |   __| | | -_|__   | | |_ -|  _| -_|     |
//	primary unit. This means that this section is under neath all of the//  |_____|_,_|_  |_____|_|_|_|___|   |_____| |_| |____/ _____|__|  |_|_|___|_____|_  |___|_| |___|_|_|_|
//	primary units that define the functions to register the sub func's  //            |___|                                     |_____|                   |___|
//////////////////////////////////////////////////////////////////////////
//
//
// This file defines all of the utilitites that this library will use or other functions to carve out information
//
//
package SkyLine_Standard_Library_File

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func SkyLine_File_Lib_CarveSector(file string, startb, endB []byte, output string) {
	in, x := ioutil.ReadFile(file)

	if x != nil {
		log.Fatal(x)
	}

	var start, end int
	for i := 0; i < len(in)-len(startb); i++ {
		if string(in[i:i+len(startb)]) == string(startb) {
			start = i + len(startb)
			break
		}
	}
	for i := len(in) - len(endB); i > 0; i-- {
		if string(in[i:i+len(endB)]) == string(endB) {
			end = i
			break
		}
	}
	if start >= end {
		os.Exit(0)
	}

	outputFile, x := os.Create(output)

	if x != nil {
		log.Fatal(x)
	}

	defer outputFile.Close()
	_, x = outputFile.Write(in[start:end])
	if x != nil {
		log.Fatal(x)
	}
}

func FileLib_Attempt_Mime_Type_Through_ShortDB(DBF, File string) (filenames []string) {
	type Database map[string][]map[string]string
	rawData, err := ioutil.ReadFile(DBF)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var db Database
	err = json.Unmarshal(rawData, &db)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	file, err := os.Open(File)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	magicBytes := make([]byte, 6)
	_, err = file.Read(magicBytes)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	hexString := ""
	for _, b := range magicBytes {
		hexString += fmt.Sprintf("%02X ", b)
	}
	hexString = hexString[:len(hexString)-1]
	if files, ok := db[hexString]; ok {
		for _, file := range files {
			filenames = append(filenames, file["description"])
		}
		return
	}
	return
}

func FileLib_GrabHeader_OfSpecificSize(Filename string, size int) (byter, ascii string) {
	file, err := os.Open(Filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	header := make([]byte, size)
	_, err = file.Read(header)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	byter = fmt.Sprintf("%X", header)
	ascii = string(header)
	return
}

func FileLib_ReadLines_InitateData(File string) []string {
	file, x := os.Open(File)
	if x != nil {
		log.Fatal(x)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if x = scanner.Err(); x != nil {
		log.Fatal(x)
	}
	return lines
}

func FileLib_Extract_Lable(lines []string, label string) []string {
	var data []string
	for _, line := range lines {
		if strings.HasPrefix(line, label+":") {
			value := strings.TrimSpace(strings.TrimPrefix(line, label+":"))
			data = append(data, value)
		}
	}
	return data
}
