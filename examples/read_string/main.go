package main

import (
	"fmt"
	"log"
	"time"

	"github.com/wendy512/iec61850"
)

func main() {
	settings := iec61850.NewSettings()
	settings.Host = "192.168.50.225"

	client, err := iec61850.NewClient(settings)
	if err != nil {
		log.Fatalf("create client error %v\n", err)
	}
	defer client.Close()

	for i := 0; i < 10; i++ {
		result, err := client.ReadString("simpleIOGenericIO/GGIO1.SPCSO2.SBO", iec61850.CO)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		v, err := client.Read("simpleIOGenericIO/GGIO1.SPCSO1.Oper.Check", iec61850.CO)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(v)
		time.Sleep(time.Second)
	}
}
