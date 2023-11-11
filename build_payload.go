package ecpri

import (
	"encoding/binary"
	"fmt"
)

func BuildHeader(protocolVersion int, messageType int, concatenationIndicatitor int, payloadLength uint16) (header []byte, err error) {
	if protocolVersion != 1 {
		fmt.Println("only protocol version 1 is supported")
		return nil, fmt.Errorf("only protocol version 1 is supported")
	}
	if concatenationIndicatitor != 0 && concatenationIndicatitor != 1 {
		fmt.Println("only 0/1 is supported as concatenation indicatior")
	}
	header = make([]byte, ECPRIHeaderSize)
	// First byte contains
	// 	4 bit: eCPRI Protocol Revision
	//	3 bit: Reserved
	//	1 bit: C is concatenation indicatior
	header[0] = byte((protocolVersion << 4 & 0xf0) | (concatenationIndicatitor & 0x0f))
	// Second byte contains message type
	header[1] = byte(messageType)
	// Third and fourth byte contains payload size
	binary.BigEndian.PutUint16(header[2:], payloadLength)

	return header, nil
}

func BuildMessageType_1(msg *MessageType_1) ([]byte, error) {
	// Build header
	var payloadLength uint16 = uint16(len(&msg.PcId) + len(&msg.SeqId) + len(*(&msg.BitSequence)))
	header, err := BuildHeader(msg.ProtocolVersion, 1, msg.ConcatenationIndicatitor, payloadLength)
	if err != nil {
		fmt.Println("unable to build header")
		return nil, err
	}

	// Build payload
	payload := make([]byte, len(msg.PcId)+len(msg.SeqId))
	payload[0] = msg.PcId[0]
	payload[1] = msg.PcId[1]
	payload[2] = msg.SeqId[0]
	payload[3] = msg.SeqId[1]
	payload = append(payload, *(&msg.BitSequence)...)

	var message []byte
	message = append(message, header...)
	message = append(message, payload...)
	return message, nil
}

func BuildMessageType_2(msg *MessageType_2) ([]byte, error) {
	// Build header
	var payloadLength uint16 = uint16(len(&msg.RtcId) + len(&msg.SeqId) + len(*(&msg.RealTimeControlData)))
	header, err := BuildHeader(msg.ProtocolVersion, 2, msg.ConcatenationIndicatitor, payloadLength)
	if err != nil {
		fmt.Println("unable to build header")
		return nil, err
	}

	// Build payload
	payload := make([]byte, len(&msg.RtcId)+len(&msg.SeqId))
	payload[0] = msg.RtcId[0]
	payload[1] = msg.RtcId[1]
	payload[2] = msg.SeqId[0]
	payload[3] = msg.SeqId[1]
	payload = append(payload, *(&msg.RealTimeControlData)...)

	var message []byte
	message = append(message, header...)
	message = append(message, payload...)
	return message, nil
}

func BuildMessageType_3(msg *MessageType_3) ([]byte, error) {
	// Build header
	var payloadLength uint16 = uint16(len(&msg.PcId) + len(&msg.SeqId) + len(*(&msg.DataTransfered)))
	header, err := BuildHeader(msg.ProtocolVersion, 3, msg.ConcatenationIndicatitor, payloadLength)
	if err != nil {
		fmt.Println("unable to build header")
		return nil, err
	}

	// Build payload
	payload := make([]byte, len(&msg.PcId)+len(&msg.SeqId))
	payload[0] = msg.PcId[0]
	payload[1] = msg.PcId[1]
	payload[2] = msg.PcId[2]
	payload[3] = msg.PcId[3]
	payload[4] = msg.SeqId[0]
	payload[5] = msg.SeqId[1]
	payload[6] = msg.SeqId[2]
	payload[7] = msg.SeqId[3]
	payload = append(payload, *(&msg.DataTransfered)...)

	var message []byte
	message = append(message, header...)
	message = append(message, payload...)
	return message, nil
}

func BuildMessageType_4(msg *MessageType_4) ([]byte, error) {
	// Build header
	var payloadLength uint16 = uint16(12 + msg.Length)
	header, err := BuildHeader(msg.ProtocolVersion, 4, msg.ConcatenationIndicatitor, payloadLength)
	if err != nil {
		fmt.Println("unable to build header")
		return nil, err
	}

	// Build payload
	payload := make([]byte, 12)
	payload[0] = msg.RemoteMemoryAccessId
	payload[1] = byte((0xf0 & (int(msg.ReadWriteIndication) << 4)) | (0x0f & int(msg.RequestResponseIndication)))
	payload[2] = msg.ElementId[0]
	payload[3] = msg.ElementId[1]
	payload[4] = msg.Address[0]
	payload[5] = msg.Address[1]
	payload[6] = msg.Address[2]
	payload[7] = msg.Address[3]
	payload[8] = msg.Address[4]
	payload[9] = msg.Address[5]
	payload[10] = uint8((msg.Length >> 8) & 0xFF) // upper byte
	payload[11] = uint8(msg.Length & 0xFF)        // lower byte
	payload = append(payload, *(&msg.Data)...)

	var message []byte
	message = append(message, header...)
	message = append(message, payload...)
	return message, nil
}

func BuildMessageType_5(msg *MessageType_5) ([]byte, error) {
	// Build header
	var payloadLength uint16 = uint16(20 + len(msg.DummyBytes))
	header, err := BuildHeader(msg.ProtocolVersion, 4, msg.ConcatenationIndicatitor, payloadLength)
	if err != nil {
		fmt.Println("unable to build header")
		return nil, err
	}

	// Build payload
	payload := make([]byte, 20)
	payload[0] = msg.MeasurementId
	payload[1] = byte(msg.ActionType)
	payload[2] = msg.TimeStamp[0]
	payload[3] = msg.TimeStamp[1]
	payload[4] = msg.TimeStamp[2]
	payload[5] = msg.TimeStamp[3]
	payload[6] = msg.TimeStamp[4]
	payload[7] = msg.TimeStamp[5]
	payload[8] = msg.TimeStamp[6]
	payload[9] = msg.TimeStamp[7]
	payload[10] = msg.TimeStamp[8]
	payload[11] = msg.TimeStamp[9]
	payload[12] = msg.CompensationValue[0]
	payload[13] = msg.CompensationValue[1]
	payload[14] = msg.CompensationValue[2]
	payload[15] = msg.CompensationValue[3]
	payload[16] = msg.CompensationValue[4]
	payload[17] = msg.CompensationValue[5]
	payload[18] = msg.CompensationValue[6]
	payload[19] = msg.CompensationValue[7]
	payload = append(payload, *(&msg.DummyBytes)...)

	var message []byte
	message = append(message, header...)
	message = append(message, payload...)
	return message, nil
}

func BuildMessageType_6(msg *MessageType_6) ([]byte, error) {
	// Build header
	var payloadLength uint16 = uint16(3 + len(*(&msg.VendorSpecificpayload)))
	header, err := BuildHeader(msg.ProtocolVersion, 6, msg.ConcatenationIndicatitor, payloadLength)
	if err != nil {
		fmt.Println("unable to build header")
		return nil, err
	}

	// Build payload
	payload := make([]byte, 3)
	payload[0] = msg.ResetId[0]
	payload[1] = msg.ResetId[1]
	payload[3] = byte(msg.ResetCodeOp)
	payload = append(payload, *(&msg.VendorSpecificpayload)...)

	var message []byte
	message = append(message, header...)
	message = append(message, payload...)
	return message, nil
}
