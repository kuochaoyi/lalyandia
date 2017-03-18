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
	client map[int16]IClientProtocol
	server map[int16]IServerProtocol
}

func NewProtocol() *Protocol {
	protocol := (&Protocol{
		client: make(map[int16]IClientProtocol),
		server: make(map[int16]IServerProtocol),
	}).registerServer(n746.NewServerInit())

	return protocol
}

func (p *Protocol) registerClient(packet IClientProtocol) *Protocol {
	p.client[packet.Opcode()] = packet

	return p
}

func (p *Protocol) registerServer(packet IServerProtocol) *Protocol {
	p.server[packet.Opcode()] = packet

	return p
}

func (p *Protocol) Decode(opcode int16, data []byte) (interface{}, error) {
	return p.client[opcode].Unmarshal(data)
}

func (p *Protocol) Encode(opcode int16, st interface{}) ([]byte, error) {
	return p.server[opcode].Marshal(st)
}
