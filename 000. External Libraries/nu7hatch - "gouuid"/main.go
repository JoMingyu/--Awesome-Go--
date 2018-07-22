package main

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
	// go get github.com/nu7hatch/gouuid
)

func main() {
	if u4, err := uuid.NewV4(); err == nil {
		fmt.Println(u4)
		// 생성되는 uuid는 *UUID 구조체기 때문에, stringify하려면 fmt를 이용하면 좋다
		fmt.Println(fmt.Sprintf("%s", u4))
	}
}
