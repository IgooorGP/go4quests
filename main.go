package main

import (
	"github.com/IgooorGP/go4quests/go4quests/config"
)

func main() {
	server, tcpSocketAddressAndPort := config.SetupApplication()
	server.Run(tcpSocketAddressAndPort)
}
