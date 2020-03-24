package vnet

// parse request data in Message
type Message struct {
	Id uint32
	DataLen uint32
	Data []byte
}

// create message pack
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// get message id
func (m *Message) GetMsgId() uint32 {
	return m.Id
}
// get message length
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}
// get message data
func (m *Message) GetData() []byte {
	return m.Data
}

// set message id
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}
// set message length
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}
// set message data
func (m *Message) SetData(data []byte) {
	m.Data = data
}