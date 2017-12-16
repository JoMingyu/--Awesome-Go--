package main

import (
	"fmt"
)

// Formatted I/O를 위해

func main() {
	var arr [5]int
	// 1차원
	arr[0] = 3
	fmt.Println(arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	// length

	arr3 := [][]int{{1, 2}, {3}}
	// 2차원. 선언과 함께 값 할당 시 array length는 생략 가능
	fmt.Println(arr3)
}
