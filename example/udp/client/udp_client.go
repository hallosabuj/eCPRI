package main

import (
	"fmt"
	"net"
	"os"
)

func SendToServer(ecpriMessage []byte) {
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

func main() {

}
