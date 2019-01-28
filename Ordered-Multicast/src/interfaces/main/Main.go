package main

import (
	"bufio"
	"distributed-systems/Ordered-Multicast/src/controller"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Application struct {
	reader *bufio.Reader
}

const (
	CLIENT_ADDR string = "localhost:1041"
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
	fmt.Println("Client: "+CLIENT_ADDR+" Requesting group")
	cntrller := controller.NewController(CLIENT_ADDR)
	m,_ := cntrller.Request(util.GROUP)

	var group model.Group
	parsed,_ := cntrller.Parse(m)
	group = parsed.(model.Group)

	//##################################################################################################################
	addr,err := net.ResolveUDPAddr("udp4",group.Address)
	checkError(err)
	cntrller.AssignGroupAddress(addr)
	iface , err := net.InterfaceByName("lo")
	checkError(err)
	conn, err := net.ListenMulticastUDP("udp4", iface , addr) // MULTICAST SOCKET
	checkError(err)

	fmt.Println("Waiting multicast messages...")
	_,err = conn.WriteToUDP([]byte("hello world"),addr)
	for {
		buf := make([]byte,util.BUFFER_SIZE) // INITIALIZE THE BUFFER
		checkError(err)
		n, addr, err := conn.ReadFromUDP(buf[0:]) // READ IT
		fmt.Println("Message received from "+addr.String())
		fmt.Println("message: "+string(buf[0:n]))
		checkError(err)
	}

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
	go app.Run()
	time.Sleep(time.Second * 3)
	app2 := NewApplication()
	app2.Run()
}

/**
* Checks whether there was an error
**/
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
