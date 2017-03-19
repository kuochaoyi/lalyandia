package main

import (
	"flag"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Nyarum/lalyandia/game/actors"
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
		props          = actor.FromInstance(&actors.GameHandle{})
	)

	_, err := servicesModule.LoadGameConfig(*resourcePath + "game_config.toml")
	if err != nil {
		fmt.Println(err)
	}

	err = srv.Run(props, ":9999")
	if err != nil {
		fmt.Println(err)
	}
}
