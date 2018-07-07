package main

func main() {
	// map은 순서가 없는 키-값 쌍의 집합을 표현

	var m map[string]int
	// map[T1]T2 는 T1을 key로, T2를 값으로 가짐
	// map도 zero value는 nil
	println(m == nil)
	// true

	m2 := map[string]int{}
	// 리터럴을 사용하여 초기값 정의
	// string을 key로, int를 값으로 가지는 길이가 0인 map
	println(m2 == nil)
	// false

	m3 := make(map[string]int)
	_ = m3
	// empty map
	// map의 길이는 key-value 매핑이 변화할 때 바뀌므로, 별도로 길이를 설정할 수 없음

	mapping := map[string]int{
		"key":  123,
		"key2": 12345,
	}
	// 리터럴을 통한 초기화

	mapping["key3"] = 132
	// 요소 삽입 또는 수정

	v, exist := m["key3"]
	// 요소 가져오기
	// map을 읽는 경우 해당 key의 값과 존재 여부를 함께 리턴함

	println(v, exist)

	// 위와 같은 특징과 Go만이 가지는 if문의 변수 초기화 구문을 통해 아래와 같은 코드를 표현할 수 있다
	if v, exist := m["key3"]; exist {
		println(v)
	}

	delete(mapping, "key3")
	// 요소 지우기
}
