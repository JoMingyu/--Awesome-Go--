package main

import (
	"fmt"
)

func main() {
	if false {
		// 조건식이 반드시 boolean 형태여야 함
		fmt.Println("false")
	} else if false {
		fmt.Println("false!!")
	} else {
		fmt.Println("else")
	}

	if val := 1 * 2; val == 2 {
		// if에서 조건식을 사용하기 전에 optional statement를 함께 실행할 수 있음
		fmt.Println("val is 2")
	}

	v := 3

	switch v {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		fallthrough
		// Go의 case는 암시적으로 break가 동작
		// 의도적으로 continue해주려면 fallthrough를 사용하면 된다
	case 100:
		fmt.Println("A case")
	}
}
