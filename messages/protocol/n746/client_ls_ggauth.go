package n746

import (
	"github.com/Nyarum/barrel"
	lsop "github.com/Nyarum/lalyandia/messages/protocol/n746/client-lsop"
)

type ClientGGAuth struct {
	IDSession int32
}

func NewClientGGAuth() ClientGGAuth {
	return ClientGGAuth{}
}

func (si ClientGGAuth) Opcode() uint16 {
	return lsop.RequestGGAuth
}

func (si ClientGGAuth) Unmarshal(data []byte) (interface{}, error) {
	bar := barrel.NewProcessor(data)
	bar.ReadInt32(&si.IDSession)

	if bar.Error() != nil {
		return nil, bar.Error()
	}

	return si, nil
}
