package main

import (
	"fmt"
	"math/rand"

	"github.com/enescakir/emoji"
)

func main() {

	var num int
	for {
		fmt.Println("enter number")

		fmt.Scanln(&num)
		randNum := rand.Intn(6)
		fmt.Println(randNum)
		if num == randNum {
			println("Random number ", randNum, " is matched with your enterd number", emoji.Parse("You Won :sunglasses:"))
			break

		}

	}
}
