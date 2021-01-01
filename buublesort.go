// https://www.coursera.org/learn/golang-functions-methods/home/welcome
// Functions, Methods, and Interfaces in Go
// week 1 : Peer-graded Assignment: Module 1 Activity: Bubble Sort Program
/* Instructions
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

code contains a function called BubbleSort() which takes a slice of integers as an argument
contains a function called Swap() function which takes two arguments, a slice of integers and an index value i,
 */

package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/*
 *
 */
func Swap(data []int,index int){
	data[index+1], data[index] = data[index], data[index+1]
}
/*
 *
 */
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < (len(data) - i); j++ {
			if data[j] < data[j-1] {
				Swap(data, j-1)
			}
		}
	}
}
/*
 *
 */
func Tests(dataSlice []int){
	//tests
	test1 := []int {3,7,2,9,0,1,2,3,8,4}
	test2 := []int {3,7,-2,9,0,1,2,3,8,4}
	test3 := []int {0,0,0,-2,0,0,0,0,-1,0}
	copy(dataSlice,test1)
	fmt.Println("test 1")
	fmt.Println(dataSlice)
	BubbleSort(dataSlice)
	fmt.Println(dataSlice)

	copy(dataSlice,test2)
	fmt.Println("test 2")
	fmt.Println(dataSlice)
	BubbleSort(dataSlice)
	fmt.Println(dataSlice)

	copy(dataSlice,test3)
	fmt.Println("test 3")
	fmt.Println(dataSlice)
	BubbleSort(dataSlice)
	fmt.Println(dataSlice)
}
func main() {
	var dataSlice = make([]int,10,10)
    //tests
	Tests(dataSlice)
	fmt.Println("Enter 10 integers separeted with space")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	response := in.Text()

	numbers := strings.Split(response," ")
	if len(numbers) != 10 {
		fmt.Println("There are not 10 integers , completed with 0")
	}
	var tmpdataSlice = make([]int, 0, 10)
	for _, value := range numbers {
		tmp, _ := strconv.Atoi(value)
		tmpdataSlice = append(tmpdataSlice, tmp)
	}
	copy(dataSlice, tmpdataSlice)
	fmt.Println("Before :")
	fmt.Println(dataSlice)
	BubbleSort(dataSlice)
	fmt.Println("After  :")
	fmt.Println(dataSlice)
}
