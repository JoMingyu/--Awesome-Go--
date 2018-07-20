package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		file, err := c.FormFile("file")
		// (c *context) FormFile(name string) (*multipart.FileHeader, error) 함수를 통해 파일 업로드를 제어한다
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		src, err := file.Open()
		// (fh *FileHeader) Open() (File, error) 함수를 통해 파일을 열어 주고
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(file.Filename)
		// 요청에서 온 파일의 이름대로 empty file을 생성
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			// 열린 파일 src를 dst로 copy
			return err
		}
		// 결과적으로, 현재 디렉토리에 요청 당시의 파일과 동일한 이름의 파일이 생성되며
		// path나 파일 이름을 임의로 수정하려면, os.Create 구문을 수정하면 된다

		return c.NoContent(http.StatusOK)
	})

	e.Start(":8000")
}
