package main

import "fmt"

func main() {
	fmt.Println(Hello("World", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const defaultHelloName = "world"
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "

// Hello : get Hello String
func Hello(name string, language string) string {
	if name == "" {
		name = defaultHelloName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
