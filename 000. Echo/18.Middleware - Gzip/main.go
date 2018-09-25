package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// response data에 대한 Gzip compress를 위한 미들웨어
	// 요청의 Accept-Encoding 헤더에 gzip이 포함되어 있을 때, 응답의 body를 gzip으로 압축한다
	e.Use(middleware.Gzip())
	// func Gzip() echo.MiddlewareFunc
	// 사용되는 GzipConfig(DefaultGzipConfig)는 다음과 같다
	/*
		GzipConfig{
			Skipper: DefaultSkipper,
			Level:   -1,
		}
	*/

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	// compress level 5로 gzip compress하도록 함

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Something")
	})

	e.Start(":8000")
}
