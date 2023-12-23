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

// Function to parse an eCPRI message
func ParseECPRIMessage(message []byte) (byte, []byte) {
	// _ = message[ECPRIHeaderVersion]
	// msgType := message[ECPRIHeaderMsgType]
	// payloadLength := binary.BigEndian.Uint16(message[ECPRIHeaderPayload:])

	// payload := message[ECPRIHeaderSize : ECPRIHeaderSize+payloadLength]

	// return msgType, payload
	return 1, nil
}
