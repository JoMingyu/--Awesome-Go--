package main

import "Awesome-Go/030._Custom_Package/some"

// 직접 정의한 라이브러리는 import [path] 형태로 import한다
// 위는 go/src 하위의 'Awesome-Go/030._Custom_Package/some'를 import하는 구문이며
// 경로는 GOPATH/src를 기준으로 한 상대경로를 사용한다
// 내장 패키지와 다르게 path가 포함되므로 이름이 겹치는 일을 피할 수 있다

func main() {
	println(some.Sum(1, 3))
	// some 패키지의 a.go에 있는 Sum 호출

	println(some.Sum2(1, 4, 3, 2, 8, 9))
	// some 패키지의 b.go에 있는 Sum 호출

	// 이처럼 동일한 패키지에 여러 go 파일이 나뉘어 있더라도, 이들은 단순히 분할을 위한 것일 뿐
	// 각 파일들의 요소는 동일한 패키지 네임스페이스에서 관리된다
}
