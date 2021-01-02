// https://www.coursera.org/learn/golang-functions-methods/home/welcome
// Functions, Methods, and Interfaces in Go
// week 2 : Peer-graded Assignment: Module 2 Activity
/* Instructions
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s =½ a t^2 + vo * t + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
 */
package main

import (
	"fmt"
)

/*
 *
 */
func GenDisplaceFn(a float64, vo float64, so float64) func(float64)  float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * t * t + vo * t + so)
	}
	return fn
}
/*
 *
 */
func main() {
	var err error
	var acceleration, initial_velocity, initial_displacement, time float64

	fmt.Printf("Enter acceleration: ")
	_, err = fmt.Scan(&acceleration)
	if err != nil {
		fmt.Printf("Error -->%s<-*\n", err)
	}
	fmt.Printf("Enter initial_velocity: ")
	_, err = fmt.Scan(&initial_velocity)
	if err != nil {
		fmt.Printf("Error -->%s<-*\n", err)
	}
	fmt.Printf("Enter initial_displacement: ")
	_, err = fmt.Scan(&initial_displacement)
	if err != nil {
		fmt.Printf("Error -->%s<-*\n", err)
	}
	// generate a function fn which will compute displacement as a function of time
	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Printf("Enter time : ")
	_, err = fmt.Scan(&time)
	if err != nil {
		fmt.Printf("Error -->%s<-*\n", err)
	}
	fmt.Println(fn(time))
}
