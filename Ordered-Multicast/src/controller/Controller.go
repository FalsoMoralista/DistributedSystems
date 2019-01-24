package controller

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/server"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
)

const (
	CLIENT_ADDR string = "debian:1041"
)

type Controller struct {
	client *model.Client
	client_address string
	server_address string
}

func NewController(client *model.Client, client_address string) *Controller {
	return &Controller{client: client, client_address: client_address}
}

func (c *Controller) Client_address() string {
	return c.client_address
}

func (c *Controller) SetClient_address(client_address string) {
	c.client_address = client_address
}

func (c *Controller) Client() *model.Client {
	return c.client
}

func (c *Controller) SetClient(client *model.Client) {
	c.client = client
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
			request := model.NewMessage(0,this.client.HostAddr,server.SERVER_ADDR,util.REQUEST,util.GROUP,this.client)
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