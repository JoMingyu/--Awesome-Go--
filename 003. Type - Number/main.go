package main

func main() {
	// 일반적인 프로그래밍 언어들처럼, 숫자형은 정수와 부동 소수점 수로 나눈다
	a := 1
	b := 1.0

	_, _ = a, b

	// 정수 타입으로는 int8/16/32/64, uint8/16/32/64를 사용할 수 있다
	var c uint64 = 18446744073709551615

	_ = c

	// uint8과 같은 byte, int32와 같은 rune이 있다
	var d byte = 255
	var e rune = 2147483647

	_, _ = d, e

	// 부동 소수점 타입으로는 float32/64(각각 단정도, 배정도 부동 소수점)
	// 복소수를 표현하기 위한 complex64/complex128을 사용할 수 있다
	var f float32 = 1.129391199
	var g complex64 = 3i + 1.1312

	_, _ = f, g
}
