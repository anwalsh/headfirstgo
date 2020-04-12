package main

import "fmt"

func Hello() string {
	return "Hello, world!"
}

func main() {
	fmt.Println(Hello())

	anon := func() {
		fmt.Println("Hello from anon!")
	}
	anon()

	anon = func() {
		fmt.Println("Another anon...")
	}
	anon()
}
