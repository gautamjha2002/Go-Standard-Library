package main

import (
	"flag"
	"fmt"
)

func main() {

	archPtr := flag.String("arch", "x86", "CPU Type")
	// arch is what we have to put with - sign
	// x86 is default value
	// CPU Type is help tag
	flag.Parse()

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	case "IA64":
		fmt.Println("Remember IA64?🤔")
	}
	fmt.Println("Process Complete")
	// command (argument) (argument)
	// command -flag
}
