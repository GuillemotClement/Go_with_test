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

	// set valeur par defaut
	prefix := prefixHelloEnglish

	// use switch for test all case
	switch lang {
	case english:
		prefix = prefixHelloEnglish
	case french:
		prefix = prefixHelloFrench
	case spanish:
		prefix = prefixHelloSpanish
	}

	return prefix + name
}
