package n746

import (
	"github.com/Nyarum/barrel"
	lsop "github.com/Nyarum/lalyandia/messages/protocol/n746/client-lsop"
)

// ServerInit struct for 0xC621 revision
type ServerInit struct {
	IDSession       int32
	LSVersion       int32
	RSAPublicKey    []byte
	Unknown         [16]byte
	BlowFishKey     [16]byte
	NullTermination byte
}

func NewServerInit() ServerInit {
	return ServerInit{
		LSVersion: 0xC621,
		Unknown: [16]byte{0x29, 0xDD, 0x95, 0x4E,
			0x77, 0xC3, 0x9C, 0xFC,
			0x97, 0xAD, 0xB6, 0x20,
			0x07, 0xBD, 0xE0, 0xF7},
		NullTermination: 0x00,
	}
}

func (si ServerInit) Opcode() uint16 {
	return lsop.RequestGGAuth
}

func (si ServerInit) Marshal() ([]byte, error) {
	bar := barrel.NewProcessor([]byte{})
	bar.WriteInt32(si.IDSession).
		WriteInt32(si.LSVersion).
		WriteBytes(si.RSAPublicKey).
		WriteBytes(si.Unknown[:]).
		WriteBytes(si.BlowFishKey[:]).
		WriteByte(si.NullTermination)

	if bar.Error() != nil {
		return nil, bar.Error()
	}

	return bar.Bytes(), nil
}
