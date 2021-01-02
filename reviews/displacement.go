package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v0, s0, t float64

	fmt.Print("acceleration: ")
	fmt.Scan(&a)

	fmt.Print("initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("initial displacement: ")
	fmt.Scan(&s0)

	fmt.Print("time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return a/2*math.Pow(t, 2) + v0*t + s0
	}
}
