package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2-unstable/bson"
)

// go get github.com/globalsign/mgo
// MongoDB ODM을 지원하는 라이브러리

type user struct {
	// 스키마를 구조체로 정의하되, MongoDB document에 들어갈 매핑의 key 이름을 설정하려면 bson 태그를 붙여준다
	// 태그를 붙이지 않으면, 필드 이름을 모두 소문자화하여 설정한다(UserName의 경우 username으로 설정)
	// PK를 설정하려면 bson 태그에 "_id"를 붙여주자
	Username string `bson:"username"`
	Pw       string
}

// EmbeddedDocument는 구조체 임베딩을 활용하며, ReferenceField는 참조할 구조체 타입의 필드를 사용한다

func main() {
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	// MongoDB에 대한 session 생성
	if err != nil {
		panic("")
	}

	db := session.DB("test")
	// 데이터베이스 선택

	userCol := db.C("users")
	// Collection 선택

	userCol.Insert(user{"PlanB", "pw"})
	userCol.Insert(user{"geni429", "pw"})
	// Insert

	var users []user
	userCol.Find(bson.M{}).All(&users)
	// Find all
	// bson.M을 통해 조건을 전달하고(MongoDB 쿼리 형태), 이후 반환되는 *Query 구조체의 메소드인 All()로
	// 조건에 맞는 document를 모두 구조체 슬라이스 포인터에 append
	fmt.Println(users)

	usr := new(user)
	if err := userCol.Find(bson.M{"username": "PlanB"}).One(usr); err == nil {
		// Find one
		// 조건에 맞는 document를 하나만 구조체 포인터에 바인딩
		// 조건에 맞는 document가 없다면 err에 nil이 아닌 값이 전달됨
		fmt.Println(*usr)
	}

	userCol.UpdateAll(bson.M{"username": "PlanB"}, bson.M{"$set": bson.M{"pw": "newPw"}})
	// UpdateAll(selector interface{}, update interface{})
	// 조건에 맞는 document에 대해 update를 실행
	userCol.Find(bson.M{"username": "PlanB"}).One(usr)
	fmt.Println(*usr)

	userCol.Remove(bson.M{"username": "geni429"})
	// 조건에 맞는 document를 하나 제거

	userCol.RemoveAll(bson.M{})
	// 조건에 맞는 document를 모두 제거
}
