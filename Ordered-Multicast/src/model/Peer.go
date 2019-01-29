package model


type Peer struct {
	HostAddr string `json:"HostAddr"`
	Listener *MulticastListener `json:"listener"`
}

func NewPeer(hostAddr string, listener *MulticastListener) *Peer {
	return &Peer{HostAddr:hostAddr, Listener: listener}
}


