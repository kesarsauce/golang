package main

import "fmt"

const spanish =  "spanish"
const french = "french"
const hindi = "hindi"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const hindiHelloPrefix = "Namaskar, "
const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Sam2", ""))
}
