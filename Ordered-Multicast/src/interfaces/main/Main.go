package main

import (
	"distributed-systems/Ordered-Multicast/src/model"
	"fmt"
)

func main(){
 	queue := model.NewQueue()
 	queue.Add(2)
	queue.Add(4)
	queue.Add(5)
	queue.Add(6)
	for queue.Size() != 0 {
		queue.Peek()
		queue.Remove()
	}
 	fmt.Println("~~~~")
 	queue.Peek()
	fmt.Println("~~~~")
 	queue.Remove()
}

