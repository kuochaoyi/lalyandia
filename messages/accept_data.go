package messages

import (
	"bytes"
	"net"
)

type AcceptData struct {
	Connection net.Conn
	Buf        bytes.Buffer
}
