package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello, "
const spanishHelloPrefix = "hola, "
const frenchHelloPrefix = "bonjour, "

// domain, available for test
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) string {
	var prefix string

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
