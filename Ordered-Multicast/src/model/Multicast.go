package model

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
	BUFFER_SIZE int = 4096
	LOCALHOST string = "localhost:1041"
)

type MulticastListener struct {
	Socket *net.UDPConn `json:"socket,omitempty"`
 	GROUP_ADDRESS *net.UDPAddr `json:"group_address,omitempty"`
	Fifo_protocol *FifoOrder `json:"fifo_protocol"`
}

/**
*	By default uses the loopback interface on the socket.
*	It can be modified by using the `Connect´ method.
**/
func NewMulticastListener(process_id int,GROUP_ADDRESS *net.UDPAddr) *MulticastListener {
	iface , _ := net.InterfaceByName("lo")
	sock, err := net.ListenMulticastUDP("udp4", iface,GROUP_ADDRESS)
	checkError(err)
	return &MulticastListener{Socket:sock, GROUP_ADDRESS: GROUP_ADDRESS, Fifo_protocol:NewFifoOrder(process_id)}  // todo fix (fifo_protocol:NewFifoOrder(-1))
}


/***************************************************************************************************************************************************************************************************/
/**
* Listens for multicast messages from the current assigned group.
**/
func (this *MulticastListener) Listen() {
	for {
		buffer := make([]byte, BUFFER_SIZE)
		n, received_addr, err := this.Socket.ReadFromUDP(buffer[0:])
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
func (this *MulticastListener) Multicast(message *Message) error{
	bArray,err := json.Marshal(message)
	_,err = this.Socket.WriteToUDP(bArray,this.GROUP_ADDRESS)
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
		this.Socket = conn
	}
	return err
}

/**
* Checks whether there was an error
**/
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
