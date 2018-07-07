package main

func main() {
	// 배열(array)는 길이가 고정(fixed)된, index 기반의 단일 타입 원소의 나열
	var x [5]int
	_ = x
	// 별도로 초기화되지 않은 배열은 모든 요소가 해당 타입의 zero value로 채워진다
	// 위의 경우 [0 0 0 0 0]

	x[0] = 13
	x[1] = 52

	println(len(x))
	// 배열의 길이

	var total int
	for _, value := range x {
		// range [arr] 표현을 사용하여 배열을 순회할 수 있다
		// 순회할 때마다 인덱스와 값을 반환하는데, Go는 사용하지 않는 변수가 있으면 컴파일 오류가 일어나므로
		// 언더스코어를 이용해 값 무시
		total += value
	}

	y := [5]int{0, 3, 8, 11, 9}
	// 리터럴을 사용한 초기값과 함께 배열을 생성
	_ = y
}
