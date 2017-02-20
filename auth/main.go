package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages"
	"github.com/Nyarum/lalyandia/server"

	"fmt"
)

type AuthHandle struct {
}

func (state *AuthHandle) Receive(context actor.Context) {
	switch context.Message().(type) {
	case messages.AcceptData:
		acceptData := context.Message().(messages.AcceptData)

		fmt.Println(acceptData.Buf.String())

		acceptData.Connection.Write([]byte("Hello, world!\n"))
	}
}

func main() {
	srv := server.NewServer()
	props := actor.FromInstance(&AuthHandle{})

	err := srv.Run(props, ":9998")
	if err != nil {
		fmt.Println(err)
	}
}
