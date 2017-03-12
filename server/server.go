package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages/local"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(props *actor.Props, service string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	fmt.Println(tcpAddr.String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[Accept] err was -", err)
			continue
		}

		fmt.Println("Accepted -", conn.RemoteAddr())

		go func() {
			var (
				pid  = actor.Spawn(props)
				data = bytes.NewBuffer(make([]byte, 0, 10384))
			)

			defer func() {
				pid.Stop()
				conn.Close()
			}()

			for {
				buf := make([]byte, 10384)
				readLen, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						fmt.Println(err)
					} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
						fmt.Println(err)
					}
					break
				}

				if readLen == 0 {
					continue
				}

				data.Write(buf[:readLen])

				if data.Len() < 2 {
					continue
				}

				lenPacket := binary.BigEndian.Uint16(data.Bytes()[0:2])

				if data.Len() < int(lenPacket) {
					continue
				}

				opcode := data.Bytes()[2]

				data.Truncate(3)

				log.Printf("Len of packet - %d and Opcode - %d\n", lenPacket, opcode)

				// Need to refactor to send only read-to-send packets
				pid.Tell(local.AcceptData{
					Connection: conn,
					Buf:        *data,
				})

				data.Truncate(int(lenPacket))
			}
		}()
	}
}
