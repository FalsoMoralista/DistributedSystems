package model

import "net"

type Message interface {
	HasAttachment() bool
	GetSender()*net.UDPAddr
	GetRecipient()*net.UDPAddr
}

// ################ UDP ################

type UdpMessage struct {
	sender *net.UDPAddr
	recipient *net.UDPAddr
	attachment interface{} // (generic object)
}

func NewUdpMessage(senderAddr *net.UDPAddr, recipientAddr *net.UDPAddr, obj interface{})*UdpMessage{
	return &UdpMessage{
		sender:senderAddr,
		recipient:recipientAddr,
		attachment:obj,
	}
}

func (this UdpMessage) HasAttatchment() bool{
	return this.attachment != nil
}

func (this UdpMessage) GetRecipient() *net.UDPAddr{
	return 	this.recipient
}

func (this UdpMessage) GetSender() *net.UDPAddr{
	return 	this.sender
}
//###############################################################################################################################################################################################################################