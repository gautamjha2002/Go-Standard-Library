package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name) // we can also use %q for quoted string check documentation
	fmt.Printf("Hello %s! nice to meet you", name)
}
