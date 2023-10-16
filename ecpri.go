package ecpri

// Function to construct an eCPRI message
func ConstructECPRIMessage(protocolVersion int, messageType int, concatenationIndicatitor int, pcId []byte, seqId []byte, bitSequence []byte) ([]byte, error) {
	var message []byte
	if messageType == 1 {
		message, _ = BuildMessageType_1(protocolVersion, messageType, concatenationIndicatitor, pcId, seqId, bitSequence)
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
