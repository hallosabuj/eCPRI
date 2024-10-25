package ecpri

import (
	"encoding/binary"
)

// This takes message buffer without header
func ParseMessageTye_0(messageBuffer []byte) *MessageType_0 {
	var messageBody MessageType_0
	// Parsing header
	payloadLength := binary.BigEndian.Uint16(messageBuffer[2:4])
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	// Parsing payload
	messageBody.PcId[0] = messageBuffer[4]
	messageBody.PcId[1] = messageBuffer[5]
	messageBody.SeqId[0] = messageBuffer[6]
	messageBody.SeqId[1] = messageBuffer[7]
	messageBody.IQData = messageBuffer[8:payloadLength]
	return &messageBody
}

func ParseMessageTye_1(messageBuffer []byte) *MessageType_1 {
	var messageBody MessageType_1
	// Parsing header
	payloadLength := binary.BigEndian.Uint16(messageBuffer[2:4])
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	// Parsing payload
	messageBody.PcId[0] = messageBuffer[4]
	messageBody.PcId[1] = messageBuffer[5]
	messageBody.SeqId[0] = messageBuffer[6]
	messageBody.SeqId[1] = messageBuffer[7]
	messageBody.BitSequence = messageBuffer[8:payloadLength]
	return &messageBody
}

func ParseMessageTye_2(messageBuffer []byte) *MessageType_2 {
	var messageBody MessageType_2
	// Parsing header
	payloadLength := binary.BigEndian.Uint16(messageBuffer[2:4])
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	// Parsing payload
	messageBody.RtcId[0] = messageBuffer[4]
	messageBody.RtcId[1] = messageBuffer[5]
	messageBody.SeqId[0] = messageBuffer[6]
	messageBody.SeqId[1] = messageBuffer[7]
	messageBody.RealTimeControlData = messageBuffer[8:payloadLength]
	return &messageBody
}
