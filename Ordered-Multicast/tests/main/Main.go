package main

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"fmt"
)

func main(){
	//udpServer := server.NewUdpServer("luciano:1041")
	//udpServer.Run()
	msg := model.NewMessage(1,"luciano","jhosnon",[]byte("007"))
	bslice,err := json.Marshal(&msg)
	msg2 := model.NewMessage(0,"","",nil)
	json.Unmarshal(bslice,&msg2)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(string(bslice))

}
