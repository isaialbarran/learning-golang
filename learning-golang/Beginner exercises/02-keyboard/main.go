package main

import (
	"bufio"
	"fmt"
	"os"
)

//Make a program that lets the user input a name

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	name, _ := reader.ReadString('\n')
	fmt.Print("Your name is ", name)

}
