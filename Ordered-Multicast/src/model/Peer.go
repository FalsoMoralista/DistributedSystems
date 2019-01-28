package model


type Peer struct {
	HostAddr string `json:"hostAddr"`
	listener *MulticastListener `json:"listener,omitempty"`
}

func NewPeer(hostAddr string, listener *MulticastListener) *Peer {
	return &Peer{HostAddr:hostAddr, listener: listener}
}

func (p *Peer) Listener() *MulticastListener {
	return p.listener
}

func (p *Peer) SetListener(listener *MulticastListener) {
	p.listener = listener
}
