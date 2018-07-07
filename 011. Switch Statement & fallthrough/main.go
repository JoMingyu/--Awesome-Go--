package main

func main() {
	i := 0

	switch i {
	// 일반적인 형태의 switch-case를 지원
	case 0:
		println("zero")
		fallthrough
		// case에 break는 기본으로 내장되어 있으므로,
		// 다음 case를 함께 실행하려 한다면 fallthrough라는 statement를 사용한다
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	default:
		println("not 0~3")
	}
}
