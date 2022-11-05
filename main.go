package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func typeStarsWithDelay(n int, ch chan bool) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("")
	}
	ch <- true
}

func main() {
	args := os.Args[1:]
	n, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("Enter a number")
	} else {
		ch1 := make(chan bool)
		go typeStarsWithDelay(int(n), ch1)
		<-ch1
	}
}
