package main

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/server"
	"fmt"
)

func main(){
	udpServer := server.NewUdpServer("luciano:1041")
	go server.Run(udpServer)
	multicast := model.NewMulticast("luciano:1041")
	_,err := model.StartConnection(multicast)
	if(err != nil){
		fmt.Print(err)
	}
}
