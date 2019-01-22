package server

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_GROUPS int = 3
	SERVICE_PORT int = 7879
	BASE_ADDRESS string = "230.0.0.0"
	SERVER_ADDR string = "debian:1041"
)

type UdpServer struct {
	address string
	udpAddr *net.UDPAddr //
	groups model.Groups // table of groups
}

func NewUdpServer(address string) *UdpServer{
	parsedAddr,err := net.ResolveUDPAddr("udp4",address) // resolve the udp address
	checkError(err) // check if there was an error
	return &UdpServer{
		address:address,
		udpAddr:parsedAddr,
		groups:make(model.Groups),
	}
}

func (u *UdpServer) Groups() model.Groups {
	return u.groups
}

func (u *UdpServer) SetGroups(groups model.Groups) {
	u.groups = groups
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


/**
* Starts the server
**/
func (this *UdpServer) Run(){
	fmt.Println("Server: Starting | Address->"+this.address+" |")
	loadGroups(this)
	conn, err := net.ListenUDP("udp",this.udpAddr) // starts listening to connections
	checkError(err)
	for{
		fmt.Println("Server: Listening...")
		handleClient(conn)
	}
}

// Loads the groups that users will use to communicate
func loadGroups(this *UdpServer){ // TODO review & finish (initialize rooms & add to a map).
	i := 0
	//m = this.groups
	for i <= MAX_GROUPS {
		name := string("Group "+strconv.Itoa(i))
		address := string(strings.Replace(BASE_ADDRESS,"0",strconv.Itoa(i),1)+":"+strconv.Itoa(SERVICE_PORT))
		fmt.Println(name +" "+ address)
		//m[name]model.NewGroup(name,10,"no owner yet",string())
		i++
	}
	//fmt.Print("...]")
}


/**
* Handle client connections
**/
func handleClient(conn *net.UDPConn){
	var buf [util.BUFFER_SIZE]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Print("Error, returning...")
		return
	}
	checkError(err)
	fmt.Println("Server: Message content:",string(buf[0:n]))
	parse(buf)
	conn.WriteToUDP([]byte("ack"), addr)
}

func parse(buf []byte){

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
