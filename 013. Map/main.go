package main

import (
	"fmt"
)

func main() {
	// Node.js의 객체, Python의 dictionary에 해당하는 Map
	// 순서가 없는 키-값 쌍의 집합을 뜻함

	var m map[string]int
	// map[T1]T2 는 T1을 key로, T2를 값으로 가지는 nil map
	fmt.Println(m)

	m2 := map[string]int{}
	// map[T1]T2{} 는 T1을 key로, T2를 값으로 가지는 empty map
	fmt.Println(m2)

	m3 := make(map[string]int)
	// empty map
	fmt.Println(m3)

	mapping := map[string]int{
		"key":  123,
		"key2": 12345,
	}
	// 리터럴을 통한 초기화. 구조체 리터럴과 비슷한 형태지만 k:v 형태로 사용

	fmt.Println(mapping)

	mapping["key3"] = 132
	// 요소 삽입 또는 수정
	fmt.Println(mapping["key3"])
	// 요소 가져오기
	delete(mapping, "key3")
	// 요소 지우기
	v, exist := m["key3"]
	// map을 읽는 경우 해당 key의 값과 존재 여부를 함께 리턴함
	fmt.Println(v, exist)
}
