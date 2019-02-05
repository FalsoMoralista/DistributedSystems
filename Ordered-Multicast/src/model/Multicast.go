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
	connected bool
}

/**
*	By default uses the loopback interface on the socket.
*	It can be modified by using the `Connect´ method.
**/
func NewMulticastListener(process_id int,GROUP_ADDRESS *net.UDPAddr) *MulticastListener {
	iface , _ := net.InterfaceByName("lo")
	sock, err := net.ListenMulticastUDP("udp4", iface,GROUP_ADDRESS)
	checkError(err)
	return &MulticastListener{Socket:sock, GROUP_ADDRESS: GROUP_ADDRESS, Fifo_protocol:NewFifoOrder(process_id)}
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
	if err == nil{
		conn, err = net.ListenMulticastUDP("udp4", inf , this.GROUP_ADDRESS) // MULTICAST SOCKET
		this.Socket = conn
		this.connected = true
	}
	return err
}
/**
* Returns whether this peer is connected.
**/
func (this *MulticastListener) isConnected() bool{
	return this.connected
}
/***************************************************************************************************************************************************************************************************/
/**
*	Multicast a message through the assigned group.
**/
func (this *MulticastListener) Multicast(obj interface{}) error{
	var msg *Message = this.Fifo_protocol.Send(obj) // STEP THROUGH THE FIFO PROTOCOL
	bArray,err := json.Marshal(msg)
	fmt.Println("Peer: sending message",msg)
	_,err = this.Socket.WriteToUDP(bArray,this.GROUP_ADDRESS)
	return err
}
/**
* Listens for multicast messages from the current assigned group.
**/
func (this *MulticastListener) Listen() {
	if this.isConnected(){
		for {
			handle(this)
		}
	}
	fmt.Println("Error: peer not connected")
}

/**
* Handles received messages.
**/
func handle(this *MulticastListener){
	buffer := make([]byte, BUFFER_SIZE)
	n, received_addr, err := this.Socket.ReadFromUDP(buffer[0:]) // LISTEN FOR CONNECTIONS
	msg,err := decode(n,buffer) // DECODES A RECEIVED MESSAGE
	protocol(msg,this) // PROTOCOL
	fmt.Println("message:",msg)
	fmt.Println("Message received from " + received_addr.String())
	if err != nil {
		fmt.Print("Peer: Error, returning...")
		return
	}
}

/**
* Decodes received messages.
**/
func decode(n int, buff []byte) (*Message, error){
	var m *Message
	err :=json.Unmarshal(buff[0:n],m)
	return m,err
}

func protocol(msg *Message, m *MulticastListener){
	m.Fifo_protocol.Receive(msg)
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
