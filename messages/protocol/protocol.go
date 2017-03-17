package protocol

import "github.com/Nyarum/lalyandia/messages/protocol/n746"

type IClientProtocol interface {
	Opcode() int16
	Unmarshal([]byte) (interface{}, error)
}

type IServerProtocol interface {
	Opcode() int16
	Marshal(interface{}) ([]byte, error)
}

type Protocol struct {
	Client map[int16]IClientProtocol
	Server map[int16]IServerProtocol
}

func NewProtocol() *Protocol {
	protocol := (&Protocol{
		Client: make(map[int16]IClientProtocol),
		Server: make(map[int16]IServerProtocol),
	}).registerServer(n746.NewServerInit())

	return &Protocol{}
}

func (p *Protocol) registerClient(packet IClientProtocol) *Protocol {
	p.Client[packet.Opcode()] = packet

	return p
}

func (p *Protocol) registerServer(packet IServerProtocol) *Protocol {
	p.Server[packet.Opcode()] = packet

	return p
}

func (p *Protocol) Decode(opcode int16, data []byte) (interface{}, error) {
	return p.Client[opcode].Unmarshal(data)
}

func (p *Protocol) Encode(opcode int16, st interface{}) ([]byte, error) {
	return p.Server[opcode].Marshal(st)
}
