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
	return &FifoOrder{PROCESS_ID: PROCESS_ID, Current_seq:0, Processes_sequences: [MAX_PEERS]int{}, Buff: make(Messages)}
}

/**
* Storage a message through the buffer. // todo implement the buffer cleaner
**/
func (this *FifoOrder) Buffer(message *Message){
	var senderId string = message.SenderAddr // GET THE SENDER ID
	var seq int = message.Seq // "" "" MESSAGE SEQUENCE
	if this.Buff == nil { // CHECK WHETHER THE BUFFER IS EMPTY todo verificar se há chance  do buffer ser resetado com mensagens dentro
		this.Buff = make(map[string]map[int]*Message)
		if this.Buff[senderId] == nil {
			this.Buff[senderId] = make(map[int]*Message)
		}
	}
	var senderMessages map[int]*Message = this.Buff[senderId] // GETS THE MESSAGES MAP FROM THE PEER
	senderMessages[seq] = message // BUFFER IT
}

/**
* Returns whether a message can be delivered to the application.
**/
func (this *FifoOrder) Receive(message *Message) bool{
	var process_id ,_ = strconv.Atoi(message.SenderAddr) // PARSES THE SENDER PROCESS`S ID
	var deliver bool = this.Processes_sequences[process_id] == (message.Seq + 1) // VERIFY WHETHER THE SEQUENCE NUMBER (FROM MESSAGE) ITS EQUAL
	if !deliver { // IF IS DIFFERENT
		fmt.Println("current sequence --->",this.Current_seq)
		if !(this.Processes_sequences[process_id] >= message.Seq) { // AND NOT MINOR THAN THE ACTUAL SEQUENCE
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
func (this *FifoOrder) Send() chan error{
	channel := make(chan error)
	go func() {
		fmt.Println("Esperando por mensagem de confirmação...")
		var err = <- channel
		if err == nil{
			fmt.Println("Registrando mensagem no protocolo")
			this.Current_seq += 1 // INCREMENT THE SEQUENCER
			this.Processes_sequences[this.PROCESS_ID] = this.Current_seq // REGISTER THE CURRENT SEQUENCE FOR THIS PEER IN THE LOGIC CLOCK
			channel <- nil // todo verify
		}else{
			fmt.Println("Erro na transmissão de mensagem")
		}
	}()
	return channel
}

