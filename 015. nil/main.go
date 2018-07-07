package main

func main() {
	// nil은 zero value의 일종으로, 초기화되지 않은 slice, map, 포인터, 인터페이스 등에서 사용된다
	var s []int
	println(s == nil)

	var m map[int]string
	println(m == nil)
}
