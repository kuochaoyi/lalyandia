package server

import (
	"bytes"
	"fmt"
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages"
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

		conn.SetWriteBuffer(2)

		fmt.Println("Accepted -", conn.RemoteAddr())

		go func() {
			var (
				pid  = actor.Spawn(props)
				data = bytes.NewBuffer([]byte{})
			)

			defer func() {
				pid.Stop()
				conn.Close()
			}()

			for {
				buf := make([]byte, 1024)
				readLen, err := conn.Read(buf)
				if err != nil {
					fmt.Println("[Read] err was -", err)
					break
				}

				if readLen == 0 {
					continue
				}

				data.Write(buf)

				fmt.Println("Len: ", readLen)

				pid.Tell(messages.AcceptData{
					Connection: conn,
					Buf:        *data,
				})

				data.Reset()
			}
		}()
	}
}
