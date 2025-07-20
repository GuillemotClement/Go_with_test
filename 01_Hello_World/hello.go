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
	fmt.Println(Hello("Jean", spanish))
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	if lang == spanish {
		return prefixHelloSpanish + name
	}

	if lang == french {
		return prefixHelloFrench + name
	}
	
	return prefixHelloEnglish + name
}
