package main

import "fmt"

// déclaration des constantes
const english = "English"
const french = "French"
const spanish = "Spanish"

const prefixHelloEnglish = "Hello, "
const prefixHelloFrench = "Bonjour, "
const prefixHelloSpanish = "Hola, "

func main() {
	fmt.Println(Hello("monde"))
}

func Hello(name string) string {
	return prefixHelloEnglish + name
}
