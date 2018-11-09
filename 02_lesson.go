package main

import (
	"fmt"
)

func firstLoop(until int) {
	for i := 0; i < until; i++ {
		fmt.Println(" ", i)
	}
}

func lesson02() {
	firstLoop(10)
}
