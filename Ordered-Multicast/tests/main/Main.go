package main

import (
	"distributed-systems/Ordered-Multicast/src/server"
)

func main(){
	udpServer := server.NewUdpServer("luciano:1041")
	udpServer.Run()
}
