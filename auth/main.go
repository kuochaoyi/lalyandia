package main

import (
	"flag"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/messages/local"
	"github.com/Nyarum/lalyandia/server"
	"github.com/Nyarum/lalyandia/services"

	"fmt"
)

type AuthHandle struct {
}

func (state *AuthHandle) Receive(context actor.Context) {
	switch context.Message().(type) {
	case local.AcceptData:
		acceptData := context.Message().(local.AcceptData)

		fmt.Println(acceptData.Buf.String())

		acceptData.Connection.Write([]byte("Hello, world!\n"))
	}
}

func main() {
	resourcePath := flag.String("resource", "resource/", "Path to resources")
	flag.Parse()

	var (
		srv            = server.NewServer()
		servicesModule = services.NewServices()
		props          = actor.FromInstance(&AuthHandle{})
	)

	_, err := servicesModule.LoadAuthConfig(*resourcePath + "auth_config.toml")
	if err != nil {
		fmt.Println(err)
	}

	err = srv.Run(props, ":9998")
	if err != nil {
		fmt.Println(err)
	}
}
