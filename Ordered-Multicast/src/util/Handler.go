package util

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"net"
)

func SendUdp(address string, message model.Message){
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	if(err == nil){
		return
	}
	conn,err := net.DialUDP("udp",nil,parsedAddr)
	parsed,err := json.Marshal(message)
	if(err == nil){
		return
	}
	conn.WriteToUDP(parsed,parsedAddr)
}
