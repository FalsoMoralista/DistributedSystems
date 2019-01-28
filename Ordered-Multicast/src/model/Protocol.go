package model

const (
	MAX_PEERS int = 4
)

type FifoOrder struct {
	PROCESS_ID int `json:"PROCESS_ID,omitempty"`
	seq int `json:"seq,omitempty"`
	message_sequences [MAX_PEERS]int `json:"message_sequences,omitempty"` // This is where the messages sequences will be stored
	buffer map[string]*Message `json:"buffer,omitempty"` // This is where the processes delayed messages will be stored
}


/**
* Storage a message through the buffer
**/
func (this *FifoOrder) Buffer(message *Message){
	//this.buffer[message.SenderAddr] = message // TODO test
}

func (this *FifoOrder) Receive(message *Message){

}

