package chatprotocol

import (
	"encoding/binary"
	"errors"
)

const (
	Uint32Length = 4
	ChatVersion  = 1
)

type ChatMeta struct {
	version int32
	msgTyp  int32
	msgLen  int32
}

func NewChatMetaV1(msgTyp int32) *ChatMeta {
	return &ChatMeta{version: ChatVersion, msgTyp: msgTyp}
}

func (cm *ChatMeta) SetMsgLen(l int32) {
	cm.msgLen = l
}

func (cm *ChatMeta) Pack() []byte {
	var (
		tmp, r []byte
	)
	tmp = make([]byte, Uint32Length)
	binary.BigEndian.PutUint32(tmp, uint32(cm.version))
	r = append(r, tmp...)

	tmp = make([]byte, Uint32Length)
	binary.BigEndian.PutUint32(tmp, uint32(cm.msgTyp))
	r = append(r, tmp...)

	tmp = make([]byte, Uint32Length)
	binary.BigEndian.PutUint32(tmp, uint32(cm.msgLen))
	r = append(r, tmp...)

	return r

}

func (cm *ChatMeta) UnPack(data []byte) (int, error) {
	var (
		offset int
	)

	if len(data) < Uint32Length {
		return 0, errors.New("less data for unpack version")
	}
	cm.version = int32(binary.BigEndian.Uint32(data[offset:]))
	offset += Uint32Length

	if len(data) < Uint32Length {
		return 0, errors.New("less data for unpack msg type")
	}
	cm.msgTyp = int32(binary.BigEndian.Uint32(data[offset:]))
	offset += Uint32Length

	if len(data) < Uint32Length {
		return 0, errors.New("less data for unpack msg length")
	}
	cm.msgLen = int32(binary.BigEndian.Uint32(data[offset:]))
	offset += Uint32Length

	return offset, nil
}
