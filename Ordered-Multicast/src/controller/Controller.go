package controller

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/server"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"net"
)

const (
	CLIENT_ADDR string = "debian:1041"
	SERVER_ADDRESS string = "debian:1041"
)

type Controller struct {
	peer *model.Peer
}


func NewController(client_address string) *Controller {
	m :=  model.NewMulticastListener(nil)
	p := model.NewPeer(client_address,m)
	return &Controller{peer:p}
}

func (c *Controller) Peer() *model.Peer {
	return c.peer
}

func (c *Controller) SetPeer(peer *model.Peer) {
	c.peer = peer
}

func (this *Controller) AssignGroupAddress(addr *net.UDPAddr){
	this.peer.Listener().AssignGroupAddress(addr)
}

func (this *Controller) ConnectToGroup(iface string) error{
	return this.Peer().Listener().Connect(iface)
}

// *************************************************************************************************************************************************************************************************************
// *************************************************************************************************************************************************************************************************************
// *************************************************************************************************************************************************************************************************************

/*
* Do a request based on its type and return back a response message or an error.
*/
func (this *Controller) Request(TYPE string) (*model.Message, error){
	switch TYPE {
		case util.GROUP:
			request := model.NewMessage(0,this.peer.HostAddr,SERVER_ADDRESS,util.REQUEST,util.GROUP,this.peer) // this.peer marshaling error breaks the application
			response := model.Message{}
			n,buffer,err := util.SendUdp(server.SERVER_ADDR,request)
			if checkError(err){
				return nil,err
			}
			err = json.Unmarshal(buffer[0:n],&response)
			return &response,err
	}
	return nil, nil
}


/**
* Parse a message making appropriate conversions (if necessary) and returns the message payload or an error.
**/
func (this *Controller) Parse(message *model.Message) (interface{},error){
	switch message.Header {  // CHECKS THE MESSAGE HEADER
	case util.RESPONSE: // WHETHER IS A RESPONSE
	switch message.Type { // CHECKS THE RESPECTIVE TYPE
		case util.GROUP: // WHETHER IS A GROUP
			b, err := json.Marshal(message.Attachment) // ENCODE THE ATTACHMENT
			var g model.Group
			err = json.Unmarshal(b,&g) // THEN DECODE IT IN ORDER TO CONVERT
			if err != nil {
				return nil,err
			}
			return g,nil // THEN RETURN IT TO THE USER
		}
	}
	return nil, nil
}

func checkError(err error)  bool{
	if(err != nil){
		return true
	}
	return false
}