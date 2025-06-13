package main

import "fmt"

func main() {
	// print string
	print("Hello Sam\n")

	// print numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}


	// wait for user input and echo it
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("You entered: %s\n", input)

}
