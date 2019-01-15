package util

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"fmt"
)

// Encode a model.UdpMessage to JSON format
func UdpMessageToJSON(message model.UdpMessage) ([]byte, error){
	byteArr,err := json.Marshal(message)
	if(err == nil){
		return nil,err
	}
	return byteArr,nil
}

// Decode a JSON format to model.Message
//func JSONToUdpMessage(json []byte) (model.UdpMessage, error){
//	var msg model.UdpMessage
//	err := json.Unmarshal(json,&msg)
//	if(err != nil){
//		return model.UdpMessage{nil,nil,nil},err
//	}
//	return msg, nil
//}


func main(){ // TODO TEST
	 msg := model.UdpMessage{"luciano:1041", "jhonson:1041","roulli molly"}
	b,err := UdpMessageToJSON(msg)
	if(err == nil){
		fmt.Println(b)
	}
}
