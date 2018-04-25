package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, c.FormValue("arg"))
		// application/x-www-form-urlencoded, application/form-data에 대응
	})

	e.POST("/2", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			// 파일이 없는 경우 err
			return err
		}

		src, err := file.Open()
		if err != nil {
			// 파일에 문제가 있는 경우 err
			return err
		}
		defer src.Close()

		dst, err := os.Create(file.Filename)
		if err != nil {
			// 파일이 제대로 저장되지 않은 경우 err
			return err
		}
		defer dst.Close()

		return c.String(http.StatusOK, "Hello!")
	})

	e.Start(":5000")
}
