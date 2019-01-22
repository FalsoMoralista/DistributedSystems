package util

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"net"
)


const (
	REQUEST string = "A"
	POST string = "B"
	ERROR string = "F"

	GROUP string = "0"
	USER string = "1"

	BUFFER_SIZE int = 1024
)


func SendUdp(address string, message *model.Message) (int, []byte, error){
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
	n,err := conn.Write(parsed)
	conn.Read(buf[0:])
	return n,buf,err
}
