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

func BuildMessageType_2() {

}

func BuildMessageType_3() {

}

func BuildMessageType_4() {

}
