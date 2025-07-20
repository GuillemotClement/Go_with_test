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

	return prefixHello(lang) + name
}

func prefixHello(lang string) (prefix string) {
	// use switch for test all case
	switch lang {
	case french:
		prefix = prefixHelloFrench
	case spanish:
		prefix = prefixHelloSpanish
	default:
		prefix = prefixHelloEnglish
	}
	// retourne automatiquement le prefix -> déclarer dans le retour de la fonction
	return
}
