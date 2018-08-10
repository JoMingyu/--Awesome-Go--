package main

import (
	"encoding/json"
	"fmt"
)

type student1 struct {
	No   int
	Name string
	// JSON marshalizer는 구조체의 exported name만을 직렬화할 수 있기 때문에 exported name 사용
}

type student2 struct {
	No   int    `json:"no"`
	Name string `json:"name"`
	// key 이름을 명시하기 위한 필드 tagging
}

func main() {
	std1 := student1{3118, "Jo Mingyu"}
	std2 := student2{3118, "Jo Mingyu"}

	if data, err := json.Marshal(std1); err == nil {
		fmt.Println(string(data))
		// {"No":3118,"Name":"Jo Mingyu"}

		emptyStd := new(student1)
		json.Unmarshal(data, emptyStd)
		fmt.Println(emptyStd)
		// &{3118 Jo Mingyu}
	}

	if data, err := json.Marshal(std2); err == nil {
		fmt.Println(string(data))
		// {"no":3118,"name":"Jo Mingyu"}
	}

	if data, err := json.Marshal(map[string]interface{}{"no": 3118, "name": "JoMingyu"}); err == nil {
		fmt.Println(string(data))
		// {"name":"JoMingyu","no":3118}
	}
}
