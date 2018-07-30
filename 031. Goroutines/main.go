package main

import (
	"fmt"
	"time"
)

func test(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 100)
		// time.Sleep(duration)은 duration만큼 고루틴을 일시정지
	}
}

func main() {
	// Go는 매우 간단하지만 가볍고 빠른 동시성 개념인 고루틴(goroutine)을 지원한다
	// 스레드에 비해 메모리 소비, 설치와 철거 비용, context switching 비용이 현저히 적다
	// 일반적으로 스레드는 guard page라는 메모리 영역과 함께 1Mb 정도로 시작하는데, 고루틴은 2kb의 스택 공간만 필요하다
	// 추후에 힙 저장 공간을 확보하여 사용하는데, 이 특징 덕분에 고루틴 수천 개를 만들어 사용해도 부담이 적다
	go test(1)
	go test(2)
	go test(3)

	time.Sleep(time.Second)

	for i := 0; i < 1000; i++ {
		go func(n int) {
			// 익명 함수로 총 1000개의 고루틴 실행
			for i := 0; i < 100; i++ {
				fmt.Println(n, "Hello")
				time.Sleep(time.Millisecond * 10)
			}
		}(i)
	}

	time.Sleep(time.Second * 2)
}
