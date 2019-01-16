package util

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"encoding/json"
	"fmt"
)

// Encode a model.Message to JSON format
func UdpMessageToJSON(message *model.Message) ([]byte, error){ // TODO verify how to encode correctly
	byteArr,err := json.Marshal(message)
	if(err == nil){
		return nil,err
	}
	return byteArr,nil
}

// Decode a JSON format to model.Message
func JSONToUdpMessage(jsn []byte) (*model.Message, error){
	var msg model.Message
	err := json.Unmarshal(jsn,&msg)
	if(err != nil){
		return model.NewMessage(0,"","",nil),err
	}
	return &msg, nil
}


func main(){ // TODO TEST
	fmt.Println("test")
	msg := model.NewMessage(1,"luciano:1041", "jhonson:1041","roulli molly")
	b,err := UdpMessageToJSON(msg)
	if(err == nil){
		fmt.Println(b)
	}
}
