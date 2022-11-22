package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", name)
	fmt.Println("Our current version of Go is " + runtime.Version())
	fmt.Printf("We are running in %v \n", runtime.GOOS)

}
