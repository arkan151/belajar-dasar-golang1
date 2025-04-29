package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")

	for true {
		fmt.Println("Hello !!")
		time.Sleep(1 * time.Second)
	}

	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	<-time.After(4 * time.Second)
	fmt.Println("expired")

	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}

}
