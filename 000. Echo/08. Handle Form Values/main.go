package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		// (c *context) FormValue(name string) string 함수를 통해 name에 해당하는 form value를 반환받을 수 있다
		// Query parameter 때와 동일하게, name에 대해 값이 전달되지 않으면 empty string이 반환된다

		if params, err := c.FormParams(); err == nil {
			fmt.Println(params)
			// Query parameter를 다루는 것과 거의 동일한 API를 가지고 있다
		}

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	e.Start(":8000")
}
