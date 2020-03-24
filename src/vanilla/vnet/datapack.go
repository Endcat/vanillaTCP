package vnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"vanilla/utils"
	"vanilla/viface"
)

// TCP data parser

type DataPack struct {}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// datalen|msgid|data

func (dp *DataPack) GetHeadLen() uint32 {
	// datalen uint32 (4byte) + id uint32 (4byte)
	return 8
}

func (dp *DataPack) Pack(msg viface.IMessage) ([]byte, error) {
	// create buffer
	dataBuff := bytes.NewBuffer([]byte{})
	// write datalen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// write msgid
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// write data
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (dp *DataPack) Unpack(binaryData []byte) (viface.IMessage, error) {
	// create io reader
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	// read datalen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// read msgid
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// max package size limit
	if (utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize) {
		return nil, errors.New("msg data too large")
	}

	return msg, nil
}