package model


// This is the structure of a Message that will be exchanged between peers on the network
type Message struct {
	Seq uint  `json:"seq"`// Sequence number
	SenderAddr string `json:"senderAddr"` // Sender address
	RecipientAddr string `json:"recipientAddr"`// Recipient address
	Attachment interface{} `json:"attachment,omitempty"` // Payload
}

func NewMessage(seq uint, senderAddr string, recipientAddr string, attachment interface{}) *Message {
	return &Message{Seq: seq, SenderAddr: senderAddr, RecipientAddr: recipientAddr, Attachment: attachment}
}

func (this *Message) HasAttachment() bool{
	return this.Attachment != nil
}
