package main

import "fmt"

func main() {
	fmt.Println(Hello("monde"))
}

func Hello(name string) string {
	return "Hello " + name
}
