package main

func sum(values ...int) int {
	// [name] ...[type]
	// 가변 인자는 JavaScript의 rest parameter와 비슷한 형식으로 표현한다
	res := 0

	for _, value := range values {
		res += value
	}

	return res
}

func main() {
	println(sum([]int{1, 2, 3, 4, 5}...))
	// [slice]...
	// 슬라이스를 인자로 풀어 전달할 수 있다
	// Python의 sequence 객체에 대한 argument unpacking과 유사하다
}
