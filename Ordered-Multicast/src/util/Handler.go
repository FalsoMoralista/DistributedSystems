package util

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"net"
)


const (
	// Methods
	REQUEST string = "A"
	POST string = "B"
	DELETE string = "C"
	ACK string = "D"
	RESPONSE string = "E"
	ERROR string = "F"
	// Types
	GROUP string = "0"
	USER string = "1"
	// Settings
	BUFFER_SIZE int = 1024
)


func SendUdp(address string, message *model.Message) (int, []byte, error){ // todo review whether this should be here
	buf := make([]byte, BUFFER_SIZE)
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	if(err != nil){
		return 0,nil,err
	}
	conn,err := net.DialUDP("udp",nil,parsedAddr)
	parsed,err := json.Marshal(message)
	if(err != nil){
		return 0,nil,err
	}
	conn.Write(parsed)
	x,err := conn.Read(buf[0:])
	return x,buf,err
}
