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
	BUFFER_SIZE int = 4096
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
		udpAddr,_:= net.ResolveUDPAddr("udp4",address)
		g := model.NewGroup(strconv.Itoa(i),udpAddr) // INITIALIZE A GROUP
		this.groups[itr] = g // PUT IT IN THE MAP
		fmt.Println(name +" "+ address)
		i++
	}
}

/**
* Handle client connections.
**/
func (this *UdpServer) handleClient(conn *net.UDPConn){
	buf := make([]byte,BUFFER_SIZE) // INITIALIZE THE BUFFER
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
		return model.NewMessage(0,SERVER_ADDR, "", util.ERROR, util.RESPONSE,err), err // TODO reduce to 1 line
	}
	switch msg.Header { // VERIFY THE MESSAGE HEADER
		case util.REQUEST: // IN CASE OF A REQUEST FROM TYPE
			switch msg.Type {
				case util.GROUP:// GROUP
					group := this.groups[strconv.Itoa(this.next_group)] // GET THE REQUESTED GROUP`S POINTER
					if len(group.Peers) < 4{ // (DEFINES A LIMIT TO THE GROUP SIZE)
						mListener := model.NewMulticastListener(len(group.Peers)+1, group.Address)
						peer := model.NewPeer(msg.SenderAddr, mListener) // PARSE THE PEER RECEIVED FROM THE MESSAGE, RESETTING NETWORK SETTINGS
						if group.Leader.HostAddr == "" && len(group.Peers) == 0 { // IF THE CURRENT GROUP HAS NO ACTIVE LEADER
							group.Leader = *peer // MAKE THE PEER THE GROUP LEADER
						}
						group.Peers[peer.HostAddr] = peer // THEN REGISTER THE PEER IN THE GROUP
						//this.next_group += 1 // todo review (next user will be conected to the next available group with this line)
						fmt.Println("Server: Client group request received, retrieving...")
						return model.NewMessage(0, SERVER_ADDR, peer.HostAddr, util.RESPONSE, util.GROUP, *group), nil // RETURNS THE GROUP TO THE USER
					}
					return model.NewMessage(0, SERVER_ADDR, msg.SenderAddr, util.RESPONSE, util.ERROR, nil), nil // RETURNS AN ERROR (FULL GROUP)
			}
	}
	return model.NewMessage(0,SERVER_ADDR, msg.SenderAddr , util.ERROR, util.RESPONSE,err), err
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
