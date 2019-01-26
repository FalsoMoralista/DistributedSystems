package multicast

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"fmt"
	"net"
)

type MulticastListener struct {
	socket *net.UDPConn
	buffer []byte
	GROUP_ADDRESS *net.UDPAddr
}

func NewMulticastListener(GROUP_ADDRESS *net.UDPAddr) *MulticastListener {
	return &MulticastListener{GROUP_ADDRESS: GROUP_ADDRESS}
}


func (this *MulticastListener) Listen() {
	for {
		this.buffer = make([]byte, util.BUFFER_SIZE)
		n, received_addr, err := this.socket.ReadFromUDP(this.buffer[0:])
		fmt.Println("Message received from " + received_addr.String())
		fmt.Println("message: " + string(this.buffer[0:n]))
		if err != nil {
			fmt.Print("Server: Error, returning...")
		}
	}
}

func (this *MulticastListener) Multicast(listener *MulticastListener,message *model.Message) error{
	bArray,err := json.Marshal(message)
	_,err = listener.socket.WriteToUDP(bArray,listener.GROUP_ADDRESS)
	if(err != nil){
		return err
	}
	return nil
}