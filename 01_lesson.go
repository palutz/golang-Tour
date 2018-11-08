package main

import (
	"fmt"
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
