package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx-copy/utils"
	"zinx-copy/ziface"
)

type DataPack struct {
}

func NewDataPack() ziface.IDataPack {
	return &DataPack{}

}

func (d *DataPack) GetHeadLen() uint32 {
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuffer := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuffer.Bytes(), nil
}

func (d *DataPack) Unpack(data []byte) (ziface.IMessage, error) {

	dataBuffer := bytes.NewReader(data)

	msg := &Message{}

	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断是否超过最大的包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg msg recv")
	}

	return msg, nil
}
