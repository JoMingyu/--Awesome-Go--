package main

import (
	"fmt"
)

func main() {
	// Go의 배열은 동일한 타입의 원소가 연속된 블록에 저장되는 fixed-length 데이터 타입

	var arr [5]int
	// 1차원
	arr[0] = 3
	fmt.Println(arr)

	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	// 리터럴로 array 선언. length는 추론됨

	arr3 := [4][2]int{1: {1, 2}, 3: {3}}
	// 특정 요소만 초기화
	fmt.Println(arr3)
}
