package ecpri

const (
	ECPRIHeaderSize     = 4
	ECPRIHeaderReserved = 3
)

type MessageType_1 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	PcId                     [2]byte
	SeqId                    [2]byte
	BitSequence              []byte
}

type Message struct {
	MessageType int
	Type_1      *MessageType_1
}
