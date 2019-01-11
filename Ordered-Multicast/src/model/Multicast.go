package model

import (
	"fmt"
	"net"
	"os"
)

type Multicast struct {
	udpAddr *net.UDPAddr

}

/**
* Returns a new *Multicast (work as a constructor)
*
**/
func NewMulticast(address string) * Multicast{

	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	checkError(err) // check if there was an error

	return &Multicast{ // otherwise return a pointer to a Multicast
		udpAddr:parsedAddr,
	}
}

/**
* Returns the udp address for a Multicast instance
*
**/
func GetUdpAddr(this *Multicast) *net.UDPAddr{
	return this.udpAddr
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
