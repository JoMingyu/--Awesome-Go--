package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// 요청 데이터를 logging하기 위한 미들웨어
	e.Use(middleware.Logger())
	// func Logger() echo.MiddlewareFunc
	// 사용되는 LoggerConfig(DefaultLoggerConfig)는 다음과 같다
	/*
		LoggerConfig{
			Skipper: DefaultSkipper,
			Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
				`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
				`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
				`"bytes_out":${bytes_out}}` + "\n",
			CustomTimeFormat: "2006-01-02 15:04:05.00000",
			Output:           os.Stdout,
			colorer:          color.New(),
		}
	*/

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// 별도의 config와 함께 미들웨어 추가

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Something")
	})

	e.Start(":8000")
}
