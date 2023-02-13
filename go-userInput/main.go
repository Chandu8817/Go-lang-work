// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type UserInfo struct {
	Id        string
	FirstName string
	LastName  string
}

func AddUserInfo(_id string, _first string, _last string) {

	var person1 UserInfo
	person1.Id = _id
	person1.FirstName = _first
	person1.LastName = _last

	fmt.Println(person1)

}

func main() {

	fmt.Println("entering user info ID: ")
	var Id string
	fmt.Scanln(&Id)
	fmt.Println("entering user info FirstName: ")
	var first string
	fmt.Scanln(&first)
	fmt.Println("entering user info LastName: ")
	var last string
	fmt.Scanln(&last)

	// var then variable name then variable type

	AddUserInfo(Id, first, last)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}
