package local

import (
	"net"

	"github.com/Nyarum/lalyandia/messages/protocol"
)

type InitData struct {
	Connection     net.Conn
	CustomProtocol protocol.Protocol
	BlowfishKey    [16]byte
	RSAKey         []byte
}
