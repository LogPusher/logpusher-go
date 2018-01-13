package main

import (
	"fmt"
	"lpusherlib/src"
	"time"
)

func main() {

	client := logpusher.New("me@amazon.com", "strongpass", "logpusherapikey")

	result, err := client.Push("My awesome log message",
		"myawesomesite.com",
		"E-commerce",
		"Notice",
		"EVENT0",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}
