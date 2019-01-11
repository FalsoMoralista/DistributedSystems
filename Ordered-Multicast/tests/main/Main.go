package main

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"fmt"
)

func main(){
	multicast := model.NewMulticast("192.168.1.17:1010")
	addr := model.GetUdpAddr(multicast)
	fmt.Print(addr.String())
}