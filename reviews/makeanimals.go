/* Make animals */

package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	fmt.Println("Make animals")

	animals := make(map[string]Animal)

	for {
		var command, name, param string

		fmt.Printf("> ")
		fmt.Scan(&command, &name, &param)

		switch command {
		case "newanimal":
			a := makeAnimal(param)
			if a != nil {
				animals[name] = a
				fmt.Println("Created it!")
			}
		case "query":
			a,p := animals[name]
			if p {
				printAttribute(a, param)
			} else {
				fmt.Println("No animal with that name is registered, try again [command name type|attribute]")
			}			
		default:
			fmt.Println("Invalid command, try again [newanimal|query].")
		}		
	}
}

func makeAnimal(anType string) Animal {
	switch anType {
	case "cow":
		return new(Cow)
	case "bird":
		return new(Bird)
	case "snake":
		return new(Snake)
	default:
		fmt.Println("Invalid animal type, try again [cow|bird|snake].")
		return nil
	}
}

func printAttribute(a Animal, attribute string) {
	switch attribute {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Invalid query, try again [eat|move|speak].")
	}

	return
}
