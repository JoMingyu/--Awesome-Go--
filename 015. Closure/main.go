package main

import (
	"fmt"
)

func valueAccumulator() func() int {
	// int를 리턴하는 함수를 리턴하므로
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Go에서 함수는 클로저로서 사용될 수도 있다
	// Node.js에서의 클로저와 동일하게 함수 바깥에 있는 변수를 참조하는 함수값을 일컫는데
	// 이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 변수를 읽거나 쓸 수 있게 된다

	next := valueAccumulator()
	// next는 함수

	fmt.Println(next())
	fmt.Println(next())
}
