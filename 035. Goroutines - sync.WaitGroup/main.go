package main

import (
	"sync"
	"time"
)

func f(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	// WaitGroup에 대해 Done()을 호출하여 WaitGroup에 고루틴이 끝났음을 알린다

	time.Sleep(time.Millisecond * 500)
}

func main() {
	var group sync.WaitGroup
	// sync.WaitGroup을 이용해 고루틴들이 끝날 때까지 기다리도록 할 수 있다
	group.Add(2)
	// 몇 개의 고루틴을 기다릴지 설정

	go f(&group)
	go f(&group)

	group.Wait()
	// 고루틴들이 끝나기를 기다림
}
