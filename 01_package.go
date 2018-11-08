package main

import (
	"fmt"
	"math/rand"
)

// function returning more than one result
func mySwap(x string, y string) (string, string) {
	return y, x
}

func notExported() {
	fmt.Println("Starting with lowercase is a not experoted function")
}

func ExportedOne(x int) {
	fmt.Println("Starting with uppercase is an experoted functioni, with one param=", x)
}

func main() {
	fmt.Println("vim-go random number ", rand.Intn(13))
}
