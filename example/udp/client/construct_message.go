package main

import (
	"fmt"

	"github.com/hallosabuj/ecpri"
)

func ConstructMessage_0() {
	////////////////////////////////////////////////////////
	// Constructing Message type 0
	var msg ecpri.Message
	msg.MessageType = 0
	msg.Type_0 = &ecpri.MessageType_0{
		ProtocolRevision:         1,
		ConcatenationIndicatitor: 0,
		PcId:                     [2]byte{18, 52},
		SeqId:                    [2]byte{13, 13},
		IQData:                   []byte("Hello"),
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}

func ConstructMessage_1() {
	//////////////////////////////////////////////////////////
	// Constructing Message type 1
	var msg ecpri.Message
	msg.MessageType = 1
	msg.Type_1 = &ecpri.MessageType_1{
		ProtocolRevision:         1,
		ConcatenationIndicatitor: 0,
		PcId:                     [2]byte{18, 52},
		SeqId:                    [2]byte{13, 13},
		BitSequence:              []byte("Hello"),
	}
	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	//////////////////////////////////////////////////////////
}

func ConstructMessage_2() {
	////////////////////////////////////////////////////////
	// Constructing Message type 2
	var msg ecpri.Message
	msg.MessageType = 2
	msg.Type_2 = &ecpri.MessageType_2{
		ProtocolRevision:         1,
		ConcatenationIndicatitor: 0,
		RtcId:                    [2]byte{18, 52},
		SeqId:                    [2]byte{13, 13},
		RealTimeControlData:      []byte("Hello"),
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}

func ConstructMessage_3() {
	////////////////////////////////////////////////////////
	// Constructing Message type 3
	var msg ecpri.Message
	msg.MessageType = 3
	msg.Type_3 = &ecpri.MessageType_3{
		ProtocolRevision:         1,
		ConcatenationIndicatitor: 0,
		PcId:                     [4]byte{18, 52, 18, 52},
		SeqId:                    [4]byte{13, 13, 13, 13},
		DataTransfered:           []byte("Hello"),
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}

func ConstructMessage_4() {
	////////////////////////////////////////////////////////
	// Constructing Message type 4
	var msg ecpri.Message
	msg.MessageType = 4
	msg.Type_4 = &ecpri.MessageType_4{
		ProtocolRevision:          1,
		ConcatenationIndicatitor:  0,
		RemoteMemoryAccessId:      byte(10),
		ReadWriteIndication:       ecpri.Type4_Read,
		RequestResponseIndication: ecpri.Type4_Request,
		ElementId:                 [2]byte{13, 13},
		Address:                   [6]byte{1, 2, 3, 4, 5, 6},
		Length:                    35,
		Data:                      []byte("HelloHelloHelloHelloHelloHelloHello"),
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}

func ConstructMessage_5() {
	////////////////////////////////////////////////////////
	// Constructing Message type 5
	var msg ecpri.Message
	msg.MessageType = 5
	msg.Type_5 = &ecpri.MessageType_5{
		ProtocolRevision:         1,
		ConcatenationIndicatitor: 0,
		MeasurementId:            byte(10),
		ActionType:               ecpri.Type5_Request,
		TimeStamp:                [10]byte{13, 13, 13, 13, 13, 13, 13, 13, 13, 13},
		CompensationValue:        [8]byte{1, 2, 3, 4, 5, 6, 7, 8},
		DummyBytes:               []byte("Hello"),
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}

func ConstructMessage_6() {
	////////////////////////////////////////////////////////
	// Constructing Message type 6
	var msg ecpri.Message
	msg.MessageType = 6
	msg.Type_6 = &ecpri.MessageType_6{
		ProtocolRevision:         1,
		ConcatenationIndicatitor: 0,
		ResetId:                  [2]byte{13, 13},
		ResetCodeOp:              ecpri.Type6_RemoteResetRequest,
		VendorSpecificpayload:    []byte("Hello"),
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}

func ConstructMessage_7() {
	////////////////////////////////////////////////////////
	// Constructing Message type 7
	var msg ecpri.Message
	msg.MessageType = 7
	msg.Type_7 = &ecpri.MessageType_7{
		ProtocolRevision:         1,
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
				RaiseOrCease:              ecpri.Type7_CeaseAFault,
				FaultOrNotificationNumber: 3214,
				AdditionalInformation:     [4]byte{3, 4, 8, 9},
			},
		},
	}

	ecpriMessage, _ := ecpri.ConstructECPRIMessage(msg)
	fmt.Println("Constructed eCPRI message:", ecpriMessage)
	SendToServer(ecpriMessage)
	////////////////////////////////////////////////////////
}
