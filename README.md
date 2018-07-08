# Awesome-Go
Supports GC, Concurrent, Compilation Language by Google

<https://golang.org/>

## Go
2007년 구글에서 개발을 시작하여 2012년에 버전 1.0을 완성한 프로그래밍 언어이며, 바둑(Go)와 구별하기 위해 Golang이라고도 부른다. 구글의 V8 엔진에 참여했던 Robert Griesemer, 유닉스 개발에 참여했던 Rob Pike, B 언어의 개발자인 Ken Thompson이 함께 개발했다.

- 바이너리 컴파일
- 타입 시스템 : 강타입, 정적 타이핑
- 타입 추론 : 명시적
- Scope : 정적, 블록 스코프
- 함수 : 일급 함수
- 패러다임 : 멀티 패러다임
- 객체지향 표현 : 구조체와 인터페이스, duck typing 기반
- 메모리 관리 : 별도의 VM 없는 가비지 컬렉션
- 컨벤션 : 코딩 스타일을 고집하지 않지만 특정 스타일을 강력하게 제안하며, 이러한 스타일의 구현으로 Go에 기본으로 내장된 go fmt라는 유틸리티 사용
- 기타 특별한 점
    - Loop를 위해 for만을 지원
    - if에도 for처럼 전처리 statement를 사용할 수 있음
    - 포인터를 지원하지만, 포인터에 관련된 복잡성을 제거하기 위해 포인터 간 연산은 불가능
    - 구조체를 지원하고, 태그 개념을 통해 재사용성을 높임
    - 초경량 스레드를 사용하는 고루틴과 채널이라는 강력한 동시성을 지원
    
수십 년간 연구해온 프로그래밍 언어론을 모두 제끼고 마이웨이로 개발했다고 일부 비판의 목소리도 있다.

## GOROOT!!! GOPATH!!!
내가 처음 Go를 시작했을 때, 이 두 개의 환경변수를 세팅하다가 안 되길래 Go를 한동안 안 배웠다. 이 둘을 천천히 알아보면,

### GOROOT
GOROOT는 Go가 설치된 위치다. Java의 JAVA_HOME같은 느낌이다. Go 홈페이지에서 .tar.gz로 된 Go를 내려받았다면 압축을 푼 경로가 될 수도 있고, Mac에서 Homebrew로 설치했다면 `/usr/local/Cellar/go/1.10.3/libexec`과 비슷하게 생긴 경로가 된다. 일반적으로 Go가 설치될 때 함께 세팅된다.

### GOPATH
대부분의 프로그래밍 언어에서 project root는 그때 그때 바뀌는데(Python이나 Node는 프로그램이 시작된 디렉토리), Go의 경우에는 project root가 GOPATH/src로 고정된다. 공식 문서에선 GOPATH를 아래와 같이 설명하고 있다.

```
The GOPATH environment variable is used to specify directories outside of $GOROOT that contain the source for Go projects and their binaries.
```

go와 관련된 모든 명령을 실행할 때 패키지와 명령어 등이 참조되는 위치가 GOPATH다. `GOPATH/pkg/`에는 다운로드된 외부 라이브러리가 모이며, `GOPATH/bin/`에는 컴파일한 바이너리 파일들이 모이고, `GOPATH/src/`는 우리가 생성한 모든 프로젝트의 root가 된다. 내가 만든 패키지 import도 GOPATH/src/를 기준으로 찾고, 어떤 디렉토리에서 go 명령어를 입력하더라도 컴파일러는 무조건 GOPATH만 참조한다.

Python처럼 '현재 디렉토리'가 프로젝트의 root로 되는 게 아니라서, 일반적으로 GOPATH/src/ 하위에만 프로젝트 디렉토리를 생성하는 식으로 운용한다. 프로젝트 이름이 'test'라면 `GOPATH/src/test/`를 만들어 프로젝트를 진행하는 식이다. 여기에 p 패키지가 있다면, import path는 `"test/p"`일 것이다. GOROOT와 GOPATH처럼 go에 관한 환경 변수는 아래 커맨드를 입력해 확인할 수 있다.

`$ go env`

나만이 아니라 많은 사람들이 이걸 맘에 안들어 하긴 하지만, go만의 문화니까 그러려니 하자.

## Go의 웹 프레임워크
루비의 Ruby on Rails와 Sinatra, 파이썬의 Django와 Flask, Node.js의 Express처럼 메인 프레임워크들이 대체적으로 정해져 있는데, Go는 딱히 그런 게 없다. Go의 대표적인 풀 스택 웹 프레임워크는 beego와 revel이 있고, 마이크로 웹 프레임워크는 martini와 iris, gin, echo가 있다. 개인적으로 풀 스택 프레임워크는 별로 좋아하지 않아서, 마이크로 웹 프레임워크들을 비교해 보자.

- [martini](https://github.com/go-martini/martini) : 초반에 인기를 끌었으나 [Go의 사상과 맞지 않는다는 의견들](https://stephensearles.com/three-reasons-you-should-not-use-martini/)이 많았다고 한다.
- [iris](https://github.com/kataras/iris) : 같은 Go 프레임워크들 중에서도 성능이 정말 좋지만, 개발자가 이슈나 PR을 고의적으로 수정하거나 닫아버리는 등 도덕적인 논란이 많았다.
- [gin](https://github.com/gin-gonic/gin) : martini와 비슷한 API를 가지지만, 더 빠르다고 소개하고 있다. 괜찮은 것 같다.
- [echo](https://github.com/labstack/echo) : gin보다 빠르다고 소개하고 있다. 소개만 보면 `martini < gin < echo`라고 보면 되겠다.

다들 속도를 계속해서 강조하지만, Go를 쓴다는 것 자체가 성능의 이점을 상당 부분 먹고 들어가게 된다. 따라서 성능이 아니라 '활발한 활동이 이뤄지는가?'를 봤을 때 현재는 echo가 가장 우위에 있다.
