package main

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"fmt"
)

func main(){
	//udpServer := server.NewUdpServer("luciano:1041")
	//udpServer.Run()
	msg := model.NewUdpMessage("luciano:1041", "jhonson:1041","roulli molly")
	b,err := util.UdpMessageToJSON(msg)
	//msg := util.JSONToUdpMessage(b)
	if(err == nil){
		fmt.Println(len(b))
	}

}
