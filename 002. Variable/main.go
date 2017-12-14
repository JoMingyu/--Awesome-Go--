package main

func main() {
	// go는 정적 타이핑 언어
	var a int
	var b, c int
	// 나열
	var d = 5
	// 타입 추론
	var f float32 = 11
	println(a, b, c, d, f)

	const e = 10
	const (
		SOME = "Some"
		THING = "Thing"
	)

	println(e, SOME, THING)

	// 아래는 Go의 25개 키워드
	/*
	break        default      func         interface    select
	case         defer        go           map          struct
	chan         else         goto         package      switch
	const        fallthrough  if           range        type
	continue     for          import       return       var
	*/
}
