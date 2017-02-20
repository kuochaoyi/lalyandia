package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages"
	"github.com/Nyarum/lalyandia/server"

	"fmt"
)

type GameHandle struct {
}

func (state *GameHandle) Receive(context actor.Context) {
	switch context.Message().(type) {
	case messages.AcceptData:
		acceptData := context.Message().(messages.AcceptData)

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
