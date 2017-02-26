package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages/local"
	"github.com/Nyarum/lalyandia/server"

	"fmt"
)

type GameHandle struct {
}

func (state *GameHandle) Receive(context actor.Context) {
	switch context.Message().(type) {
	case local.AcceptData:
		acceptData := context.Message().(local.AcceptData)

		fmt.Println(acceptData.Buf.String())

		acceptData.Connection.Write([]byte{0x01, 0x02})
	}
}

func main() {
	srv := server.NewServer()
	props := actor.FromInstance(&GameHandle{})

	err := srv.Run(props, ":9999")
	if err != nil {
		fmt.Println(err)
	}
}
