package util

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"fmt"
	"net"
)


const (
	REQUEST string = "A"
	POST string = "B"

	GROUP string = "0"
	USER string = "1"

)


func SendUdp(address string, message model.Message) (int, []byte, error){
	var buffer [1024]byte
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	if(err != nil){
		return 0,nil,err
	}
	conn,err := net.DialUDP("udp",nil,parsedAddr)
	parsed,err := json.Marshal(message)
	if(err != nil){
		return 0,nil,err
	}
	n,err := conn.WriteToUDP(parsed,parsedAddr)
	conn.Read(buffer[0:])
	if n {
		
	}
	//return n,buffer,err
}
