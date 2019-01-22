package main

import (
	"bufio"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/server"
	"distributed-systems/Ordered-Multicast/src/util"
	"fmt"
	"os"
	"strconv"
)

type Application struct {
	reader *bufio.Reader
	client *model.Client
}

func NewApplication() *Application {
	return &Application{reader:bufio.NewReader(os.Stdin) }
}

func (this *Application) showMenu(){
	fmt.Println("| ######## Main menu ######## | ")
	fmt.Println("| 1 - Retrieve available groups | ")
	this.reader = bufio.NewReader(os.Stdin)
	str,_ := this.readKeyboard()
	var option int64
	option,_ = strconv.ParseInt(str,10,64)
	fmt.Println(option)
	if option == 1{
		fmt.Println("Retrieving lobbies")
	}
}

func (this *Application) Before(){
	fmt.Println("Please inform your computer hostname:")
	host, _ := this.readKeyboard()
	this.client = model.NewClient(host)
	fmt.Println("Welcome " + host)
}

func (this *Application) readKeyboard() (string,error){
	str, err := this.reader.ReadString('\n')
	return str,err
}

func (this *Application) Run()  {
	this.Before()
	fmt.Println("Client: Requesting group")
	m := model.NewMessage(0,this.client.HostAddr,server.SERVER_ADDR,util.REQUEST,util.GROUP,nil)
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