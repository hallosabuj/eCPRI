package ecpri

import (
	"fmt"
)

// Function to construct an eCPRI message
func ConstructECPRIMessage(msg Message) ([]byte, error) {
	var message []byte
	if msg.MessageType == 1 {
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
