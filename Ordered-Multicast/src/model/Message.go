package model

type Message interface {
	HasAttachment() bool
	GetSender() string
	GetRecipient()string
}

//################################################################ UDP ######################################################################

type UdpMessage struct {  // TODO Refactor: create message struct
	senderAddr string `json:"senderAddr,omitempty"`
	recipientAddr string `json:"recipientAddr,omitempty"`
	attachment interface{} `json:"attachment"`
}

func NewUdpMessage(senderAddr string, recipientAddr string, attachment interface{}) *UdpMessage {
	return &UdpMessage{senderAddr: senderAddr, recipientAddr: recipientAddr, attachment: attachment}
}

func (this *UdpMessage) RecipientAddr() string {
	return this.recipientAddr
}

func (this *UdpMessage) SetRecipientAddr(recipientAddr string) {
	this.recipientAddr = recipientAddr
}

func (this *UdpMessage) SenderAddr() string {
	return this.senderAddr
}

func (this *UdpMessage) SetSenderAddr(senderAddr string) {
	this.senderAddr = senderAddr
}


func (this UdpMessage) HasAttachment() bool{
	return this.attachment != nil
}
//###############################################################################################################################################################################################################################