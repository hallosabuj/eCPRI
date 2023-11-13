package ecpri

const (
	ECPRIHeaderSize     = 4
	ECPRIHeaderReserved = 3
)

// /////////////////////////////////////////////////////////
// IQ Data
type MessageType_0 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	PcId                     [2]byte
	SeqId                    [2]byte
	IQData                   []byte
}

// /////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////
// Bit Sequence
type MessageType_1 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	PcId                     [2]byte
	SeqId                    [2]byte
	BitSequence              []byte
}

// /////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////
// Real-Time Control Data
type MessageType_2 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	RtcId                    [2]byte
	SeqId                    [2]byte
	RealTimeControlData      []byte
}

// /////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////
// Generic Data Transfer
type MessageType_3 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	PcId                     [4]byte
	SeqId                    [4]byte
	DataTransfered           []byte
}

// /////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////
// Remorte Memory Access
type Type4_ReadOrWrite int
type Type4_ReqOrResp int

const (
	Type4_Read          Type4_ReadOrWrite = 0
	Type4_Write         Type4_ReadOrWrite = 1
	Type4_Write_No_Resp Type4_ReadOrWrite = 2

	Type4_Request  Type4_ReqOrResp = 0
	Type4_Response Type4_ReqOrResp = 1
	Type4_Failure  Type4_ReqOrResp = 2
)

type MessageType_4 struct {
	ProtocolVersion           int
	ConcatenationIndicatitor  int
	RemoteMemoryAccessId      byte
	ReadWriteIndication       Type4_ReadOrWrite
	RequestResponseIndication Type4_ReqOrResp
	ElementId                 [2]byte
	Address                   [6]byte
	Length                    uint16
	Data                      []byte
}

// /////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////
// One-way Delay Measurement
type Type5_ActionType byte

const (
	Type5_Request                   Type5_ActionType = 0x00
	Type5_RequestwithFollowUp       Type5_ActionType = 0x01
	Type5_Response                  Type5_ActionType = 0x02
	Type5_RemoteRequest             Type5_ActionType = 0x03
	Type5_RemoteRequestWithFollowUp Type5_ActionType = 0x04
	Type5_FollowUp                  Type5_ActionType = 0x05
)

type MessageType_5 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	MeasurementId            byte
	ActionType               Type5_ActionType
	TimeStamp                [10]byte
	CompensationValue        [8]byte
	DummyBytes               []byte
}

// /////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////
// Remote Reset
type Type6_ResetCodeOp byte

const (
	Type6_RemoteResetRequest  Type6_ResetCodeOp = 0x01
	Type6_RemoteResetResponse Type6_ResetCodeOp = 0x01
)

type MessageType_6 struct {
	ProtocolVersion          int
	ConcatenationIndicatitor int
	ResetId                  [2]byte
	ResetCodeOp              Type6_ResetCodeOp
	VendorSpecificpayload    []byte
}

// /////////////////////////////////////////////////////////

type Message struct {
	MessageType int
	Type_1      *MessageType_1
	Type_2      *MessageType_2
}
