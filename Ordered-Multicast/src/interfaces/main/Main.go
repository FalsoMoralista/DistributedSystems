package main

import (
	"bufio"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/server"
	"fmt"
	"net"
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

	parsedAddr,err := net.ResolveUDPAddr("udp4",this.client.HostAddr) // resolve the udp address

	fmt.Println("Client: Requesting group")

	conn,_:= net.DialUDP("udp",nil,parsedAddr)

	int ,err := conn.Write([]byte("testing communication"))

	if(int == 0){
		fmt.Println("Error")
	}
	n,err := conn.Read(this.buffer[0:]) // TODO: Read documentation about read/send functions

	fmt.Println("Client: Message received from server:",string(this.buffer[0:n]))
	return conn,err

}

func main(){
	app := NewApplication()
	app.Run()
}