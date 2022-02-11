package main

import (
	"fmt"
	"rsc.io/quote"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
	fmt.Println(quote.Go())
}
