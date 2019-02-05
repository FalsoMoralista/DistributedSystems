package model

import "strconv"

const (
	MAX_PEERS int = 4
)

type FifoOrder struct {
	PROCESS_ID int `json:"PROCESS_ID"`
	Current_seq int `json:"current_seq"`
	Processes_sequences [MAX_PEERS]int `json:"message_sequences"` // This is where the messages sequences will be stored
	Buff map[string]*Message `json:"buffer,omitempty"` // This is where the processes delayed messages will be stored // TODO fix
}

func NewFifoOrder(PROCESS_ID int) *FifoOrder {
	return &FifoOrder{PROCESS_ID: PROCESS_ID, Current_seq:0, Processes_sequences: [MAX_PEERS]int{}, Buff:make(map[string]*Message)}
}

/**
* Storage a message through the buffer
**/
func (this *FifoOrder) Buffer(message *Message){
	//this.buffer[message.SenderAddr] = message // TODO test
}

/**
* Returns whether a message can be delivered to the application
**/
func (this *FifoOrder) Receive(message *Message) bool{
	var process_id ,_ = strconv.Atoi(message.SenderAddr) // PARSES THE SENDER PROCESS`S ID
	var deliver bool = this.Processes_sequences[process_id] == (message.Seq + 1) // VERIFY WHETHER THE SEQUENCE NUMBER (FROM MESSAGE) ITS EQUAL
	if !deliver { // IF IS DIFFERENT
		if !(this.Processes_sequences[process_id] >= message.Seq) { // AND NOT MINOR THAN THE ACTUAL SEQUENCE
			this.Buffer(message) // BUFFER UNTIL IT`S TRUE
		}
	}else{ // OTHERWISE: INCREMENT ITS SEQUENCE AND DELIVER
		this.Processes_sequences[process_id] += 1
	}
	return deliver
}

func (this *FifoOrder) Send(obj interface{}) *Message{
	this.Current_seq += 1
	this.Processes_sequences[this.PROCESS_ID] = this.Current_seq
	var id string = strconv.Itoa(this.PROCESS_ID)
	var msg *Message = NewMessage(this.Current_seq,id,"","","",obj)
	return msg
}

