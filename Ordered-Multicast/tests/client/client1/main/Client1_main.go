package main

import (
	"distributed-systems/Ordered-Multicast/src/interfaces"
	"fmt"
)

func main(){
	app := interfaces.NewApplication(interfaces.CL1)
	app.Run()
	fmt.Scanln()
}
