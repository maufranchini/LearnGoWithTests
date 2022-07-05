package main

import "fmt"

const englishHelloPrefish = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	} 
	return englishHelloPrefish + name
}

func main() {
	fmt.Println(Hello("world"))
}
