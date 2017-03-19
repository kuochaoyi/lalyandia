package actors

import (
	"fmt"
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages/local"
	"github.com/Nyarum/lalyandia/messages/protocol"
)

type GameHandle struct {
	connection     net.Conn
	customProtocol protocol.Protocol
}

func (state *GameHandle) Receive(context actor.Context) {
	switch context.Message().(type) {
	case local.InitData:
		initData := context.Message().(local.InitData)
		state.connection = initData.Connection
		state.customProtocol = initData.CustomProtocol

		context.SetBehavior(state.Client)
	}
}

func (state *GameHandle) Client(context actor.Context) {
	switch context.Message().(type) {
	default:
		fmt.Println("Test receive")
	}
}
