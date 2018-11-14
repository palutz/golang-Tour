package main 

import (
	"fmt"
	"math"
)

// interface and methods on tpye

type Vertex struct {
	X, Y float64
}

// method on vertex
// A method is a function with a "receiver"
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// same abs code but as a normal function - without receiver 
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func lesson03() {
	v := Vertex { 3, 4 }
	fmt.Println("abs ", v.Abs())
	// calling the normal function 
	// same functionality and result but different syntax
	fmt.Println("abs ", Abs(v))
}