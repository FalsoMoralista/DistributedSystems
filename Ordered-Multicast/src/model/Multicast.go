package model

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	BUFFER_SIZE int = 1024
)

type MulticastListener struct {
	socket *net.UDPConn
	GROUP_ADDRESS *net.UDPAddr
	fifo_protocol *FifoOrder `json:"fifo_protocol,omitempty"`
}

func NewMulticastListener(GROUP_ADDRESS *net.UDPAddr) *MulticastListener {
	return &MulticastListener{GROUP_ADDRESS: GROUP_ADDRESS}
}
/**
* Listens for multicast messages from the current assigned group.
**/
func (this *MulticastListener) Listen() {
	for {
		buffer := make([]byte, BUFFER_SIZE)
		n, received_addr, err := this.socket.ReadFromUDP(buffer[0:])
		fmt.Println("Message received from " + received_addr.String())
		fmt.Println("message: " + string(buffer[0:n]))
		if err != nil {
			fmt.Print("Server: Error, returning...") // todo replace
		}
	}
}

/**
*	Multicast a message through the assigned group.
**/
func (this *MulticastListener) Multicast(listener *MulticastListener, message *Message) error{
	bArray,err := json.Marshal(message)
	_,err = listener.socket.WriteToUDP(bArray,listener.GROUP_ADDRESS)
	if(err != nil){
		return err
	}
	return nil
}

/**
*	Assigns an address to this group.
**/
func(this *MulticastListener) AssignGroupAddress(addr *net.UDPAddr){
	this.GROUP_ADDRESS = addr
}

/**
*	Connects to the group given an interface.
**/
func (this *MulticastListener) Connect(iface string) error{
	inf , err := net.InterfaceByName(iface)
	var conn *net.UDPConn
	if err != nil{
		conn, err = net.ListenMulticastUDP("udp4", inf , this.GROUP_ADDRESS) // MULTICAST SOCKET
		this.socket = conn
	}
	return err
}