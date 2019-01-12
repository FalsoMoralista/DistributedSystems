package model

import (
	"fmt"
	"net"
	"os"
)

type Multicast struct {
	serverAddr *net.UDPAddr

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

func StartConnection(this *Multicast)(*net.UDPConn, error){
	conn,err := net.DialUDP("udp",nil,this.serverAddr)
	checkError(err)
	int ,err := conn.Write([]byte("init"))
	if(int == 0){
		fmt.Print("Connection error")
	}
	checkError(err)
	var buffer[512]byte
	n,err := conn.Read(buffer[0:])
	checkError(err)
	fmt.Print(string(buffer[0:n]))
	return conn,err
}
