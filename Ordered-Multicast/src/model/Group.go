package model


type Group struct {
	Name string 	`json:"name"`		// group name
	MaxPeers int 	`json:"maxPeers"`	// max amount of peers that can be connected
	Client Clients		`json:"clients"`		// peers connected
	Owner string 	`json:"owner"` // group leader
	Address  string 	`json:"address"` // group address
}

func NewGroup(name string, maxPeers int, owner string, address string) *Group {
	return &Group{Name: name, MaxPeers: maxPeers, Client: make(Clients), Owner: owner, Address: address}
}


