package src

import (
	"fmt"
	"time"
)

/**
*	Right away after the boring function is called a timeout has been inserted so that the synchronization between channels
* 	becomes more visible. This way, although the anonymous function inside the boring function was called to run concurrently
*	and indefinitely through the "go" command, it becomes locked until the "Printf" function is called to receive the val-
*	ues outputted from the channel. That is, the communication only occurs when the the process is available to receive the
*	values from the channel.
**/
func Generate(){
	c := boring("redbone !") // Function that returns a channel
	time.Sleep(time.Second * 60)
	for i:= 0;i < 5 ; i++ {
		fmt.Printf("You say + %q\n",<-c) // getting the value output from the channel
	}
	fmt.Println("Vous êtes ennuyeux, à plus tard!")
}


func boring(msg string) <-chan string{
	c := make(chan string)
	go func() {
		for i:= 0; ; i++ {
			c <- fmt.Sprintf("%s %d",msg , i)
			time.Sleep(time.Second *1)
		}
	}()
	return c
}