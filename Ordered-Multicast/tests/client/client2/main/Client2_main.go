package main

import (
	"distributed-systems/Ordered-Multicast/src/interfaces"
	"fmt"
)

func main() {
	app := interfaces.NewApplication(interfaces.CL2)
	app.Run()
	fmt.Scanln()
}