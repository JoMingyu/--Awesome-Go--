package main

import (
	"fmt"
)

func main() {
	// Go에서 연속적인 데이터를 저장하기 위해 array 또는 slice를 사용
	// Go의 array는 고정된 크기 안에 동일한 타입의 데이터를 연속적으로 저장하지만
	// 배열의 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등의 기능을 가지고 있지 않음

	// Slice는 내부적으로 동적 배열의 개념으로 구현되었기 때문에 필요에 따라 크기를 늘리거나 줄일 수 있음
	// append와 같이 빠르게 효과적으로 그 크기를 조절할 수 있는 내장 함수를 제공하기 때문에 크기 측면에서는 유연함
	// 게다가 메모리가 연속적인 블록으로 할당되기 때문에 인덱싱, iteration과 GC에서의 장점도 제공함

	// 슬라이스가 관리하는 세개의 필드는 내부 배열의 포인터, 슬라이스가 접근할 수 있는 요소의 개수 또는 길이, 슬라이스가 최대한 저장할 수 있는 용량
	// Array와 슬라이스는 문법적으로 굉장히 비슷한데, 대괄호 사이에 길이가 없다는 특징을 알고 있으면 구별하기 쉽다

	var s []int
	// []T 는 타입 T를 가지는 요소의 nil slice(길이와 최대 크기가 0이며 아무 값도 할당되지 않아 nil인 slice)
	fmt.Println(s == nil)
	fmt.Println(len(s), cap(s))

	s2 := []int{}
	// []T{} 는 타입 T를 가지는 empty slice(요소의 길이와 최대 크기가 0이지만 slice의 값 자체는 {}로 할당됨)
	fmt.Println(s2 == nil)

	s3 := make([]string, 5)
	// make 함수로도 만들 수 있음. 이렇게 생성된 슬라이스는 0을 할당한 배열을 생성하고, 그것을 참조함
	// 슬라이스의 길이만 지정하면 최대 용량도 같은 값으로 설정됨. 따라서 길이와 용량이 모두 5인 slice
	fmt.Println(len(s3), cap(s3))

	s4 := make([]string, 0)
	// 길이와 용량이 0인 empty slice
	fmt.Println(len(s4), cap(s4))

	s5 := make([]int, 1, 5)
	// 슬라이스의 길이와 용량을 함께 지정할 수 있음
	// 길이는 1, 최대 용량은 5인 slice

	s5 = append(s5, 5)
	// 슬라이스에 값 append
	fmt.Println(s5)

	s6 := []int{1, 2, 3, 4}
	// 리터럴을 사용하는 슬라이스의 보편적인 선언 방식
	fmt.Println(len(s6), cap(s6))
	fmt.Println(s6[1:3])
	// 파이썬의 슬라이싱 개념과 같음

	for i, v := range s6 {
		// range를 사용해 slice 순회
		fmt.Println(i, v)
	}

	for _, v := range s6 {
		// 언더스코어를 이용해 인덱스 무시
		fmt.Println(v)
	}
}
