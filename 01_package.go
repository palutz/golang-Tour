package main

import (
	"fmt"
	"math/rand"
)

func notExported() {
	fmt.Println("Starting with lowercase is a not experoted function")
}

func ExportedOne(x int) {
	fmt.Println("Starting with uppercase is an experoted functioni, with one param=", x)
}

// function returning more than one result
func mySwap(x string, y string) (string, string) {
	return y, x
}

// naked return function#: usign named param for return (USE ONLY WITH SMALL FUNCTIONS)
func addRes(x, y int) (r int, err string) {
	if x > 0 && y > 0 {
		r = x + y
		err = ""
	} else {
		r = 0
		err = "negative number"
	}
	return // returning the named param r and err
}

const myPi = 3.14

// VAriable declaration
var x int           // outside func we need to use alway var in front
var y, z int = 0, 1 // variable declaration with initializer

func main() {
	k := 1 // declaration and init of a varialbe (inside the func)

	fmt.Println("vim-go random number ", rand.Intn(13))
	fmt.Println("variablesssss  ", x, y, z, k)
	// type conversion T(v) - convert v to type T
	fmt.Println("type conversion ", float64(x))
	fmt.Println("just susing a constant here: ", myPi)
}
