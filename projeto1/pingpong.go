package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func pong(c chan string) {
	for {
		msg := <-c
		if msg == "ping" {
			fmt.Println(msg)
			fmt.Println("pong")
		}
	}
}

func main() {
	var canal chan string = make(chan string)
	go ping(canal)
	go pong(canal)
	time.Sleep(time.Second * 10)
	fmt.Println("FIM")
}
