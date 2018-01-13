package main

import (
	"fmt"
	logpusher "logpusher-go/src"
	"time"
)

func main() {

	client := logpusher.New("xx@xx.com", "xXXXx", "xxxXXXxxx")

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

	fmt.Println(result.Message)
}
