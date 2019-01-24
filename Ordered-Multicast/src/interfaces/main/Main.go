package main

import (
	"bufio"
	"distributed-systems/Ordered-Multicast/src/controller"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Application struct {
	reader *bufio.Reader
	client *model.Client
}

const (
	CLIENT_ADDR string = "debian:1041"
)

func NewApplication() *Application {
	return &Application{reader:bufio.NewReader(os.Stdin) }
}


func (this *Application) readKeyboard() (string,error){
	str, err := this.reader.ReadString('\n')
	strings.Replace(str,"\n","",-1) // todo review replace (" \n ")
	return str,err
}

func (this *Application) Run()  {
	this.client = model.NewClient(CLIENT_ADDR)
	fmt.Println("Client: "+this.client.HostAddr+" Requesting group")
	cntrller := controller.NewController(this.client,this.client.HostAddr)
	m,_ := cntrller.Request(util.GROUP)
	group,_ := cntrller.Parse(m)
	fmt.Println(group.(model.Group).Clients["debian:1041"]) // (testing conversion)
}


func parse(n int, buffer []byte){
	fmt.Println("Client: Message received from server")
	message := model.Message{}
	err := json.Unmarshal(buffer[0:n],&message)
	if err != nil {
		return
	}
	switch message.Header {
	case util.RESPONSE:
		switch message.Type {
			case util.GROUP:
				b, err := json.Marshal(message.Attachment)
				var g model.Group
				err = json.Unmarshal(b,&g)
				fmt.Println(g)
				if err != nil {
					fmt.Print(err)
				}
		}
	}
}

func main(){
	app := NewApplication()
	app.Run()
}