package n746

import (
	"errors"

	"github.com/Nyarum/barrel"
	lsop "github.com/Nyarum/lalyandia/messages/protocol/n746/server-lsop"
)

// ServerInit struct for 0xC621 revision
type ServerInit struct {
	IDSession       int32
	LSVersion       int32
	RSAPublicKey    [128]byte
	Unknown         [16]byte
	BlowFishKey     [16]byte
	NullTermination byte
}

func NewServerInit() *ServerInit {
	return &ServerInit{
		Unknown: [16]byte{0x29, 0xDD, 0x95, 0x4E,
			0x77, 0xC3, 0x9C, 0xFC,
			0x97, 0xAD, 0xB6, 0x20,
			0x07, 0xBD, 0xE0, 0xF7},
		NullTermination: 0x00,
	}
}

func (si *ServerInit) Opcode() int16 {
	return lsop.Init
}

func (si *ServerInit) Marshal(st interface{}) ([]byte, error) {
	siNew, ok := st.(*ServerInit)
	if !ok {
		return nil, errors.New("I don't know that struct")
	}

	siNew.Unknown = si.Unknown
	siNew.NullTermination = si.NullTermination

	bar := barrel.NewProcessor([]byte{})
	bar.WriteInt32(siNew.IDSession).
		WriteInt32(siNew.LSVersion).
		WriteBytes(siNew.RSAPublicKey[:]).
		WriteBytes(siNew.Unknown[:]).
		WriteBytes(siNew.BlowFishKey[:]).
		WriteByte(siNew.NullTermination)

	return nil, nil
}
