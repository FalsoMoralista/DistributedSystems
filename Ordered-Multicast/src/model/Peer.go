package model

import "distributed-systems/Ordered-Multicast/src/model/multicast"

type Peer struct {
	client *Client
	listener *multicast.MulticastListener
}

func (p *Peer) Listener() *multicast.MulticastListener {
	return p.listener
}

func (p *Peer) SetListener(listener *multicast.MulticastListener) {
	p.listener = listener
}

func (p *Peer) Client() *Client {
	return p.client
}

func (p *Peer) SetClient(client *Client) {
	p.client = client
}

func NewPeer(client *Client, listener *multicast.MulticastListener) *Peer {
	return &Peer{client: client, listener: listener}
} 


