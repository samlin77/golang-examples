package main


import (
	"fmt"
	"os"
)

func main() {
	
	// print numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	// wait for user input and echo it
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("You entered: %s\n", input)

	// read file content and print it
	data, err := os.ReadFile("lorem")
	if err != nil {
		fmt.Println("lorem doesn't exist")
	} else {
		fmt.Println(string(data))
	}

}
