package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			// (c *context) JSON(code int, i interface{}) (err error) 함수를 통해 JSON payload를 response할 수 있다
			"name": c.QueryParam("name"),
		})
	})

	e.Start(":8000")
}
