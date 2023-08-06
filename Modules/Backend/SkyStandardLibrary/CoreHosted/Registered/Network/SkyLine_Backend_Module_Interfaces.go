package SkyLine_Network

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func CallInterfaces() []net.Interface {
	interfaces, x := net.Interfaces()
	if x != nil {
		log.Fatal(x)
	}
	return interfaces
}

func RetrieveInterfaceNames() ([]string, error) {
	var names []string
	for _, iface := range CallInterfaces() {
		names = append(names, iface.Name)
	}
	return names, nil
}

func RetireveAllNetworkAddresses() ([]string, error) {
	var network []string
	for _, iface := range CallInterfaces() {
		addresses, x := iface.Addrs()
		if x != nil {
			return nil, x
		}
		for _, addr := range addresses {
			if ipn, ok := addr.(*net.IPNet); ok {
				network = append(network, ipn.IP.String())
			}
		}
	}
	return network, nil
}

func GetInterfaceIPByName(name string) (string, error) {
	var nosuchinterface bool
	for _, iface := range CallInterfaces() {
		if iface.Name == name {
			addresses, x := iface.Addrs()
			if x != nil {
				return "", x
			}
			for _, addr := range addresses {
				if ipn, ok := addr.(*net.IPNet); ok {
					return ipn.IP.String(), nil
				}
			}
		}
		nosuchinterface = true
	}
	if nosuchinterface {
		return "Unknown", nil
	} else {
		return "error?", nil
	}
}

func ParseMACAddress(mac string) (net.HardwareAddr, bool) {
	byter, x := hex.DecodeString(CleanCharacter(mac, ':'))
	if x != nil {
		fmt.Println("Error when decoding...")
		return nil, false // error during decoding
	}
	if len(byter) != 6 {
		return nil, false // invalid mac
	}
	return net.HardwareAddr(byter), true
}
