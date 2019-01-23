package model


type Client struct {
	HostAddr string `json:"hostAddr"`
}

func NewClient(hostAddr string) *Client {
	return &Client{HostAddr: hostAddr}
}

