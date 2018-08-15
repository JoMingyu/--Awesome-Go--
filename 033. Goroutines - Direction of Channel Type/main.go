package main

import (
	"fmt"
	"time"
)

func producer(c chan<- int) {
	// 채널 타입에 방향을 지정할 수 있는데, chan<- 타입은 해당 채널에 값을 적재만 하도록 제한한다
	for i := 0; i < 100; i++ {
		c <- i
		time.Sleep(time.Millisecond * 45)
	}
}

func consumer(c <-chan int) {
	// <-chan 타입은 해당 채널에서 값을 가져오기만 하도록 제한한다
	for {
		msg := <-c

		fmt.Println(msg)
	}
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	time.Sleep(time.Second * 5)
}
