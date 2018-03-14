package main

import (
	"fmt"
)

func main() {
	if false {
		// For문과 비슷하게, 조건 표현을 위한 소괄호는 사용하지 않고 전체적인 구조는 C/Java와 비슷함
		fmt.Println("false")
	} else if false {
		fmt.Println("false!!")
	} else {
		fmt.Println("else")
	}

	if val := 1 * 2; val == 2 {
		// for의 초기화 구문처럼, if에서도 조건문 앞에 짧은 문장을 실행할 수 있음
		// 해당 실행문을 통해 선언된 변수는 if 블럭 안쪽 scope, 해당 statement의 else 블럭에서만 사용할 수 있음
		fmt.Println("val is 2")
	}
}
