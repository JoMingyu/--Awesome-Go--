package main

func main() {
	var b bool
	// Boolean

	var s string
	// Go에서 string은 immutable

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

	println(b, s, i, u, by, r, f32, f64, comp64, comp128)
	println(int(f32))
	// 타입 캐스팅
}
