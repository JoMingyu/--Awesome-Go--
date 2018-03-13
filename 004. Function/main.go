package main

import (
	"fmt"
)

func add(x int, y int) int {
	// 함수는 매개변수(인자)를 가질 수 있음. 여기서는 두 개의 int 타입 매개변수를 받고 있음
	// 매개변수의 타입은 변수명 뒤에 명시하며, 좌에서 우로 읽었을 때 읽기 쉽도록 하기 위함이라고 함
	return x + y
}

func add2(x, y int) int {
	// 두 개 이상의 매개변수가 같은 타입일 때, 같은 타입을 취하는 마지막 매개변수에만 타입을 명시하고 나머진 생략 가능
	return x + y
}

func add3(x, y int) (sum int) {
	// 함수 시그니처의 반환에 이름을 부여하면 변수처럼 사용할 수도 있음
	// 반환 값을 지정하지 않은 return 문장으로 결과 변수의 현재 값을 알아서 반환함
	// Named result라고 하는데, 그리 좋지는 않음. 커뮤니티 측에서도 비추천하는 방식
	sum = x + y
	return
}

func swap(x, y int) (int, int) {
	// 함수는 여러 개의 결과를 반환할 수 있음
	return y, x
}

func some(msg *string) {
	// Call by reference
	*msg = "new!"
}

func main() {
	str := "msg"

	some(&str)
	fmt.Println(str)

	fmt.Println(some2("a", "bc", "d"))
}

func some2(msg ...string) (a int, b int) {
	// 가변인자 전달과 복수 개의 값 리턴
	// 위는 Named Return Parameter
	for _, m := range msg {
		fmt.Println(m)
	}

	a = 1
	b = 2

	return
}
