package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	time.Sleep(5 * time.Second)

	time2 := time.Now()

	elapsedTime := time2.Sub(time1).Seconds()
	fmt.Println(elapsedTime)
}
