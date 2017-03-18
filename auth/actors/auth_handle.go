package actors

import (
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
