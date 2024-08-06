package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("What is your year of birth?")
	var age int
	fmt.Scanln(&age)
	// Get the current time
	now := time.Now()
	age = now.Year() - age
	fmt.Println("Hello, ", name, "! You are currently ", age, "years old.")
}
