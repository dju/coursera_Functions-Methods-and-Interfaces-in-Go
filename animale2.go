// https://www.coursera.org/learn/golang-functions-methods/home/welcome
// Functions, Methods, and Interfaces in Go
// week 4 : Peer-graded Assignment: Module 4 Activity
/* Instructions
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
	or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user.
Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomotion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
	The first string is “newanimal”.
	The second string is an arbitrary string which will be the name of the new animal.
	The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings.
	The first string is “query”.
	The second string is the name of the animal.
	The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
	The Eat() method should print the animal’s food,
	the Move() method should print the animal’s locomotion, and
	the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//
type Cow struct {
	food       string
	locomotion string
	noise      string
}
type Bird struct {
	food       string
	locomotion string
	noise      string
}
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// func / method
//
func (c Cow) Eat()     { fmt.Println(c.food) }
func (c Cow) Move()    { fmt.Println(c.locomotion) }
func (c Cow) Speak()   { fmt.Println(c.noise) }
func (b Bird) Eat()    { fmt.Println(b.food) }
func (b Bird) Move()   { fmt.Println(b.locomotion) }
func (b Bird) Speak()  { fmt.Println(b.noise) }
func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomotion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func main() {
	var command, param1, param2 string
	var animal, cow, bird, snake Animal
	var animals map[string]Animal
	cow = Cow{"grass", "walk", "moo"}
	bird = Bird{"worms", "fly", "peep"}
	snake = Snake{"mice", "slither", "hsss"}
	animals = map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Usage: <newanimal|query> <param1> <param2>")
	fmt.Println("to add an animal")
	fmt.Println("newanimal animal_name <animal_type>")
	fmt.Println("Or to query animal information")
	fmt.Println("query <animal_type> <information>")
	fmt.Println("with possible value for <animal_type> cow,bird,snake")
	fmt.Println("with possible value for <information> eat,move,speak")
	for {
		command, param1, param2 = "", "", ""
		fmt.Printf(">")
		in.Scan()
		response := strings.ToLower(in.Text())
		// check if there are 3 words
		words := strings.Fields(response)
		if len(words) == 3 {
			command = (strings.Split(response, " "))[0]
			param1 = (strings.Split(response, " "))[1]
			param2 = (strings.Split(response, " "))[2]
			switch command {
			case "newanimal":
				{
					_, ok := animals[param1]
					if ok {
						fmt.Println("animal name " + param1 + " already known")
					} else {
						switch param2 {
						case "cow":
							{
								animals[param1] = cow
								fmt.Println("Created it!")
							}
						case "bird":
							{
								animals[param1] = bird
								fmt.Println("Created it!")
							}
						case "snake":
							{
								animals[param1] = snake
								fmt.Println("Created it!")
							}
						default:
							fmt.Println("unknow animal type " + param2)
						}
					}
				}
			case "query":
				{
					_, ok := animals[param1]
					switch ok {
					case true:
						{
							animal = animals[param1]
						}
					case false:
						{
							fmt.Println("unknow animal " + param1)
						}
					}
					switch param2 {
					case "eat":
						animal.Eat()
					case "move":
						animal.Move()
					case "speak":
						animal.Speak()
					default:
						fmt.Println("unknow information " + param2)
					}
				}
			}
		} else {
			fmt.Println("invalid request, try again")

		}
	}
}
