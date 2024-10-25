package ecpri

import (
	"fmt"
)

// Function to construct an eCPRI message
func ConstructECPRIMessage(msg Message) ([]byte, error) {
	var message []byte
	if msg.MessageType == 0 {
		if msg.Type_0 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_0(msg.Type_0)
	} else if msg.MessageType == 1 {
		if msg.Type_1 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_1(msg.Type_1)
	} else if msg.MessageType == 2 {
		if msg.Type_2 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_2(msg.Type_2)
	} else if msg.MessageType == 3 {
		if msg.Type_3 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_3(msg.Type_3)
	} else if msg.MessageType == 4 {
		if msg.Type_4 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_4(msg.Type_4)
	} else if msg.MessageType == 5 {
		if msg.Type_5 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_5(msg.Type_5)
	} else if msg.MessageType == 6 {
		if msg.Type_6 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_6(msg.Type_6)
	} else if msg.MessageType == 7 {
		if msg.Type_7 == nil {
			fmt.Println("not enough information to build message")
			return nil, fmt.Errorf("not enough information to build message")
		}
		message, _ = BuildMessageType_7(msg.Type_7)
	} else {
		return nil, fmt.Errorf("unknown message type")
	}
	return message, nil
}

// This takes message buffer with header
func ParseECPRIMessage(messageBuffer []byte) Message {
	var msg Message
	if messageBuffer[1] == byte(0) {
		fmt.Println("received message type is 0")
		msg.MessageType = 0
		msg.Type_0 = ParseMessageTye_0(messageBuffer)
	} else if messageBuffer[1] == byte(1) {
		fmt.Println("received message type is 1")
		msg.MessageType = 1
		msg.Type_1 = ParseMessageTye_1(messageBuffer)
	} else if messageBuffer[1] == byte(2) {
		fmt.Println("received message type is 2")
		msg.MessageType = 2
		msg.Type_2 = ParseMessageTye_2(messageBuffer)
	} else if messageBuffer[1] == byte(3) {
		fmt.Println("received message type is 3")
		msg.MessageType = 3
		msg.Type_3 = ParseMessageTye_3(messageBuffer)
	} else {
		fmt.Println("received message type is undefined")
		msg.MessageType = -1
	}
	return msg
}
