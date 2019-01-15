package model

type Udp struct {
	client *Client
	buffer [1024]byte
}

func (this *Udp) Client() *Client {
	return this.client
}

func (u *Udp) SetClient(client *Client) {
	u.client = client
}

func NewUdp(client *Client) *Udp {
	return &Udp{client: client}
}

