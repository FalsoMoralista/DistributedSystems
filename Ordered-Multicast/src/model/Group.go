package model


type Group struct {
	ID string 	`json:"id"`		// group name
	Address  string 	`json:"address"` // group address
	Clients Clients		`json:"clients"`		// peers connected
	Leader Client 	`json:"leader"` // group leader
}

func NewGroup(id string, address string) *Group {
	return &Group{ID: id, Clients: make(Clients), Address: address}
}

