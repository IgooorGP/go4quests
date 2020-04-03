package main

import (
	"github.com/IgooorGP/go4quests/go4quests/infra"
)

func main() {
	server, tcpSocketAddressAndPort := infra.SetupApplication()
	server.Run(tcpSocketAddressAndPort)
}
