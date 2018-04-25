package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// 1. Query parameter
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, c.QueryParam("arg"))
		// 전달된 파라미터가 없으면 빈 문자열을 리턴
	})

	// 2. Path parameter
	e.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})

	e.Start(":5000")
}
