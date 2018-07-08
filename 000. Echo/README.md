# [Echo](https://github.com/labstack/echo)
martini, iris, gin과 함께 Go의 대표적인 마이크로 웹 프레임워크

```
go get -u github.com/labstack/echo/...
```

- Ruby의 웹 프레임워크인 Sinatra를 기반으로 설계
- 대부분의 프레임워크들처럼, 콜백에 전달되는 context 인자로 요청에 대한 문맥을 관리
- Echo 구조체에서 파생되는 Group과 Router 구조체를 이용한 어플리케이션 구조화 지원
- 미들웨어 개념 지원
- 내부적으로 라우팅 테이블을 radix tree로 관리하여 빠른 경로 탐색 가능

[공식 문서](https://echo.labstack.com/guide)는 그리 친절하지 않다. 웬만하면 소스 코드 뜯어가며 배워야 한다(..)