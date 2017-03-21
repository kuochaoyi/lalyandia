package protocol

import (
	"encoding/binary"
	"errors"

	"github.com/Nyarum/lalyandia/messages/protocol/n746"
	"github.com/davecgh/go-spew/spew"
)

type IClientProtocol interface {
	Opcode() uint16
	Unmarshal([]byte) (interface{}, error)
}

type IServerProtocol interface {
	Opcode() uint16
	Marshal() ([]byte, error)
}

type Protocol struct {
	packets map[uint16]IClientProtocol
}

func NewProtocol() *Protocol {
	protocol := (&Protocol{
		packets: make(map[uint16]IClientProtocol),
	}).registerPacket(n746.NewClientGGAuth())

	return protocol
}

func (p *Protocol) registerPacket(packet IClientProtocol) *Protocol {
	p.packets[packet.Opcode()] = packet

	return p
}

func (p *Protocol) Decode(opcode uint16, data []byte) (interface{}, error) {
	packet, ok := p.packets[opcode]
	if !ok {
		return nil, errors.New("Not found this opcode")
	}

	return packet.Unmarshal(data)
}

func (p *Protocol) Encode(sp IServerProtocol) ([]byte, error) {
	result, err := sp.Marshal()
	if err != nil {
		return nil, err
	}

	var (
		sizeOfPacket = len(result) + 3
		encodeBytes  = make([]byte, 0, sizeOfPacket)
		lenBytes     = make([]byte, 2)
	)

	binary.LittleEndian.PutUint16(lenBytes, uint16(sizeOfPacket))
	encodeBytes = append(encodeBytes, lenBytes...)
	encodeBytes = append(encodeBytes, byte(sp.Opcode()))
	encodeBytes = append(encodeBytes, result...)

	spew.Dump(encodeBytes)

	return encodeBytes, nil
}
