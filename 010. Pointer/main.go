package main

import (
	"fmt"
)

func main() {
	n := 15
	p := &n
	// Go에는 포인터가 있지만 포인터 연산은 불가능
	// 포인터는 해당 변수의 주소값을 가르킴

	*p = 20
	// 포인터를 이용한 간접적인 접근은 해당 주소의 변수에도 영향을 미침

	fmt.Println(n)
}
