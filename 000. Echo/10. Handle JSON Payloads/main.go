package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type binder struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Echo에서 JSON payload를 요청으로 받아 관리하려면, bind라는 개념을 사용한다
// 첫 번째 방법은 구조체이며, exported name들로 필드가 정의된 구조체에 각 필드마다 json이라는 태그로 key name을 명시한다
// Echo binder가 구조체에 리플렉션으로 접근해야 하기 때문에, 필드들이 exported name으로 정의되어 있지 않으면 문제가 생길 수 있다

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		payload := &binder{}
		// 1. 구조체 포인터를 생성하고

		if err := c.Bind(payload); err != nil {
			// 2. (c *context) Bind(i interface{}) error 함수에 구조체 포인터를 넘겨 바인딩을 시도
			// 대표적으로 구조체에 정의된 타입과 JSON payload의 타입이 맞지 않는 경우에 에러가 발생(400 bad request)
			// key가 존재하지 않으면 zero value로 초기화된 대로 두면 되므로, key exist는 별도로 체크하지 않는다
			return err
		}

		fmt.Println(payload.Name, payload.Age)

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", payload.Name))
	})

	e.POST("/2", func(c echo.Context) error {
		payload := map[string]interface{}{}
		// 두 번째는 map을 사용하는 방법이다
		// 1. empty map을 생성한 후

		if err := c.Bind(&payload); err != nil {
			// 2. Bind 함수에 map의 포인터를 넘겨준다
			return err
		}

		fmt.Println(payload["name"], payload["age"])

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", payload["name"]))
	})
	// map을 사용하는 방식은 구조체를 사용하는 경우에 비해 type validation을 별도로 해야 하기 때문에, 구조체를 사용하는 편이 나을 것 같다

	e.Start(":8000")
}
