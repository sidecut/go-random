package main

import (
	"fmt"
	"time"
)

func main() {
	var xs [32]int64
	xs[0] = time.Now().UTC().UnixNano()
	for i := 1; i < 32; {
		t := time.Now().UTC().UnixNano()
		if t != xs[i-1] {
			xs[i] = t
			i++
		}
	}
	for i := 1; i < 32; i++ {
		fmt.Println(xs[i] - xs[i-1])
	}
}
