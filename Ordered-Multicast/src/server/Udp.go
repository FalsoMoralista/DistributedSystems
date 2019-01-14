package server

import (
	"fmt"
	"net"
	"os"
)

type UdpServer struct {
	address *net.UDPAddr
}

/**
* Returns the address from a given server
**/
func GetAddress(this *UdpServer) *net.UDPAddr{
	return this.address
}

/**
* Constructor
**/
func NewUdpServer(address string) *UdpServer{
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	checkError(err) // check if there was an error
	return &UdpServer{
		address:parsedAddr,
	}
}

/**
* Starts the server
**/
func Run(this *UdpServer){
	fmt.Println("Server: Starting | Address :"+this.address.String()+" listening...")
	conn, err := net.ListenUDP("udp",this.address) // allocates a connection
	checkError(err)
	for{
		fmt.Println("Server: Done")
		handleClient(conn)
	}
}

/**
* Handle client connections
**/
func handleClient(conn *net.UDPConn){
	fmt.Println("Server: Message arrived")
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Print("Error, returning...")
		return
	}
	checkError(err)
	fmt.Println("Server: Content:",string(buf[0:n]))
	conn.WriteToUDP([]byte("ack"), addr)
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
