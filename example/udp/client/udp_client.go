package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hallosabuj/ecpri"
)

func main() {
	//////////////////////////////////////////////////////////
	// Creating eCPRI payload
	pcId := []byte{18, 52}
	seqId := []byte{13, 13}
	bitSequence := []byte("Hello")
	protocolVersion := 1
	messageType := 1
	concatenationIndicatitor := 0

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(protocolVersion, messageType, concatenationIndicatitor, pcId, seqId, bitSequence)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)

	//////////////////////////////////////////////////////////

	udpServer, err := net.ResolveUDPAddr("udp", ":1053")

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	_, err = conn.Write(ecpriMessage)
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println("eCPRI packet sent successfully!")

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println(string(received))
}
