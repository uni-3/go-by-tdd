package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "hello, "
const spanishHelloPrefix = "hola, "

// domain, available for test
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
