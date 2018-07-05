# Awesome-Go
Supports GC, Concurrent, Compilation Language by Google

<https://golang.org/>

## Go
2007년 구글에서 개발을 시작하여 2012년에 버전 1.0을 완성한 프로그래밍 언어이며, 바둑(Go)와 구별하기 위해 Golang이라고도 부른다. 구글의 V8 엔진에 참여했던 Robert Griesemer, 유닉스 개발에 참여했던 Rob Pike, B 언어의 개발자인 Ken Thompson이 함께 개발했다. 단순하고 간결한 프로그래밍 언어를 지향하여, Java의 절반에 해당하는 25개의 키워드만으로 프로그래밍이 가능하게 했다. CSP(Communicating Sequential Processes) 스타일의 강력한 동시성 프로그래밍, 정적 타이핑과 강타입, 가비지 컬렉션, 헤더 파일 관련의 의존관계 문제를 해결하여 매우 빠른 컴파일, 높은 생산성과 성능이 특징이다.

Go는 많은 프로그래밍 언어들의 장점들을 최대한 많이 수용하였고, 따라서 아래와 같은 특징을 가진다.

- 정적 타이핑 강타입 언어지만 타입 추론을 지원한다.
- for만으로 다른 언어의 while를 표현할 수 있다.
- if에도 for의 초기값 설정처럼 statement를 사용할 수 있다.
- 함수에서 여러 개의 값을 한 번에 반환할 수 있다.
- Python의 argument unpacking처럼, slice를 함수의 인자에 넘길 수 있다.
- Go에서 함수는 1급 함수(first class function)이며, 따라서 런타임 생성과 익명 생성이 가능하다.
- 포인터를 지원하지만, 포인터에 관련된 복잡성을 제거하기 위해 포인터 간 연산은 불가능하다.
- 구조체를 지원하고, 태그 개념을 통해 재사용성을 높인다.
- 인터페이스와 duck typing 기반의 자연스러운 추상화를 지원한다.
- 고루틴과 채널이라는 강력한 동시성을 지원한다.
- go fmt라는 내장 기능 덕분에, 컨벤션이 통일되어 있다.
- 별도의 VM 없이 가비지 컬렉션을 제공한다.
- 클래스와 상속, 예외, 제네릭, assertion, 오버로딩을 지원하지 않는다.

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

## gofmt
