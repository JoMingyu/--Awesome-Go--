package main

import (
	"fmt"
)

func main() {
	v := 3

	switch v {
	// switch의 각 조건은 위에서 아래로 평가되므로, 참인 조건이 여럿 있더라도 맨 위에 있는 case가 동작
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		fallthrough
		// Go의 case는 코드 실행 후 알아서 break
		// 의도적으로 continue해주려면 fallthrough를 사용하면 된다
	case 100:
		fmt.Println("A case")
	}
}
