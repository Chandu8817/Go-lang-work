// Program to pass pointer as a function argument

package main

import "fmt"

// function definition with a pointer argument
func update(num *int) {

	fmt.Println("enter update number: ")
	var _num int
	fmt.Scanln(&_num)
	// dereference the pointer
	*num = _num

}

func main() {

	var number = 55

	// function call
	update(&number)

	fmt.Println("The number is", number)

}
