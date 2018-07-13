package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		// (c *context) QueryParam(name string) string 함수를 통해 name에 해당하는 query parameter를 받아올 수 있다
		// 해당 name이 query parameter에 존재하지 않는 경우, empty string이 반환된다

		params := c.QueryParams()
		// (c *context) QueryParams() url.Values 함수를 통해 query parameter들을 url.Values 타입으로 반환받을 수 있다
		// url.Values는 map 타입이므로(type Values map[string][]string) map처럼 사용 가능하다
		fmt.Println(params)

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	e.Start(":8000")
}
