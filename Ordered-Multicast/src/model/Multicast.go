package model

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	BUFFER_SIZE int = 4096
	LOCALHOST string = "localhost:1041"
)

type MulticastListener struct {
	Socket *net.UDPConn `json:"socket,omitempty"`
 	GROUP_ADDRESS *net.UDPAddr `json:"group_address,omitempty"`
	Fifo_protocol *FifoOrder `json:"fifo_protocol"` // TODO IMPLEMENT COMMUNICATION BETWEEN MULTICAST AND PROTOCOL USING CHANNELS
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
*	Multicast a message through the assigned group.
*	Tries to multicast a message and then register the event synchronously in the protocol.
**/
func (this *MulticastListener) Multicast(messsage *Message) error{
	fmt.Printf("%s: Trying to multicast: ",messsage.SenderAddr)
	time.Sleep(time.Second * 1)
	var channel chan error = this.Fifo_protocol.Send()
	bArray,err := json.Marshal(messsage)
	fmt.Printf("message -> %s \n", string(bArray))
	time.Sleep(time.Second * 2)
	_,err = this.Socket.WriteToUDP(bArray,this.GROUP_ADDRESS)
	channel <- err
	if <- channel == nil{
		fmt.Printf("%s: Mensagem enviada com sucesso\n.",messsage.SenderAddr)
		time.Sleep(time.Second * 2)
	}
	return err
}

/**
* Handles received messages.
**/
func handle(this *MulticastListener){
	buffer := make([]byte, BUFFER_SIZE)
	n, _, err := this.Socket.ReadFromUDP(buffer[0:]) // LISTEN FOR CONNECTIONS
	msg,err := decode(n,buffer) // DECODES A RECEIVED MESSAGE
	if  !(msg.SenderAddr == strconv.Itoa(this.Fifo_protocol.PROCESS_ID)) { // VERIFY WHETHER THE MESSAGE IS FROM THE OWN PROCESS
		fmt.Printf("%d: Mensagem recebida \n", this.Fifo_protocol.PROCESS_ID)
		fmt.Printf("%d: Sender id-> %s\n", this.Fifo_protocol.PROCESS_ID, msg.SenderAddr)
		protocol(msg,this) // PROTOCOL
	}
	if err != nil {
		fmt.Println("Peer: Error, returning...")
		return
	}
}

/**
* Decodes received messages.
**/
func decode(n int, buff []byte) (*Message, error){
	var m = Message{}
	err :=json.Unmarshal(buff[0:n],&m)
	return &m,err
}

func protocol(msg *Message, m *MulticastListener){
	var channel chan bool = m.Fifo_protocol.Receive(msg) // CHECKS WHETHER THE PROTOCOL AUTHORIZES THE DELIVERY OF THIS MESSAGE TO THE APPLICATION
	var deliver bool = <- channel
	fmt.Println("Can the message",msg.Seq,"from "+msg.SenderAddr+" be delivered to the Appplication? ->",deliver)
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
