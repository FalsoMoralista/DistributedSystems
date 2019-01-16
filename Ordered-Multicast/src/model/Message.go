package model


// This is the structure of a Message that will be exchanged between peers on the network
type Message struct {
	seq uint // Sequence number
	senderAddr string // Sender address
	recipientAddr string // Recipient address
	attachment interface{} // Payload
}

func NewMessage(seq uint, senderAddr string, recipientAddr string, attachment interface{}) *Message {
	return &Message{seq: seq, senderAddr: senderAddr, recipientAddr: recipientAddr, attachment: attachment}
}

func (this *Message) Attachment() interface{} {
	return this.attachment
}

func (this *Message) SetAttachment(attachment interface{}) {
	this.attachment = attachment
}

func (this *Message) RecipientAddr() string {
	return this.recipientAddr
}

func (this *Message) SetRecipientAddr(recipientAddr string) {
	this.recipientAddr = recipientAddr
}

func (this *Message) SenderAddr() string {
	return this.senderAddr
}

func (this *Message) SetSenderAddr(senderAddr string) {
	this.senderAddr = senderAddr
}

func (this *Message) Seq() uint {
	return this.seq
}

func (this *Message) SetSeq(seq uint) {
	this.seq = seq
}

func (this *Message) HasAttachment() bool{
	return this.attachment != nil
}
