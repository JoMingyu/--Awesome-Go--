package main

// 모든 Go 프로그램은 패키지로 구성되어 있음. 현재 패키지는 main
// 패키지 이름은 디렉토리 경로의 마지막 이름을 사용하는 것이 규칙 path1/path2 하위에 파일이 있다면, 패키지명은 path2

import (
	"fmt"
	"math"
)

// 이 프로그램은 fmt와 math라는 빌트인 패키지를 import하고 있음
// import 문장을 여러 번 사용할 수도 있고, 위처럼 여러 패키지들을 소괄호로 감싸서 표현할 수도 있음

func main() {
	fmt.Println("Pi is", math.Pi)
}
