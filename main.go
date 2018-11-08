package main

import (
	"fmt"
	"math/rand"
)

// lesson 01 - general intro
func lesson01() {

	const (
		myPi = 3.14
		bigv = 1 << 100
	)

	// VAriable declaration
	// var x int           // outside func we need to use alway var in front
	// var y, z int = 0, 1 // variable declaration with initializer

	// k := 1 // declaration and init of a varialbe (inside the func)

	fmt.Println("vim-go random number ", rand.Intn(13))
	// fmt.Println("variablesssss  ", x, y, z, k)
	fmt.Println("just susing a constant here: ", myPi)

	v := 1
	fmt.Println("start v=", v)
	v = v << 1
	fmt.Println("v one shift", v)
	v = v << 1
	fmt.Println("v one shift", v)

	// type conversion T(v) - convert v to type T
	fmt.Println("v one shift", float64(bigv))
}

func lesson02() {
}

// main module for the go tour
func main() {
	lesson01()
}
