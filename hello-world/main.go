package main

import "fmt"

const englishHelloPrefix = "hello, "

// domain, available for test
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
