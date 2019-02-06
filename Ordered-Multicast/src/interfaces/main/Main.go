package main

import (
	"distributed-systems/Ordered-Multicast/src/controller"
	"distributed-systems/Ordered-Multicast/src/model"
	"distributed-systems/Ordered-Multicast/src/util"
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
	fmt.Println("Client: "+this.CLIENT_ADDR+" Requesting group")
	cntrller := controller.NewController(this.CLIENT_ADDR)
	m,_ := cntrller.Request(util.GROUP) // REQUEST A GROUP

	var group model.Group
	parsed,_ := cntrller.Parse(m)
	group = parsed.(model.Group) // GET THE REQUEST AND CONVERT

	peer := group.Peers[this.CLIENT_ADDR] // GET A PEER
	cntrller.SetPeer(peer) // REPLACE THE OLD ONE IN THE CONTROLLER
	cntrller.ConnectPeer("lo")  // CONNECT HIM
	go cntrller.Peer().Listener.Listen() // START LISTENING FOR CONNECTIONS
	time.Sleep(time.Second*4)
	cntrller.Peer().Listener.Multicast(model.NewMessage(0,"eu","voce","e o zubumafoo","",nil))
}

func main(){
	app := NewApplication(CL1)
	go app.Run()
	fmt.Scanln()
	time.Sleep(time.Second * 3)
	//app2 := NewApplication(CL2)
	//app2.Run()
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
