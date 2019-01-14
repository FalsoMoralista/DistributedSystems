package model

import (
	"fmt"
	"net"
	"os"
)

type Multicast struct {
	serverAddr *net.UDPAddr // TODO remove from here, create a Message struct with sender and recipient addresses
	buffer [1024]byte
}

/**
* Returns a new *Multicast (work as a constructor)
*
**/
func NewMulticast(address string) * Multicast{
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	checkError(err) // check if there was an error

	return &Multicast{ // otherwise return a pointer to a Multicast
		serverAddr:parsedAddr,
	}
}

/**
* Returns the udp address for a Multicast instance
*
**/
func GetServerAddr(this *Multicast) *net.UDPAddr{
	return this.serverAddr
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

/**
*  Tests the connection with a server
**/
func TestConnection(this *Multicast)(*net.UDPConn, error){
	fmt.Println("Client: Testing connection")
	conn,err := net.DialUDP("udp",nil,this.serverAddr)
	checkError(err)
	int ,err := conn.Write([]byte("testing communication"))
	if(int == 0){
		fmt.Println("Error")
	}
	checkError(err)
	n,err := conn.Read(this.buffer[0:]) // TODO: Read documentation about read/send functions
	checkError(err)
	fmt.Println("Client: Message received from server:",string(this.buffer[0:n]))
	return conn,err
}

// TODO Create a file to define the protocol messages (as enums)
/**
*  Send a message through a udp socket
**/
func SendUdp(message UdpMessage){
	conn,err := net.DialUDP("udp",nil,message.GetRecipient())
	checkError(err)
	int ,err := conn.Write([]byte("MESSAGE")) // TODO: Verify how to convert from generic type into byte array
	if(int == 0){
		fmt.Println("Error")
	}
	checkError(err)
}