package main

import (
	"fmt"
)

func main() {
	sum := func(a, b int) int {
		return a + b
	}
	// Go의 함수는 JS처럼 일급 함수로 분류되므로, 변수에 '값'으로 할당할 수 있음

	fmt.Println(sum(3, 4))

	// 일급 객체와 일급 함수의 차이는 런타임 생성과 익명 생성 가능 여부에 의해 결정됨
	fmt.Println(calc(func(a, b int) int {
		// 런타임에 익명으로 생성해서 함수를 값처럼 전달
		return a - b
	}, 3, 5))
}

func calc(executor func(int, int) int, a int, b int) int {
	res := executor(a, b)

	return res
}
