package model

import "net"

type Group struct {
	ID string 	`json:"id"` // group name
	Address  *net.UDPAddr 	`json:"address"` // group address
	Peers Peers `json:"peers"` // peers connected
	Leader Peer 	`json:"leader"` // group leader
}

func NewGroup(groupID string, groupAddress *net.UDPAddr) *Group {
	return &Group{ID: groupID, Peers : make(Peers), Address: groupAddress}
}

