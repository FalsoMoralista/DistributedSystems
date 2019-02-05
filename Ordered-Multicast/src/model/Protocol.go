package model

import (
	"fmt"
	"strconv"
)

const (
	MAX_PEERS int = 4
)

type FifoOrder struct {
	PROCESS_ID int `json:"PROCESS_ID"`
	Current_seq int `json:"current_seq"`
	Processes_sequences [MAX_PEERS]int `json:"message_sequences"` // This is where the messages sequences will be stored
	Buff Messages `json:"buffer,omitempty"` // This is where the processes delayed messages will be stored
}

func NewFifoOrder(PROCESS_ID int) *FifoOrder {
	return &FifoOrder{PROCESS_ID: PROCESS_ID, Current_seq:0, Processes_sequences: [MAX_PEERS]int{}, Buff:make(Messages)}
}

/**
* Storage a message through the buffer
**/
func (this *FifoOrder) Buffer(message *Message){
	var senderId string = message.SenderAddr // GET THE SENDER ID
	var seq int = message.Seq // "" "" MESSAGE SEQUENCE
	var senderMessages map[int]*Message = this.Buff[senderId] // GETS THE MESSAGES MAP FROM THE PEER
	fmt.Println("mensagem a ser bufferizada ->",message)
	fmt.Println("buffer status:",senderMessages[seq]) // Breaking application
	senderMessages[seq] = message // BUFFER IT todo verify why does it is nil
}

/**
* Returns whether a message can be delivered to the application
**/
func (this *FifoOrder) Receive(message *Message) bool{
	var process_id ,_ = strconv.Atoi(message.SenderAddr) // PARSES THE SENDER PROCESS`S ID
	var deliver bool = this.Processes_sequences[process_id] == (message.Seq + 1) // VERIFY WHETHER THE SEQUENCE NUMBER (FROM MESSAGE) ITS EQUAL
	if !deliver { // IF IS DIFFERENT
		fmt.Println("entrou auqui")
		fmt.Println("current sequence --->",this.Current_seq)
		if !(this.Processes_sequences[process_id] >= message.Seq) { // AND NOT MINOR THAN THE ACTUAL SEQUENCE
			fmt.Print("buffer current state -> ")
			fmt.Println(this.Buff)
			this.Buffer(message) // BUFFER UNTIL IT`S TRUE
		}
	}else{ // OTHERWISE: INCREMENT ITS SEQUENCE AND DELIVER
		this.Processes_sequences[process_id] += 1
	}
	return deliver
}

/**
* Register the send of a message in the protocol.
**/
func (this *FifoOrder) Send(obj interface{}) *Message{
	this.Current_seq += 1 // INCREMENT THE SEQUENCER
	this.Processes_sequences[this.PROCESS_ID] = this.Current_seq // REGISTER THE CURRENT SEQUENCE FOR THIS PEER IN THE LOGIC CLOCK
	var id string = strconv.Itoa(this.PROCESS_ID) // PARSE THE PEER ADDRESS
	var msg *Message = NewMessage(this.Current_seq,id,"","","",obj) // RETURNS A MESSAGE WITH THE INFO ABOVE
	return msg
}

