package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/:name", func(c echo.Context) error {
		// Path parameter는 :[name] 형식으로 표현한다

		name := c.Param("name")
		// (c *context) Param(name string) string 함수를 통해 path parameter를 받아올 수 있다
		// 요청 URI에 Path parameter 자리가 비어있는 경우, 프레임워크 레벨에서 404가 반환된다

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	e.Start(":8000")
}
