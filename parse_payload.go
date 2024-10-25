package ecpri

import (
	"encoding/binary"
)

// This takes message buffer without header
func ParseMessageTye_0(messageBuffer []byte) *MessageType_0 {
	payloadLength := binary.BigEndian.Uint16(messageBuffer[2:4])
	var messageBody MessageType_0
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	messageBody.PcId[0] = messageBuffer[4]
	messageBody.PcId[1] = messageBuffer[5]
	messageBody.SeqId[0] = messageBuffer[6]
	messageBody.SeqId[1] = messageBuffer[7]
	messageBody.IQData = messageBuffer[4:payloadLength]
	return &messageBody
}
