package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wendy512/iec61850"
)

func main() {
	TestGetLogicalDeviceList()
}

func TestGetLogicalDeviceList() {
	settings := iec61850.NewSettings()
	settings.Host = "192.168.50.225"

	client, err := iec61850.NewClient(settings)
	if err != nil {
		log.Fatalf("create client error %v\n", err)
	}
	defer client.Close()

	deviceList := client.GetLogicalDeviceList(true)
	marshal, err := json.MarshalIndent(deviceList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshal))
}
