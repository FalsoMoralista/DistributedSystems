package model

import (
	"fmt"
	"strconv"
	"time"
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
* Storage a message through the buffer.
**/
func (this *FifoOrder) Buffer(message *Message){
	fmt.Printf("Buffering message (ID = %v) sent by: Process %v\n",message.Seq,message.SenderAddr)
	time.Sleep(time.Second * 2)
	var senderId string = message.SenderAddr // GET THE SENDER ID
	var seq int = message.Seq // "" "" MESSAGE SEQUENCE
	if this.Buff == nil { // CHECK WHETHER THE BUFFER HAS NOT BEEN INITIALIZED
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
func (this *FifoOrder) Receive(message *Message) chan *Queue{
	var channel chan *Queue = make(chan *Queue)
	go func() {
		var process_id ,_ = strconv.Atoi(message.SenderAddr) // PARSES THE SENDER PROCESS`S ID
		var deliver bool = (this.Processes_sequences[process_id] + 1) == message.Seq // VERIFY WHETHER THE MESSAGE SEQUENCE NUMBER IS EQUAL
		if !deliver { // IF IS DIFFERENT
			time.Sleep(time.Second * 2)
			if (this.Processes_sequences[process_id] + 1 < message.Seq) { // AND NOT MINOR THAN THE ACTUAL SEQUENCE
				this.Buffer(message) // BUFFER UNTIL IT`S TRUE
				channel <- nil //
			}
		}else{ // OTHERWISE: INCREMENT ITS SEQUENCE AND DELIVER
			this.Processes_sequences[process_id] += 1
			var cleaned_buffer *Queue = this.cleanBuffer(message)
			channel <- cleaned_buffer
		}
	}()
	return channel
}

/**
* Register the send of a message in the protocol.
* First it instances the channel that will be delivered back to this function caller, then it calls a "thread" to run
* an anonymous function which soon get locked waiting for an error message from the channel. When the message arrives,
* check if no error has occurred, if so, register the sending of the message, then return back the error message whether
* it is empty or not.
**/ // TODO refactor: the message has to be registered before
func (this *FifoOrder) Send() chan error{
	channel := make(chan error)
	go func() {
		var err = <- channel // STARTS TO WAIT FOR AN ERROR MESSAGE
		if err == nil{ // IF IT IS NIL,
			this.Current_seq += 1 // INCREMENT THE SEQUENCER
			this.Processes_sequences[this.PROCESS_ID] = this.Current_seq // REGISTER THE CURRENT SEQUENCE FOR THIS PEER IN THE LOGIC CLOCK
			channel <- err // SENDS BACK THE EMPTY ERROR MESSAGE
		}else{ // OTHERWISE SENDS BACK THE ERROR MESSAGE AND UNDO THE REGISTER OPERATION
			this.Current_seq -= 1 // INCREMENT THE SEQUENCER
			this.Processes_sequences[this.PROCESS_ID] = this.Current_seq // REGISTER THE CURRENT SEQUENCE FOR THIS PEER IN THE LOGIC CLOCK
			channel <- err
		}
	}()
	return channel
}

func ( this *FifoOrder) cleanBuffer(auth *Message) *Queue{
	var queue *Queue = NewQueue()
	queue.Add(auth)
	var current_seq int = auth.Seq
	usr_buffer := this.Buff[auth.SenderAddr]
	for usr_buffer[current_seq + 1] != nil {
			queue.Add(usr_buffer[current_seq +1])
			current_seq ++
	}
	var id,_ = strconv.Atoi(auth.SenderAddr)
	this.Processes_sequences[id] = current_seq
	return queue
}

