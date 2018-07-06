package main

func main() {
	// 초기화 없이 선언된 변수들은 기본값으로 채워지며, 이들을 zero value라고 부른다
	// C에서 존재하는 쓰레기 값에 관한 문제를 제거하기 위함이다

	var a int
	// int의 zero value는 0

	var b float32
	// float의 zero value는 +0.000000e+000

	var c complex64
	// complex의 zero value는 (+0.000000e+000+0.000000e+000i)

	var d string
	// string의 zero value는 ""

	var e bool
	// bool의 zero value는 false

	_, _, _, _, _ = a, b, c, d, e
}
