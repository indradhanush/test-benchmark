package main

import (
	"fmt"
	_ "net/http/pprof"
	"time"
	"math/rand"
)

func Calculate(x, y int64) int64 {
	// get random integer between 0 and 100
	n := rand.Intn(100)
	// sleep for a random duration between 0 and 100 milliseconds
	time.Sleep(time.Duration(n) * time.Millisecond)

	return x + y
}

func main() {
	fmt.Println(Calculate(1, 2))
}
