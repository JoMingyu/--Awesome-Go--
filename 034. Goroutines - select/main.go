package main

import (
	"fmt"
	"time"
)

func producer(id int, c chan<- int, sleepDuration time.Duration) {
	for {
		c <- id
		time.Sleep(sleepDuration)
	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go producer(1, c1, time.Second*2)
	// 2초마다 c1에 1을 적재
	go producer(2, c2, time.Second*3)
	// 3초마다 c2에 2를 적재

	go func() {
		for {
			select {
			// switch와 비슷하게 동작하지만, 채널에 대해서만 동작하는 select 구문
			// 여러 채널에 대해 한 번에 메시지를 listen할 수 있다
			case v1 := <-c1:
				fmt.Println(v1)
			case v2 := <-c2:
				fmt.Println(v2)
			default:
				// case에
				fmt.Println("Waiting")
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()

	time.Sleep(time.Second * 10)
}
