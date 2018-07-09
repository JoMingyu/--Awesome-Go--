package main

func main() {
	// Go는 많은 언어들의 장점들을 최대한 수용했고, 따라서 JavaScript의 대표적인 특징인 익명 함수를 지원한다
	func(a, b int, callback func(int)) {
		callback(a + b)
	}(1, 3, func(result int) {
		println(result)
	})
	// 두 수를 더해 콜백에 전달하는 함수를 익명으로 생성하고, 즉시 실행
}
