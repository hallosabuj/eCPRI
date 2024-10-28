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
	messageBody.IQData = messageBuffer[8:(4 + payloadLength)]
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
	messageBody.BitSequence = messageBuffer[8:(4 + payloadLength)]
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
	messageBody.RealTimeControlData = messageBuffer[8:(4 + payloadLength)]
	return &messageBody
}

func ParseMessageTye_3(messageBuffer []byte) *MessageType_3 {
	var messageBody MessageType_3
	// Parsing header
	payloadLength := binary.BigEndian.Uint16(messageBuffer[2:4])
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	// Parsing payload
	messageBody.PcId[0] = messageBuffer[4]
	messageBody.PcId[1] = messageBuffer[5]
	messageBody.PcId[2] = messageBuffer[6]
	messageBody.PcId[3] = messageBuffer[7]
	messageBody.SeqId[0] = messageBuffer[8]
	messageBody.SeqId[1] = messageBuffer[9]
	messageBody.SeqId[2] = messageBuffer[10]
	messageBody.SeqId[3] = messageBuffer[11]
	messageBody.DataTransfered = messageBuffer[12:(4 + payloadLength)]
	return &messageBody
}

func ParseMessageTye_4(messageBuffer []byte) *MessageType_4 {
	var messageBody MessageType_4
	// Parsing header
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	// Parsing payload
	messageBody.RemoteMemoryAccessId = messageBuffer[4]
	messageBody.ReadWriteIndication = Type4_ReadOrWrite((messageBuffer[5] >> 4) & 0x0f)
	messageBody.RequestResponseIndication = Type4_ReqOrResp(messageBuffer[5] & 0x0f)
	messageBody.ElementId[0] = messageBuffer[6]
	messageBody.ElementId[1] = messageBuffer[7]
	messageBody.Address[0] = messageBuffer[8]
	messageBody.Address[1] = messageBuffer[9]
	messageBody.Address[2] = messageBuffer[10]
	messageBody.Address[3] = messageBuffer[11]
	messageBody.Address[4] = messageBuffer[12]
	messageBody.Address[5] = messageBuffer[13]
	messageBody.Length = binary.BigEndian.Uint16(messageBuffer[14:16])
	messageBody.Data = messageBuffer[16 : 16+messageBody.Length]
	return &messageBody
}

func ParseMessageTye_5(messageBuffer []byte) *MessageType_5 {
	var messageBody MessageType_5
	// Parsing header
	payloadLength := binary.BigEndian.Uint16(messageBuffer[2:4])
	messageBody.ProtocolRevision = int((messageBuffer[0] >> 4) & 0x0f)
	messageBody.ConcatenationIndicatitor = int(messageBuffer[0] & 0x01)
	// Parsing payload
	messageBody.MeasurementId = messageBuffer[4]
	messageBody.ActionType = Type5_ActionType(messageBuffer[5])
	messageBody.TimeStamp[0] = messageBuffer[6]
	messageBody.TimeStamp[1] = messageBuffer[7]
	messageBody.TimeStamp[2] = messageBuffer[8]
	messageBody.TimeStamp[3] = messageBuffer[9]
	messageBody.TimeStamp[4] = messageBuffer[10]
	messageBody.TimeStamp[5] = messageBuffer[11]
	messageBody.TimeStamp[6] = messageBuffer[12]
	messageBody.TimeStamp[7] = messageBuffer[13]
	messageBody.TimeStamp[8] = messageBuffer[14]
	messageBody.TimeStamp[9] = messageBuffer[15]
	messageBody.CompensationValue[0] = messageBuffer[16]
	messageBody.CompensationValue[1] = messageBuffer[17]
	messageBody.CompensationValue[2] = messageBuffer[18]
	messageBody.CompensationValue[3] = messageBuffer[19]
	messageBody.CompensationValue[4] = messageBuffer[20]
	messageBody.CompensationValue[5] = messageBuffer[21]
	messageBody.CompensationValue[6] = messageBuffer[22]
	messageBody.CompensationValue[7] = messageBuffer[23]
	messageBody.DummyBytes = append(messageBody.DummyBytes, messageBuffer[24:payloadLength+4]...)
	return &messageBody
}
