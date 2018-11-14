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
 
type MyFloat float64
// can declare method also on not struct type
// ONLY on type in the same package!
func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

func lesson03() {
	v := Vertex { 3, 4 }
	fmt.Println("abs ", v.Abs())
	// calling the normal function 
	// same functionality and result but different syntax
	fmt.Println("abs ", Abs(v))
}