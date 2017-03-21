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
	"github.com/Nyarum/lalyandia/messages/protocol"
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

	fmt.Println("Server is starting... - ", tcpAddr.String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[Accept] err was -", err)
			continue
		}

		fmt.Println("Accepted -", conn.RemoteAddr())

		go func() {
			var (
				pid            = actor.Spawn(props)
				data           = bytes.NewBuffer(make([]byte, 0, 10384))
				customProtocol = protocol.NewProtocol()
			)

			defer func() {
				pid.Stop()
				conn.Close()
			}()

			pid.Tell(local.InitData{
				Connection: conn,
			})

			for {
				buf := make([]byte, 10384)
				readLen, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						fmt.Println("Client is disconnected -", conn.RemoteAddr())
					} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
						fmt.Println(err)
					}
					break
				}

				if readLen == 0 {
					continue
				}

				fmt.Println("Accept any bytes -", buf[:readLen])

				data.Write(buf[:readLen])

				if data.Len() < 2 {
					continue
				}

				lenPacket := binary.LittleEndian.Uint16(data.Bytes()[0:2])

				if data.Len() < int(lenPacket) {
					continue
				}

				opcode := data.Bytes()[2]
				// For other packets with two sub opcode
				/*
					if opcode == 0xFE {
						opcode =
					}
				*/

				data.Truncate(3)

				log.Printf("Len of packet - %d and Opcode - %d\n", lenPacket, opcode)

				st, err := customProtocol.Decode(uint16(opcode), data.Bytes())
				if err != nil {
					fmt.Println(err)
					break
				}

				pid.Tell(st)

				data.Truncate(int(lenPacket))
			}
		}()
	}
}
