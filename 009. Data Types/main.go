package main

import (
	"fmt"
)

func main() {
	var b bool
	// Boolean

	var s string = "qw"
	// Go에서 string은 immutable
	// 문자 하나는 작은따옴표. 문자열은 큰따옴표로 감싸 줌

	var i int
	// int, int8, int16, int32, int64

	var u uint
	// unsigned int
	// uint, uint8, uint16, uint32, uint64, uintptr

	var by byte
	// uint8과 동일하며 바이트 코드에 사용

	var r rune
	// int32와 동일하며 유니코드 코드포인트에 사용

	var f32 float32
	var f64 float64
	// float32, float64

	var comp64 complex64
	var comp128 complex128
	// 복소수 타입
	// complex64, complex128

	fmt.Println(b, s, i, u, by, r, f32, f64, comp64, comp128)
	fmt.Println(int(f32))
	// 타입 캐스팅
}
