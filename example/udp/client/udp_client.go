package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hallosabuj/ecpri"
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
	// ////////////////////////////////////////////////////////
	// // Constructing Message type 0
	// var msg ecpri.Message
	// msg.MessageType = 0
	// msg.Type_0 = &ecpri.MessageType_0{
	// 	ProtocolVersion:          1,
	// 	ConcatenationIndicatitor: 0,
	// 	PcId:                     [2]byte{18, 52},
	// 	SeqId:                    [2]byte{13, 13},
	// 	IQData:                   []byte("Hello"),
	// }

	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// ////////////////////////////////////////////////////////

	// //////////////////////////////////////////////////////////
	// // Constructing Message type 1
	// var msg ecpri.Message
	// msg.MessageType = 1
	// msg.Type_1 = &ecpri.MessageType_1{
	// 	ProtocolVersion:          1,
	// 	ConcatenationIndicatitor: 0,
	// 	PcId:                     [2]byte{18, 52},
	// 	SeqId:                    [2]byte{13, 13},
	// 	BitSequence:              []byte("Hello"),
	// }
	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// //////////////////////////////////////////////////////////

	// ////////////////////////////////////////////////////////
	// // Constructing Message type 2
	// var msg ecpri.Message
	// msg.MessageType = 2
	// msg.Type_2 = &ecpri.MessageType_2{
	// 	ProtocolVersion:          1,
	// 	ConcatenationIndicatitor: 0,
	// 	RtcId:                    [2]byte{18, 52},
	// 	SeqId:                    [2]byte{13, 13},
	// 	RealTimeControlData:      []byte("Hello"),
	// }

	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// ////////////////////////////////////////////////////////

	// ////////////////////////////////////////////////////////
	// // Constructing Message type 3
	// var msg ecpri.Message
	// msg.MessageType = 3
	// msg.Type_3 = &ecpri.MessageType_3{
	// 	ProtocolVersion:          1,
	// 	ConcatenationIndicatitor: 0,
	// 	PcId:                     [4]byte{18, 52, 18, 52},
	// 	SeqId:                    [4]byte{13, 13, 13, 13},
	// 	DataTransfered:           []byte("Hello"),
	// }

	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// ////////////////////////////////////////////////////////

	// ////////////////////////////////////////////////////////
	// // Constructing Message type 4
	// var msg ecpri.Message
	// msg.MessageType = 4
	// msg.Type_4 = &ecpri.MessageType_4{
	// 	ProtocolVersion:           1,
	// 	ConcatenationIndicatitor:  0,
	// 	RemoteMemoryAccessId:      byte(10),
	// 	ReadWriteIndication:       ecpri.Type4_Read,
	// 	RequestResponseIndication: ecpri.Type4_Request,
	// 	ElementId:                 [2]byte{13, 13},
	// 	Address:                   [6]byte{1, 2, 3, 4, 5, 6},
	// 	Length:                    5,
	// 	Data:                      []byte("Hello"),
	// }

	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// ////////////////////////////////////////////////////////

	// ////////////////////////////////////////////////////////
	// // Constructing Message type 5
	// var msg ecpri.Message
	// msg.MessageType = 5
	// msg.Type_5 = &ecpri.MessageType_5{
	// 	ProtocolVersion:          1,
	// 	ConcatenationIndicatitor: 0,
	// 	MeasurementId:            byte(10),
	// 	ActionType:               ecpri.Type5_Request,
	// 	TimeStamp:                [10]byte{13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
	// 	CompensationValue:        [8]byte{1, 2, 3, 4, 5, 6, 7, 8},
	// 	DummyBytes:               []byte("Hello"),
	// }

	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// ////////////////////////////////////////////////////////

	// ////////////////////////////////////////////////////////
	// // Constructing Message type 6
	// var msg ecpri.Message
	// msg.MessageType = 6
	// msg.Type_6 = &ecpri.MessageType_6{
	// 	ProtocolVersion:          1,
	// 	ConcatenationIndicatitor: 0,
	// 	ResetId:                  [2]byte{13, 13},
	// 	ResetCodeOp:              ecpri.Type6_RemoteResetRequest,
	// 	VendorSpecificpayload:    []byte("Hello"),
	// }

	// ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	// fmt.Println("Constructed eCPRI message:", ecpriMessage)
	// SendToServer(ecpriMessage)
	// ////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////
	// Constructing Message type 7
	var msg ecpri.Message
	msg.MessageType = 7
	msg.Type_7 = &ecpri.MessageType_7{
		ProtocolVersion:          1,
		ConcatenationIndicatitor: 0,
		EventId:                  byte(10),
		EventType:                ecpri.Type7_FaultIndication,
		SequenceNumber:           byte(11),
		NumberOfFaultOrNotif:     2,
		ElementDetails: []ecpri.Type7_ElementDetails{
			ecpri.Type7_ElementDetails{
				ElementId:                 [2]byte{1, 1},
				RaiseOrCease:              ecpri.Type7_RaiseAFault,
				FaultOrNotificationNumber: 1234,
				AdditionalInformation:     [4]byte{6, 7, 6, 7},
			},
			ecpri.Type7_ElementDetails{
				ElementId:                 [2]byte{2, 2},
				RaiseOrCease:              ecpri.Type7_RaiseAFault,
				FaultOrNotificationNumber: 1234,
				AdditionalInformation:     [4]byte{6, 7, 6, 7},
			},
		},
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////

}
