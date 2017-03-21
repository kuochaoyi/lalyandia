package server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages/local"
	"github.com/Nyarum/lalyandia/messages/protocol"
	"github.com/davecgh/go-spew/spew"
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

			// Init crypt things
			blowfishRandom := [16]byte{}
			_, err := rand.Read(blowfishRandom[:])
			if err != nil {
				fmt.Println("Can't add random bytes to blowfish key -", err)
				return
			}

			rsaKey, err := rsa.GenerateKey(rand.Reader, 1024)
			if err != nil {
				fmt.Println(err)
				return
			}

			rsaKeyBytes := rsaKey.N.Bytes()
			protocol.ScrambleMod(rsaKeyBytes)

			// Send init data to an auth actor
			pid.Tell(local.InitData{
				Connection:     conn,
				CustomProtocol: *customProtocol,
				BlowfishKey:    blowfishRandom,
				RSAKey:         rsaKeyBytes,
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

				data.Write(buf[:readLen])

				if data.Len() < 2 {
					continue
				}

				lenPacket := binary.LittleEndian.Uint16(data.Bytes()[0:2])

				if data.Len() < int(lenPacket) {
					continue
				}

				spew.Dump(data.Bytes()[2:])

				dataDecrypted, err := protocol.BlowfishDecrypt(data.Bytes()[2:], blowfishRandom[:])
				if err != nil {
					fmt.Println(err)
					break
				}

				spew.Dump(dataDecrypted)

				opcode := dataDecrypted[0]
				// For other packets with two sub opcode
				/*
					if opcode == 0xFE {
						opcode =
					}
				*/

				log.Printf("Len of packet - %d and Opcode - %d\n", lenPacket, opcode)

				st, err := customProtocol.Decode(uint16(opcode), dataDecrypted[1:])
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
