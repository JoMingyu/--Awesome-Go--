package main

import "fmt"

func main() {
	// 함수명을 갖지 않는 함수를 익명함수라 부른다
	// 일반적으로 함수를 변수에 할당하거나 다른 함수의 파라미터에 정의하여 사용하곤 한다

	sum := func(a int, b int) int {
		return a + b
	}

	res := sum(1, 2)
	fmt.Println(res)

	// Go에서 함수는 일급함수로서 기본 타입과 동일하게 취급되므로 파라미터로 전달하거나 리턴값이 될 수 있다
	fmt.Println(calc(sum, 3, 4))

	// 런타임에 생성해서 전달할 수도 있다
	fmt.Println(calc(func(a int, b int) int {
		return a - b
	}, 1, 3))
}

func calc(executor func(int, int) int, a int, b int) int {
	res := executor(a, b)

	return res
}
