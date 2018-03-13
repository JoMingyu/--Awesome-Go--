package main

import (
	"fmt"
	"math"
)
// 패키지를 import하면 패키지가 외부로 export한 것들에 접근할 수 있음
// Go에서는 첫 문자가 대문자로 시작하면, 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됨
// Node.js에서 명시적으로 exports.***나 module.exports를 사용하는 것처럼, Go에서는 첫 문자를 대문자로 둬서 exported name을 표현

func main() {
	fmt.Println(math.Pi)
	// 따라서 타 패키지에선 이름이 대문자로 시작하는 요소들에만 접근할 수 있음
}
