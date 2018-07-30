package some

// 폴더 구조가 깊더라도, 패키지 이름은 현재 디렉토리의 이름으로 설정하는 것이 관례다

func Sum(a, b int) int {
	// Exported name 사용
	return a + b
}

func init() {
	// 패키지가 import되면 해당 패키지에 존재하는 모든 init 함수가 호출된다
	// 패키지 요소들의 초기화 구문에 유용하게 사용될 수 있다
	println("a.go initialized!")
}
