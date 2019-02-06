package main

import (
	"fmt"
	"time"
)

func main(){
	//udpServer := server.NewUdpServer("debian:1041")
	//udpServer.Run()
	c := boring("redbone !")
	fmt.Println("waiting...")
	time.Sleep(time.Second * 10)
	fmt.Println("done")
	for i:= 0;i < 5 ; i++ {
		fmt.Printf("You say + %q\n",<-c)
	}
	fmt.Println("too boring, à plus tard!")
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