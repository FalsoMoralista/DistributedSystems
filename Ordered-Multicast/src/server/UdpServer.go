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
	conn, err := net.ListenUDP("udp",this.udpAddr) // STARTS LISTENING TO CONNECTIONS
	checkError(err)
	fmt.Println("Server: Listening...")
	for{
		this.handleClient(conn)
	}
}

/*
* Loads the groups that users will use to communicate and add them to a map.
*/
func (this *UdpServer) loadGroups(){
	i := 0
	for i <= MAX_GROUPS {
		itr := strconv.Itoa(i)
		name := string("Group "+itr)
		address := string(strings.Replace(BASE_ADDRESS,"0",itr,1)+":"+strconv.Itoa(MULTICAST_PORT)) // PARSE AN ADDRESS
		g := model.NewGroup(strconv.Itoa(i),address) // INITIALIZE A GROUP
		this.groups[itr] = g // PUT IT IN THE MAP
		fmt.Println(name +" "+ address)
		i++
	}
}

/**
* Handle client connections.
**/
func (this *UdpServer) handleClient(conn *net.UDPConn){
	buf := make([]byte,util.BUFFER_SIZE) // INITIALIZE THE BUFFER
	n, addr, err := conn.ReadFromUDP(buf[0:]) // READ IT
	if err != nil {
		fmt.Print("Server: Error, returning...")
		return
	}
	var msg *model.Message
	msg,err = this.parse(buf,n) // PARSES THE MESSAGE
	throwback,err := json.Marshal(msg)
	conn.WriteToUDP(throwback, addr) // RETURN A RESPONSE TO THE CLIENT
}

/*
*	Parses client messages from a buffer, returns either a *model.Message or an error.
*	TODO Implement logs
*/
func(this *UdpServer) parse(buf []byte , to int) (*model.Message, error){
	fmt.Print("Server: Message arrived,")
	fmt.Println(" content: "+string(buf[0:to]))
	msg := model.Message{}
	err := json.Unmarshal(buf[0:to],&msg) // DECODE THE MESSAGE
	if(err != nil){ // IN CASE OF ERROR, SEND BACK AN ERROR MESSAGE
		return model.NewMessage(0,"", SERVER_ADDR, util.ERROR, util.RESPONSE,nil), err
	}
	switch msg.Header { // VERIFY THE MESSAGE HEADER
		case util.REQUEST: // IN CASE OF A REQUEST FROM TYPE
			switch msg.Type {
				case util.GROUP:// GROUP
					var usrInfo  = msg.Attachment.(map[string]interface {}) // DO THE ATTACHMENT CONVERSION
					hostAddr := usrInfo["hostAddr"].(string)
					usr := model.NewClient(hostAddr) // PARSE THE USER RECEIVED FROM THE MESSAGE
					id := strconv.Itoa(this.next_group)
					group := this.groups[id] // GET THE REQUESTED GROUP
					group.Leader = *usr // MAKE THE USER THE GROUP LEADER
					group.Clients[usr.HostAddr] = usr // INSERT IN THE MAP
					this.next_group = 1 // todo review
					return model.NewMessage(0,usr.HostAddr,SERVER_ADDR,util.RESPONSE,util.GROUP,group), nil // RETURNS THE GROUP TO THE USER
			}
	}
	return model.NewMessage(0,"", SERVER_ADDR, util.ERROR, util.RESPONSE,nil), err
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
