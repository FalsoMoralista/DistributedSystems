package server

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_GROUPS int = 3
	SERVICE_PORT int = 1041
	MULTICAST_PORT int = 7879
	BASE_ADDRESS string = "230.0.0.0"
	SERVER_ADDR string = "debian:1041"
)

type UdpServer struct {
	address string
	udpAddr *net.UDPAddr //
	groups model.Groups // table of groups
	next_group int
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

//#####################################################################################################################################################################################################################################

/**
* Starts the server
**/
func (this *UdpServer) Run(){
	fmt.Println("Server: Starting | Address->"+this.address+" |")
	this.loadGroups()
	conn, err := net.ListenUDP("udp",this.udpAddr) // starts listening to connections
	checkError(err)
	for{
		fmt.Println("Server: Listening...")
		this.handleClient(conn)
	}
}

// Loads the groups that users will use to communicate and add them to a map
func (this *UdpServer) loadGroups(){
	i := 0
	//m = this.groups
	for i <= MAX_GROUPS {
		itr := strconv.Itoa(i)
		name := string("Group "+itr)
		address := string(strings.Replace(BASE_ADDRESS,"0",itr,1)+":"+strconv.Itoa(MULTICAST_PORT))
		g := model.NewGroup(strconv.Itoa(i),address)
		this.groups[itr] = g
		fmt.Println(name +" "+ address)
		i++
	}
}


/**
* Handle client connections
**/
func (this *UdpServer) handleClient(conn *net.UDPConn){ // todo comment
	buf := make([]byte,util.BUFFER_SIZE)
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Print("Error, returning...")
		return
	}
	checkError(err)
	var msg *model.Message
	msg,err = this.parse(buf,n)
	if err != nil{
		msg = model.NewMessage(0,"",SERVER_ADDR,util.RESPONSE,util.ERROR,nil)
	}
	throwback,err := json.Marshal(msg)
	conn.WriteToUDP(throwback, addr)
}

// Parse client messages
func(this *UdpServer) parse(buf []byte , to int) (*model.Message, error){
	fmt.Println("Server: Message arrived")
	fmt.Println("Server: Message content: "+string(buf))
	msg := model.Message{}
	err := json.Unmarshal(buf[0:to],&msg)
	fmt.Println(err)
	if(err != nil){
		return &model.Message{},err
	}
	switch msg.Header {
		case util.REQUEST:
			switch msg.Type {
				case util.GROUP:// TODO comment

					var usrInfo  = msg.Attachment.(map[string]interface {})
					hostAddr := usrInfo["hostAddr"].(string)
					usr := model.NewClient(hostAddr)
					id := strconv.Itoa(this.next_group)
					group := this.groups[id]
					group.Leader = *usr
					group.Clients[usr.HostAddr] = usr
					this.next_group = 1 // todo review
					return model.NewMessage(0,usr.HostAddr,SERVER_ADDR,util.RESPONSE,util.GROUP,group), nil
			}
	}
	return &model.Message{},err
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
