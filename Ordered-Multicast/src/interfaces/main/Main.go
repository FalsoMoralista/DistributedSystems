package main

import (
	"bufio"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/server"
	"distributed-systems/Ordered-Multicast/src/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Application struct {
	reader *bufio.Reader
	client *model.Client
}

func NewApplication() *Application {
	return &Application{reader:bufio.NewReader(os.Stdin) }
}

func (this *Application) Before(){
	fmt.Println("Please inform your computer hostname:")
	host, _ := this.readKeyboard()
	port := strconv.Itoa(server.SERVICE_PORT)
	this.client = model.NewClient(host +":"+port)
	fmt.Println("Welcome!")
}

func (this *Application) readKeyboard() (string,error){
	str, err := this.reader.ReadString('\n')
	strings.Replace(str,"\n","",1) // todo review replace (" \n ")
	return str,err
}

func (this *Application) Run()  {
	this.Before()
	fmt.Println("Client: "+this.client.HostAddr+" Requesting group")
	m := model.NewMessage(0,this.client.HostAddr,server.SERVER_ADDR,util.REQUEST,util.GROUP,this.client)
	n,buffer,err := util.SendUdp(server.SERVER_ADDR,m)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println("Client: Message received from server: "+string(buffer[0:n]))
}

func main(){
	app := NewApplication()
	app.Run()
}