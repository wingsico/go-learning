package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

// Hello : get Hello String
func Hello(s string) string {
	return "Hello, " + s
}
