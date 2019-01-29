package main

import (
	"distributed-systems/Ordered-Multicast/src/controller"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	CL1 string = "jhonson:1041"
	CL2 string = "baby:1041"
)

type Application struct {
	CLIENT_ADDR string
}

func NewApplication(CLIENT_ADDR string) *Application {
	return &Application{CLIENT_ADDR: CLIENT_ADDR}
}


func (this *Application) Run()  {
	// TODO COMPARE CONTROLLER'S PEER AND PEER RECEIVED FROM THE GROUP
	fmt.Println("Client: "+this.CLIENT_ADDR+" Requesting group")
	cntrller := controller.NewController(this.CLIENT_ADDR)
	m,_ := cntrller.Request(util.GROUP)

	var group model.Group
	parsed,_ := cntrller.Parse(m)
	group = parsed.(model.Group)
	// todo set back controller peer (received from group)
	fmt.Println("Group:", group.Peers[this.CLIENT_ADDR])
	//peer := group.Peers[this.CLIENT_ADDR]
	//peer.Listener.Multicast(model.NewMessage(0,"eu","voce","empty","empty",nil))
	//peer.Listener.Listen()

	//##################################################################################################################
	//addr,err := net.ResolveUDPAddr("udp4",group.Address)
	//checkError(err)
	//cntrller.AssignGroupAddress(addr)
	//iface , err := net.InterfaceByName("lo")
	//checkError(err)
	//conn, err := net.ListenMulticastUDP("udp4", iface , addr) // MULTICAST SOCKET
	//checkError(err)
	//
	//fmt.Println("Waiting multicast messages...")
	//_,err = conn.WriteToUDP([]byte("hello world"),addr)
	//for {
	//	buf := make([]byte,util.BUFFER_SIZE) // INITIALIZE THE BUFFER
	//	checkError(err)
	//	n, addr, err := conn.ReadFromUDP(buf[0:]) // READ IT
	//	fmt.Println("Message received from "+addr.String())
	//	fmt.Println("message: "+string(buf[0:n]))
	//	checkError(err)
	//}

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
	app := NewApplication(CL1)
	go app.Run()
	time.Sleep(time.Second * 3)
	app2 := NewApplication(CL2)
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
