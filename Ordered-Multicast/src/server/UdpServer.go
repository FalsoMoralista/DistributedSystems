package server

import (
	"fmt"
	"net"
	"os"
)

type UdpServer struct {
	address string
	udpAddr *net.UDPAddr
}

func (u *UdpServer) UdpAddr() *net.UDPAddr {
	return u.udpAddr
}

func (u *UdpServer) SetUdpAddr(udpAddr *net.UDPAddr) {
	u.udpAddr = udpAddr
}

func (u *UdpServer) Address() string {
	return u.address
}

func (u *UdpServer) SetAddress(address string) {
	u.address = address
}

func NewUdpServer(address string) *UdpServer{
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	checkError(err) // check if there was an error
	return &UdpServer{
		address:address,
		udpAddr:parsedAddr,
	}
}

/**
* Starts the server
**/
func (this *UdpServer) Run(){
	fmt.Println("Server: Starting | Address->"+this.address+" |")
	conn, err := net.ListenUDP("udp",this.udpAddr) // starts listening to connections
	checkError(err)
	for{
		fmt.Println("Server: Listening...")
		handleClient(conn)
	}
}

/**
* Handle client connections
**/
func handleClient(conn *net.UDPConn){
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Print("Error, returning...")
		return
	}
	checkError(err)
	fmt.Println("Server: Message content:",string(buf[0:n]))
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
