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

func (s *Server) Run(pid *actor.PID, service string) error {
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
			continue
		}

		go func() {
			defer conn.Close()

			data := bytes.NewBuffer([]byte{})

			for {
				readLen, err := conn.Read(data.Bytes())
				if err != nil {
					break
				}

				if readLen == 0 {
					continue
				}

				fmt.Println("Len: ", readLen)

				pid.Tell(messages.AcceptData{
					Connection: conn,
					Buf:        *data,
				})

				data.Truncate(readLen)
			}
		}()
	}
}
