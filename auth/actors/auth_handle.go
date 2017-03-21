package actors

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages/local"
	"github.com/Nyarum/lalyandia/messages/protocol"
	"github.com/Nyarum/lalyandia/messages/protocol/n746"
)

type AuthHandle struct {
	connection     net.Conn
	customProtocol protocol.Protocol
}

func (state *AuthHandle) Receive(context actor.Context) {
	switch context.Message().(type) {
	case local.InitData:
		initData := context.Message().(local.InitData)
		state.connection = initData.Connection
		state.customProtocol = initData.CustomProtocol

		blowfishRandom := [16]byte{}
		_, err := rand.Read(blowfishRandom[:])
		if err != nil {
			fmt.Println("Can't add random bytes to blowfish key -", err)
			state.connection.Close()
		}

		rsaKey, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			fmt.Println(err)
			state.connection.Close()
		}

		rsaKeyBytes := rsaKey.N.Bytes()
		protocol.ScrambleMod(rsaKeyBytes)

		serverInit := n746.NewServerInit()
		serverInit.IDSession = 0
		serverInit.RSAPublicKey = rsaKeyBytes
		serverInit.BlowFishKey = blowfishRandom

		newPacket, err := state.customProtocol.Encode(serverInit)
		if err != nil {
			fmt.Println("Can't create packet -", err)
			state.connection.Close()
		}

		state.connection.Write(newPacket)

		context.SetBehavior(state.Client)
	}
}

func (state *AuthHandle) Client(context actor.Context) {
	switch context.Message().(type) {
	case n746.ClientGGAuth:
		ggAuthPacket := context.Message().(n746.ClientGGAuth)

		fmt.Println(ggAuthPacket)
	}
}
