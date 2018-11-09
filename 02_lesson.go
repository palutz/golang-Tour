package main

import (
	"fmt"
	"math"
	"runtime"
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
	}
	return z
}

func switcF() {
	fmt.Println("Running on...")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// any other os...
		fmt.Println("%s", os)
	}
}

func defer01() {
	defer fmt.Println("Stefano")

	fmt.Printf("Hello, ")
}

func defer02() {
	// counting reverse
	fmt.Println("counting defer....")
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// defer are pushed in a stack (LIFO)
}

func lesson02() {
	firstLoop(10)
	fmt.Println(
		myPow(3, 2, 10),
		myPow(3, 3, 20),
	)
	a := []int{1, 9, 23, 49, 144}
	for _, e := range a {
		fmt.Println("r= ", mySqrt(float64(e)))
		fmt.Println("sqrt = ", math.Sqrt(float64(e)))
	}
	switcF()
	defer01()
	defer02()
}
