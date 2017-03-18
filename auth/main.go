package main

import (
	"flag"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/auth/actors"
	"github.com/Nyarum/lalyandia/server"
	"github.com/Nyarum/lalyandia/services"

	"fmt"
)

func main() {
	resourcePath := flag.String("resource", "resource/", "Path to resources")
	flag.Parse()

	var (
		srv            = server.NewServer()
		servicesModule = services.NewServices()
		props          = actor.FromInstance(&actors.AuthHandle{})
	)

	_, err := servicesModule.LoadAuthConfig(*resourcePath + "auth_config.toml")
	if err != nil {
		fmt.Println(err)
	}

	err = srv.Run(props, ":7777")
	if err != nil {
		fmt.Println(err)
	}
}
