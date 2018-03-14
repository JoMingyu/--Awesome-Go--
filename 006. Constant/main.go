package main

import (
	"fmt"
)

func main() {
	const g = 10
	// 상수 선언. const 키워드를 명시적으로 사용해야 함
	// 상수의 타입은 Character, String, Boolean, Int 중 하나가 될 수 있음
	const (
		SOME  = "Some"
		THING = "Thing"
	)
	// 나열식 상수 선언

	fmt.Println(g, SOME, THING)
}
