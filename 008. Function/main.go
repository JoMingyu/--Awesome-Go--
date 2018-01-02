package main

import (
	"fmt"
)

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
