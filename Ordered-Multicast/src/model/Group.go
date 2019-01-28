package model


type Group struct {
	ID string 	`json:"id"` // group name
	Address  string 	`json:"address"` // group address
	Peers Peers `json:"peers"` // peers connected
	Leader Peer 	`json:"leader"` // group leader
}

func NewGroup(groupID string, groupAddress string) *Group {
	return &Group{ID: groupID, Peers : make(Peers), Address: groupAddress}
}

