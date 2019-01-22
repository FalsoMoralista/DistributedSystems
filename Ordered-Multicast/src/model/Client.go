package model


type Client struct {
	HostAddr string
}

func NewClient(hostAddr string) *Client {
	return &Client{HostAddr: hostAddr}
}

