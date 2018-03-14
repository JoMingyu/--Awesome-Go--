package main

import (
	"fmt"
)

func main() {
	// go는 정적 타이핑 언어
	// 변수 선언을 위해 var를 사용하며, 함수의 매개변수처럼 타입은 문장 긑에 명시
	var a int
	var b, c int
	// 나열

	var d int = 5
	// 선언과 함께 초기화
	var e = 5
	// 타입 추론

	var f float32 = 11
	fmt.Println(a, b, c, d, e, f)

	h := 10
	// :=을 사용하면 var와 타입 명시 생략 가능, 그러나 함수 밖에서는 := 선언을 사용할 수 없음
	i, j := 11, 12
	fmt.Println(h, i, j)

	// 아래는 Go의 25개 키워드
	/*
		break        default      func         interface    select
		case         defer        go           map          struct
		chan         else         goto         package      switch
		const        fallthrough  if           range        type
		continue     for          import       return       var
	*/
}
