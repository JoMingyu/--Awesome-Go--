package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// 미들웨어는 다른 프레임워크들처럼, Context에 액세스 가능한 HTTP 요청과 응답 사이클에 체인된(chained) 함수다
	// 모든 요청을 로깅하거나, 요청 수를 제한하는 등 특정 작업을 수행하는 데 사용되며
	// Echo는 요청에 대해 모든 미들웨어 실행을 마친 후 handler를 실행시킨다

	// CORS는 CORS처리를 위한 미들웨어
	e.Use(middleware.CORS())
	// func CORS() echo.MiddlewareFunc
	// 사용되는 CORSConfig(DefaultCORSConfig)는 아래와 같다
	/*
		CORSConfig{
			Skipper:      DefaultSkipper,
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}
	*/

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAccept},
	}))
	// func CORSWithConfig(config CORSConfig) echo.MiddlewareFunc
	// 별도의 config 구조체 변수를 받아 CORS 처리

	// 기본적으로 echo의 미들웨어는 미들웨어 함수() + 미들웨어 함수WithConfig() + 미들웨어Config 구조체 + DefaultConfig 변수가 한 세트로 제공된다
	// middleware.CORS() + middleware.CORSWithConfig() + middleware.CORSConfig + middleware.DefaultCORSConfig

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8000")
}
