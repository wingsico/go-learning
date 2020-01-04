package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

const enHelloPrefix = "Hello, "

// Hello : get Hello String
func Hello(s string) string {
	return enHelloPrefix + s
}
