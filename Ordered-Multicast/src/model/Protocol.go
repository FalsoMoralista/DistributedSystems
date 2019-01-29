package model

const (
	MAX_PEERS int = 4
)

type FifoOrder struct {
	PROCESS_ID int `json:"PROCESS_ID"`
	Current_seq int `json:"current_seq"`
	Processes_sequences [MAX_PEERS]int `json:"message_sequences"` // This is where the messages sequences will be stored
	Buff map[string]*Message `json:"buffer,omitempty"` // This is where the processes delayed messages will be stored
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

func (this *FifoOrder) Receive(message *Message){

}

