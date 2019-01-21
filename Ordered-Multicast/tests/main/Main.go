package main

import (
	"distributed-systems/Ordered-Multicast/src/server"
)

func main(){
	udpServer := server.NewUdpServer("luciano:1041")
	udpServer.Run()
	//msg := model.NewMessage(1,"luciano","jhosnon",[]byte("007"))
	//var msg2 = model.Message{}
	//json.Unmarshal(bslice,&msg2)
	//if(err != nil){
	//	fmt.Println(err)
	//}

}
