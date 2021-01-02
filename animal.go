// https://www.coursera.org/learn/golang-functions-methods/home/welcome
// Functions, Methods, and Interfaces in Go
// week 3 : Peer-graded Assignment: Module 3 Activity
/* Instructions
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
	1) the food that it eats,
	2) its method of locomotion, and
	3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss
Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
	The Eat() method should print the animal’s food,
	the Move() method should print the animal’s locomotion, and
	the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}
func (a Animal) Eat() {
	fmt.Println(a.food)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var name string
	var info string
	var animal Animal
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Usage:<name> <information>")
	fmt.Println("with possible value for <name> cow,bird,snake")
	fmt.Println("with possible value for <information> eat,move,speak")

	for {
		animal = Animal{"", "", ""}
		name = ""
		info = ""
		fmt.Printf(">")
		in.Scan()
		response := strings.ToLower(in.Text())
		// check if there are 2 words
		words := strings.Fields(response)
		if len(words) == 2 {
			name = (strings.Split(response, " "))[0]
			info = (strings.Split(response, " "))[1]
			switch name {
			case "cow":
				animal = Animal{"grass", "walk", "moo"}
			case "bird":
				animal = Animal{"worms", "fly", "peep"}
			case "snake":
				animal = Animal{"mice", "slither", "hsss"}
			default:
				fmt.Println("unknow animal " + name )
			}
			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("unknow information " + info)
			}
		} else {
			fmt.Println("invalid request, try again")
		}
	}
}
