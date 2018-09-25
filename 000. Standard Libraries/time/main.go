package main

import (
	"fmt"
	"time"
	// https://golang.org/pkg/time/
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Weekday(), now.Hour(), now.Minute())
	fmt.Println(now.Date())

	now.Add(time.Hour + 6)
	// func (Time) Add
	now.AddDate(0, 7, 11)
	// func (Time) AddDate

	time.Sleep(time.Millisecond * 20)
}
