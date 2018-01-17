package main

import "fmt"

func main() {
	// Node.js의 객체, Python의 dictionary에 해당하는 Map
	// 순서가 없는 키-값 쌍의 집합을 뜻함

	var m map[string]int
	// nil map
	fmt.Println(m)

	m2 := make(map[string]int)
	// empty map
	fmt.Println(m2)

	m3 := map[string]int{}
	// 리터럴을 이용한 empty map
	fmt.Println(m3)

	mapping := map[string]int{
		"key":  123,
		"key2": 12345,
	}
	// 리터럴을 통한 초기화

	fmt.Println(mapping)

	mapping["key3"] = 132
	fmt.Println(mapping["key3"])
}
