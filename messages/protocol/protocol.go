package protocol

type IProtocol interface {
	Opcode() uint16
	Unmarshal([]byte) (interface{}, error)
	Marshal(interface{}) ([]byte, error)
}

type Protocol struct {
	Client map[uint16]IProtocol
	Server map[uint16]IProtocol
}

func NewProtocol() *Protocol {
	return &Protocol{}
}

func (p *Protocol) Decode(opcode uint16, data []byte) (interface{}, error) {
	return p.Client[opcode].Unmarshal(data)
}

func (p *Protocol) Encode(opcode uint16, st interface{}) ([]byte, error) {
	return p.Server[opcode].Marshal(st)
}
