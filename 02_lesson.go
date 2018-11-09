package main

import (
	"fmt"
	"math"
)

// if with statement
func myPow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func firstLoop(until int) {
	for i := 0; i < until; i++ {
		fmt.Println(" ", i)
	}
	j := 10
	for j < 20 {
		j++
		fmt.Println("j= ", j)
	}
}

func mySqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("z= ", z)
	}
	return z
}

func lesson02() {
	firstLoop(10)
	fmt.Println(
		myPow(3, 2, 10),
		myPow(3, 3, 20),
	)
	mySqrt(10)
}
