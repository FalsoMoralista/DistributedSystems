package controller

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"net"
)

const (
	SERVER_ADDRESS string = "debian:1041"
)

type Controller struct {
	peer *model.Peer
}


func NewController(client_address string) *Controller {
	p := model.NewPeer(client_address,nil)
	return &Controller{peer:p}
}

func (c *Controller) Peer() *model.Peer {
	return c.peer
}

func (c *Controller) SetPeer(peer *model.Peer) {
	c.peer = peer
}

func (this *Controller) AssignGroupAddress(addr *net.UDPAddr){
	this.peer.Listener.AssignGroupAddress(addr)
}

func (this *Controller) ConnectPeer(iface string) error{
	return this.peer.Listener.Connect(iface)
}

/*********************************************************************************************************************************************************************************************/
/*
* Do a request based on its type and return back a response message or an error.
*/
func (this *Controller) Request(TYPE string) (*model.Message, error){
	switch TYPE {
		case util.GROUP:
			request := model.NewMessage(0,this.peer.HostAddr,SERVER_ADDRESS,util.REQUEST,util.GROUP,nil)
			response := model.Message{}
			n,buffer,err := util.SendUdp(SERVER_ADDRESS,request)
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